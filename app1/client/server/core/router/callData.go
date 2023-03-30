package router

import (
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"reflect"
	"strconv"
	"strings"
	"time"
	"github.com/denisbrodbeck/machineid"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	"github.com/syndtr/goleveldb/leveldb"
	"gitlab.com/meta-node/client/models"

	"gitlab.com/meta-node/client/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/holiman/uint256"
	hdlTransaction "gitlab.com/meta-node/client/handlers/transaction_handler"
	"gitlab.com/meta-node/client/command"
	cm "gitlab.com/meta-node/meta-node/pkg/common"

	. "github.com/ethereum/go-ethereum/accounts/abi"
	hdlDapp "gitlab.com/meta-node/client/handlers/dapp_handler"
	"gitlab.com/meta-node/meta-node/pkg/logger"
	pb "gitlab.com/meta-node/meta-node/pkg/proto"
	"gitlab.com/meta-node/meta-node/pkg/state"
	"gitlab.com/meta-node/meta-node/pkg/transaction"
	"gitlab.com/meta-node/client/network"


)
type ActionListenerCallback map[string]interface{}
var levelDbDeviceKey *leveldb.DB
var levelDbSettingdapp *leveldb.DB
var levelDbMysetting *leveldb.DB

// var db *sqlx.DB
var defaultRelatedAddress [][]byte
// type blstSignature = blst.P2Affine
// var dstMinPk = []byte("BLS_SIG_BLS12381G2_XMD:SHA-256_SSWU_RO_POP_")
var (
	ErrorGetAccountStateTimedOut = errors.New("get account state timed out")
	ErrorInvalidAction           = errors.New("invalid action")
)

type CallData struct {
	server *Server
	client *Client
}
type WalletKey struct {
    PriKey []byte
    PubKey []byte
}
func (caller *CallData) TryCall(callMap map[string]interface{},) interface{} {
	i := 0
	var result interface{}
	result = "TimeOut"

	for {
		if i >= 3 {
			break
		}
		if i != 0 {
			time.Sleep(time.Second)
		}
		result = caller.call(callMap)

		if result != "TimeOut" {
			log.Info("Success time - ", i)
			log.Info(" - Result: ", result)
			return result
		}
		i++
	}

	return result
}

func (caller *CallData) call(callMap map[string]interface{},) interface{} {
	fromAddress, _ := callMap["from-address"].(string)
	caller.SendTransaction1(callMap)

	for {

		select {
		case receiver := <-caller.client.tcpServerMap[fromAddress].GetHandler():
			// log.Info("Hash on server", common.BytesToHash(hash.([]byte)))
			// log.Info("Hash from chain", (receiver).(network.Receipt).Hash)
			// if (receiver).(network.Receipt).Hash != common.BytesToHash(hash.([]byte)) {
			// 	continue
			// }
			return (receiver).(network.Receipt1).Value
		case <-time.After(5 * time.Second):
			return "TimeOut"
		}
	}

}


func toHexInt(n *big.Int) string {
    return fmt.Sprintf("%x", n) // or %x or upper case
}

func (caller *CallData) SendTransaction1(call map[string]interface{}) error {
	inputArray, _ := call["inputArray"].([]interface{})
    if inputArray == nil {
        inputArray = []interface{}{}
    }

	inputStr, _ := call["input"].(string)

	imageSc, _ := call["image"].(string)
	feeType, _ := call["feeType"].(string)
    if feeType == "" {
        feeType = "user"
    }
	    fee, _ := call["fee"].(string)
    if fee == "" {
        fee = "1"
    }
    name, _ := call["name"].(string)
    if name == "" {
        name = ""
    }
    abiData, _ := call["abiData"].(string)
    bin, _ := call["bin"].(string)

	functionName, _ := call["function-name"].(string)
	receiveInfo, _ := call["receive-info"].(string)
	message, _ := call["message"].(string)
	tip, _ := call["tip"].(string)
    if tip == "" {
        tip = "0"
    }
	toAddressStr, _ := call["to-address"].(string)
	toAddress:=common.HexToAddress(toAddressStr)
	relatedAddress:=caller.EnterRelatedAddress(call)
	hexAmount, _ := call["amount"].(string)
    if hexAmount == "" {
        hexAmount = "0"
    }
	amount := uint256.NewInt(0).SetBytes(common.FromHex(hexAmount))
	isDeploy, _ := call["is-deploy"].(bool)
    isCall, _ := call["is-call"].(bool)
	fromAddress, _ := call["from-address"].(string)
	var maxGas uint64
	maxGaskq, ok:=call["maxGas"].(string)
	if!ok{
		maxGas = 500000
	}
    maxGas, err1 := strconv.ParseUint(maxGaskq,10,64)
	if err1!=nil  {
		maxGas = 500000
	}
	var maxGasPriceGwei uint64
	maxGasPriceGweikq, ok:=call["maxGas"].(string)
	if!ok{
		maxGasPriceGwei = 10
	}
	maxGasPriceGwei, err1 = strconv.ParseUint(maxGasPriceGweikq,10,64)
	if err1 !=nil  {
		maxGasPriceGwei = 10
	}
	maxGasPrice := 1000000000 * maxGasPriceGwei
	var maxTimeUse uint64
	maxTimeUsekq, ok:=call["maxGas"].(string)
	if!ok{
		maxTimeUse = 60000
	}
	maxTimeUse, err1 = strconv.ParseUint(maxTimeUsekq,10,64)
	if err1!=nil {
		maxTimeUse = 60000
	}

	var action pb.ACTION
	if isDeploy{
		action = pb.ACTION_DEPLOY_SMART_CONTRACT
	}else if isCall{
		action = pb.ACTION_CALL_SMART_CONTRACT
	}

	var data []byte
	// leveldb, err := leveldb.OpenFile("./db/wallets", nil)
	// if err != nil {
 	// 	panic(err)
	// }
	// defer leveldb.Close()

	sign:= GetSignGetAccountState(fromAddress)

	as, err := caller.GetAccountState(fromAddress,sign)
	if err != nil {
		return err
	}
	if action == pb.ACTION_DEPLOY_SMART_CONTRACT {
		data, err = caller.GetDataForDeploySmartContract(call)
		if err != nil {
			panic(err)
		}
		toAddress = common.BytesToAddress(
			crypto.Keccak256(
				append(
					as.GetAddress().Bytes(),
					as.GetLastHash().Bytes()...),
			)[12:],
		)
	}

	if action == pb.ACTION_CALL_SMART_CONTRACT {
		if len(inputStr) > 0 {
			data = common.FromHex(inputStr)
		} else {
		// 	mInput := getSmartContractInput(functionName, inputArray)
		// 	callData = _convertToByteString(mInput[2:])
		// }
		data, err = caller.GetDataForCallSmartContract(call)
		if err != nil {
			panic(err)
		}
	}
}
	// transferFee := caller.client.config.TransactionFee
	// lastBalance := as.GetBalance()
	// u256PendingUse := as.GetPendingBalance()
	// balance = lastbalance + pendinguse - amount - fee
	// balance := uint256.NewInt(0).Add(lastBalance, u256PendingUse)
	// balance = uint256.NewInt(0).Sub(balance, amount)
	// balance = uint256.NewInt(0).Sub(balance, transferFee)
	fmt.Println("GetLastHash:",as.GetLastHash())
	transaction, err := caller.client.transactionControllerMap[fromAddress].SendTransaction(
		as.GetLastHash(),
		toAddress,
		as.GetPendingBalance(),
		amount,
		maxGas,
		maxGasPrice,
		maxTimeUse,
		action,
		data,
		relatedAddress,
	)
	logger.Debug("Sending transaction", transaction)
	if err != nil {
		logger.Warn(err)
	}
	// Save transaction to database
	// Add transaction with status 0(pending)
		time := time.Now().Unix()
	tipBigInt, _ := new(big.Int).SetString(tip, 10)
	// deviceKey:=CreateDeviceKey()
	db, err := sqlx.Connect("sqlite3", "./database/doc_2022-12-26_09-16-03.meta_findsdk.db")
	if err != nil {
		logger.Error(fmt.Sprintf("error when connect sqlite %", err))
		panic(fmt.Sprintf("error when connect sqlite %v", err))
	}

	transCtrl := hdlTransaction.NewTransactionController(db)

	lastTransaction,num := transCtrl.GetLastTransaction(fromAddress)

	if num != 0 {

		if lastTransaction.Status == 0 || lastTransaction.Status == 1 {

			lastTimeTrans := lastTransaction.Time
			distanceTime := time - lastTimeTrans
	
			if distanceTime > 30 {

				transCtrl.UpdateStatusTransactionByHash(4, lastTransaction.Hash)
			}
			// return map[string]interface{}{
			// 	"success": false,
			// 	"message": "You have 1 transaction in progress",
			// }
		}
		if lastTransaction.Status == 3 || lastTransaction.Status == 4 {

			lastTransaction,_ = transCtrl.GetLastTransactionWithStatus(fromAddress, 2)
		}
		// if lastTransaction != nil {
		// preDataTransaction = &pb.Transaction{
		// 	Hash: common.HexToHash(lastTransaction.Hash),
		// }
	}

	lastHash := lastTransaction.Hash
	if lastHash == "" {
		lastHash = "0000000000000000000000000000000000000000000000000000000000000000"
	}
	//get devicekey from leveldb

	callmap1 :=map[string]interface{}{
		"key":"deviceKey",
	}

	result2:=ReadValueStorage(callmap1,levelDbDeviceKey)
	// fmt.Println("deviceKey:",result2["value"])
	//
	deviceKey:=hex.EncodeToString(result2["value"].([]uint8))
	lastKey := deviceKey + lastHash
	lastKeyByteArray:= common.FromHex(lastKey)

	hashLastKey := crypto.Keccak256(lastKeyByteArray)
	amountBigInt, _ := new(big.Int).SetString(toHexInt(amount.ToBig()), 10)
	currentAmount := new(big.Int).Add(tipBigInt, amountBigInt)
	feeBigInt, _ := new(big.Int).SetString(fee, 10)

	transactionInsert := hdlTransaction.TransactionModel{
	// Id: 0,
	Hash: (transaction.GetHash()).Hex(),
	Address: fromAddress,
	ToAddress: hex.EncodeToString(toAddress.Bytes()),
	PubKey: hex.EncodeToString(transaction.GetPubkey().Bytes()),
	Amount: hex.EncodeToString(transaction.GetAmount().Bytes()),
	PendingUse: hex.EncodeToString(transaction.GetPendingUse().Bytes()),
	// Balance: hex.EncodeToString(transaction.GetBalance().Bytes()),
	// Fee: hex.EncodeToString(transaction.GetFee().Bytes()),
	Time: time,
	Status: 0,
	Type: "send",
	PrevHash: hex.EncodeToString(transaction.GetLastHash().Bytes()),
	Sign: hex.EncodeToString(transaction.GetSign().Bytes()),
	Tip: toHexInt(tipBigInt),
	Message: message,
	ReceiveInfo: receiveInfo,
	IsCall: isCall,
	IsDeploy: isDeploy,
	Data: hex.EncodeToString(transaction.GetData()),
	FunctionCall: "",
	TotalBalance: "",
	LastDeviceKey: hex.EncodeToString(hashLastKey),
	}

	if isCall && functionName != "" {
	transactionInsert.FunctionCall = functionName
	}

	if feeType == "user" {

	transactionInsert.TotalBalance = (big.NewInt(0).Add(currentAmount, feeBigInt)).String()
	} else {
	transactionInsert.TotalBalance = (currentAmount).String()
	}

	kq1:=transCtrl.InsertTransaction(&transactionInsert)
	fmt.Println("InsertTransaction:",kq1)
	// add smart contract to db
	dappCtrl := hdlDapp.NewDappController(db)

	if isDeploy {
	// 	scInDB := dappCtrl.GetSmartContractByAddress(hex.EncodeToString(toAddress.Bytes()))
	// 	var id int
	// 	if len(scInDB) > 0 {
	// 		id = scInDB["id"].(int)
	// 	} else {
	// 		id = 0
	// 	}
	sc := hdlDapp.DappModel{
		// Id: id,
		Name: name,
		BundleId: hex.EncodeToString(toAddress.Bytes()),
		AbiData: abiData,
		BinData: bin,
		Status: 0,
		Image: imageSc,
		Type: 2,
		IsInstalled: 1,
	}
	page := dappCtrl.GetLastPage()
    dApps := dappCtrl.GetAllDAppsByGroupIdAndPage(0, page)
    position := 0
    for _, dApp := range dApps {
        if dApp.Type == 3 {
			positionObj := make(map[string]interface{})
			err =json.Unmarshal([]byte(dApp.PositionObj),&positionObj)
			if err != nil{
				logger.Error(fmt.Sprintf("error when Unmarshal insert Dapp %", err))
				panic(fmt.Sprintf("error when Unmarshal insert Dapp %v", err))
		
			}
			width := positionObj["width"].(int)
			height := positionObj["height"].(int)
			position += width * height
        } else {
            position++
        }
    }
    limit := 24
    if position >= limit {
        page++
        position -= limit
    }
    x := position % 4
    y := position / 4
    positionObj := map[string]int{
        "x":      x,
        "y":      y,
        "width":  1,
        "height": 1,
    }
    sc.Page = page
    sc.Position = position
	bPositionObj,_ := json.Marshal(positionObj)
	sc.PositionObj = string(bPositionObj)
	dappCtrl.InsertSmartContract(&sc)
	}
	return err
}




func (caller *CallData) GetAccountState(address string ,sign cm.Sign) (state.IAccountState, error) {
	parentConn := caller.client.connectionsManager.GetParentConnection()
	caller.client.messageSenderMap[address].SendBytes(parentConn, command.GetAccountState, common.FromHex(address), sign)

	select {
	case accountState := <-caller.client.accountStateChan:
		return accountState, nil
	case <-time.After(5 * time.Second):
		return nil, ErrorGetAccountStateTimedOut
	}
	
}


func (caller *CallData)GetWalletInfo(addressString string) {

	sign:= GetSignGetAccountState(addressString)
	as, err :=caller.GetAccountState(addressString,sign)
	if err != nil {
		logger.Error(fmt.Sprintf("error when GetAccountState %", err))
		panic(fmt.Sprintf("error when GetAccountState %v", err))
	}
	call:= map[string]interface{}{
		"address":as.GetAddress(),
		"last_hash":as.GetLastHash(),
		"balance":as.GetBalance(),
		"pending_balance":as.GetPendingBalance(),
	}
	header := models.Header{ Success:true,Data: call}
	kq := utils.NewResultTransformer(header)

	go caller.sentToClient("desktop","get-wallet-info", false,kq)
}

func (caller *CallData)InitAppService() map[string]interface{} {
	result := make(map[string]interface{})
	defer func() {
		if r := recover(); r != nil {
			result["success"] = false
			result["message"] = r
		}
		
	}()

	// Code to initialize the database
	leveldb1, err := leveldb.OpenFile("./db/device_info", nil)
	if err != nil {
		panic(err)
	}
	levelDbDeviceKey = leveldb1
	CreateDeviceKey()
	leveldb2, err := leveldb.OpenFile("./db/wallets", nil)
	levelDbWallet =leveldb2
	if err != nil {
		panic(err)
	}
	leveldb3, err := leveldb.OpenFile("./db/mysetting", nil)
	if err != nil {
		panic(err)
	}
	levelDbSettingdapp =leveldb3
	leveldb4, err := leveldb.OpenFile("./db/setting_d_app", nil)
	if err != nil {
		panic(err)
	}
	levelDbSettingdapp =leveldb4


	// _skClient := createTcpSocket()
	// Code to reset transactions with status 0 or 1 to 4
	iTime := time.Now().Unix() - (1000 * 60 * 5)
	db, err := sqlx.Connect("sqlite3", "./database/doc_2022-12-26_09-16-03.meta_findsdk.db")
	if err != nil {
		log.Fatal(err)
	}

	transactionCtrl := hdlTransaction.NewTransactionController(db)
	err=transactionCtrl.UpdateTransactionWithStatusPending(iTime)
	if err!=nil{
		result["success"] =false
	}else{
		result["success"] = true
	}

	header := models.Header{ Success:true, Data: ""}
	kq := utils.NewResultTransformer(header)

	go caller.sentToClient("desktop","init-app",false, kq)

	return result
}

func CreateDeviceKey()string {
	callmap1 :=map[string]interface{}{
		"key":"deviceKey",
	}

	result2:=ReadValueStorage(callmap1,levelDbDeviceKey)
    if result2["value"] == nil {
        desktopId,err := machineid.ID()
		if err!=nil{
			logger.Error(fmt.Sprintf("error when ReadValueStorage deviceKey %", err))
			panic(fmt.Sprintf("error when ReadValueStorage deviceKey %v", err))
		}

		hash := fmt.Sprintf("%x", crypto.Keccak256([]byte(desktopId)))
		callmap3 :=map[string]interface{}{
			"key":"deviceKey",
			"data":hash,
	
		}
		result3:= WriteValueStorage(callmap3,levelDbDeviceKey)
		fmt.Println("write deviceKey to storage success:",result3["success"])
		// callmap1 :=map[string]interface{}{
		// 	"key":"deviceKey",
		// }
	
		// result4:=ReadValueStorage(callmap1,leveldb)
	
		// fmt.Println("deviceKey:",result4["value"])
    // }else{
	// 	deviceKey:= result2["value"].([]uint8)
		// fmt.Println("deviceKey:",deviceKey)

	}
	fmt.Println("create devicekey success")
	// return result2["value"].(string)
	return hex.EncodeToString(result2["value"].([]uint8))

}
// get priKey &pubKey by address in levelDB database
func GetWalletKeyFromAddress(address string) map[string]interface{} {
	// address, _ := call["address"].(string)

	callmap1 :=map[string]interface{}{
		"key":address,
	}

	result2:=ReadValueStorage(callmap1,levelDbWallet)

	// obj,err := leveldb.Get([]byte(address), nil)
	// 	if err != nil {
	if result2["value"]==nil{
		fmt.Println(111111111)
		return map[string]interface{}{
			"success": false,
			"msg": "",
		}

	} else {
		fmt.Println(22222222)
	var walletKey map[string]interface{}
	err := json.Unmarshal(result2["value"].([]byte),&walletKey)
	if err != nil {
		panic(err)
	}

	priKey:= walletKey["priKey"].(string)
	pubKey:= walletKey["pubKey"].(string)

		return map[string]interface{}{
		"priKey": priKey,
		"pubKey": pubKey,
		}
	}
}                 


func (caller *CallData) sentToClient(platform string,command string,isSocket bool, data *utils.ResultTransformer) {
	caller.client.sendChan <- Message1{platform,command,isSocket, data}
	// sendQueue[caller.client.ws] <- Message{msgType, value}
}


func (caller *CallData) EnterRelatedAddress(call map[string]interface{}) [][]byte {
	var arrmap []map[string]interface{}
	arr,_:=call["relatedAddresses"].([]interface{})
	if call["relatedAddresses"]==nil || len(arr) ==0 {
		return defaultRelatedAddress
	}else{
		for _,v := range arr{
			arrmap= append(arrmap,v.(map[string]interface{}))
		}
	
		var relatedAddStr []string
	
		for _,v := range arrmap{
			relatedAddStr=append(relatedAddStr,v["address"].(string))
		}
		var relatedAddress [][]byte
	
		// temp := strings.Split(relatedAddStr, ",")
		logger.Info("Temp Related Address")
		for _, addr := range relatedAddStr {
			addressHex := common.HexToAddress(addr)
			logger.Info(addressHex)
			relatedAddress = append(relatedAddress, addressHex.Bytes())
		}
		defaultRelatedAddress = append(defaultRelatedAddress, relatedAddress...)
		return relatedAddress
	
	}
}
func (caller *CallData) GetDataForCallSmartContract(call map[string]interface{}) ([]byte, error) {
	 kq:=caller.EncodeAbi(call)
	callData := transaction.NewCallData(kq)
	return callData.Marshal()
}

func (caller *CallData)EncodeAbi(call map[string]interface{}) []byte {
	inputArray, _ := call["inputArray"].([]interface{})
    if inputArray == nil {
        inputArray = []interface{}{}
    }
	functionName, _ := call["function-name"].(string)

	abiData, _ := call["abiData"].(string)
	abiJson, err := JSON(strings.NewReader(abiData))
	if err != nil {
		panic(err)
	}
	

	 var abiTypes []interface{}
	for _, item := range inputArray {
		itemArr:=encodeAbiItem(item)
		for _,v:=range itemArr{
			abiTypes=append(abiTypes,v)
		}

	}
	out, err := abiJson.Pack(functionName,abiTypes[:]...)

	if err != nil {
		panic(err)
	}
	fmt.Println("out:",hex.EncodeToString(out))
	return out
}

func encodeAbiItem(item interface{}) []interface{} {
	var result []interface{}
	var itemMap map[string]interface{}
	if err := json.Unmarshal([]byte(item.(string)), &itemMap); err != nil {
		log.Fatal(err)
	}
	itemType, _ := itemMap["type"].(string)
	switch itemType {
	case "tuple":
		var value []interface{}
		if err := json.Unmarshal([]byte(fmt.Sprintf("%v", itemMap["value"])), &value); err != nil {
			log.Fatal(err)
		}

		var components []interface{}
		if err := json.Unmarshal([]byte(fmt.Sprintf("%v", itemMap["components"])), &components); err != nil {
			log.Fatal(err)
		}

		var abiTypes []interface{}
		for i, component := range components {
			componentBytes, _ := json.Marshal(component)
			componentType, _ := component.(map[string]interface{})["type"].(string)
			if componentType == "tuple" || componentType == "tuple[]" {
				components[i].(map[string]interface{})["value"] = value[i]
				abiTypes = append(abiTypes, encodeAbiItem(componentBytes))
			} else {
				abiTypes = append(abiTypes, getAbiType(componentType, value[i]))
			}
		}
		result = abiTypes
	case "tuple[]":
		var value []interface{}
		if err := json.Unmarshal([]byte(fmt.Sprintf("%v", itemMap["value"])), &value); err != nil {
			log.Fatal(err)
		}

		var components []interface{}
		if err := json.Unmarshal([]byte(fmt.Sprintf("%v", itemMap["components"])), &components); err != nil {
			log.Fatal(err)
		}

		var tuples []interface{}
		for _, v := range value {
			vArray := v.([]interface{})
			var abiTypes []interface{}
			for j, component := range components {
				componentBytes, _ := json.Marshal(component)
				componentType, _ := component.(map[string]interface{})["type"].(string)
				components[j].(map[string]interface{})["value"] = vArray[j]
				if componentType == "tuple" || componentType == "tuple[]" {
					abiTypes = append(abiTypes, encodeAbiItem(componentBytes))
				} else {
					abiTypes = append(abiTypes, getAbiType(componentType, vArray[j]))
				}
			}
			tuples = append(tuples, abiTypes...)
		}
		result =  tuples
	default:
		value := itemMap["value"]

		var arr []interface{}

		result1 := getAbiType(itemType, value)
		result = append(arr, result1)

	}
	return result
}
func getAbiType(dataType string, data interface{}) interface{} {
    if strings.Contains(dataType, "int") {
		params:=big.NewInt(0)
		params, ok := params.SetString(fmt.Sprintf("%v",int64(data.(float64))), 10)

		if !ok {
			log.Warn("Format big int: error")
			return nil
		}		
		return params

    } else {
	switch dataType {
	case "string":
		return data.(string)
	case "bool":
		return data.(bool)
	case "address":
		return common.HexToAddress(data.(string))
	case "uint8":
		intVar, err := strconv.Atoi(data.(string))
		if err != nil {
			log.Warn("Conver Uint8 fail", err)
			return nil
		}
		return uint8(intVar)
	// case "uint", "uint256":
	// 	nubmer := big.NewInt(0)
	// 	nubmer, ok := nubmer.SetString(data.(string), 10)
	// 	if !ok {
	// 		log.Warn("Format big int: error")
	// 		return nil
	// 	}
	// 	return nubmer
	case "array","slice" :
		fmt.Println("array nè")
	rv := reflect.ValueOf(data)
	var out []interface{}
		for i := 0; i < rv.Len(); i++ {
			out = append(out, rv.Index(i).Interface())
		}

	return out
	default:
		return data
	}
}
}

func (caller *CallData) GetDataForDeploySmartContract(call map[string]interface{}) ([]byte, error) {
	contractFileName, _ := call["name"].(string)
    if contractFileName == "" {
        contractFileName = ""
    }
	// b, _ := os.ReadFile("./contracts/" + contractFileName)
    bin, _ := call["bin"].(string)
	contractStorageHost := "34.75.117.76:3051"
	contractStorageAddress := "7fa14c0f8e22e13dfd54726207c88728429a7246"
	deployData := transaction.NewDeployData(common.FromHex(bin), contractStorageHost, common.HexToAddress(contractStorageAddress))
	return deployData.Marshal()
}

func (caller *CallData)DeploySmartContract(call map[string]interface{}) *Result {
	result := &Result{}
    dataDeploy := map[string]interface{}{
        "from-address": call["address"],
        "to-address": "",
        "is-deploy": true,
        "amount": "0",
        "name": call["name"],
        "abiData": call["abiData"],
        "bin": call["binData"],
    }
    fmt.Println("dataDeploy", dataDeploy)
    kq:=caller.TryCall(dataDeploy)
    defer func() {
        if r := recover(); r != nil {
			result = &Result{
				Success: false,
				Data:    r.(error).Error(),
			}
			
        }else{
			result = &Result{
				Success: true,
				Data: kq.(map[string]interface{})["data"].(map[string]interface{})["data"],
			}
		}
		// header := models.Header{Success: true, Data: result}
		kq := utils.NewResultTransformer(result)

		go caller.sentToClient("desktop", "deploy-sc", false, kq)

    }()

	return result

}
func (caller *CallData)DeployDApp(call map[string]interface{})*Result{
    callData := map[string]interface{}{
        "name": call["name"],
        "author": call["author"],
        "hash": call["hash"],
        "size": call["size"],
        "sign": call["sign"],
        "version": call["version"],
        "pathStorage": call["pathStorage"],
        "image": call["image"],
        "time": call["time"],
        "total-transaction": call["total-transaction"],
        "total-wallet": call["total-wallet"],
        "bundle-id": call["bundle-id"],
        "is-update": call["is-update"],
        "orientation": call["orientation"],
        "isLocal": call["isLocal"],
        "fullScreen": call["fullScreen"],
        "statusBar": call["statusBar"],
        "page": call["page"],
        "isInstalled": call["isInstalled"],
        "position": call["position"],
        "id": call["id"],
        "urlWeb": call["urlWeb"],
        "type": call["type"],
        "positionObj": call["positionObj"],
        "style": call["style"],
    }
    kq :=caller.InsertDapp(callData)
    return kq
}

func (caller *CallData) InsertDapp(call map[string]interface{})  *Result {
	result:= &Result{}
    name, _ := call["name"].(string)
    author, _ := call["author"].(string)
    hash, _ := call["hash"].(string)
    size, _ := call["size"].(string)
    sign, _ := call["sign"].(string)
	version, _ := call["version"].(string)
    pathStorage, _ := call["pathStorage"].(string)
    image, _ := call["image"].(string)
    time, _ := call["time"].(int64)
    totalTransaction, _ := call["total-transaction"].(int)
    totalWallet, _ := call["total-wallet"].(int)
    bundleId, _ := call["bundle-id"].(string)
    isUpdate, _ := call["is-update"].(bool)
    orientation, _ := call["orientation"].(string)
    isLocal, _ := call["isLocal"].(int)
    fullScreen, _ := call["fullScreen"].(int)
    statusBar, _ := call["statusBar"].(string)
    page, _ := call["page"].(int)
    isInstalled, _ := call["isInstalled"].(int)
    position, _ := call["position"].(int)
    // id, _ := call["id"].(int)
    urlWeb, _ := call["urlWeb"].(string)
    typeT, _ := call["type"].(int)
    
    if version == "" || pathStorage == "" || image == "" || time == 0 || totalTransaction == 0 || totalWallet == 0 || bundleId == "" || orientation == "" || isLocal == 0 || fullScreen == 0 || statusBar == "" || page == -1 {
        return &Result{
            Success: false,
            Data: "name, author, hash, size, sign, version, path, image, time, total-transaction, total-wallet, bundle-id, orientation, isLocal, fullScreen, statusBar, page is required",
        }
    }
	db, err := sqlx.Connect("sqlite3", "./database/doc_2022-12-26_09-16-03.meta_findsdk.db")
	if err != nil {
		log.Fatal(err)
	}

	dappCtrl := hdlDapp.NewDappController(db)

	dApps:= dappCtrl.GetAllDAppsByGroupIdAndPage(0, page)

	if position == -1 {
		position = 0
		for _, dApp := range dApps {
			if dApp.Type == 3 {
				positionObj := make(map[string]interface{})
				err :=json.Unmarshal([]byte(dApp.PositionObj),&positionObj)
				if err != nil{
					logger.Error(fmt.Sprintf("error when Unmarshal insert Dapp %", err))
					panic(fmt.Sprintf("error when Unmarshal insert Dapp %v", err))
			
				}
				width := positionObj["width"].(int)
				height := positionObj["height"].(int)
				position += width * height
			} else {
				position++
			}
		}
		limit := 24

		if position >= limit {
			page++
			position -= limit
		}
	}
	x := position % 4
	y := position / 4
	positionObj := make(map[string]interface{})
	positionObj["x"] = x
	positionObj["y"] = y
	positionObj["width"] = 1
	positionObj["height"] = 1
	positionObject,_:=json.Marshal(positionObj)

	// positionJSON, err := json.Marshal(positionObj)
	// if err != nil {
	// 	fmt.Println("Error marshalling JSON:", err)
	// 	return
	// }
	dApp := hdlDapp.DappModel{
        // Id: id,
        Name: name,
        Author: author,
        Hash: hash,
        Size: size,
        Sign: sign,
        Version: version,
        Image: image,
        Time: time,
        TotalTransaction: totalTransaction,
        TotalWallet: totalWallet,
        PathStorage: pathStorage,
        BundleId: bundleId,
        IsShowInApp: 1,
        Orientation: orientation,
        IsLocal: isLocal,
        FullScreen: fullScreen,
        StatusBar: statusBar,
        GroupId: 0,
        Page: page,
        Position: position,
        PositionObj: string(positionObject),
        IsInstalled: isInstalled,
        Type: typeT,
        UrlWeb: urlWeb,
    }
	insertDApp(&dApp, isUpdate)
    // if err != nil {
    //     return map[string]interface{}{
    //         "success": false,
    //         "message": err.Error(),
    //     }
    // }
	 result = &Result{
		Success: true,
	}
	header := models.Header{ Success:true, Data: result}
	kq := utils.NewResultTransformer(header)

	go caller.sentToClient("desktop","insert-dapp", false, kq)

	return result
}

func insertDApp(dApp *hdlDapp.DappModel, isUpdate bool) int64 {
	db, err := sqlx.Connect("sqlite3", "./database/doc_2022-12-26_09-16-03.meta_findsdk.db")
	if err != nil {
		log.Fatal(err)
	}

	dappCtrl := hdlDapp.NewDappController(db)


    if dappCtrl.IsExistDApp(dApp.BundleId) || dappCtrl.IsExistSmartContract(dApp.BundleId) {
        if isUpdate {
            return int64(dappCtrl.UpdateDapp(dApp))
        } else {
            return -1
        }
    } else {
        return int64(dappCtrl.InsertDapp1(dApp))
    }
}
