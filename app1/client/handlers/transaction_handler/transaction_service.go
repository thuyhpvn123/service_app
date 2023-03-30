package transactionhandler

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"gitlab.com/meta-node/client/models"
	"gitlab.com/meta-node/client/utils"
	"gitlab.com/meta-node/meta-node/pkg/logger"

)

// TransactionService struct
type TransactionService struct {
	db *sqlx.DB
}

// newTransactionService return new TransactionService object.
func newTransactionService(db *sqlx.DB) *TransactionService {
	return &TransactionService{db}
}

// getTransactions return all Transactions.
func (ts *TransactionService) getAllTransaction() []TransactionModel  {

	transactions := []TransactionModel{}

	// err := ts.db.Select(&transactions, "select * from transactionTB order by id asc limit 500")
	err := ts.db.Select(&transactions, "SELECT * FROM transactionTB ORDER BY time DESC")

	if err != nil {
		// panic(err)
		logger.Error(fmt.Sprintf("error when getAllTransaction %", err))

	}

	// header := models.Header{ Success:true, Data: transactions}

	// result := utils.NewResultTransformer(header)

	return transactions
}
// getTransactionByHash return Transactions by Hash.
func (ts *TransactionService) getTransactionByHash(hash string) (*utils.ResultTransformer ,error) {
	transaction := TransactionModel{}
	err := ts.db.Get(&transaction, "select * from transactionTB where hash = ? order by id desc",hash)
	if err != nil {
		// panic(err)
		logger.Error(fmt.Sprintf("error when getTransactionByHash %", err))
		
	}

	header := models.Header{ Success:true, Data: transaction}

	result := utils.NewResultTransformer(header)

	return result,err
}
// getTransactionPagination return Transactions by Pagination.
func (ts *TransactionService) getTransactionPagination(offset int,limit int) *utils.ResultTransformer  {
	transactions := []TransactionModel{}

	err := ts.db.Select(&transactions, "SELECT * FROM transactionTB  ORDER BY id DESC LIMIT ? OFFSET ? ;",limit,offset)
	if err != nil {
		logger.Error(fmt.Sprintf("error when getTransactionPagination %", err))
	}
	
	header := models.Header{Success:true,  Data: transactions}

	result := utils.NewResultTransformer(header)

	return result
}
// getTransactionPagination return Transactions by Pagination.
func (ts *TransactionService) getTransactionPaginationByAddress(offset int,limit int,address string) *utils.ResultTransformer  {
	transactions := []TransactionModel{}

	err := ts.db.Select(&transactions, "SELECT * FROM transactionTB WHERE (address = ? AND type = 'send') OR (toAddress = ? AND type = 'receive') ORDER BY id DESC LIMIT ? OFFSET ? ;",address,address,limit,offset)
	if err != nil {
		logger.Error(fmt.Sprintf("error when getTransactionPaginationByAddress %", err))
	}
	
	header := models.Header{Success:true,  Data: transactions}

	result := utils.NewResultTransformer(header)

	return result
}

// getTransactionAtAddress return all Transactions.
func (ts *TransactionService) getTransactionByAddress(offset int,limit int,address string, status int) *utils.ResultTransformer {

	transactions := []TransactionModel{}

	err := ts.db.Select(&transactions, "SELECT * FROM transactionTB WHERE (address= ? AND type = 'send' AND status = ?) OR (toAddress= ? AND type = 'receive' AND status = ?) ORDER BY id DESC LIMIT ? OFFSET ? ", address, status, address, status, limit, offset)
	if err != nil {
		logger.Error(fmt.Sprintf("error when getTransactionByAddress %", err))
	}

	header := models.Header{Success:true, Data: transactions}

	result := utils.NewResultTransformer(header)

	return result
}
// getLastTransaction return last Transaction .
func (ts *TransactionService) getLastTransaction(address string) (TransactionModel,int){

	transaction := TransactionModel{}

	err := ts.db.Get(&transaction, "SELECT * FROM transactionTB WHERE address =  ? ORDER BY id DESC LIMIT 1",address)

	if err != nil {
		// panic(err)
		logger.Error(fmt.Sprintf("error when getLastTransaction %", err))

		return transaction,0
	}

	// header := models.Header{Success:true, Data: transaction}

	// result := utils.NewResultTransformer(header)

	return transaction,1
}
func (ts *TransactionService) getLastTransactionSmartContract(address string) (TransactionModel,error){

	transaction := TransactionModel{}

	err := ts.db.Get(&transaction, "SELECT * FROM transactionTB WHERE address = ? and status= 2 and isCall = 1 ORDER BY id DESC LIMIT 1",address)

	if err != nil {
		// panic(err)
		logger.Error(fmt.Sprintf("error when getLastTransactionSmartContract %", err))

	}

	// header := models.Header{Success:true, Data: transaction}

	// result := utils.NewResultTransformer(header)

	return transaction,err
}
// getLastTransaction return last Transaction .
func (ts *TransactionService) getLastTransactionWithStatus(address string,status int64) (TransactionModel,int) {

	transaction := TransactionModel{}

	err := ts.db.Get(&transaction, "SELECT * FROM transactionTB WHERE address =  ? and status=? ORDER BY id DESC LIMIT 1",address,status)
	if err != nil {
		// panic(err)
		logger.Error(fmt.Sprintf("error when getLastTransactionWithStatus %", err))

		return transaction,0
	}

	// header := models.Header{Success:true, Data: transaction}

	// result := utils.NewResultTransformer(header)

	// return result
	return transaction,1

}

//  updateStatusTransactionByHash update status of Transaction by hash .
func (ts *TransactionService) updateStatusTransactionByHash(status int,hash string) error {
	fmt.Println("hash là:",hash)

	// updateSQL := fmt.Sprintf("UPDATE transactionTB SET status = ? where hash = ? ", status, hash)

	_,err := ts.db.Exec("UPDATE transactionTB SET status = ? where hash = ? ", status, hash)
	if err != nil {
		logger.Error(fmt.Sprintf("error when updateStatusTransactionByHash %", err))

	}

	return nil
}

func (ts *TransactionService) getTotalTransactionSuccess(address string) *utils.ResultTransformer {

	var count int

	err := ts.db.Get(&count, "SELECT COUNT(*) FROM transactionTB WHERE type = 'send' and status = 2 and address = ? ", address)
	if err != nil {
		logger.Error(fmt.Sprintf("error when getTotalTransactionSuccess %", err))
	}

	header := models.Header{Success:true, Data: count}

	result := utils.NewResultTransformer(header)

	return result
}

//  updateStatusTransactionByHash update transactions with status 0 or 1 to 4 .
// status: Int = -1, // 0: pending, 1 = sent, 2 = success, 3: fail, 4: cancel
func (ts *TransactionService) updateTransactionWithStatusPending(time int64) error {
	// updateSQL := fmt.Sprintf("UPDATE transactionTB SET status = 4 WHERE time >= ? AND status != 2", time)

	_,err := ts.db.Exec("UPDATE transactionTB SET status = 4 WHERE time >= ? AND status != 2", time)

	if err != nil {
		logger.Error(fmt.Sprintf("error when updateTransactionWithStatusPending %v", err))
		panic(fmt.Sprintf("error when updateTransactionWithStatusPending %v", err))

		return err
	}
	// if rowsAffected == 0 {
	// 	return err
	// }

	return nil
}
func (ts *TransactionService) insertTransaction(tx *TransactionModel)int{
	_, err := ts.db.NamedExec("INSERT INTO transactionTB( hash,address,toAddress,pubKey,amount,pendingUse,tip,message,time,status,type,prevHash,sign,receiveInfo,isDeploy,isCall,functionCall,data,totalBalance,lastDeviceKey ) values(:hash,:address,:toAddress,:pubKey,:amount,:pendingUse,:tip,:message,:time,:status,:type,:prevHash,:sign,:receiveInfo,:isDeploy,:isCall,:functionCall,:data,:totalBalance,:lastDeviceKey)",
	map[string]interface{}{
	"hash":tx.Hash,     
	"address":tx.Address, 
	"toAddress":tx.ToAddress, 
	"pubKey":tx.PubKey, 
	"amount":tx.Amount, 
	"pendingUse":tx.PendingUse ,
	// "balance":tx.Balance ,
	// "fee":tx.Fee ,
	"tip":tx.Tip ,
	"message":tx.Message ,
	"time":tx.Time ,
	"status":tx.Status ,
	"type":tx.Type ,
	"prevHash":tx.PrevHash ,
	"sign":tx.Sign,
	"receiveInfo":tx.ReceiveInfo ,
	"isDeploy":tx.IsDeploy ,
	"isCall":tx.IsCall ,
	"functionCall":tx.FunctionCall ,
	"data":tx.Data ,
	"totalBalance":tx.TotalBalance ,
	"lastDeviceKey":tx.LastDeviceKey ,
		
	})

		if err != nil {
			logger.Error(fmt.Sprintf("error when insertTransaction %", err))
			panic(fmt.Sprintf("error when insertTransaction %v", err))
			return -1
		}
	
		fmt.Println("Insert Transaction in database successed")
		return 1
}
	
	
func (ts *TransactionService) isExistTransaction(hash string) bool {
	_,err := ts.db.Exec("SELECT EXISTS(SELECT hash FROM transactionTB WHERE hash =?)",hash)
	if err != nil {
		logger.Error(fmt.Sprintf("error when check isExistRecentNode %", err))

		return false
	}
	return true
}
func (ts *TransactionService) deleteAllTrans()error{
	_, err := ts.db.Exec("DELETE FROM transactionTB")
		
	if err != nil {
		logger.Error(fmt.Sprintf("error when DeleteAllTrans %", err))
		panic(fmt.Sprintf("error when DeleteAllTrans %v", err))
		return err
	}
	fmt.Println("DeleteAllTrans in database successed")
	return nil
}
	
	
