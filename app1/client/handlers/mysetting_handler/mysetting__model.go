package mysettinghandler

// MysettingModel struct
type MysettingModel struct {
	ID       int64     `db:"id" json:"id"` 
	ReceiveNotificationFromNode bool    `db:"receiveNotificationFromNode" json:"receiveNotificationFromNode"` 
	ReceiveTransactionStatus bool    `db:"receiveTransactionStatus" json:"receiveTransactionStatus"` 
	NodeAddress string    `db:"nodeAddress" json:"nodeAddress"`  
	NetworkType string    `db:"networkType" json:"networkType"` 
	IsHasPincode bool    `db:"isHasPincode" json:"isHasPincode"`  
	IsEnableFaceOrTouchId bool    `db:"isEnableFaceOrTouchId" json:"isEnableFaceOrTouchId"` 
	IsWatchConfirm bool    `db:"isWatchConfirm" json:"isWatchConfirm"` 
	AccountInfo string    `db:"accountInfo" json:"AccountInfo"`        

}
