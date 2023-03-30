package whitelisthandler

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.com/meta-node/meta-node/pkg/logger"
)

// WhiteListService struct
type WhiteListService struct {
	db *sqlx.DB
}

// newWhiteListService return new WhiteListService object.
func newWhiteListService(db *sqlx.DB) *WhiteListService {
	return &WhiteListService{db}
}



func (ts *WhiteListService) getWhitelistPagination(limit int,offset int) []WhiteListModel {
	whitelist := []WhiteListModel{}
	err := ts.db.Select(&whitelist, "SELECT * FROM whiteListTB LIMIT ? OFFSET ?",limit,offset)
	if err != nil {
		panic(err)
	}
	return whitelist
}

func (ts *WhiteListService) getAllWhitelist() []WhiteListModel {
	whitelist := []WhiteListModel{}
	err := ts.db.Select(&whitelist, "SELECT * FROM whiteListTB")
	if err != nil {
		panic(err)
	}
	return whitelist
}
func (ts *WhiteListService) insertWhitelist(whitelist *WhiteListModel)int{
	_, err := ts.db.NamedExec("INSERT INTO whiteListTB( image,name,email,user_name,phoneNumber,address ) values(:image,:name,:email,:user_name,:phoneNumber,:address)",
	map[string]interface{}{
	"image":whitelist.Image,     
	"name":whitelist.Name, 
	"email":whitelist.Email, 
	"user_name":whitelist.UserName, 
	"phoneNumber":whitelist.PhoneNumber, 
	"address":whitelist.Address ,		
	})

		if err != nil {
			logger.Error(fmt.Sprintf("error when insertTransaction %", err))
			panic(fmt.Sprintf("error when insertTransaction %v", err))
			return -1
		}
	
		fmt.Println("Insert Transaction in database successed")
		return 1
}
func (ts *WhiteListService) isExistWhitelist(id int) bool {
	_,err := ts.db.Exec("SELECT EXISTS(SELECT id FROM whiteListTB WHERE id =?",id)
	if err != nil {
		logger.Error(fmt.Sprintf("error when check isExistWhitelist %", err))

		return false
	}
	return true
}
func (ts *WhiteListService) deleteAllWhitelists()error{
	_, err := ts.db.Exec("DELETE FROM whiteListTB")
		
	if err != nil {
		logger.Error(fmt.Sprintf("error when deleteAllWhitelists %", err))
		panic(fmt.Sprintf("error when deleteAllWhitelists %v", err))
		return err
	}
	fmt.Println("deleteAllWhitelists in database successed")
	return nil
}

	