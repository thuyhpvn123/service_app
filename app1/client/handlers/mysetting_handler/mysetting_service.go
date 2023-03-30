package mysettinghandler

import (
	// "fmt"

	// "encoding/json"
	"fmt"
	// "log"

	"github.com/jmoiron/sqlx"

	// "gitlab.com/meta-node/client/models"
	// "gitlab.com/meta-node/client/utils"
	"gitlab.com/meta-node/meta-node/pkg/logger"
)

// MysettingService struct
type MysettingService struct {
	db *sqlx.DB
}

// newMysettingService return new MysettingService object.
func newMysettingService(db *sqlx.DB) *MysettingService {
	return &MysettingService{db}
}


// getMysettingPagination return Mysettings by Pagination.
func (ts *MysettingService) getMysetting() (MysettingModel,error) {
	mysetting := MysettingModel{}

	err := ts.db.Get(&mysetting, "SELECT * FROM mysetting ")
	if err != nil {
		logger.Error(fmt.Sprintf("error when getMysetting %", err))
		panic(fmt.Sprintf("error when getMysetting %v", err))
		return mysetting,err
	}
	

	return mysetting,nil
}

func (ts *MysettingService) updatePinCode(kq bool)error{
	_, err := ts.db.Exec("UPDATE mysetting SET isHasPincode=? WHERE id=0",kq)
		
	if err != nil {
		logger.Error(fmt.Sprintf("error when updatePinCode %", err))
		panic(fmt.Sprintf("error when updatePinCode %v", err))
		return err
	}
	fmt.Println("updatePinCode in database successed")
	return nil
}

func (ts *MysettingService) insertMysetting(setting *MysettingModel)int{
	mysetting:=map[string]interface{}{
		"receiveNotificationFromNode":setting.ReceiveNotificationFromNode ,
		"receiveTransactionStatus":setting.ReceiveTransactionStatus  ,
		"nodeAddress":setting.NodeAddress ,
		"networkType":setting.NetworkType,
		"isHasPincode":setting.IsHasPincode,
		"isEnableFaceOrTouchId":setting.IsEnableFaceOrTouchId,
		"isWatchConfirm":setting.IsWatchConfirm,
		"accountInfo":setting.AccountInfo ,
	}
	fmt.Sprintf("MySetting map", mysetting)

_, err := ts.db.NamedExec(`
INSERT INTO mysetting(
	receiveNotificationFromNode, 
	receiveTransactionStatus,
	nodeAddress,
	networkType,
	isHasPincode,
	isEnableFaceOrTouchId,
	isWatchConfirm,
	accountInfo) values (:receiveNotificationFromNode,:receiveTransactionStatus,:nodeAddress ,:networkType,:isHasPincode,:isEnableFaceOrTouchId,:isWatchConfirm,:accountInfo)`,
mysetting)
	
	if err != nil {
		logger.Error(fmt.Sprintf("error when insertMysetting %", err))
		panic(fmt.Sprintf("error when insertMysetting %v", err))
		return -1
	}

	fmt.Println("Insert Mysetting in database successed")
	return 1
}


