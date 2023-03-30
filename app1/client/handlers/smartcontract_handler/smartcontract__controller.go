package smartContracthandler

import (
	// "fmt"
	// "log"
	// "net/http"

	"github.com/jmoiron/sqlx"
	// "gitlab.com/meta-node/client/models"

	"gitlab.com/meta-node/client/utils"
)

// SmartContractController struct
type SmartContractController struct {
	service *SmartContractService
}

// NewSmartContractController return new SmartContractController object.
func NewSmartContractController(db *sqlx.DB) *SmartContractController {
	return &SmartContractController{newSmartContractService(db)}
}

// GetSmartContracts return all SmartContracts.
func (tc *SmartContractController) GetSmartContracts() *utils.ResultTransformer {

	smartContracts:= tc.service.getSmartContracts()
	return smartContracts
}
func (tc *SmartContractController) GetLastTransactionSmartContract(address string) []SmartContractModel{

	smartContracts:= tc.service.getLastTransactionSmartContract(address)
	return smartContracts
}
