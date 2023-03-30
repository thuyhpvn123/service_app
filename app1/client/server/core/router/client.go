package router

import (
	"fmt"
	// "log"

	// "log"

	// "strings"
	// "net/http"
	// "strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
	hdlDapp "gitlab.com/meta-node/client/handlers/dapp_handler"
	// hdlGroupDapp "gitlab.com/meta-node/client/handlers/groupDapp_handler"
	hdlSmartContract "gitlab.com/meta-node/client/handlers/smartcontract_handler"
	hdlTransaction "gitlab.com/meta-node/client/handlers/transaction_handler"
	hdlWallet "gitlab.com/meta-node/client/handlers/wallet_handler"
	"gitlab.com/meta-node/client/models"
	"gitlab.com/meta-node/client/utils"
	// cc "gitlab.com/meta-node/core/controllers"
	// "gitlab.com/meta-node/client/server/core/router"
	// "gitlab.com/meta-node/client/model"
	// cn "gitlab.com/meta-node/core/network"
	"gitlab.com/meta-node/client/config"
	"gitlab.com/meta-node/client/controllers"
	"gitlab.com/meta-node/meta-node/pkg/bls"
	"gitlab.com/meta-node/meta-node/pkg/network"
	"gitlab.com/meta-node/meta-node/pkg/state"
	// "github.com/ethereum/go-ethereum/common"

)

type Client struct {
	ws     *websocket.Conn
	server *Server
	caller CallData
	sync.Mutex
	sendChan chan Message1
	hub      Hub
	uid      string

	keyPairMap          map[string]*bls.KeyPair
	config             *config.Config
	messageSenderMap      map[string]network.IMessageSender
	connectionsManager network.IConnectionsManager
	tcpServerMap          map[string]network.ISocketServer

	transactionControllerMap map[string]controllers.ITransactionController
	accountStateChan      chan state.IAccountState
}

func (client *Client) init() {
	// send init message
	// client.ws.WriteJSON(
	// 	Message{Type: "message", Msg: "Here is new client"})
	client.caller = CallData{server: client.server, client: client}
	go client.handleMessage()
	go client.sendInitData()
	log.Info("End init client")
}

// Broadcast struct
type Broadcast struct {
	uid     string
	address string
	message []byte
}

// // NewConnection return new Connection object.
// func NewConnection(ws *websocket.Conn, uid string, hub *Hub) *Connection {
// 	return &Connection{
// 		ws:   ws,
// 		uid:  uid,
// 		send: make(chan []byte, 256),
// 		hub:  hub,
// 	}
// }

func (client *Client) handleListen() {
	for {
		// Read in a new message as JSON and map it to a Message object
		var msg map[string]interface{}
		err := client.ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			client.ws.Close()
			break
		}
		// log.Info("Message from client: ", msg)
		client.handleCallChain(msg)
	}
}


// handle message struct tu chain tra ve va chuyen qua dang JSON gui toi cac client
func (client *Client) handleMessage() {
	for {
		msg := <-client.sendChan
		// msg1 := <-sendDataC
		log.Info(msg)
		err := client.ws.WriteJSON(msg)

		if err != nil {
			log.Printf("error: %v", err)
			client.ws.Close()
		}
	}
}
func (client *Client) handleCallChain(msg map[string]interface{}) {
	// var message []string
	db, err := sqlx.Connect("sqlite3", "./database/doc_2022-12-26_09-16-03.meta_findsdk.db")
	if err != nil {
		log.Fatal(err)
	}
	dappCtrl := hdlDapp.NewDappController(db)
	walletCtrl := hdlWallet.NewWalletController(db)
	smartContractCtrl := hdlSmartContract.NewSmartContractController(db)
	// groupDappCtrl := hdlGroupDapp.NewGroupDappController(db)
	transactionCtrl := hdlTransaction.NewTransactionController(db)



	// if msg.Msg != nil {
		// message = strings.Split((msg.Msg).(string), "\n")
	// }
	// handle call
	switch msg["command"] {
		//wallet
	case "get-all-wallet":
		kq := walletCtrl.GetAllWallets()
		header := models.Header{ Success:true, Data: kq}
		kq1 := utils.NewResultTransformer(header)

		go client.caller.sentToClient("desktop","get-all-wallet", false, kq1)
	case "get-wallet-pagination":
		limit:= int(msg["limit"].(float64))
		page:= int(msg["page"].(float64))
		kq := walletCtrl.GetWalletPagination(page,limit)
		go client.caller.sentToClient("desktop","get-wallet-pagination", false, kq)
	case "get-balance-wallet":
		call:=msg["value"].(map[string]interface{}) 
		kq:= client.caller.GetBalanceWallet(call)
		header := models.Header{ Success:true, Data: kq}
		kq1 := utils.NewResultTransformer(header)

		go client.caller.sentToClient("desktop","getBalanceWallet", false, kq1)
	case "getWalletAtAddress":
		call:=msg["value"].(map[string]interface{}) 
		kq:= client.caller.GetWalletAtAddress(call)
		header := models.Header{ Success:true, Data: kq}
		kq1 := utils.NewResultTransformer(header)

		go client.caller.sentToClient("desktop","getWalletAtAddress", false, kq1)

	case "get-wallet-info":
		address:= msg["address"].(string)
		client.caller.GetWalletInfo(address)
	case "get-wallet":
		kq := walletCtrl.GetAllWallets()
		header := models.Header{ Success:true, Data: kq}
		kq1 := utils.NewResultTransformer(header)

		go client.caller.sentToClient("desktop","get-wallet", false, kq1)
	case "update-wallet-position":
		call:=msg["value"].(map[string]interface{}) //chua test

		go client.caller.UpdateWalletPosition(call)
	// case "update-wallet-by-address":
	// 	call:=msg["value"].(map[string]interface{}) //chua test . proto thay đổi ko có balance nên ko cần luu

	// 	go client.caller.UpdateWalletByteAddress(call)


		// smart contract
	case "get-all-smart-contract":
		kq := smartContractCtrl.GetSmartContracts()
		go client.caller.sentToClient("desktop","get-all-smart-contract", false, kq)
	case "get-last-smart-contract-excute":
		address:= msg["address"].(string)

		kq := smartContractCtrl.GetLastTransactionSmartContract(address)
		header := models.Header{ Success:true, Data: kq}
		kq1 := utils.NewResultTransformer(header)

		go client.caller.sentToClient("desktop","get-last-smart-contract-excute", false, kq1)  //chua test
	case "excute-smart-contract":
		call:=msg["value"].(map[string]interface{}) //chua test

		client.caller.ExecuteSmartContract(call)
		// go client.caller.sentToClient("desktop","excute-smart-contract", false, kq)  //chua test

		// group dapp
	case "get-all-group-d-app":
		call:=msg["value"].(map[string]interface{}) //chua test
		go client.caller.GetAllGroupDApps(call)

	case "delete-group-d-app":
		call:=msg["value"].(map[string]interface{}) //chua test

		go client.caller.DeleteGroupDApp(call)  
		//whitelist
	case "get-white-list-pagination":
		call:=msg["value"].(map[string]interface{}) //chua test
		go client.caller.GetWhitelistPagination(call)
		//dapp
	case "get-all-d-app-no-group":
		kq := dappCtrl.GetAllDappNoGroup()
		go client.caller.sentToClient("desktop","get-all-d-app-no-group", false, kq)
	case "insert-dapp":
		call:=msg["value"].(map[string]interface{})

		go client.caller.InsertDapp(call)
	case "update-d-app-page":
		call:=msg["value"].(map[string]interface{})

		go client.caller.UpdateDAppPage(call)
	case "update-d-app-group-id":
		call:=msg["value"].(map[string]interface{})

		go client.caller.UpdateDAppGroupId(call)

	case "delete-d-app":
		call:=msg["value"].(map[string]interface{}) //chua test

		go client.caller.DeleteDApp(call)
	case "get-all-d-app-for-browser":
		call:=msg["value"].(map[string]interface{}) //chua test

		go client.caller.GetAllDAppsForBrowser(call) //chua test
	case "get-all-d-app":
		kq := dappCtrl.GetAllDApps()   //chua test
		header := models.Header{ Success:true, Data: kq}
		kq1 := utils.NewResultTransformer(header)
		go client.caller.sentToClient("desktop","get-all-d-app", false, kq1)

	case "update-d-app-position":
		call:=msg["value"].(map[string]interface{}) //chua test

		go client.caller.UpdateDAppPosition(call)
	case "update-d-app-page-and-position":
		call:=msg["value"].(map[string]interface{}) //chua test

		go client.caller.UpdateDAppPageAndPosition(call)

	case "get-d-app-by-bundle-id":
		bundleId:= msg["bundle-id"].(string)
		kq := dappCtrl.GetDappByBundleId(bundleId)
		go client.caller.sentToClient("desktop","get-d-app-by-bundle-id", false, kq) // chua test 
	case "get-all-d-app-by-group-id":
		call:=msg["value"].(map[string]interface{}) //chua test

		go client.caller.GetAllDAppsByGroupId(call)
	case "rename-group-d-app":
		call:=msg["value"].(map[string]interface{}) //chua test

		go client.caller.RenameGroupDApp(call)
	case "update-group-d-app-position":
		call:=msg["value"].(map[string]interface{}) //chua test

		go client.caller.UpdateGroupDAppPosition(call)
		//transaction
	case "get-all-trans":
		kq := transactionCtrl.GetAllTransaction()
		header := models.Header{ Success:true, Data: kq}
		kq1 := utils.NewResultTransformer(header)

		go client.caller.sentToClient("desktop","get-all-trans", false, kq1)  //test fail
	case "get-transaction-by-hash":
		hash:= msg["hash"].(string)
		kq,_ := transactionCtrl.GetTransactionByHash(hash)
		go client.caller.sentToClient("desktop","get-transaction-by-hash", false, kq)
	case "get-transaction-pagination":
		limit:= int(msg["limit"].(float64))
		page:= int(msg["page"].(float64))
		kq := transactionCtrl.GetTransactionPagination(page,limit)
		go client.caller.sentToClient("desktop","get-transaction-pagination", false, kq)   
	case "get-transaction-by-address-wallet":
		address:= msg["address"].(string)
		limit:= int(msg["limit"].(float64))
		page:= int(msg["page"].(float64))
		status:= int(msg["status"].(float64))
		kq := transactionCtrl.GetTransactionByAddress(page,limit,address,status)
		go client.caller.sentToClient("desktop","get-transaction-by-address-wallet", false, kq)

		//other
	case "check-performance":
		go client.caller.checkperformance()
	case "get-person-info":
		header := models.Header{ Success:true, Data: "hard code"}
		kq1 := utils.NewResultTransformer(header)
	
		go client.caller.sentToClient("desktop","console-log", false, kq1)

	case "download":
		message:= msg["message"].(string)

		go client.caller.CheckVersion(message)
	// case "GetBalance":
	// 	go client.caller.getBalance(message[0], message[1]) //chua test

	case "send-transaction":
		fmt.Println("send-transaction")
		call:=msg["value"].(map[string]interface{})

		result:=client.caller.TryCall(call) 
		 header := models.Header{ Success:true,Data: result}
		 kq := utils.NewResultTransformer(header)
	 
		 go client.caller.sentToClient("desktop","send-transaction", false,kq)
	 
	case "deploy-sc":
		fmt.Println("deploy-sc")
		call:=msg["value"].(map[string]interface{})

		client.caller.DeploySmartContract(call) 
	case "deploy-d-app":
		fmt.Println("deploy-d-app")
		call:=msg["value"].(map[string]interface{})

		go client.caller.DeployDApp(call) 

	case "get-my-setting":
		header := models.Header{ Success:true, Data: "hard code"}
		kq1 := utils.NewResultTransformer(header)
	
		go client.caller.sentToClient("desktop","get-my-setting", false, kq1)
	case "show-bottom":
		header := models.Header{ Success:true, Data: "hard code"}
		kq1 := utils.NewResultTransformer(header)
	
		go client.caller.sentToClient("desktop","show-bottom", false, kq1)

	case "get-raw-seed":

		client.caller.GetRawSeed(msg)
	case "console-log":
		header := models.Header{ Success:true, Data: "hard code"}
		kq1 := utils.NewResultTransformer(header)
	
		go client.caller.sentToClient("desktop","console-log", false, kq1)
	case "has-device-notch":
		header := models.Header{ Success:true, Data: "hard code"}
		kq1 := utils.NewResultTransformer(header)
	
		go client.caller.sentToClient("desktop","has-device-notch", false, kq1)
	case "get-biometric-type":
		header := models.Header{ Success:true, Data: "hard code"}
		kq1 := utils.NewResultTransformer(header)
	
		go client.caller.sentToClient("desktop","get-biometric-type", false, kq1)

	case "create-wallet":
		fmt.Println("CreateWallet")
		call:=msg["value"].(map[string]interface{})
		client.caller.CreateWallet(call)
	case "init-app":
		client.caller.InitAppService()
	case "encode":
		call:=msg["value"].(map[string]interface{})

		client.caller.EncodeAbi(call)

	// case "create-deviceKey":
	// 	client.caller.CreateDeviceKey()
	case "connect-node":
		fmt.Println("connect-node")
		call:=msg["value"].(map[string]interface{})
		client.caller.ConnectSocket(call)
	case "check-amount":
		call:=msg["value"].(map[string]interface{}) //chua test
		go client.caller.CheckAmount(call)
	case "get-node":
		call:=msg["value"].(map[string]interface{}) //chua test
		go client.caller.GetNode(call)
	case "set-wallet-active-d-app":
		call:=msg["value"].(map[string]interface{}) //chua test
		go client.caller.SetWalletActiveDApp(call)
	case "get-setting-d-app":
		call:=msg["value"].(map[string]interface{}) //chua test
		go client.caller.GetSettingDApp(call)
	case "get-status-connected":
		call:=msg["value"].(map[string]interface{}) //chua test
		go client.caller.GetStatusConnected(call)
	case "get-public-key":
		call:=msg["value"].(map[string]interface{}) //chua test
		go client.caller.GetPublicKey(call)
	case "get-sign":
		call:=msg["value"].(map[string]interface{}) //chua test
		go client.caller.GetSign(call)
	case "delete-list-d-app-and-smart-contract":
		call:=msg["value"].(map[string]interface{}) //chua test
		go client.caller.DeleteListDAppAndSmartContract(call)
	case "get-password-from-seedphrase":
		call:=msg["value"].(map[string]interface{}) //chua test
		go client.caller.GetPasswordFromSeedPhrase(call)
	case "delete-d-app-table":
		go client.caller.DeleteDAppTable()
	case "disconnect-wallet":
		call:=msg["value"].(map[string]interface{}) //chua test
		go client.caller.DisconnectWallet(call)
	case "edit-wallet-ui":
		call:=msg["value"].(map[string]interface{}) //chua test
		go client.caller.EditWalletUI(call)
	case "create-hash-from-file":
		call:=msg["value"].(map[string]interface{}) //chua test
		go client.caller.CreateHashFromFile(call)
	case "verify-json-app":
		call:=msg["value"].(map[string]interface{}) //chua test
		go client.caller.VerifyJSONApp(call)
	case "deleteAllStorage":
		go client.caller.DeleteAllStorage()


	// case "get-content-backup":
	// 	call:=msg["value"].(map[string]interface{}) //chua test
	// 	go client.caller.GetContentBackup(call)

	// case "init-connection":
	// 	fmt.Println("init-connection")
	// 	call:=msg["value"].(map[string]interface{})
	// 	client.caller.InitCallConnection(call)
	// case "test":
	// 	fmt.Println("test")
	// 	address:=msg["address"].(string)
	// 	client.caller.GetWalletKeyFromAddress(address)
	// case "init-connection":
	// 	fmt.Println("init-connection")
	// 	address:=msg["value"].(string)
	// 	client.caller.InitCallConnection(address)

	default:
		log.Warn("Require call not match: ", msg)
	}

	// switch msg["method"] {
	// case "connect-node":

	// }
	// }
}

func (client *Client) sendInitData() {
	go client.server.database.transferToChan(client.sendChan)
}

func convertToArrayInterface(s []string) []interface{} {
	array := make([]interface{}, len(s))
	for i, v := range s {
		array[i] = v
	}
	return array
}
// func parseJson(str *Message1)string{
// 	strjson,err:= json.Marshal(str)
// 	if err != nil{
// 		log.Fatal()
// 	}
// 	return strjson
// }