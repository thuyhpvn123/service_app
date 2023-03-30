package router

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
	"gitlab.com/meta-node/client/config"
	// config "gitlab.com/meta-node/client/config"
	// "gitlab.com/meta-node/client/controllers"
	"gitlab.com/meta-node/client/utils"
	"gitlab.com/meta-node/meta-node/pkg/bls"
	"gitlab.com/meta-node/meta-node/pkg/logger"
	"gitlab.com/meta-node/meta-node/pkg/network"
	// "gitlab.com/meta-node/meta-node/pkg/state"
	c_config "gitlab.com/meta-node/client/config"
	"gitlab.com/meta-node/client/controllers"
	// "github.com/ethereum/go-ethereum/common"

)

var accounts = [...]Account{
	{
		Address: "aa39344b158f4004cac70bb4ace871a9b54baa1e",
		Private: "5808195a0d285c98dd942b7602f180ebbaa57ba15622786147a924d8e29daf4a",
	},
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ClientList struct {
	sync.RWMutex
	data map[*websocket.Conn]Client
}
type Database struct {
	sync.RWMutex
	data []Message1
}

type ListString struct {
	sync.Mutex
	data []string
}
type QueueLock struct {
	sync.Mutex
	queue map[string]bool
}

type Server struct {
	sync.Mutex
	clients           ClientList
	broadcast         chan interface{}
	database          Database
	subscribe         Subscribe
	availableAccounts chan Account
	contractABI       map[string]*ContractABI
	config            *config.Config
	hub               Hub
}

// connect to chain
var queue = QueueLock{queue: make(map[string]bool)}

// NewWsController return new WSController object.
// func NewWsController(hub *Hub) *WSController {
// 	return &WSController{hub}
// }
const (
	CONFIG_FILE_PATH = "config/conf.json"
)

func (clients *ClientList) Remove(ws *websocket.Conn) {
	clients.Lock()
	delete(clients.data, ws)
	clients.Unlock()
	log.Warn("Client disconnection")
}
func (server *Server) GiveBackAccount(account Account) {
	queue.Lock()
	fmt.Println("give back account ")
	queue.queue[account.Address] = false
	fmt.Println("end give back account ")
	queue.Unlock()
	server.availableAccounts <- account
}
func (server *Server) Init(config *config.Config) *Server {
	// init subscriber
	server.config = config
	server.contractABI = make(map[string]*ContractABI)
	var wg sync.WaitGroup
	for _, contract := range contracts {
		wg.Add(1)
		go server.getABI(&wg, contract)
	}
	wg.Wait()

	// connected clients

	server.clients.data = make(map[*websocket.Conn]Client)

	// broadcast channel
	server.broadcast = make(chan interface{})
	// available account map
	server.availableAccounts = make(chan Account, len(accounts))

	for _, account := range accounts {
		go server.GiveBackAccount(account)
	}
	// account1 := <-server.availableAccounts
	// fmt.Println("account day:",account1)
	server.subscribe = Subscribe{server: server}

	// server.subscribe.initSub()
	server.hub = Hub{}
	fmt.Println("the end")
	return &Server{
		clients:           ClientList{},
		broadcast:         server.broadcast,
		database:          Database{},
		subscribe:         server.subscribe,
		availableAccounts: server.availableAccounts,
		contractABI:       server.contractABI,
		config:            config,
		hub:               Hub{},
	}
}
func (server *Server) Start(port int) {

	//start subscriber
	// go server.subscribe.handleSubscribeMessage()

	// sure subscribeConn close
	// defer server.subscribe.subscribeConn.Close()

	// log.Info("http server started on :", port)
	// err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	// if err != nil {
	// 	log.Fatal("ListenAndServe: ", err)
	// }
}
func (server *Server) getABI(wg *sync.WaitGroup, contract Contract) {
	var temp ContractABI
	temp.initContract(contract)
	server.Lock()
	server.contractABI[contract.Name] = &temp
	server.Unlock()
	wg.Done()
}

func (server *Server) WebsocketHandler(w http.ResponseWriter, r *http.Request) {

	upgrader.CheckOrigin = func(r *http.Request) bool {
		return true

	}
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	// Register our new client
	// clients[conn] = true
	// sendQueue[conn] = make(chan Message)
	// generate new id with size as md5 string
	uid := utils.GenerateRandomMD5String(16)
	config, err := c_config.LoadConfig(CONFIG_FILE_PATH)
	if err != nil {
		logger.Error(fmt.Sprintf("error when loading config %v", err))
		panic(fmt.Sprintf("error when loading config %v", err))
	}
	cConfig := config.(*c_config.Config)

	// accountStateChan := make(chan state.IAccountState, 1)
	client := Client{ws: conn, sendChan: make(chan Message1), server: server, hub: Hub{}, uid: uid,
	keyPairMap  : make(map[string]*bls.KeyPair),
	config  :cConfig,
	messageSenderMap : make(map[string]network.IMessageSender),
	// connectionsManager :network.IConnectionsManager{},
	tcpServerMap      : make(map[string]network.ISocketServer),
	// stop              : false,
	// commands          : make(map[int]string),

	transactionControllerMap :make(map[string]controllers.ITransactionController),
	// accountStateChanMap      :make(map[string](chan accountStateChan)),

	}
	client.init()
	// conn.WriteJSON(

	// 	Message{Command: "test-message", Msg: " welcome to Metanode"})

	log.Println("Client Connected successfully") //write on server terminal
	// // Make sure we close the connection when the function returns
	// defer conn.Close()
	// client.caller = CallData{server: client.server, client: conn}
	// go client.server.database.transferToChan(client.sendChan) alo
	// log.Info("End init client")
	// add client into list to listen broadcast
	// server.clients.Lock()
	// server.clients.data[conn] = client
	// server.clients.Unlock()
	//make sure remove client
	defer server.clients.Remove(conn)

	//listen websocket
	client.handleListen()

}

// functiont to add new message to database
func (database *Database) add(msg Message1) {
	database.Lock()
	database.data = append(database.data, msg)
	database.Unlock()
}

// function to pass all message in database to chan to send to client
func (database *Database) transferToChan(reciever chan Message1) {
	database.RLock()
	for _, item := range database.data {
		reciever <- item
	}
	database.RUnlock()
}
