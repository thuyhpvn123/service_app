package smartContracthandler

// import "time"

// SmartContractModel struct
type SmartContractModel struct {
	ID       int64     `db:"id" json:"id"`
	Name     string    `db:"name" json:"name"`
	Address string `db:"address" json:"address"`
	AbiData string `db:"abiData" json:"abiData"`
	BinData string `db:"binData" json:"binData"`
	Image string `db:"image" json:"image"`
	Status int `db:"status" json:"status"` // 0:pending; 1 = success; 2 = fail
}
