package smartContracthandler

import (

	"github.com/jmoiron/sqlx"

	"gitlab.com/meta-node/client/models"
	"gitlab.com/meta-node/client/utils"

)

// SmartContractService struct
type SmartContractService struct {
	db *sqlx.DB
}

// newSmartContractService return new SmartContractService object.
func newSmartContractService(db *sqlx.DB) *SmartContractService {
	return &SmartContractService{db}
}

// getSmartContracts return all SmartContracts.
func (ts *SmartContractService) getSmartContracts() *utils.ResultTransformer  {

	smartContracts := []SmartContractModel{}

	err := ts.db.Select(&smartContracts, "select * from smartContractTB order by id asc limit 500")
	if err != nil {
		panic(err)
	}

	header := models.Header{Success:true,  Data: smartContracts}

	result := utils.NewResultTransformer(header)

	return result
}
func (ts *SmartContractService) getLastTransactionSmartContract(address string) []SmartContractModel {

	smartContracts := []SmartContractModel{}

	err := ts.db.Select(&smartContracts, "SELECT * FROM transactionTB WHERE address = ? and status= 2 and isCall = 1 ORDER BY id DESC LIMIT 1",address)
	if err != nil {
		panic(err)
	}

	// header := models.Header{Success:true,  Data: smartContracts}

	// result := utils.NewResultTransformer(header)

	return smartContracts
}
