package main

import (
	// "net/http"

	// "github.com/gorilla/mux"
	// log "github.com/sirupsen/logrus"
	"fmt"
	// "log"

	// "gitlab.com/meta-node/client/config"
	server_app "gitlab.com/meta-node/client/server/app"

	// "gitlab.com/meta-node/client/database"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	hdlWs "gitlab.com/meta-node/client/server/core/router"
	"gitlab.com/meta-node/meta-node/pkg/logger"
	c_config "gitlab.com/meta-node/client/config"

)

// Define object


const (
	CONFIG_FILE_PATH = "config/conf.json"
)


func main() {
	config, err := c_config.LoadConfig(CONFIG_FILE_PATH)
	if err != nil {
		logger.Error(fmt.Sprintf("error when loading config %v", err))
		panic(fmt.Sprintf("error when loading config %v", err))
	}
	cConfig := config.(*c_config.Config)

	// open connection to database
	db, err := sqlx.Connect("sqlite3","./database/doc_2022-12-26_09-16-03.meta_findsdk.db")
	if err != nil {
		logger.Error(fmt.Sprintf("error when connect sqlite %", err))
		panic(fmt.Sprintf("error when connect sqlite %v", err))
	}
	fmt.Println("Connection opened to database")
	db.SetMaxIdleConns(100)
	defer db.Close()
	ConnectDb(db)
	// MySettingDataBase()
	// websocket hub
	hub:= hdlWs.NewHub(db)
	// go hub.Run()
	app := server_app.InitApp(cConfig,db,hub)
	app.Run()

}
//database store my-setting
// func MySettingDataBase(){
// 	db, err := sqlx.Connect("sqlite3","./database/my-setting.db")
// 	if err != nil {
// 		logger.Error(fmt.Sprintf("error when connect sqlite %", err))
// 		panic(fmt.Sprintf("error when connect sqlite %v", err))
// 	}
// 	fmt.Println("Connection opened to mysetting database")
// 	createSettingQuery := `CREATE TABLE IF NOT EXISTS mysetting(
// 		id INTEGER PRIMARY KEY AUTOINCREMENT, 
// 		receiveNotificationFromNode TEXT, 
// 		receiveTransactionStatus TEXT, 
// 		nodeAddress TEXT, 
// 		networkType TEXT, 
// 		isHasPincode TEXT, 
// 		isEnableFaceOrTouchId TEXT,
// 		isWatchConfirm TEXT,
// 		accountInfo TEXT)`       
// 	_, err = db.Exec(createSettingQuery)
// 	if err != nil {
// 		fmt.Println(err)             
// 	}
	
// }	