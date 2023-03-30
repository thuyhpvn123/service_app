package router

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"

	// "math/big"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
	"gitlab.com/meta-node/client/controllers"
	"google.golang.org/protobuf/proto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jmoiron/sqlx"

	hdlDapp "gitlab.com/meta-node/client/handlers/dapp_handler"
	hdlMysetting "gitlab.com/meta-node/client/handlers/mysetting_handler"
	hdlNode "gitlab.com/meta-node/client/handlers/node_handler"
	hdlWallet "gitlab.com/meta-node/client/handlers/wallet_handler"

	"gitlab.com/meta-node/client/models"
	"gitlab.com/meta-node/client/utils"

	"gitlab.com/meta-node/meta-node/pkg/bls"

	"gitlab.com/meta-node/client/command"
	c_network "gitlab.com/meta-node/client/network"

	c_config "gitlab.com/meta-node/client/config"
	blst "gitlab.com/meta-node/meta-node/pkg/bls/blst/bindings/go"
	cm "gitlab.com/meta-node/meta-node/pkg/common"
	p_common "gitlab.com/meta-node/meta-node/pkg/common"
	"gitlab.com/meta-node/meta-node/pkg/logger"
	"gitlab.com/meta-node/meta-node/pkg/network"
	pb "gitlab.com/meta-node/meta-node/pkg/proto"
	"gitlab.com/meta-node/meta-node/pkg/state"
)

type blstPublicKey = blst.P1Affine

type blstSecretKey = blst.SecretKey

type Result struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
var dbSQL *sqlx.DB
var listWallets []map[string]interface{}
var objSocketInfo = make(map[string]interface{})
var socketId = 1
var rawSeed []string
var ipHost string
var levelDbWallet *leveldb.DB
var (
	connectionTypesForClient = []string{p_common.NODE_CONNECTION_TYPE}
)

// func createWallet(call map[string]interface{}, database *DB, securityDb *SecurityDbModules) {
func (caller *CallData) CreateWallet(call map[string]interface{}) *Result {
		// leveldb, err := leveldb.OpenFile("./db/device_info", nil)
	// if err != nil {
	// 	panic(err)
	// }

	result := &Result{}
	seedPhrase, ok := call["raw-seed-restore"].([]string)
	fmt.Println("seedPhrase:", seedPhrase)
	if ok && len(seedPhrase) != 0 {
		rawSeed = seedPhrase
	}

	bg, _ := call["bg"].(string)
	color, _ := call["color"].(string)
	idSymbol, _ := call["idSymbol"].(string)
	pattern, _ := call["pattern"].(string)

	lastNodeConnected, _ := call["last-node-connected"].(map[string]interface{})

	if len(rawSeed) == 0 {
		result = &Result{
			Success: false,
			Data:    "rawSeed is empty",
		}
		return result
	}

	// Compute SHA-256 hash of seed phrase
	h := sha256.New()
	rawSeedStr := strings.Join(rawSeed, "-")
	// fmt.Println("rawSeedStr:", rawSeedStr)
	rawSeedStr = strings.Replace(rawSeedStr, "\"", "", -1)
	h.Write([]byte(rawSeedStr))
	digestRootFile := h.Sum(nil)
	// fmt.Println("digestRootFile:",digestRootFile)
	// Generate private-public key pair using BLS algorithm
	priKey, pubKey := blsGenPriKey(digestRootFile)
	// if len(priPubKey) == 80 {
	// 	priKey := priPubKey[0:32]
	// 	pubKey := priPubKey[32:80]

	// Hash public key
	// hashPub := sha3.Sum256(pubKey)
	fmt.Println("priKey:", hex.EncodeToString(priKey))

	hashPub := crypto.Keccak256([]byte(pubKey))

	fmt.Println("hash:", hex.EncodeToString(hashPub))

	// Calculate wallet address from hash of public key
	walletAddress := hashPub[12:]
	fmt.Println("walletAddress:", hex.EncodeToString(walletAddress))

	// Store wallet to database
	db, err := sqlx.Connect("sqlite3", "./database/doc_2022-12-26_09-16-03.meta_findsdk.db")
	if err != nil {
		log.Fatal(err)
	}
	dbSQL =db
	name := getUniqueWalletName(db)
	dappCtrl := hdlDapp.NewDappController(db)

	maxPosition := dappCtrl.GetMaxPositionWallet()
	// maxPosition := 100

	wallet := &hdlWallet.WalletModel{
		0,
		hex.EncodeToString(walletAddress),
		name,
		"0000000000000000000000000000000000000000000000000000000000000000",
		"0000000000000000000000000000000000000000000000000000000000000000",
		"0",
		bg,
		color,
		idSymbol,
		pattern,
		maxPosition + 1,
	}
	fmt.Println("wallet:", wallet)
	walletCtrl := hdlWallet.NewWalletController(db)

	walletCtrl.InsertWallet1(wallet)
	// result = &Result{
	// 	Success: true,
	// 	Data:    wallet,
	// }
	// header := models.Header{Success: true, Data: result}
	// kq := utils.NewResultTransformer(header)

	// go caller.sentToClient("desktop", "create-wallet", false, kq)

	// Store priKey and pubKey to security storage
	// leveldb, err := leveldb.OpenFile("./db/wallets", nil)
	// levelDbWallet =leveldb
	// if err != nil {
	// 	panic(err)
	// }
	// defer leveldb.Close()
	objPriPub := map[string]interface{}{
		"priKey": hex.EncodeToString(priKey),
		"pubKey": hex.EncodeToString(pubKey),
	}
	b, err := json.Marshal(objPriPub)
	if err != nil {
		fmt.Println("error:", err)
	}
	callmap := map[string]interface{}{
		"key":  hex.EncodeToString(walletAddress),
		"data": string(b),
	}
	// fmt.Println("inserted key là address:",hex.EncodeToString(walletAddress))
	// fmt.Println("inserted priKey &pubKey là:", string(b))
	 WriteValueStorage(callmap, levelDbWallet)
	// fmt.Println("result1 là:", result1)
	//test get data từ leveldb với key là address ví
	callmap5 := map[string]interface{}{
		"key": hex.EncodeToString(walletAddress),
	}

	result10 := ReadValueStorage(callmap5, levelDbWallet)
	var kq1 map[string]interface{}
	err = json.Unmarshal(result10["value"].([]byte), &kq1)

	// fmt.Println("Get kq tu leveldb:", kq1)

	// Store password-backup to security storage

	callmap1 := map[string]interface{}{
		"key": "password-backup",
	}

	result2 := ReadValueStorage(callmap1, levelDbWallet)
	var hash string
	if result2["value"] == nil {
		hash = fmt.Sprintf("%x", crypto.Keccak256([]byte(pubKey)))
		password := hash[:32]
		callmap3 := map[string]interface{}{
			"key":  "password-backup",
			"data": password,
		}
		result3 := WriteValueStorage(callmap3, levelDbWallet)
		fmt.Println("write password-backup to storage success:", result3["sucess"])

	}

	// response success
	rawSeed = []string{}
	if lastNodeConnected != nil {
		ip, ok := lastNodeConnected["ip"].(string)
		if !ok {
			log.Panic()
		}

		port, ok := lastNodeConnected["port"].(string)
		if !ok {
			log.Panic()
		}

		if ip != "" {
			wallets := call["wallets"].([]interface{})
			wallets = append(wallets, map[string]interface{}{
            "address": hex.EncodeToString(walletAddress),
        })
			//connect node
			// fmt.Println("wallet address day ne:", wallets)

			params := map[string]interface{}{
				"method": "connect-node",
				"node": map[string]interface{}{
					"ip":   ip,
					"port": port,
				},
				"wallets": wallets,
			}
			resultkq := caller.ConnectSocket(params)
			header := models.Header{Success: true, Data: resultkq}
			kq := utils.NewResultTransformer(header)

			go caller.sentToClient("desktop", "connect-node", false, kq)

		}

		result = &Result{
			Success: true,
			Data: map[string]interface{}{
				"name":    name,
				"address": hex.EncodeToString(walletAddress),
				"hash":    hashPub,
			},
		}
	} else {
		result = &Result{
			Success: false,
			Data:    "[1x1000] Gen bls fail",
		}
	}
	header := models.Header{Success: true, Data: result}
	kq := utils.NewResultTransformer(header)

	go caller.sentToClient("desktop", "create-wallet", false, kq)

	return result
}

//connect-node
func (caller *CallData) ConnectSocket(call map[string]interface{}) *Result {
	fmt.Println("ConnectSocket")
	result := &Result{}
	node, ok := call["node"].(map[string]interface{})
	if !ok {
		result = &Result{
			Success: false,
			Data:    "node, wallets is required",
		}
		return result
	}

	if node["ip"] == nil || node["port"] == nil {
		result = &Result{
			Success: false,
			Data:    "node, wallets is required",
		}
		return result
	}

	wallets, ok := call["wallets"].([]interface{})
	if ok {
		for _,v := range wallets{
			listWallets= append(listWallets,v.(map[string]interface{}))
		}
		// listWallets = wallets
	}
	// dbsetting, err := sqlx.Connect("sqlite3","./database/my-setting.db")
	// db, err := sqlx.Connect("sqlite3", "./database/doc_2022-12-26_09-16-03.meta_findsdk.db")

	// if err != nil {
	// 	logger.Error(fmt.Sprintf("error when connect sqlite %", err))
	// 	panic(fmt.Sprintf("error when connect sqlite %v", err))
	// }
	//store node-address ,network-type in sql database
	mysettingCtrl := hdlMysetting.NewMysettingController(dbSQL)
	str:=[]string{node["ip"].(string),node["port"].(string)}
	data:=strings.Join(str,":")
	mysetting := &hdlMysetting.MysettingModel{
		ID:   2,
		ReceiveNotificationFromNode :false,
		ReceiveTransactionStatus :false,
		NodeAddress :data,  
		NetworkType :"TestNet", 
		IsHasPincode :false,
		IsEnableFaceOrTouchId :false,
		IsWatchConfirm :false,
		AccountInfo :"",
		// 0,
		// false,
		// false,
		// data,  
		// "TestNet", 
		// false,
		// false,
		// false,
		// "",
	}
	
	mysettingCtrl.InsertMysetting(mysetting)
// fmt.Println(mysetting)
	// callmap1 := map[string]interface{}{
	// 	"key": "node-address",
	// }

	// result2 := ReadValueStorage(callmap1, levelDbWallet)
	// if result2["value"] == nil {
	// 	str:=[]string{node["ip"].(string),node["port"].(string)}
	// 	callmap3 := map[string]interface{}{
	// 		"key":  "node-address",
	// 		"data": strings.Join(str,":"),
	// 		// "data":node["ip"] ,

	// 	}
	// 	result3 := WriteValueStorage(callmap3, levelDbWallet)
	// 	fmt.Println("write node-address to storage success:", result3["sucess"])

	// }
	// callmap4 := map[string]interface{}{
	// 	"key": "network-type",
	// }

	// result4 := ReadValueStorage(callmap4, levelDbWallet)
	// if result4["value"] == nil {
	// 	callmap4 := map[string]interface{}{
	// 		"key":  "network-type",
	// 		"data": "TestNet",

	// 	}
	// 	result5 := WriteValueStorage(callmap4, levelDbWallet)
	// 	fmt.Println("write network-type to storage success:", result5["sucess"])

	// }


	// fmt.Println("listWallets:",listWallets)
	ip := node["ip"].(string)
	port := node["port"].(string)
	ipHost = ip
	// connection to parent
	config, err := c_config.LoadConfig("config/conf.json")
	if err != nil {
		logger.Error(fmt.Sprintf("error when loading config %v", err))
		panic(fmt.Sprintf("error when loading config %v", err))
	}
	cConfig := config.(*c_config.Config)
	parentPort,_ :=strconv.Atoi(port)
	// connect to parent
	connectionsManager := network.NewConnectionsManager(connectionTypesForClient)

	parentConn := network.NewConnection(
		common.HexToAddress(cConfig.ParentAddress),
		ip,
		parentPort,
		cConfig.ParentConnection.Type,
	)
	accountStateChan := make(chan state.IAccountState)
	chData :=make(chan interface{})
	handler := c_network.NewHandler(accountStateChan,chData)


	err = parentConn.Connect()
	if err != nil {
		logger.Error(fmt.Sprintf("error when connect to parent %v", err))
		// panic(fmt.Sprintf("error when connect to parent %v", err))
	} else {
		// init connection
		connectionsManager.AddParentConnection(parentConn)

		for _, wallet := range listWallets {

			objSocketInfo[wallet["address"].(string)] = socketId

			addressString, ok := wallet["address"].(string)
			if !ok {
				result = &Result{
					Success: false,
					Data:    "node, wallets is required",
				}
			}

			// code của Quang chỉ là init client :createSocketClient
			//sign
			walletKey := GetWalletKeyFromAddress(addressString)

			if walletKey["priKey"] == nil {
				logger.Error(fmt.Sprintf("error when GetWalletKeyFromAddress %", err))
				panic(fmt.Sprintf("error when GetWalletKeyFromAddress %v", err))
			}else{
				priKey := common.FromHex(walletKey["priKey"].(string))
				keyPair := bls.NewKeyPair(priKey)
			// 	keyPair := bls.NewKeyPair(common.FromHex(walletKey["priKey"].(string)))

				logger.Info("Running with key pair: " + "\n" + keyPair.String())
				// // init message sender
				messageSender := network.NewMessageSender(keyPair, config.GetVersion())

				// caller.client.accountStateChanMap =make(map[string](chan state.IAccountState))

				tcpServer := network.NewSockerServer(config, keyPair, connectionsManager, handler)
				tcpServer.OnConnect(parentConn)
				sign:=GetSignInit(addressString)
				// caller.InitConnection( addressString)
				messageSender.SendMessage(parentConn, command.InitConnection, &pb.InitConnection{
					Address: keyPair.GetAddress().Bytes(),
					Type:    cConfig.GetNodeType(),
				}, sign)
				// fmt.Println("keyPair.GetAddress():",keyPair.GetAddress())

				go tcpServer.HandleConnection(parentConn)
		
				// init controller
				transactionCtl := controllers.NewTransactionController(keyPair, messageSender, connectionsManager)
				// init and start client
				// fmt.Println("addressString:",addressString)
				caller.client.keyPairMap[addressString] = keyPair
				caller.client.messageSenderMap[addressString] = messageSender
				caller.client.transactionControllerMap[addressString]= transactionCtl
				caller.client.tcpServerMap[addressString]=tcpServer
				caller.client.accountStateChan=accountStateChan
				// fmt.Println("caller.client.accountStateChan là:",caller.client)

			}
			
			caller.client.connectionsManager= connectionsManager
			caller.client.config = cConfig
			header := models.Header{Success: true, Data: addressString}
			kq := utils.NewResultTransformer(header)

			go caller.sentToClient("desktop", "init-connection", false, kq)	


		}
				// caller.InitConnection( addressString, parentConn)
			// }
	}
	//insert node into recentNodeTB table
	// db, err = sqlx.Connect("sqlite3", "./database/doc_2022-12-26_09-16-03.meta_findsdk.db")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	nodeCtrl := hdlNode.NewNodeController(dbSQL)
	time := time.Now().Unix()
	portCnv, _ := strconv.Atoi(port)
	recentNode := &hdlNode.NodeModel{
		ID:   0,
		IP:   ip,
		Port: portCnv,
		Time: time,
	}

	nodeCtrl.InsertRecentNode(recentNode)
	if err != nil {
		result = &Result{
			Success: false,
			Data:    err.Error(),
		}
	}
	result = &Result{
		Success: true,
		Data:    "",
	}
	return result
}

func getUniqueWalletName(db *sqlx.DB) string {
	fmt.Println("getUniqueWalletName")
	listData := []string{
		"Sirius",
		"Canopus",
		"Arcturus",
		"Alpha Centauri A",
		"Vega",
		"Rigel",
		"Procyon",
		"Achernar",
		"Betelgeuse",
		"Hadar (Agena)",
		"Capella A",
		"Altair",
		"Aldebaran",
		"Capella B",
		"Spica",
		"Antares",
		"Pollux",
		"Fomalhaut",
		"Deneb",
		"Mimosa",
	}

	size := len(listData)
	rand.Seed(time.Now().UnixNano())
	randomValue := listData[rand.Intn(size)]
	// fmt.Println("randomValue:",randomValue)
	walletCtrl := hdlWallet.NewWalletController(db)
	count := walletCtrl.CountWalletByName(randomValue)
	if count == 0 {
		return randomValue
	}
	countWallet := walletCtrl.CountWalletTable()
	fmt.Println("countWallet:", countWallet)
	if countWallet < 20*count {
		return getUniqueWalletName(db)
	} else {
		name := randomValue + "_" + string(count)
		if walletCtrl.CountWalletByName(name) == 0 {
			return name
		} else {
			return getUniqueWalletName(db)
		}
	}
}
// init sign
func GetSignInit(addressString string) cm.Sign {
	addressS := common.FromHex(addressString)

	initConnectCommand := &pb.InitConnection{
		Address: addressS,
		Type:    "Client",
	}

	walletKey := GetWalletKeyFromAddress(hex.EncodeToString(addressS))
	if walletKey == nil {
		logger.Error(fmt.Sprintf("error when get wallet key "))
	}

	dataHash, _ := proto.Marshal(initConnectCommand)
	hash := crypto.Keccak256(dataHash)
	keyPair := bls.NewKeyPair(common.FromHex(walletKey["priKey"].(string)))
	prikey:= keyPair.GetPrivateKey()
	sign := bls.Sign(prikey, hash)
	return sign
}
// init sign
func GetSignGetAccountState(addressString string) cm.Sign {
	addressS := common.FromHex(addressString)
	hash := crypto.Keccak256(addressS)
	walletKey := GetWalletKeyFromAddress(addressString)
	if walletKey["priKey"] == nil {
		logger.Error(fmt.Sprintf("error when get wallet key "))
	}

	keyPair := bls.NewKeyPair(common.FromHex(walletKey["priKey"].(string)))
	prikey:= keyPair.GetPrivateKey()
	sign := bls.Sign(prikey, hash)
	return sign
}

