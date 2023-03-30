package main

import (
	// "database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectDb(db *sqlx.DB) {	
	
        createSCQuery := `CREATE TABLE IF NOT EXISTS smartContractTB(
            id INTEGER PRIMARY KEY AUTOINCREMENT, 
            name TEXT, 
            address TEXT, 
            abiData TEXT, 
            binData TEXT, 
            image TEXT, 
            status INTEGER)`
        createTransactionQuery := `CREATE TABLE IF NOT EXISTS transactionTB(
            id INTEGER PRIMARY KEY AUTOINCREMENT, 
            hash TEXT, 
            address TEXT, 
            toAddress TEXT, 
            pubKey TEXT, 
            amount TEXT, 
            pendingUse TEXT, 
            tip TEXT, 
            message TEXT, 
            time INTEGER, 
            status INTEGER, 
            type TEXT, 
            prevHash TEXT, 
            sign TEXT, 
            receiveInfo TEXT, 
            isDeploy INTEGER, 
            isCall INTEGER, 
            functionCall TEXT, 
            data TEXT, 
            totalBalance TEXT, 
            lastDeviceKey TEXT)`
        createDAppQuery := `CREATE TABLE IF NOT EXISTS decentralizedApplicationTB(
            id INTEGER PRIMARY KEY AUTOINCREMENT, 
            name TEXT, 
            author TEXT, 
            hash TEXT, 
            sign TEXT, 
            version TEXT, 
            image TEXT, 
            pathStorage TEXT, 
            time INTEGER, 
            totalWallet INTEGER, 
            totalTransaction INTEGER, 
            size TEXT, 
            bundleId TEXT,  
            orientation TEXT, 
            urlWeb TEXT, 
            isLocal INTEGER, 
            fullScreen INTEGER,  
            statusBar TEXT, 
            groupId INTEGER, 
            isShowInApp INTEGER, 
            page INTERGER, 
            position INTERGER, 
            positionObj TEXT,
            isInstalled INTERGER, 
            abiData TEXT, 
            binData TEXT, 
            status INTEGER, 
            type INTEGER)`
        createWalletQuery := `CREATE TABLE IF NOT EXISTS walletTB(
            id INTEGER PRIMARY KEY AUTOINCREMENT, 
            address TEXT, 
            name TEXT, 
            pendingBalance TEXT, 
            balance TEXT, 
            totalBalance TEXT, 
            bg TEXT, 
            color TEXT, 
            idSymbol TEXT, 
            pattern TEXT, 
            position INTEGER)`
        createWhiteListQuery := `CREATE TABLE IF NOT EXISTS whiteListTB(
            id INTEGER PRIMARY KEY AUTOINCREMENT, 
            image TEXT, 
            name TEXT, 
            email TEXT, 
            user_name TEXT, 
            phoneNumber TEXT, 
            address TEXT)`
        createRecentNodeQuery := `CREATE TABLE IF NOT EXISTS recentNodeTB(
            id INTEGER PRIMARY KEY AUTOINCREMENT, 
            ip TEXT, 
            port INTEGER, 
            time INTEGER)`
        createGroupDAppQuery := `CREATE TABLE IF NOT EXISTS GroupDApp(
            id INTEGER PRIMARY KEY AUTOINCREMENT, 
            name TEXT, 
            position INTEGER)`
        createSettingQuery := `CREATE TABLE IF NOT EXISTS mysetting(
            id INTEGER PRIMARY KEY AUTOINCREMENT, 
            receiveNotificationFromNode TEXT, 
            receiveTransactionStatus TEXT, 
            nodeAddress TEXT, 
            networkType TEXT, 
            isHasPincode TEXT, 
            isEnableFaceOrTouchId TEXT,
            isWatchConfirm TEXT,
            accountInfo TEXT)`       
        
        queries := []string{createSCQuery,createTransactionQuery,createWalletQuery,createDAppQuery, createWhiteListQuery, createRecentNodeQuery, createGroupDAppQuery,createSettingQuery}
        
        for _, query := range queries {
            _, err := db.Exec(query)
            if err != nil {
                fmt.Println(err)
                
            }
        }
		
    
}





           
