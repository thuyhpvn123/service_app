package whitelisthandler

// import "time"

// WhiteListModel struct
type WhiteListModel struct {
	ID       int64     `db:"id" json:"id"` 
	Image string `db:"image" json:"image"`
	Name     string    `db:"name" json:"name"` 
	Email string `db:"email" json:"email"`
	UserName string `db:"user_name" json:"user_name"`
	PhoneNumber string `db:"phoneNumber" json:"phoneNumber"`
	Address string `db:"address" json:"address"`
}
