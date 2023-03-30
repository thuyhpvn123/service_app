package transactionhandler

import (
	// "encoding/json"
	"fmt"
	// "net/http"
	// "log"
	// "strconv"

	// "gitlab.com/meta-node/client/models"

	"log"

	"github.com/jmoiron/sqlx"
	// "gitlab.com/meta-node/client/server/core/request"
	"gitlab.com/meta-node/client/utils"
)

// TransactionController struct
type TransactionController struct {
	service *TransactionService
}

// NewTransactionController return new TransactionController object.
func NewTransactionController(db *sqlx.DB) *TransactionController {
	return &TransactionController{newTransactionService(db)}
}

// GetTransactions return all Transactions.
func (tc *TransactionController) GetAllTransaction() []TransactionModel {
	transactions := tc.service.getAllTransaction()
	return transactions
}
// getTransactionByHash return Transactions by Hash.
func (tc *TransactionController) GetTransactionByHash(hash string)(*utils.ResultTransformer,error)   {
	transaction,err := tc.service.getTransactionByHash(hash)
	return transaction,err
}
func (tc *TransactionController) GetLastTransactionSmartContract(address string) (TransactionModel,error) {
	transaction,err := tc.service.getLastTransactionSmartContract(address)
	return transaction,err
}

// GetTransactions return all Transaction  by limit, offset,.
func (tc *TransactionController) GetTransactionPagination(offset int,limit int) *utils.ResultTransformer {

	transactions:= tc.service.getTransactionPagination(offset,limit)
	return transactions
}
// GetTransactions return all Transaction  by Address, limit, offset,.
func (tc *TransactionController) GetTransactionPaginationByAddress(offset int,limit int,address string) *utils.ResultTransformer {

	transactions:= tc.service.getTransactionPaginationByAddress(offset,limit,address)
	return transactions
}

// getTransactionByAddress return Transactions by Address, limit, offset, status.
func (tc *TransactionController) GetTransactionByAddress(offset int,limit int,address string, status int)*utils.ResultTransformer  {
	transactions := tc.service.getTransactionByAddress(offset,limit,address,status)
	return transactions
}
// GetLastTransaction return the last Transactions by Address.
func (tc *TransactionController) GetLastTransaction(address string)(TransactionModel,int) {
	transaction,num := tc.service.getLastTransaction(address)
	return transaction,num
}
// GetLastTransaction return the last Transactions by Address and status.
func (tc *TransactionController) GetLastTransactionWithStatus(address string,status int64)(TransactionModel,int) {
	transaction,num := tc.service.getLastTransactionWithStatus(address,status)
	return transaction,num
}

func (tc *TransactionController) UpdateStatusTransactionByHash(status int,hash string) {

	err := tc.service.updateStatusTransactionByHash(status,hash)
   if err != nil {
	fmt.Println("gsdsgsgsdg")
	   log.Fatal()
	//    return
   }
}
// GetTotalTransactionSuccess return number of Transactions success by address.
func (tc *TransactionController) GetTotalTransactionSuccess(address string)*utils.ResultTransformer  {
	count := tc.service.getTotalTransactionSuccess(address)
	return count
}
//  updateStatusTransactionByHash update transactions with status 0 or 1 to 4 .
func (tc *TransactionController) UpdateTransactionWithStatusPending(time int64) error{

	err := tc.service.updateTransactionWithStatusPending(time)
	fmt.Println("UpdateTransactionWithStatusPending success")

	return err
}

func (tc *TransactionController) InsertTransaction(dapp *TransactionModel) int {

	kq := tc.service.insertTransaction(dapp)
	return kq
}

func (tc *TransactionController) IsExistTransaction(hash string) bool {

	kq := tc.service.isExistTransaction(hash)
	return kq
}
func (tc *TransactionController) DeleteAllTrans()error {

	kq := tc.service.deleteAllTrans()
	return kq
}
