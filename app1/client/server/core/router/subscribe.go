package router
import (
	// "fmt"
	// "strings"

	// "github.com/ethereum/go-ethereum/common"
	// log "github.com/sirupsen/logrus"
	// "gitlab.com/meta-node/client/network"
	// "gitlab.com/meta-node/client/network/messages"
	cn "gitlab.com/meta-node/meta-node/pkg/network"
)
type Subscribe struct {
	server        *Server
	subscribeChan chan interface{}
	connRoot      *cn.Connection
	subscribeConn *cn.Connection

}
// func (subscribe *Subscribe) initSub() {

// 	subscribe.subscribeChan = make(chan interface{})
// 	subscribe.connRoot = network.ConnectToServer(STORAGEHOST, subscribe.subscribeChan)

// 	subscribe.subscribeChain(subscribe.connRoot)
// }
// func (subscribe *Subscribe) subscribeChain(connRoot *cn.Connection) {
// 	// fmt.Println("Storage host:", STORAGEHOST)
// 	contractsub := subscribe.server.contractABI
// 	cn.SendBytes(subscribe.server.config, connRoot, messages.SubscribeToAddress, common.FromHex(contractsub["token0"].Address))
// 	log.Info("Listen address: ", contractsub["token0"].Address)
// }
// func (subscribe *Subscribe) handleSubscribeMessage() {
// 	server := subscribe.server
// 	broadcast := subscribe.subscribeChan
// 	cake := server.contractABI["token0"]

// 	for {
// 		fmt.Println("start handleSubscribe")
// 		// capture event from chain
// 		msg := (<-broadcast).(network.EventI)
// 		// handle format event
// 		sendData := Message1{}
// 		switch msg.Address {
// 		case strings.ToLower(cake.Address):
// 			// cake.handleCakeMessage(msg, &sendData)
// 		default:
// 			// contract.handleAllTokenMessage(msg, &sendData)
// 		}

// 		//   Send it out to every client that is currently connected
// 		// for client := range clients {
// 		// 	sendQueue[client] <- sendData
// 		// }
// 		//   Send it out to every client that is currently connected
// 		log.Info(" - Send to all player - ")
// 		for _, client := range server.clients.data {
// 			client.sendChan <- sendData
// 			// client.caller.getFarmPoolInfoUpdate(msg.Data, msg.Topics)
// 			// conn := subscribe.connRoot
// 		}

// 	}
// }



