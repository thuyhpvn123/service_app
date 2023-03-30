package groupDapphandler

import (
	// "fmt"

	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.com/meta-node/meta-node/pkg/logger"

	// "gitlab.com/meta-node/client/models"
	// "gitlab.com/meta-node/client/utils"
)

// DappService struct
type GroupDappService struct {
	db *sqlx.DB
}

// newDappService return new DappService object.
func newGroupDappService(db *sqlx.DB) *GroupDappService {
	return &GroupDappService{db}
}

func (ts *GroupDappService) getAllGroupDApps() []GroupDappModel{
	groupDapps := []GroupDappModel{}

	err := ts.db.Select(&groupDapps, "SELECT * FROM GroupDApp ORDER BY position")
	
	if err != nil {
		panic(err)
	}
	// header := models.Header{Success:true, Data: groupDapp}

	// result := utils.NewResultTransformer(header)

	return groupDapps
}

func (ts *GroupDappService) renameGroupDApp(name string,id int)error{
	_, err := ts.db.Exec("UPDATE GroupDApp SET name=? WHERE id=?",name,id)
		
	if err != nil {
		logger.Error(fmt.Sprintf("error when renameGroupDApp %", err))
		panic(fmt.Sprintf("error when renameGroupDApp %v", err))
		return err
	}
	fmt.Println("renameGroupDApp in database successed")
	return nil
}
func (ts *GroupDappService) updateGroupDAppPosition(position int,id int)error{
	_, err := ts.db.Exec("UPDATE GroupDApp SET position=? WHERE id=?",position,id)
		
	if err != nil {
		logger.Error(fmt.Sprintf("error when updateGroupDAppPosition %", err))
		panic(fmt.Sprintf("error when updateGroupDAppPosition %v", err))
		return err
	}
	fmt.Println("updateGroupDAppPosition in database successed")
	return nil
}
//decentralizedApplicationTB có có groupid, id lấy từ bảng GroupDApp
func (ts *GroupDappService) updateDAppGroupId(groupId int64,id int64)error{
	_, err := ts.db.Exec("UPDATE decentralizedApplicationTB SET groupId=? WHERE id=?",groupId,id)
		
	if err != nil {
		logger.Error(fmt.Sprintf("error when updateDAppGroupId %", err))
		panic(fmt.Sprintf("error when updateDAppGroupId %v", err))
		return err
	}
	fmt.Println("updateDAppGroupId in database successed")
	return nil
}

func (ts *GroupDappService) deleteGroupDApp(id int64)error{
	_, err := ts.db.Exec("DELETE FROM GroupDApp WHERE id=?",id)
		
	if err != nil {
		logger.Error(fmt.Sprintf("error when deleteGroupDApp %", err))
		// panic(fmt.Sprintf("error when deleteGroupDApp %v", err))
		return err
	}
	fmt.Println("deleteGroupDApp in database successed")
	return nil
}

// func (ts *GroupDappService) insertGroupDApp( name string, position int) int64 {
func (ts *GroupDappService) insertGroupDApp(groupDapp *GroupDappModel) int64 {
		_, err := ts.db.NamedExec("INSERT INTO GroupDApp( name, position) values (:name,:position)",
	map[string]interface{}{
	 	"name":groupDapp.Name,
		"position":groupDapp.Position,
	})
	lastGroupDapp := GroupDappModel{}
	err = ts.db.Get(&lastGroupDapp, "SELECT * FROM GroupDApp WHERE name =  ? and position=? ORDER BY id DESC LIMIT 1",groupDapp.Name,groupDapp.Position)
	if err != nil {
		return -1
	}

	fmt.Println("Insert GroupDApp successed")
	return lastGroupDapp.ID
}

func (ts *GroupDappService) getMaxPositionGroupDApp() int{
	var position int

	err := ts.db.Select(&position, "SELECT position FROM GroupDApp ORDER BY position DESC LIMIT 1")
	
	if err != nil {
		panic(err)
	}
	// header := models.Header{Success:true, Data: groupDapp}

	// result := utils.NewResultTransformer(header)

	return position
}
func (ts *GroupDappService) deleteAllGroupDApps()error{
	_, err := ts.db.Exec("DELETE FROM GroupDApp")
		
	if err != nil {
		logger.Error(fmt.Sprintf("error when deleteAllGroupDApps %", err))
		panic(fmt.Sprintf("error when deleteAllGroupDApps %v", err))
		return err
	}
	fmt.Println("deleteAllGroupDApps in database successed")
	return nil
}
