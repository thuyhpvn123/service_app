package mysettinghandler

import (
	// "encoding/json"
	// "fmt"
	// "gitlab.com/meta-node/client/models"

	"github.com/jmoiron/sqlx"
	// "gitlab.com/meta-node/client/utils"
)

// MysettingController struct
type MysettingController struct {
	service *MysettingService
}

// NewMysettingController return new MysettingController object.
func NewMysettingController(db *sqlx.DB) *MysettingController {
	return &MysettingController{newMysettingService(db)}
}


func (tc *MysettingController) GetMysetting() (MysettingModel,error){
	mysetting, err:= tc.service.getMysetting()
	return mysetting,err
}

func (tc *MysettingController) UpdatePinCode(kq bool) error {

	err := tc.service.updatePinCode(kq)
	return err
}
func (tc *MysettingController) InsertMysetting(setting *MysettingModel) int {

	kq := tc.service.insertMysetting(setting)
	return kq
}
