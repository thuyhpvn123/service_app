package router

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"

	// log "github.com/sirupsen/logrus"
	hdlDapp "gitlab.com/meta-node/client/handlers/dapp_handler"
	hdlGroupDapp "gitlab.com/meta-node/client/handlers/groupDapp_handler"

	hdlNode "gitlab.com/meta-node/client/handlers/node_handler"
	// hdlSc "gitlab.com/meta-node/client/handlers/smartcontract_handler"
	hdlMysetting "gitlab.com/meta-node/client/handlers/mysetting_handler"
	hdlTransaction "gitlab.com/meta-node/client/handlers/transaction_handler"
	hdlWallet "gitlab.com/meta-node/client/handlers/wallet_handler"
	hdlWhiteList "gitlab.com/meta-node/client/handlers/whitelist_handler"
	"gitlab.com/meta-node/meta-node/pkg/bls"
	cm "gitlab.com/meta-node/meta-node/pkg/common"

	// "gitlab.com/meta-node/client/models"

	"gitlab.com/meta-node/meta-node/pkg/logger"
)
var statusSocket ="none"
func OpenSQL()*sqlx.DB{	
	db, err := sqlx.Connect("sqlite3", "./database/doc_2022-12-26_09-16-03.meta_findsdk.db")
	if err != nil {
		logger.Error(fmt.Sprintf("error when connect sqlite %", err))
		panic(fmt.Sprintf("error when connect sqlite %v", err))
	}

	return db

}
func(caller *CallData) DeleteDApp(call map[string]interface{})map[string]interface{} {
	db:=OpenSQL()
	dappCtrl := hdlDapp.NewDappController(db)
	var err error
	result:=make(map[string]interface{})
    id, ok := call["id"].(int)
    if !ok {
        id = -1
    }
    if id == -1 {
        result=(map[string]interface{}{
            "success": false,
            "message": "Item not exists",
        })
    } else {
        success := dappCtrl.DeleteDAppAndSmartContract(id) == 1
        result=(map[string]interface{}{
            "success": success,
        })
    }
    if err != nil {
        result=(map[string]interface{}{
            "success": false,
            "message": err.Error(),
        })
    }
	return result
}
func(caller *CallData) GetAllDAppsForBrowser(call map[string]interface{}) map[string]interface{}{
	db:=OpenSQL()
	result:=make(map[string]interface{})

	dappCtrl := hdlDapp.NewDappController(db)
    data, err := dappCtrl.GetAllDAppsForBrowser()
    if err != nil {
        result=(map[string]interface{}{
            "success": false,
            "message": err.Error(),
        })
    }else{
		result=(map[string]interface{}{
			"success": true,
			"data":    data,
		})
	
	}
	return result
}
func (caller *CallData) GetAllDAppsByGroupId(call map[string]interface{})map[string]interface{} {
	db:=OpenSQL()
	result:=make(map[string]interface{})

	dappCtrl := hdlDapp.NewDappController(db)
    groupId := call["groupId"].(int64)
    data,err := dappCtrl.GetAllDAppsByGroupId(groupId)
	if err != nil {
		result=(map[string]interface{}{
			"success": true,
			"data": data,
		})
	}else{
		result=(map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})	
	}
	return result

}
func (caller *CallData)RenameGroupDApp(call map[string]interface{})map[string]interface{} {
	db:=OpenSQL()
	result:=make(map[string]interface{})
	groupDappCtrl := hdlGroupDapp.NewGroupDappController(db)
    name := call["name"].(string)
    id := call["id"].(int)
    err:=groupDappCtrl.RenameGroupDApp(name, id)
	if err != nil {
		result=(map[string]interface{}{
        "success": true,
		})
    }else{
		result=(map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})	
	}
    return result
}
func(caller *CallData) UpdateGroupDAppPosition(call map[string]interface{}) map[string]interface{}{
	db:=OpenSQL()
	result:=make(map[string]interface{})
	groupDappCtrl := hdlGroupDapp.NewGroupDappController(db)
    position := call["position"].([]int)
    for i := 0; i < len(position); i++ {
        err:=groupDappCtrl.UpdateGroupDAppPosition(i, position[i])
		if err != nil {
			result=(map[string]interface{}{
				"success": true,
			})
		}else{
			result=(map[string]interface{}{
				"success": false,
				"message": err.Error(),
			})
		}			
    }
	return result
}
func (caller *CallData) UpdateWalletPosition(call map[string]interface{}) map[string]interface{}{
	db:=OpenSQL()
	result:=make(map[string]interface{})
	walletCtrl := hdlWallet.NewWalletController(db)

    position := call["position"].([]int)
    for i := 0; i < len(position); i++ {
        err:=walletCtrl.UpdateWalletPosition(i, position[i])
		if err != nil {
			result=(map[string]interface{}{
				"success": true,
			})
		}else{
			result=(map[string]interface{}{
				"success": false,
				"message": err.Error(),
			})
		}			
    }
	return result
}
func (caller *CallData) DeleteGroupDApp(call map[string]interface{}) map[string]interface{}{
	db:=OpenSQL()
	result:=make(map[string]interface{})
	dappCtrl := hdlDapp.NewDappController(db)
	groupDappCtrl := hdlGroupDapp.NewGroupDappController(db)

    id := call["id"].(int64)
    dApps,_ := dappCtrl.GetAllDAppsByGroupId(id)
    for i := 0; i < len(dApps); i++ {
        item := dApps[i]
        dAppId := item.ID
        groupDappCtrl.UpdateDAppGroupId(0, dAppId)
    }
    err:=groupDappCtrl.DeleteGroupDApp(id)
	if err != nil {
		result=(map[string]interface{}{
			"success": true,
		})
	}else{
		result=(map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
	}	
	return result		
}
func (caller *CallData) UpdateDAppGroupId(call map[string]interface{}) map[string]interface{}{
	db:=OpenSQL()
	result:=make(map[string]interface{})
	groupDappCtrl := hdlGroupDapp.NewGroupDappController(db)
	dappCtrl := hdlDapp.NewDappController(db)

    defer func() {
        if r := recover(); r != nil {
			result=(map[string]interface{}{
                "success": false,
                "message": r.(error).Error(),
            })
        }
    }()
    groupId := int64(call["groupId"].(float64))
    data,_ := dappCtrl.GetAllDAppsByGroupId(groupId)
    maxPosition := -1
    for _, item := range data {
        if item.Position > maxPosition {
            maxPosition = item.Position
        }
    }
    dAppIds := call["dApps"].([]interface{})
    for _, id := range dAppIds {
        groupDappCtrl.UpdateDAppGroupId(groupId, int64(id.(float64)))
        dappCtrl.UpdateDAppPosition(maxPosition+1, int64(id.(float64)))
        maxPosition++
    }
	result=(map[string]interface{}{
        "success": true,
    })
	return result		

}
func (caller *CallData) UpdateDAppPosition(call map[string]interface{}) map[string]interface{}{
	db:=OpenSQL()
	result:=make(map[string]interface{})
	dappCtrl := hdlDapp.NewDappController(db)

    defer func() {
        if r := recover(); r != nil {
			result=(map[string]interface{}{
                "success": false,
                "message": r.(error).Error(),
            })
        }
    }()
    position := call["position"].([]interface{})
    for i, pos := range position {
        dappCtrl.UpdateDAppPosition(i, int64(pos.(float64)))
    }
	result=(map[string]interface{}{
        "success": true,
    })
	return result		
}
func (caller *CallData) UpdateDAppPageAndPosition(call map[string]interface{})map[string]interface{} {
	db:=OpenSQL()
	result:=make(map[string]interface{})
	dappCtrl := hdlDapp.NewDappController(db)

    id := call["id"].(int64)
    page := call["page"].(int)
    position := call["position"].(int)

    dApps := dappCtrl.GetAllDAppsByGroupIdAndPage(0, page)
    maxPosition := 0
    for i := 0; i < len(dApps); i++ {
        // item := dApps[i].(map[string]interface{})
        item := dApps[i]
        itemId := item.ID
        itemPosition := item.Position
        if itemPosition > maxPosition {
            maxPosition = itemPosition
        }
        if itemPosition >= position {
            dappCtrl.UpdateDAppPosition(itemPosition + 1, itemId)
        }
    }
    dappCtrl.UpdateDAppPage(page, id)
    dappCtrl.UpdateDAppPosition(position, id)
    if position > maxPosition {
        maxPosition = position
    }

    for i := 0; i < len(dApps); i++ {
        item := dApps[i]
        itemId := item.ID
        itemPosition := item.Position
        if itemPosition == -1 {
            maxPosition++
            dappCtrl.UpdateDAppPosition(maxPosition, itemId)
        }
    }
	result=(map[string]interface{}{
        "success": true,
    })
	return result
}

// func (caller *CallData) InsertGroupDApp(call map[string]interface{}) map[string]interface{}{
// 	db:=OpenSQL()
// 	result:=make(map[string]interface{})
// 	groupDappCtrl := hdlGroupDapp.NewGroupDappController(db)
// 	dappCtrl := hdlDapp.NewDappController(db)

//     name := call["name"].(string)
//     dAppIds := call["dApps"].([]interface{})
//     maxPosition := groupDappCtrl.GetMaxPositionGroupDApp()
//     groupDApp := &hdlGroupDapp.GroupDappModel{Name: name, Position: maxPosition + 1}
//     groupId := groupDappCtrl.InsertGroupDApp(groupDApp)
//     if groupId == -1 {
// 		result=(map[string]interface{}{
//             "success": false,
//         })
//     }
//     for i := 0; i < len(dAppIds); i++ {
//         dAppId := dAppIds[i].(int64)
//         groupDappCtrl.UpdateDAppGroupId(groupId, dAppId)
//         dappCtrl.UpdateDAppPosition(i, dAppId)
//     }
//     dApps,_ := dappCtrl.GetAllDAppsByGroupId(groupId)
// 	result=(map[string]interface{}{
//         "success": true,
//         "data": map[string]interface{}{
//             "id": groupId,
//             "name": name,
//             "dApps": dApps,
//         },
//     })
// 	return result

// }
// func insert(groupDApp GroupDApp) int64 {
//     isExist := groupDAppDao.isExist(groupDApp.id)
//     if isExist {
//         return int64(groupDAppDao.update(groupDApp))
//     }
//     return groupDAppDao.insert(groupDApp)
// }

func(caller *CallData) UpdateDAppPage(call map[string]interface{}) map[string]interface{}{
	db:=OpenSQL()
	result:=make(map[string]interface{})
	dappCtrl := hdlDapp.NewDappController(db)

    page := call["page"].(int)
    dApps := call["dApps"].([]interface{})
    for i := 0; i < len(dApps); i++ {
        id := dApps[i].(int64)
        dappCtrl.UpdateDAppPage(page, id)
    }
	result=(map[string]interface{}{
        "success": true,
    })
	return result
}
func (caller *CallData)GetLastSmartContractExcute(call map[string]interface{})map[string]interface{} {
    result:=make(map[string]interface{})

    defer func() {
        if r := recover(); r != nil {
            result=(map[string]interface{}{
                "success": false,
                "message": r.(error).Error(),
            })
        }
    }()
	db:=OpenSQL()

    scAddress, ok := call["address"].(string)
    if !ok || scAddress == "" {
        result=(map[string]interface{}{
            "success": false,
            "message": "address is required",
        })
        return result
    }
	transactionCtrl := hdlTransaction.NewTransactionController(db)

    lastScExecuteResult, err := transactionCtrl.GetLastTransactionSmartContract(scAddress)
    if err != nil {
        result=(map[string]interface{}{
            "success": true,
            "message": "Smart contract never executed",
        })
    return result
    }

    result=(map[string]interface{}{
        "success": true,
        "data": map[string]interface{}{
            "function":     lastScExecuteResult.FunctionCall,
            "amount":       lastScExecuteResult.Amount,
            "from-address": lastScExecuteResult.Address,
            "time":         lastScExecuteResult.Time,
        },
    })
    return result
}

func (caller *CallData) ExecuteSmartContract(call map[string]interface{}) map[string]interface{}{
	result:=make(map[string]interface{})

    defer func() {
        if r := recover(); r != nil {
			result=(map[string]interface{}{
                "success": false,
                "message": fmt.Sprint(r),
            })
        }
    }()

    callData := map[string]interface{}{
        "from-address":    call["fromAddress"],
        "to-address":      call["toAddress"],
        "is-call":         true,
        "amount":          call["amount"],
        "isOfflineMode":   call["isOfflineMode"],
        "feeType":         call["feeType"],
        "relatedAddresses": call["relatedAddresses"],
    }
    functionName, _ := call["functionName"].(string)
    if functionName != "" {
        callData["function-name"] = functionName
        inputArray, ok := call["inputArray"]
        if ok {
            callData["inputArray"] = inputArray
        }
    }

    input, _ := call["input"].(string)
    if input != "" {
        callData["input"] = input
    }

    kq:=caller.TryCall(callData)
	result=(map[string]interface{}{
		"success": true,
		"message": kq,
	})
	return result
}
func (caller *CallData) GetWhitelistPagination(call map[string]interface{})map[string]interface{} {
	db:=OpenSQL()
	result:=make(map[string]interface{})
	whitelistCtrl := hdlWhiteList.NewWhiteListController(db)
    page := call["page"].(int)
    limit := call["limit"].(int)
    data := whitelistCtrl.GetWhitelistPagination(limit, page)
    result["success"] = true
    result["data"] = data

    defer func() {
        if r := recover(); r != nil {
            errorMsg := fmt.Sprintf("%v", r)
            result["success"] = false
            result["message"] = errorMsg
        }
    }()
	return result
}
func (caller *CallData)GetWalletByAddress(address string) *hdlWallet.WalletModelShort {
    db:=OpenSQL()
    walletCtrl := hdlWallet.NewWalletController(db)
    myData,err := walletCtrl.GetWalletByAddress(address)
    if err != nil {
        return nil
    }
    return &myData
}

func (caller *CallData) CheckAmount(call map[string]interface{}) map[string]interface{}{
	// db:=OpenSQL()
	result:=make(map[string]interface{})

    defer func() {
        if r := recover(); r != nil {
            errorMsg := fmt.Sprintf("%v", r)
            result["success"] = false
            result["message"] = errorMsg
        }
    }()

    fromAddress := call["from-address"].(string)
    amount := call["amount"].(string)
    fee, ok := call["fee"].(string)
    if !ok {
        fee = "1"
    }
    tip, ok := call["tip"].(string)
    if !ok {
        tip = "0"
    }

    walletInDB := caller.GetWalletByAddress(fromAddress)
    // walletInDB := kq.Data.(models.Header).Data.(hdlWallet.WalletModelShort)
    if walletInDB == nil {
		result=(map[string]interface{}{
            "success": false,
            "message": "Wallet invalid",
        })
    }

    pendingWallet := walletInDB.PendingBalance
    balanceWallet := walletInDB.Balance

    balanceOfWalletBigInt := big.NewInt(0)
    if balanceWallet != "" {
        balanceOfWalletBigInt.SetString(balanceWallet, 16)
    }
    pendingOfWalletBigInt := big.NewInt(0)
    if pendingWallet != "" {
        pendingOfWalletBigInt.SetString(pendingWallet, 16)
    }

    currentBalanceOfWallet := big.NewInt(0).Add(balanceOfWalletBigInt, pendingOfWalletBigInt)
    feeBigInt := big.NewInt(0)
    feeBigInt.SetString(fee, 10)
    tipBigInt := big.NewInt(0)
    tipBigInt.SetString(tip, 10)
    amountBigInt := big.NewInt(0)
    amountBigInt.SetString(amount, 10)

    currentAmount := big.NewInt(0).Add(feeBigInt, tipBigInt)
    currentAmount = currentAmount.Add(currentAmount, amountBigInt)

    if currentBalanceOfWallet.Cmp(currentAmount) < 0 {
		result=(map[string]interface{}{
            "success": false,
            "message": "Amount is not valid",
        })
		return result
    }

	result=(map[string]interface{}{
        "success": true,
        "data": map[string]interface{}{
            "total": currentAmount.String(),
        },
    })
	return result
}

func (caller *CallData) CreateFolder(call map[string]interface{}) map[string]interface{}{
	result:=make(map[string]interface{})
    folderName, ok1 := call["name"].(string)
    path, ok2 := call["path"].(string)

    if !ok1 || !ok2 {
		result=(map[string]interface{}{
            "success": false,
            "message": "Please check path and name of folder",
        })
		return result
    }

    folder := filepath.Join(path, folderName)

    err := os.MkdirAll(folder, 0755)

    if err == nil {
        // Do something on success
		result=(map[string]interface{}{
            "success": true,
            "message": "Create folder " + folderName + " success",
            "data": map[string]interface{}{
                "path": folder,
            },
        })
		return result
    } else {
        // Do something else on failure
		result=(map[string]interface{}{
            "success": false,
            "message": "Create folder " + folderName + " fail",
        })
		return result
    }
}
func (caller *CallData) GetAllGroupDApps1() []map[string]interface{} {
	db:=OpenSQL()
    groupDappCtrl := hdlGroupDapp.NewGroupDappController(db)

    data := groupDappCtrl.GetAllGroupDApps()
    result := make([]map[string]interface{}, 0)
    if len(data) > 0 {
        for _, item := range data {
            var kq map[string]interface{}
            bitem,_:=json.Marshal(item)
            err:=json.Unmarshal(bitem,&kq)
            if err != nil{
				logger.Error(fmt.Sprintf("error when Unmarshal GetAllGroupDApps1 %", err))
				panic(fmt.Sprintf("error when Unmarshal GetAllGroupDApps1 %v", err))
		
			}

            result = append(result, kq)
        }
    }
    return result
}
func (caller *CallData) GetAllGroupDApps(call map[string]interface{}) map[string]interface{}{
	db:=OpenSQL()
	result:=make(map[string]interface{})
	dappCtrl := hdlDapp.NewDappController(db)

    defer func() {
        if r := recover(); r != nil {
			result=(map[string]interface{}{
                "success": false,
                "message": fmt.Sprint(r),
            })
        }
    }()
    
    data := caller.GetAllGroupDApps1()
    if len(data) > 0 {
        for i := range data {
            item := data[i]
            dApps,_ := dappCtrl.GetAllDAppsByGroupId(item["id"].(int64))
            item["dApps"] = dApps

            data[i] = item
        }
    }
    
	result=(map[string]interface{}{
        "success": true,
        "data":    data,
    })
    return result
}

func (caller *CallData) GetContentBackup(call map[string]interface{})map[string]interface{} {
	result:=make(map[string]interface{})
    contentBackupJSON := make(map[string]interface{})

    walletJSONArray := caller.GetTableJSONArray("walletTB")

    databaseJSON := make(map[string]interface{})
    databaseJSON["dApp"] = caller.GetTableJSONArray("decentralizedApplicationTB")
    databaseJSON["groupDApp"] = caller.GetTableJSONArray("GroupDApp")
    databaseJSON["transactions"] = caller.GetTableJSONArray("transactionTB")
    databaseJSON["wallets"] = walletJSONArray
    databaseJSON["whiteLists"] = caller.GetTableJSONArray("whiteListTB")
    databaseJSON["recentNode"] = caller.GetTableJSONArray("recentNodeTB")
    contentBackupJSON["database"] = databaseJSON

    contentBackupJSON["setting"] ,_= getMySettingJSONObject()

    contentBackupJSON["keychain"] = getKeyChainJSONObject(walletJSONArray)

    result=(map[string]interface{}{
        "success":        true,
        "content-backup": contentBackupJSON,
    })
	return result
}
func (caller *CallData) GetAllDApps() []map[string]interface{} {
    db:=OpenSQL()
	dappCtrl := hdlDapp.NewDappController(db)
    data := dappCtrl.GetAllDApps()
    dApps := make([]map[string]interface{}, 0)
    if len(data) != 0 {
        for _, item := range data {
            kq:=make(map[string]interface{})
            bitem,_ := json.Marshal(item)
            // strItem:=string(bitem)
            err:=json.Unmarshal(bitem,&kq)
            if err!=nil{
                logger.Error(fmt.Sprintf("error when Unmarshal toMap %", err))
				panic(fmt.Sprintf("error when Unmarshal toMap %v", err))	
            }
            dApps = append(dApps, kq)
        }
    }
    return dApps
}
func (caller *CallData) GetAllGroupDApps2() []map[string]interface{} {
    db:=OpenSQL()
	groupDappCtrl := hdlGroupDapp.NewGroupDappController(db)
    data := groupDappCtrl.GetAllGroupDApps()
    dApps := make([]map[string]interface{}, 0)
    if len(data) != 0 {
        for _, item := range data {
            kq:=make(map[string]interface{})
            bitem,_ := json.Marshal(item)
            // strItem:=string(bitem)
            err:=json.Unmarshal(bitem,&kq)
            if err!=nil{
                logger.Error(fmt.Sprintf("error when Unmarshal toMap %", err))
				panic(fmt.Sprintf("error when Unmarshal toMap %v", err))	
            }
            dApps = append(dApps, kq)
        }
    }
    return dApps
}
func (caller *CallData) GetAllRecentNode() []map[string]interface{} {
    db:=OpenSQL()
	nodeCtrl :=hdlNode.NewNodeController(db)
    data := nodeCtrl.GetAllRecentNode()
    dApps := make([]map[string]interface{}, 0)
    if len(data) != 0 {
        for _, item := range data {
            kq:=make(map[string]interface{})
            bitem,_ := json.Marshal(item)
            // strItem:=string(bitem)
            err:=json.Unmarshal(bitem,&kq)
            if err!=nil{
                logger.Error(fmt.Sprintf("error when Unmarshal toMap %", err))
				panic(fmt.Sprintf("error when Unmarshal toMap %v", err))	
            }
            dApps = append(dApps, kq)
        }
    }
    return dApps
}
func (caller *CallData) GetAllTransaction() []map[string]interface{} {
    db:=OpenSQL()
	transactionCtrl := hdlTransaction.NewTransactionController(db)
    data := transactionCtrl.GetAllTransaction()
    dApps := make([]map[string]interface{}, 0)
    if len(data) != 0 {
        for _, item := range data {
            kq:=make(map[string]interface{})
            bitem,_ := json.Marshal(item)
            // strItem:=string(bitem)
            err:=json.Unmarshal(bitem,&kq)
            if err!=nil{
                logger.Error(fmt.Sprintf("error when Unmarshal toMap %", err))
				panic(fmt.Sprintf("error when Unmarshal toMap %v", err))	
            }
            dApps = append(dApps, kq)
        }
    }
    return dApps
}
func (caller *CallData) GetAllWallets() []map[string]interface{} {
    db:=OpenSQL()
	walletCtrl := hdlWallet.NewWalletController(db)
    data := walletCtrl.GetAllWallets()
    dApps := make([]map[string]interface{}, 0)
    if len(data) != 0 {
        for _, item := range data {
            kq:=make(map[string]interface{})
            bitem,_ := json.Marshal(item)
            // strItem:=string(bitem)
            err:=json.Unmarshal(bitem,&kq)
            if err!=nil{
                logger.Error(fmt.Sprintf("error when Unmarshal toMap %", err))
				panic(fmt.Sprintf("error when Unmarshal toMap %v", err))	
            }
            dApps = append(dApps, kq)
        }
    }
    return dApps
}
func (caller *CallData) GetAllWhitelist() []map[string]interface{} {
    db:=OpenSQL()
	whitelistCtrl := hdlWhiteList.NewWhiteListController(db)
    data := whitelistCtrl.GetAllWhitelist()
    dApps := make([]map[string]interface{}, 0)
    if len(data) != 0 {
        for _, item := range data {
            kq:=make(map[string]interface{})
            bitem,_ := json.Marshal(item)
            // strItem:=string(bitem)
            err:=json.Unmarshal(bitem,&kq)
            if err!=nil{
                logger.Error(fmt.Sprintf("error when Unmarshal toMap %", err))
				panic(fmt.Sprintf("error when Unmarshal toMap %v", err))	
            }
            dApps = append(dApps, kq)
        }
    }
    return dApps
}
func toMap(data []interface{})[]map[string]interface{}{
    dApps := make([]map[string]interface{}, 0)
    if len(data) != 0 {
        for _, item := range data {
            kq:=make(map[string]interface{})
            bitem,_ := json.Marshal(item)
            // strItem:=string(bitem)
            err:=json.Unmarshal(bitem,&kq)
            if err!=nil{
                logger.Error(fmt.Sprintf("error when Unmarshal toMap %", err))
				panic(fmt.Sprintf("error when Unmarshal toMap %v", err))	
            }
            dApps = append(dApps, kq)
        }
    }
    return dApps
}
func (caller *CallData) GetTableJSONArray(tableName string) []interface{} {
	// db:=OpenSQL()
	// walletCtrl := hdlWallet.NewWalletController(db)
	// dappCtrl := hdlDapp.NewDappController(db)
	// groupDappCtrl := hdlGroupDapp.NewGroupDappController(db)
	// whitelistCtrl := hdlWhiteList.NewWhiteListController(db)
	// transactionCtrl := hdlTransaction.NewTransactionController(db)
	// nodeCtrl :=hdlNode.NewNodeController(db)
	// jsonArray := NewJSONArray()
	var jsonArray []interface{}
	var data []map[string]interface{}

	switch tableName {
	case "decentralizedApplicationTB":
		data = caller.GetAllDApps()
	case "GroupDApp":
		data = caller.GetAllGroupDApps2()
	case "recentNodeTB":
		data = caller.GetAllRecentNode()
	case "transactionTB":
		data = caller.GetAllTransaction()
	case "walletTB":
		data = caller.GetAllWallets()
	case "whiteListTB":
		data = caller.GetAllWhitelist()
	default:
		data = []map[string]interface{}{}
	}

	for _, item := range data {
		// json := NewJSONObject()
		bjson,_ := json.Marshal(item)
		json:=string(bjson)
		// for key, value := range item {
			// json.Put(key, value)
			jsonArray=append(jsonArray,json)
		// }
		// jsonArray.Put(json)
	}

	return jsonArray
}
func getMySettingJSONObject()( map[string]interface{},error) {  
    mySettingJSONObject := make(map[string]interface{})
    var isReceiveFromNode  bool
    var isReceiveTransactionStatus bool
    var nodeConnected string
    var networkType string
    var isHasPinCode bool
    var isEnableFaceOrTouchID bool
    var isWatchConfirm bool
    var accountInfo string
    db, err := sqlx.Connect("sqlite3","./database/my-setting.db")
    if err != nil {
        logger.Error(fmt.Sprintf("error when connect sqlite %", err))
        panic(fmt.Sprintf("error when connect sqlite %v", err))
    }

    mysettingCtrl := hdlMysetting.NewMysettingController(db)
	mysetting,err:=mysettingCtrl.GetMysetting()
    if err!=nil || mysetting==(hdlMysetting.MysettingModel{}){
        isReceiveFromNode = false
        isReceiveTransactionStatus =  false
        nodeConnected = ""
        networkType =  ""
        isHasPinCode =  false
        isEnableFaceOrTouchID =  false
        isWatchConfirm =  false
        accountInfo = ""
    }else{
        isReceiveFromNode = mysetting.ReceiveNotificationFromNode
        isReceiveTransactionStatus =  mysetting.ReceiveTransactionStatus
        nodeConnected = mysetting.NodeAddress
        networkType = mysetting.NetworkType
        isHasPinCode =  mysetting.IsHasPincode
        isEnableFaceOrTouchID =  mysetting.IsEnableFaceOrTouchId
        isWatchConfirm =  mysetting.IsWatchConfirm
        accountInfo = mysetting.AccountInfo
    }
    
    callmap1 := map[string]interface{}{
		"key": "password-backup",
	}

	result2 := ReadValueStorage(callmap1, levelDbWallet)
    passwordBackup := result2["value"]
    // passwordBackup := securyDb.Read("password-backup")
    accountInfoJSON := make(map[string]interface{})
    if accountInfo != "" {
        err := json.Unmarshal([]byte(accountInfo), &accountInfoJSON)
        if err != nil {
            logger.Error("Failed to unmarshal account info json: %v", err)
            panic(fmt.Sprintf("Failed to unmarshal account info json %v", err))
        }
    }

    mySettingJSONObject["receive-notification-from-node"] = isReceiveFromNode
    mySettingJSONObject["receive-transaction-status"] = isReceiveTransactionStatus
    mySettingJSONObject["node-address"] = nodeConnected
    mySettingJSONObject["network-type"] = networkType
    mySettingJSONObject["is-has-pincode"] = isHasPinCode
    mySettingJSONObject["is-enable-face-or-touch-id"] = isEnableFaceOrTouchID
    mySettingJSONObject["is-watch-confirm"] = isWatchConfirm
    mySettingJSONObject["password-backup"] = passwordBackup
    mySettingJSONObject["account-info"] = accountInfoJSON

    return mySettingJSONObject,err
}
func getKeyChainJSONObject(walletJSONArray []interface{}) string {
	keyChainJSONObjectmap := make(map[string]interface{})
	// securityDB := SecurityDbModules(context)
	for i := 0; i < len(walletJSONArray); i++ {
        walletMap:=make(map[string]interface{})
		wallet := walletJSONArray[i].(string)
        err:=json.Unmarshal([]byte(wallet),&walletMap)
        if err!=nil{
            logger.Error(fmt.Sprintf("error when Unmarshal getKeyChainJSONObject %", err))
            panic(fmt.Sprintf("error when Unmarshal getKeyChainJSONObject %v", err))
        }
		address := walletMap["address"].(string)
		// key := securityDB.read(address)
        keyJSON:= GetWalletKeyFromAddress(address)
		// keyJSON := make(map[string]interface{})
		// if key != nil {
		// 	_ = json.Unmarshal([]byte(key), &keyJSON)
		// }
		keyChainJSONObjectmap[address] = keyJSON
	}
    bkeyChainJSONObject,err:=json.Marshal(keyChainJSONObjectmap)
    if err!=nil{
        logger.Error(fmt.Sprintf("error when marshal getKeyChainJSONObject %", err))
        panic(fmt.Sprintf("error when marshal getKeyChainJSONObject %v", err))
    }
	return string(bkeyChainJSONObject)
}
func DAppTableFromJSON(call map[string]interface{})hdlDapp.DappModel{
    positionObj := make(map[string]interface{})
    err:=json.Unmarshal([]byte(call["positionObj"].(string)),&positionObj)
    if err!=nil{
        logger.Error(fmt.Sprintf("error when unmarshal DAppTableFromJSON %", err))
        panic(fmt.Sprintf("error when unmarshal DAppTableFromJSON %v", err))
    }
     dApp := hdlDapp.DappModel{
        // Id: id,
        Name: call["name"].(string),
        Author: call["author"].(string),
        Hash: call["hash"].(string),
        Size: call["size"].(string),
        Sign: call["sign"].(string),
        Version: call["version"].(string),
        Image: call["image"].(string),
        Time: call["time"].(int64),
        TotalTransaction: call["totalTransaction"].(int),
        TotalWallet: call["totalWallet"].(int),
        PathStorage: call["pathStorage"].(string),
        BundleId: call["bundleId"].(string),
        IsShowInApp: call["isShowInApp"].(int),
        Orientation: call["orientation"].(string),
        IsLocal: call["isLocal"].(int),
        FullScreen: call["fullScreen"].(int),
        StatusBar: call["statusBar"].(string),
        GroupId: call["groupId"].(int),
        Page: call["page"].(int),
        Position: call["position"].(int),
        PositionObj: string(call["positionObj"].(int)),
        IsInstalled: 1,
        Type: 1,
        UrlWeb: call["urlWeb"].(string),
    }
    return dApp
}
func GroupDAppFromJSON(call map[string]interface{})hdlGroupDapp.GroupDappModel{
    groupDApp := hdlGroupDapp.GroupDappModel{
        Name: call["name"].(string),
        Position: call["position"].(int),
    }
    return groupDApp

}
func TransactionsFromJSON(call map[string]interface{})hdlTransaction.TransactionModel{
    transaction := hdlTransaction.TransactionModel{
        Hash: call["hash"].(string),
        Address: call["address"].(string),
        ToAddress: call["toAddress"].(string),
        PubKey: call["pubKey"].(string),
        Amount: call["amount"].(string),
        PendingUse: call["pendingUse"].(string),
        Time: call["time"].(int64),
        Status: call["status"].(int),
        Type: call["type"].(string),
        PrevHash: call["prevHash"].(string),
        Sign: call["sign"].(string),
        Tip: call["tip"].(string),
        Message: call["message"].(string),
        ReceiveInfo: call["receiveInfo"].(string),
        IsCall: call["isCall"].(bool),
        IsDeploy: call["isDeploy"].(bool),
        Data: call["data"].(string),
        FunctionCall: call["functionCall"].(string),
        TotalBalance: call["totalBalance"].(string),
        LastDeviceKey: call["lastDeviceKey"].(string),
	}
    return transaction
}
func WalletsFromJSON(call map[string]interface{})hdlWallet.WalletModel{
    wallet := hdlWallet.WalletModel{
        Address :call["address"].(string),
        Name : call["name"].(string),
        PendingBalance: call["pendingBalance"].(string),
        Balance: call["balance"].(string),
        TotalBalance :call["totalBalance"].(string),
        Bg :call["bg"].(string),
        Color :call["color"].(string),
        IdSymbol :call["idSymbol"].(string),
        Pattern: call["pattern"].(string),
        Position:call["position"].(int),
    
	}
    return wallet
}
func WhitelistFromJSON(call map[string]interface{})hdlWhiteList.WhiteListModel{
    whitelist := hdlWhiteList.WhiteListModel{
        Image: call["image"].(string),
        Name:call["name"].(string),
        Email :call["email"].(string),
        UserName:call["user_name"].(string),
        PhoneNumber: call["phoneNumber"].(string),
        Address: call["address"].(string),    
	}
    return whitelist
}
func RecentNodeFromJSON(call map[string]interface{})hdlNode.NodeModel{
    node := hdlNode.NodeModel{
        IP : call["ip"].(string),
        Port: call["port"].(int),
        Time :call["time"].(int64),
	}
    return node
}
func (caller *CallData)RestoreByFile(call map[string]interface{})map[string]interface{} {
    result:=make(map[string]interface{})
    contentBackupJSON:=make(map[string]interface{})
    go func() {
        defer func() {
            if r := recover(); r != nil {
                result=(map[string]interface{}{
                    "success": false,
                    "message": r.(error).Error(),
                })
            }
        }()
        
        contentBackup := call["content-backup"].(string)
        // contentBackupJSON, err := json.Parse(contentBackup)
        // if err != nil {
        //     panic(err)
        // }
        // contentBackupJSON := json.NewDecoder(strings.NewReader(contentBackup))
        db:=OpenSQL()
        walletCtrl := hdlWallet.NewWalletController(db)
        dappCtrl := hdlDapp.NewDappController(db)
        groupDappCtrl := hdlGroupDapp.NewGroupDappController(db)
        whitelistCtrl := hdlWhiteList.NewWhiteListController(db)
        transactionCtrl := hdlTransaction.NewTransactionController(db)
        nodeCtrl :=hdlNode.NewNodeController(db)

        dappCtrl.DeleteDAppTable()
        groupDappCtrl.DeleteAllGroupDApps()
        transactionCtrl.DeleteAllTrans()
        walletCtrl.DeleteAllWallets()
        whitelistCtrl.DeleteAllWhitelists()
        nodeCtrl.DeleteAllRecentNodes()
        err:=json.NewDecoder(strings.NewReader(contentBackup)).Decode(&contentBackupJSON)
        databaseJSON := contentBackupJSON["database"].(map[string]interface{})
        dAppJSONArray := databaseJSON["dApp"].([]interface{})
        
        for _, dAppJSON := range dAppJSONArray {
            dappVar:=DAppTableFromJSON(dAppJSON.(map[string]interface{}))
            dappCtrl.InsertDapp1(&dappVar)
        }
        
        groupDAppJSONArray := databaseJSON["groupDApp"].([]interface{})
        for _, groupDAppJSON := range groupDAppJSONArray {
            groupDappVar:=GroupDAppFromJSON(groupDAppJSON.(map[string]interface{}))
            groupDappCtrl.InsertGroupDApp(&groupDappVar)
        }
        
        transactionJSONArray := databaseJSON["transactions"].([]interface{})
        for _, transactionJSON := range transactionJSONArray {
            if !transactionCtrl.IsExistTransaction(transactionJSON.(map[string]interface{})["hash"].(string)) {
                transactionVar:=TransactionsFromJSON(transactionJSON.(map[string]interface{}))
                transactionCtrl.InsertTransaction(&transactionVar)
            }
        }
        
        walletJSONArray := databaseJSON["wallets"].([]interface{})
        for _, walletJSON := range walletJSONArray {
            if !walletCtrl.IsExistWallet(walletJSON.(map[string]interface{})["address"].(string)) {
                walletVar:=WalletsFromJSON(walletJSON.(map[string]interface{}))
                walletCtrl.InsertWallet1(&walletVar)
            }
        }
        
        whitelistJSONArray := databaseJSON["whiteLists"].([]interface{})
        for _, whitelistJSON := range whitelistJSONArray {
            if !whitelistCtrl.IsExistWhitelist(int(whitelistJSON.(map[string]interface{})["id"].(float64))) {
                whitelistVar:=WhitelistFromJSON(whitelistJSON.(map[string]interface{}))
                whitelistCtrl.InsertWhitelist(&whitelistVar)
            }
        }
        
        recentNodeJSONArray := databaseJSON["recentNode"].([]interface{})
        for _, recentNodeJSON := range recentNodeJSONArray {
            ip := recentNodeJSON.(map[string]interface{})["ip"].(string)
            port := int(recentNodeJSON.(map[string]interface{})["port"].(float64))
            if !nodeCtrl.IsExistRecentNode(ip, port) {
                recentNodeVar:=RecentNodeFromJSON(recentNodeJSON.(map[string]interface{}))
                nodeCtrl.InsertRecentNode(&recentNodeVar)
            }
        }
        
        mySettingJSON := contentBackupJSON["setting"].(map[string]interface{})
        // mySettingJSON :=databaseJSON["setting"].(map[string]interface{})
        isReceiveFromNode := mySettingJSON["receive-notification-from-node"].(bool)
        isReceiveTransactionStatus := mySettingJSON["receive-transaction-status"].(bool)
        nodeConnected := mySettingJSON["node-address"].(string)
        networkType := mySettingJSON["network-type"].(string)
        isHasPinCode := mySettingJSON["is-has-pincode"].(bool)
        isEnableFaceOrTouchID := mySettingJSON["is-enable-face-or-touch-id"].(bool)
        isWatchConfirm := mySettingJSON["is-watch-confirm"].(bool)
        passwordBackup := mySettingJSON["password-backup"].(bool)
        // accountInfo := make(map[string]interface{})

        accountInfo := mySettingJSON["account-info"].(string)
        db, err = sqlx.Connect("sqlite3","./database/my-setting.db")
        if err != nil {
            logger.Error(fmt.Sprintf("error when connect sqlite %", err))
            panic(fmt.Sprintf("error when connect sqlite %v", err))
        }
    
        mysettingCtrl := hdlMysetting.NewMysettingController(db)
        mysetting := &hdlMysetting.MysettingModel{
            ReceiveNotificationFromNode:isReceiveFromNode ,
            ReceiveTransactionStatus: isReceiveTransactionStatus ,
            NodeAddress:"",
            NetworkType : "",
            IsHasPincode: isHasPinCode ,
            IsEnableFaceOrTouchId: isEnableFaceOrTouchID ,
            IsWatchConfirm : isWatchConfirm,
            AccountInfo: "" ,       
        }
        mysettingkq,err:=mysettingCtrl.GetMysetting()
        if err!=nil || mysettingkq.NodeAddress ==""{
            mysetting.NodeAddress=nodeConnected
        }
        if err!=nil || mysettingkq.NetworkType ==""{
            mysetting.NetworkType=networkType
        }
        if err!=nil || mysettingkq.AccountInfo ==""{
            mysetting.AccountInfo=accountInfo
        }
        mysettingCtrl.InsertMysetting(mysetting)

        callmap := map[string]interface{}{
			"key":  "password-backup",
			"data": passwordBackup,
		}
		result3 := WriteValueStorage(callmap, levelDbWallet)
		fmt.Println("write password-backup to storage restoreByFile success:", result3["sucess"])

        keyChainJSONObject := contentBackupJSON["keychain"].(string)
        item:=make(map[string]interface{})
        err=json.Unmarshal([]byte(keyChainJSONObject),&item)
        if err!=nil{
            logger.Error("Failed to unmarshal keyChainJSONObject: %v", err)
            panic(fmt.Sprintf("Failed to unmarshal keyChainJSONObject %v", err))
        }
        for i,v:=range item{
            fmt.Println("address:",i)
            strV, _ := json.Marshal(v)
            // keyar:=make(map[string]interface{})
            // err:=json.Unmarshal([]byte(strV),&keyar)
            // if err!=nil{
            //     logger.Error("Failed to unmarshal keyChainJSONObject json: %v", err)
            //     panic(fmt.Sprintf("Failed to unmarshal keyChainJSONObject json %v", err))
            //     }
            // for in,va:=range keyar{
            //     fmt.Println("Key:",in)
            //     strVa, _ := json.Marshal(va)
            //     fmt.Println(string(strVa))
                //store keychain in l
                callmap := map[string]interface{}{
                    "key":  i,
                    "data": string(strV),
                }
                result3 := WriteValueStorage(callmap, levelDbWallet)
                fmt.Println("write keychain to storage restoreByFile success:", result3["sucess"])
        
            // }           
        }
		// keys := keyChainJSONObject.keys()
		// for keys.HasNext() {
		// 	address := keys.Next()
		// 	if ey.read(address) == nil {
		// 		keyJSON := keyChainJSONObject[address]
		// 		securyDb.write(address, keyJSON.String())
		// 	}
		// }
		
        result=(map[string]interface{}{
			"success": true,
		})
		
	}()
    return result
}        
func(caller *CallData) SetPinCode(call map[string]interface{})map[string]interface{} {
    result:=make(map[string]interface{})
    pinCode := call["pin-code"].(string)
    pinCodeConfirm := call["pin-code-confirm"].(string)
    db, err := sqlx.Connect("sqlite3","./database/my-setting.db")
	if err != nil {
		logger.Error(fmt.Sprintf("error when connect sqlite %", err))
		panic(fmt.Sprintf("error when connect sqlite %v", err))
	}

	mysettingCtrl := hdlMysetting.NewMysettingController(db)
	// mysetting := &hdlMysetting.MysettingModel{
	// 	IsHasPincode :true,  
	// }

	mysettingCtrl.UpdatePinCode(true)

    // sharedPreferences := context.GetSharedPreferences("my-setting", Context.MODE_PRIVATE)
    // editor := sharedPreferences.Edit()
    // editor.PutBoolean("is-has-pincode", true)
    // editor.Apply()
    callmap1 := map[string]interface{}{
        "key": "pin-code",
	}

	result1 := ReadValueStorage(callmap1, levelDbMysetting)
	if result1["value"] == nil {
		callmap2 := map[string]interface{}{
			"key":  "pin-code",
			"data": pinCode,
		}
        result2 := WriteValueStorage(callmap2, levelDbMysetting)
        fmt.Println("write pin-code to storage success:", result2["sucess"])   
    }

    callmap3 := map[string]interface{}{
        "key": "pin-code-confirm",
	}
	result3 := ReadValueStorage(callmap3, levelDbMysetting)
	if result3["value"] == nil {
		callmap4 := map[string]interface{}{
			"key":  "pin-code-confirm",
			"data": pinCodeConfirm,
		}
        result4 := WriteValueStorage(callmap4, levelDbMysetting)
        fmt.Println("write pin-code-confirm to storage success:", result4["sucess"])   
    }

    // securyDb.Write("pin-code", pinCode)
    // securyDb.Write("pin-code-confirm", pinCodeConfirm)
	result=(map[string]interface{}{
        "success": true,
    })
    // } catch e := Exception {
	// result=(map[string]interface{}{
    //     "success": false,
    //     "message": e.Message,
    // })
    return result

}
func (caller *CallData)GetMySetting(call map[string]interface{})map[string]interface{}  {
    result:=make(map[string]interface{})

    mySetting, err := getMySettingJSONObject()
    if err != nil {
        result=(map[string]interface{}{
            "success": false,
            "message": err.Error(),
        })
        return result
    }

	result=(map[string]interface{}{
        "success": true,
        "data":    mySetting,
    })
    return result
}
		
func (caller *CallData)GetBalanceWallet(call map[string]interface{})map[string]interface{} {
    result:=make(map[string]interface{})
    db, err := sqlx.Connect("sqlite3", "./database/doc_2022-12-26_09-16-03.meta_findsdk.db")
	if err != nil {
		logger.Error(fmt.Sprintf("error when connect sqlite %", err))
		panic(fmt.Sprintf("error when connect sqlite %v", err))
	}
	walletCtrl := hdlWallet.NewWalletController(db)
   
    wallets := call["wallets"].([]interface{})
    data := []map[string]interface{}{}
    for i := 0; i < len(wallets); i++ {
        item := wallets[i].(map[string]interface{})
        walletInfo,err :=  walletCtrl.GetWalletByAddress1(item["address"].(string))
        if err!=nil{
            result=(map[string]interface{}{
                "success": false,
                "message": err.Error(),
            })   
            return result     
        }
        if walletInfo != (hdlWallet.WalletModel{}) {
            data = append(data, map[string]interface{}{
                // "id": walletInfo.I,
                "name": walletInfo.Name,
                "address": walletInfo.Address,
                "pendingBalance": walletInfo.PendingBalance,
                "balance": walletInfo.Balance,
                "totalBalanceString": walletInfo.TotalBalance,
            })
        }
    }
    result=(map[string]interface{}{
        "success": true,
        "data": data,
    })
    
    return result
}
func (caller *CallData)GetWalletAtAddress(call map[string]interface{})map[string]interface{} {

    result:=make(map[string]interface{})
    db, err := sqlx.Connect("sqlite3", "./database/doc_2022-12-26_09-16-03.meta_findsdk.db")
	if err != nil {
		logger.Error(fmt.Sprintf("error when connect sqlite %v", err))
		panic(fmt.Sprintf("error when connect sqlite %v", err))
	}
	walletCtrl := hdlWallet.NewWalletController(db)

    addressWallet, ok := call["data"].(string)
    if !ok {
        result=(map[string]interface{}{
            "success": false,
            "message":"addressWallet is required",
        })
        return result
    }
    data ,_:= walletCtrl.GetWalletByAddress1(addressWallet)
    if data == (hdlWallet.WalletModel{}) {
        result=(map[string]interface{}{
            "success": false,
            "message":"Not found wallets",
        })
        return result
    }

    result=(map[string]interface{}{
        "data": map[string]interface{}{
            "address":        data.Address,
            "pendingBalance": data.PendingBalance,
            "balance":        data.Balance,
        },
    })
    return result
}
func(caller *CallData) GetStatusConnected(call map[string]interface{})map[string]interface{} {
    result:=make(map[string]interface{})

    defer func() {
        if r := recover(); r != nil {
            result=(map[string]interface{}{
                "success": false,
                "message": fmt.Sprintf("%v", r),
            })
        }
    }()
    status := 0
    if statusSocket == "open" {
        status = 1
    }
    result=(map[string]interface{}{
        "success": true,
        "data": map[string]interface{}{
            "status": status,
        },
    })
    return result
}
func (caller *CallData)WriteToStorage(call map[string]interface{}) map[string]interface{}{
    result:=make(map[string]interface{})
    rawText, ok := call["text"].(string)
    if !ok {
        result=(map[string]interface{}{
            "success": false,
            "message":"error when parse string WriteToStorage",
        })
            return result
    }

    data, err := base64.URLEncoding.DecodeString(rawText)
    if err != nil {
        result=(map[string]interface{}{
            "success": false,
            "message":err.Error(),
        })
            return result
    }

    downloadFolderPath := filepath.Join(
        filepath.Dir("/Downloads"), "yyy_gz.gz")
    newImage, err := os.Create(downloadFolderPath)
    if err != nil {
        result=(map[string]interface{}{
            "success": false,
            "message":err.Error(),
        })
            return result
    }

    _, err = newImage.Write(data)
    if err != nil {
        newImage.Close()
        result=(map[string]interface{}{
        "success": false,
        "message":err.Error(),
        })
    return result
    }
    result=(map[string]interface{}{
        "success": true,
    })

    newImage.Close()
    // downloadFolderPath := "/Downloads"
    // err = ioutil.WriteFile(downloadFolderPath+"/yyy_gz.gz", data, 0644)
    // if err != nil {
    //     return err
    // }
    // return nil

    // result=(map[string]interface{}{
    //     "success": true,
    // })
    return result

}

func (caller *CallData)CreateHashFromFile(call map[string]interface{}) map[string]interface{} {
    data, _ := call["data"].([]byte)
    hash := crypto.Keccak256(data)
    return map[string]interface{}{
        "success": true,
        "data": hex.EncodeToString(hash[:]),
    }
}
func (caller *CallData) VerifyJSONApp(call map[string]interface{})  map[string]interface{}{
    result:=make(map[string]interface{})
    hash, ok := call["hash"].([]byte)
    if !ok {
        result=(map[string]interface{}{
            "success": true,
            "data":    false,
        })
        return result
    }

    sign, ok := call["sign"].([]byte)
    if !ok {
        result=(map[string]interface{}{
            "success": true,
            "data":    false,
        })
        return result
    }

    pubKey, ok := call["pubKey"].([]byte)
    if !ok {
        result=(map[string]interface{}{
            "success": true,
            "data":    false,
        })
        return result
    }
    pubKeyCm:=cm.PubkeyFromBytes(pubKey)
    signCm:=cm.SignFromBytes(sign)    
    // success := cryptoHelper.GetInstance().BlsVerify(pubKey, sign, hash)
    success := bls.VerifySign(pubKeyCm, signCm, hash)

    result=(map[string]interface{}{
        "success": true,
        "data":    success,
    })
    return result
}

func (caller *CallData) DeleteDAppTable() map[string]interface{} {
    db:=OpenSQL()
    dappCtrl := hdlDapp.NewDappController(db)
    err:=dappCtrl.DeleteDAppTable()
    if err!=nil{
        logger.Error(fmt.Sprintf("error when connect sqlite %", err))
		panic(fmt.Sprintf("error when connect sqlite %v", err))

    }
    return map[string]interface{}{
        "success": true,
    }
}

func (caller *CallData) GetPasswordFromSeedPhrase(call map[string]interface{}) map[string]interface{} {
    seedPhrase, ok := call["seedPhrase"].([]interface{})
    if !ok {
        return map[string]interface{}{
            "success": false,
            "message": "seedPhrase is required",
        }
    }
    var seedPhraseStringSlice []string
    for _, phrase := range seedPhrase {
        seedPhraseStringSlice = append(seedPhraseStringSlice, phrase.(string))
    }
    dataHash := []byte(strings.Join(seedPhraseStringSlice, "-"))
    // hash := sha3.Sum256(dataHash)
    hash := crypto.Keccak256(dataHash)

    password := hex.EncodeToString(hash[:32])
    return map[string]interface{}{
        "success": true,
        "data":    password,
    }
}

func (caller *CallData)DeleteListDAppAndSmartContract(call map[string]interface{}) map[string]interface{}{
    db:=OpenSQL()
    result:=make(map[string]interface{})
    dApps := call["dApps"].([]interface{})
    dappCtrl := hdlDapp.NewDappController(db)

    for i := 0; i < len(dApps); i++ {
        id := dApps[i].(int)
        kq:=dappCtrl.DeleteDAppAndSmartContract(id)
    
        if kq == 1 {
            result=(map[string]interface{}{
                "success": true,
            })
        } else {
            result=(map[string]interface{}{
                "success": false,
                "message": "fail on DeleteListDAppAndSmartContract",
            })
        }
    }
    return result

}

func (caller *CallData) GetSign(call map[string]interface{}) map[string]interface{}{
    result:=make(map[string]interface{})
    address := call["address"].(string)
    hash := call["hash"].(string)

    walletKey := GetWalletKeyFromAddress(address)
    if walletKey == nil {
    result=(map[string]interface{}{
            "success": false,
            "message": "Wallet invalid",
        })
    }
    keyPair := bls.NewKeyPair(common.FromHex(walletKey["priKey"].(string)))
	prikey:= keyPair.GetPrivateKey()
	sign := bls.Sign(prikey, common.FromHex(hash))
    // sign := CryptoHelper.getInstance().blsSign(
    //     hex.DecodeString(hash),
    //     walletKey["priKey"].([]byte),
    // )
    result=(map[string]interface{}{
        "success": true,
        "data": map[string]interface{}{
            "sign": sign,
        },
    })
    return result
}
func (caller *CallData) GetPublicKey(call map[string]interface{}) map[string]interface{}{
    result:=make(map[string]interface{})
    address := call["address"].(string)
    walletKey := GetWalletKeyFromAddress(address)
    if walletKey == nil {
        result=(map[string]interface{}{
            "success": false,
            "message": "Wallet invalid",
        })
    }

    pubKey := walletKey["pubKey"].([]byte)
    result=(map[string]interface{}{
        "success": true,
        "data": map[string]interface{}{
            "pubKey": hex.EncodeToString(pubKey),
        },
    })
    return result
}

		
// func (caller *CallData) UpdateWalletByteAddress(call map[string]interface{}) map[string]interface{}{
// 	db:=OpenSQL()
// 	result:=make(map[string]interface{})
// 	walletCtrl := hdlWallet.NewWalletController(db)


//     walletAddress, ok := call["address"].(string)
//     if !ok || walletAddress == "" {
// 		result=(map[string]interface{}{
//             "success": false,
//             "message": "address is required",
//         })
//     }

//     balance := int64(0)
//     if call["balance"] != nil {
//         balance = call["balance"].(int64)
//     }

//     pendingBalance := int64(0)
//     if call["pendingBalance"] != nil {
//         pendingBalance = call["pendingBalance"].(int64)
//     }

//     // walletCtrl.updateBalance(balance, pendingBalance, walletAddress)

// 	result=(map[string]interface{}{
//         "success": true,
//         "data": map[string]interface{}{
//             "address":        walletAddress,
//             "balance":        balance,
//             "pendingBalance": pendingBalance,
//         },
//     })
//     return result
// }
func (caller *CallData) EditWalletUI(call map[string]interface{}) map[string]interface{}{
    db:=OpenSQL()
	result:=make(map[string]interface{})
	walletCtrl := hdlWallet.NewWalletController(db)
    bg := ""
    if bgObj, ok := call["bg"].(map[string]interface{}); ok {
        bgBytes, err := json.Marshal(bgObj)
        if err != nil {
            panic(err)
        }
        bg = string(bgBytes)
    } else {
        bg = call["bg"].(string)
    }

    color := ""
    if colorObj, ok := call["color"].(map[string]interface{}); ok {
        colorBytes, err := json.Marshal(colorObj)
        if err != nil {
            panic(err)
        }
        color = string(colorBytes)
    } else {
        color = call["color"].(string)
    }

    idSymbol := call["idSymbol"].(int)
    pattern := call["pattern"].(string)
    address := call["address"].(string)

    walletCtrl.EditWalletUI(bg, color, idSymbol, pattern, address)
    result=(map[string]interface{}{
        "success": true,
    })
    defer func() {
        if r := recover(); r != nil {
            result=(map[string]interface{}{
                "success": false,
                "message": fmt.Sprint(r),
            })
        }
    }()
    return result
}

func (caller *CallData)DisconnectWallet(call map[string]interface{})map[string]interface{} {
    result:=make(map[string]interface{})
    bundleId := call["bundle-id"].(string)
    // sharedPreferences := context.GetSharedPreferences("setting_d_app", Context.MODE_PRIVATE)
    // editor := sharedPreferences.Edit()
    // editor.Remove("wallet-active-" + bundleId)
    // editor.Commit()
    callmap3 := map[string]interface{}{
        "key": "wallet-active-" + bundleId,
	}
	result3 := ReadValueStorage(callmap3, levelDbSettingdapp)
    walletActiveStr:=result3["value"] 
	if walletActiveStr == nil {
        result=(map[string]interface{}{
            "success": true,
        })
        return result
    }else{
		// json.Unmarshal([]byte(walletActiveStr.(string)), &walletActive) 
        DeleteKeyStorage(callmap3, levelDbSettingdapp)
        result=(map[string]interface{}{
            "success": true,
        })
    
    }

    return result
}

func (caller *CallData) SetWalletActiveDApp(call map[string]interface{})map[string]interface{} {
    result:=make(map[string]interface{})
    id, ok := call["id"].(string)
    if !ok {
        result=(map[string]interface{}{
            "success": true,
            "message": "id not found",
        })
    
        return result
    }

    wallet, ok := call["wallet"].(map[string]interface{})
    if !ok {
        result=(map[string]interface{}{
            "success": true,
            "message": "wallet not found",
        })
    
        return result
    }

    // sharedPreferences := context.GetSharedPreferences("setting_d_app", context.MODE_PRIVATE)
    // editor := sharedPrefeettingdapprences.Edit()
    walletJson, err := json.Marshal(wallet)
    if err != nil {
        result=(map[string]interface{}{
            "success": true,
            "data": err.Error(),
        })
    
        return result

    }
    callmap := map[string]interface{}{
        "key":  "wallet-active-" + id,
        "data": string(walletJson),
    }
    result3 := WriteValueStorage(callmap, levelDbSettingdapp)
    fmt.Println("write wallet-active to storage SetWalletActiveDApp success:", result3["sucess"])


    // editor.PutString("wallet-active-" + id, string(walletJson))
    // editor.Apply()

    result=(map[string]interface{}{
		"success": true,
	})
    return result
}
func (caller *CallData) GetSettingDApp(call map[string]interface{}) map[string]interface{}{
    result:=make(map[string]interface{})
    walletActive := make(map[string]interface{})

    id := call["id"].(string)
    // sharedPreferences := context.GetSharedPreferences("setting_d_app", Context.MODE_PRIVATE)
    // walletActiveStr := sharedPreferences.GetString("wallet-active-" + id, "")
    callmap3 := map[string]interface{}{
        "key": "wallet-active-" + id,
	}
	result3 := ReadValueStorage(callmap3, levelDbSettingdapp)
    walletActiveStr:=result3["value"] 
	if walletActiveStr == nil {
        result=(map[string]interface{}{
            "success": false,
            "message": "can not get wallet-active",
        })
        return result
    }else{
		json.Unmarshal([]byte(walletActiveStr.(string)), &walletActive) 
    }
    

    // walletActive := make(map[string]interface{})
    // if !strings.isNullOrEmpty(walletActiveStr) {
    //     json.Unmarshal([]byte(walletActiveStr), &walletActive)
    // }
    result=(map[string]interface{}{
        "success": true,
        "data": map[string]interface{}{
            "wallet-active": walletActive,
        },
    })
    return result
}

func(caller *CallData) GetNode(call map[string]interface{}) map[string]interface{}{
    result:=make(map[string]interface{})
    db, err := sqlx.Connect("sqlite3", "./database/doc_2022-12-26_09-16-03.meta_findsdk.db")
	if err != nil {
		logger.Error(fmt.Sprintf("error when connect sqlite %", err))
		panic(fmt.Sprintf("error when connect sqlite %v", err))
	}
	nodeCtrl :=hdlNode.NewNodeController(db)

    defer func() {
        if r := recover(); r != nil {
            result=(map[string]interface{}{
                "success": false,
                "message": fmt.Sprint(r),
            })
        }
    }()

    listNode := []map[string]interface{}{
        {
            "ip":   "34.138.137.194",
            "port": 3011,
        },
        {
            "ip":   "34.75.70.255",
            "port": 3011,
        },
        {
            "ip":   "35.185.63.163",
            "port": 3011,
        },
        {
            "ip":   "34.73.73.35",
            "port": 3011,
        },
        {
            "ip":   "104.196.68.241",
            "port": 3011,
        },
        {
            "ip":   "35.229.32.194",
            "port": 3011,
        },
        {
            "ip":   "34.148.29.59",
            "port": 3011,
        },
        {
            "ip":   "34.139.50.177",
            "port": 3011,
        },
        {
            "ip":   "34.73.211.146",
            "port": 3011,
        },
        {
            "ip":   "34.73.156.60",
            "port": 3011,
        },
        {
            "ip":   "35.196.152.49",
            "port": 3011,
        },
        {
            "ip":   "35.243.233.132",
            "port": 3011,
        },
    }

    recentNode := nodeCtrl.GetAllRecentNode()

    result=(map[string]interface{}{
        "success": true,
        "data": map[string]interface{}{
            "listNode":   listNode,
            "recentNode":recentNode,
        },
    })
    return result
}
func (caller *CallData)DeleteAllStorage() map[string]interface{} {
    success := false
    defer func() {
        if r := recover(); r != nil {
            log.Println("Error:", r)
        }
    }()

    // Delete all storage from securyDb
    err := os.RemoveAll("/db")
    if err != nil {
        log.Fatal(err)
    }

    success = true
    return map[string]interface{}{
        "success": success,
    }
}


