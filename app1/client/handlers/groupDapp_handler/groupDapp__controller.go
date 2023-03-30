package groupDapphandler

import (
	// "encoding/json"
	// "fmt"
	// "log"
	// "log"
	// "net/http"

	// "log"
	// "gitlab.com/meta-node/client/models"

	"github.com/jmoiron/sqlx"
	// "gitlab.com/meta-node/client/server/core/request"
	// "gitlab.com/meta-node/client/utils"
)

// DappController struct
type GroupDappController struct {
	service *GroupDappService
}

// NewDappController return new DappController object.
func NewGroupDappController(db *sqlx.DB) *GroupDappController {
	return &GroupDappController{newGroupDappService(db)}
}

func (tc *GroupDappController) GetAllGroupDApps() []GroupDappModel {
	groupDapp:= tc.service.getAllGroupDApps()

	return groupDapp
}
func (tc *GroupDappController) GetMaxPositionGroupDApp() int {
	position:= tc.service.getMaxPositionGroupDApp()

	return position
}

func (tc *GroupDappController) UpdateGroupDAppPosition(position int,id int) error {

	err := tc.service.updateGroupDAppPosition(position,id)
	return err
}

func (tc *GroupDappController) RenameGroupDApp(name string,id int) error {

	err := tc.service.renameGroupDApp(name,id)
	return err
}
func (tc *GroupDappController) UpdateDAppGroupId(groupId int64,id int64) error {

	err := tc.service.updateDAppGroupId(groupId,id)
	return err
}

func (tc *GroupDappController) DeleteGroupDApp(id int64) error {

	err := tc.service.deleteGroupDApp(id)
	return err
}
func (tc *GroupDappController) InsertGroupDApp( groupDapp *GroupDappModel) int64 {

	groupID := tc.service.insertGroupDApp(groupDapp)
	return groupID
}
func (tc *GroupDappController) DeleteAllGroupDApps()error {

	kq := tc.service.deleteAllGroupDApps()
	return kq
}




