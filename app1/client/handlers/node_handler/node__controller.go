package nodehandler

import (
	// "encoding/json"
	// "fmt"
	// "log"
	// "net/http"

	"github.com/jmoiron/sqlx"
	// "gitlab.com/meta-node/client/server/core/request"

	// "gitlab.com/meta-node/client/utils"
)

// NodeController struct
type NodeController struct {
	service *NodeService
}

// NewNodeController return new NodeController object.
func NewNodeController(db *sqlx.DB) *NodeController {
	return &NodeController{newNodeService(db)}
}

// GetNodes return all Nodes.
func (tc *NodeController) GetAllRecentNode() []NodeModel{
	nodes:= tc.service.getAllRecentNode()
	return nodes
}
// InsertNode 
func (tc *NodeController) InsertRecentNode(request *NodeModel)int {

	kq := tc.service.insertRecentNode(request.IP,request.Port,request.Time)
	return kq
}
func (tc *NodeController) IsExistRecentNode(ip string,port int) bool {

	kq := tc.service.isExistRecentNode(ip,port)
	return kq
}
func (tc *NodeController) DeleteAllRecentNodes()error {

	kq := tc.service.deleteAllRecentNodes()
	return kq
}


