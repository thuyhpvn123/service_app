package transactionhandler

// import "time"

// TransactionModel struct
type TransactionModel struct {
	ID       int64     `db:"id" json:"id"`
	Hash     string    `db:"hash" json:"hash"`
	Address string `db:"address" json:"address"`
	ToAddress string `db:"toAddress" json:"toAddress"`
	PubKey string `db:"pubKey" json:"pubKey"`
	Amount string `db:"amount" json:"amount"`
	PendingUse string `db:"pendingUse" json:"pendingUse"`
	Balance string `db:"balance" json:"balance"`
	Fee string `db:"fee" json:"fee"`
	Tip string `db:"tip" json:"tip"`
	Message string `db:"message" json:"message"`
	Time int64 `db:"time" json:"time"`
	Status int `db:"status" json:"status"` // 0: pending, 1 = sent, 2 = success, 3: fail, 4: cancel
	Type string `db:"type" json:"type"`
	PrevHash string `db:"prevHash" json:"prevHash"`
	Sign string `db:"sign" json:"sign"`
	ReceiveInfo string `db:"receiveInfo" json:"receiveInfo"`
	IsDeploy bool `db:"isDeploy" json:"isDeploy"`
	IsCall bool `db:"isCall" json:"isCall"`
	FunctionCall string `db:"functionCall" json:"functionCall"`
	Data string `db:"data" json:"data"`
	TotalBalance string `db:"totalBalance" json:"totalBalance"`
	LastDeviceKey string `db:"lastDeviceKey" json:"lastDeviceKey"`

}
