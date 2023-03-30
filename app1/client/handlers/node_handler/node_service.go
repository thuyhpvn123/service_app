package nodehandler

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"gitlab.com/meta-node/meta-node/pkg/logger"

	// "gitlab.com/meta-node/client/models"
	// "gitlab.com/meta-node/client/utils"
)

// NodeService struct
type NodeService struct {
	db *sqlx.DB
}

// newNodeService return new NodeService object.
func newNodeService(db *sqlx.DB) *NodeService {
	return &NodeService{db}
}

// getAllRecentNode return all Nodes.
func (ts *NodeService) getAllRecentNode() []NodeModel {

	nodes := []NodeModel{}

	err := ts.db.Select(&nodes, "SELECT * FROM recentNodeTB ORDER BY time DESC LIMIT 10 OFFSET 0")
	if err != nil {
		panic(err)
	}

	// header := models.Header{Status: "ok", Count: len(nodes), Data: nodes}
	// result := utils.NewResultTransformer(header)

	return nodes
}
// // getNodeByHash return Nodes by Hash.
// func (ts *NodeService) getNodeByHash(hash string) *utils.ResultTransformer {
// 	transaction := NodeModel{}
// 	err := ts.db.Get(&transaction, "select * from transactionTB where hash = ? order by id desc",hash)
// 	if err != nil {
// 		panic(err)
// 	}

// 	header := models.Header{Status: "ok", Count: 1, Data: transaction}
// 	result := utils.NewResultTransformer(header)

// 	return result
// }
// // getNodePagination return Nodes by Pagination.
// func (ts *NodeService) getNodePagination(offset int,limit int,address string) *utils.ResultTransformer {
// 	nodes := []NodeModel{}

// 	err := ts.db.Select(&nodes, "SELECT * FROM transactionTB WHERE (address = ? AND type = 'send') OR (toAddress = ? AND type = 'receive') ORDER BY id DESC LIMIT ? OFFSET ? ;",address,address,limit,offset)
// 	if err != nil {
// 		panic(err)
// 	}
	
// 	header := models.Header{Status: "ok", Count: len(nodes), Data: nodes}
// 	result := utils.NewResultTransformer(header)

// 	return result
// }
// // getNodeAtAddress return all Nodes.
// func (ts *NodeService) getNodeByAddress(offset int,limit int,address string, status int) *utils.ResultTransformer {

// 	nodes := []NodeModel{}

// 	err := ts.db.Select(&nodes, "SELECT * FROM transactionTB WHERE (address= ? AND type = 'send' AND status = ?) OR (toAddress= ? AND type = 'receive' AND status = ?) ORDER BY id DESC LIMIT ? OFFSET ? ", address, status, address, status, limit, offset)
// 	if err != nil {
// 		panic(err)
// 	}

// 	header := models.Header{Status: "ok", Count: len(nodes), Data: nodes}
// 	result := utils.NewResultTransformer(header)

// 	return result
// }
// insertNodeAtAddress return all Nodes.
func (ts *NodeService) insertRecentNode(ip string, port int, time int64) int {
	_, err := ts.db.NamedExec("INSERT INTO recentNodeTB(ip, port, time) values (:ip, :port, :time)",	
	map[string]interface{}{
		"ip":ip,
	 	"port":port,
		"time":time,
	})

	if err != nil {
		logger.Error(fmt.Sprintf("error when insertDapp %", err))
		panic(fmt.Sprintf("error when insertDapp %v", err))
		return -1
	}

	fmt.Println("Insert Dapp in database successed")
	return 1
}
func (ts *NodeService) isExistRecentNode(ip string,port int) bool {
	_,err := ts.db.Exec("SELECT EXISTS (SELECT * FROM recent_node_table WHERE ip=? AND port=?)",ip,port)
	if err != nil {
		logger.Error(fmt.Sprintf("error when check isExistRecentNode %", err))

		return false
	}
	return true
}
func (ts *NodeService) deleteAllRecentNodes()error{
	_, err := ts.db.Exec("DELETE FROM decentralizedApplicationTB")
		
	if err != nil {
		logger.Error(fmt.Sprintf("error when deleteAllRecentNodes %", err))
		panic(fmt.Sprintf("error when deleteAllRecentNodes %v", err))
		return err
	}
	fmt.Println("deleteAllRecentNodes in database successed")
	return nil
}





