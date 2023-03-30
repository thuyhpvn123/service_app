package wallethandler

// import "time"

// TopicModel struct
type WalletModel struct {
	ID       int64     `db:"id" json:"id"`
	Address string `db:"address" json:"address"`
	Name     string    `db:"name" json:"name"`
	PendingBalance string `db:"pendingBalance" json:"pendingBalance"`
	Balance string `db:"balance" json:"balance"`
	TotalBalance string `db:"totalBalance" json:"totalBalance"`
	Bg string `db:"bg" json:"bg"`
	Color string `db:"color" json:"color"`
	IdSymbol string `db:"idSymbol" json:"idSymbol"`
	Pattern string `db:"pattern" json:"pattern"`
	Position int`db:"position" json:"position"`
}
type WalletModelShort struct {
	Address string `db:"address" json:"address"`
	PendingBalance string `db:"pendingBalance" json:"pendingBalance"`
	Balance string `db:"balance" json:"balance"`

}
