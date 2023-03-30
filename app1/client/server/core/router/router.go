package router

import (
	// "html/template"
	"net/http"

	"github.com/gorilla/mux"
	// "gitlab.com/meta-node/client/server/app"
	// hdlLog "gitlab.com/meta-node/client/handlers/log_handler"
	// hdlSubscriber "gitlab.com/meta-node/client/handlers/subscriber_handler"
	// hdlTopic "gitlab.com/meta-node/client/handlers/topic_handler"
	// hdlWs "gitlab.com/meta-node/client/server/core/router"
	"github.com/jmoiron/sqlx"
	// hdlWallet "gitlab.com/meta-node/client/handlers/wallet_handler"
	// hdlSmartContract "gitlab.com/meta-node/client/handlers/smartcontract_handler"
	// hdlTransaction "gitlab.com/meta-node/client/handlers/transaction_handler"
	// hdlDapp "gitlab.com/meta-node/client/handlers/dapp_handler"
)

// var tmpl *template.Template

func InitRouter(
	server *Server,db *sqlx.DB,hub *Hub,
) *mux.Router {
	r := mux.NewRouter()
	// controllers
	// logCtrl := hdlLog.NewLogController(db)
	// topicCtrl := hdlTopic.NewTopicController(db)
	// subscriberCtrl := hdlSubscriber.NewSubscriberController(db)
	// walletCtrl := hdlWallet.NewWalletController(db)
	// smartContractCtrl := hdlSmartContract.NewSmartContractController(db)
	// transactionCtrl := hdlTransaction.NewTransactionController(db)
	// dappCtrl := hdlDapp.NewDappController(db)
	// static
	r.PathPrefix("/public").Handler(http.FileServer(http.Dir("frontend")))
	r.HandleFunc("/ws", server.WebsocketHandler)
	// // user handler
	// r.HandleFunc("/api/v1/logs", logCtrl.GetLogs).Methods("GET")
	// r.HandleFunc("/api/v1/topics", topicCtrl.GetTopics).Methods("GET")
	// r.HandleFunc("/api/v1/subscribers", subscriberCtrl.GetSubscribers).Methods("GET")
	// //wallet handler
	// r.HandleFunc("/api/v2/wallets", walletCtrl.GetWallets).Methods("GET")
	// r.HandleFunc("/api/v2/walletPagination", walletCtrl.GetWalletPagination).Methods("POST")
	// r.HandleFunc("/api/v2/walletAtAddress", walletCtrl.GetWalletAtAddress).Methods("POST")
	// //insert wallet
	// r.HandleFunc("/api/v5/insertWallet", walletCtrl.InsertWallet).Methods("POST")

	// //smart contract handler
	// r.HandleFunc("/api/v3/smartContracts", smartContractCtrl.GetSmartContracts).Methods("GET")
	// //transaction handler
	// r.HandleFunc("/api/v4/transactions", transactionCtrl.GetTransactions).Methods("GET")
	// r.HandleFunc("/api/v4/transactionByHash", transactionCtrl.GetTransactionByHash).Methods("POST")
	// r.HandleFunc("/api/v4/transactionPagination", transactionCtrl.GetTransactionPagination).Methods("POST")
	// r.HandleFunc("/api/v4/transactionByAddress", transactionCtrl.GetTransactionByAddress).Methods("POST")
	// //dapp handler
	// r.HandleFunc("/api/v5/dappPagination", dappCtrl.GetDappPagination).Methods("POST")

	return r
}
