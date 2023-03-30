package nodehandler

// import "time"

// NodeModel struct
type NodeModel struct {
	ID       int64     `db:"id" json:"id"`
	IP     string    `db:"ip" json:"ip"`
	Port int `db:"port" json:"port"`
	Time int64 `db:"time" json:"time"`

}
