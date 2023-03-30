package groupDapphandler

// import "time"

// GroupDappModel struct
type GroupDappModel struct {
	ID       int64    `db:"id" json:"id"`
	Name     string    `db:"name" json:"name"`
	Position int `db:"position" json:"position"`
}
