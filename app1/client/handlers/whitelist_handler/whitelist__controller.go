package whitelisthandler

import (
// "gitlab.com/meta-node/client/models"

	"github.com/jmoiron/sqlx"
	// "gitlab.com/meta-node/client/utils"
)

// WhiteListController struct
type WhiteListController struct {
	service *WhiteListService
}

// NewWhiteListController return new WhiteListController object.
func NewWhiteListController(db *sqlx.DB) *WhiteListController {
	return &WhiteListController{newWhiteListService(db)}
}


func (tc *WhiteListController) GetWhitelistPagination(limit int,offset int) []WhiteListModel{
	whitelist:= tc.service.getWhitelistPagination(limit,offset)
	return whitelist
}
func (tc *WhiteListController) GetAllWhitelist() []WhiteListModel{
	whitelist:= tc.service.getAllWhitelist()
	return whitelist
}
func (tc *WhiteListController) InsertWhitelist(whitelist *WhiteListModel) int {

	kq := tc.service.insertWhitelist(whitelist)
	return kq
}
func (tc *WhiteListController) IsExistWhitelist(id int) bool {

	kq := tc.service.isExistWhitelist(id)
	return kq
}
func (tc *WhiteListController) DeleteAllWhitelists()error {

	kq := tc.service.deleteAllWhitelists()
	return kq
}
