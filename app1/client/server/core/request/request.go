package request

type GetAccountAddressRequest struct {
	PrivateKey string `json:"private_key"`
}

type TransferRequest struct {
	PrivateKey                  string `json:"private_key"`
	ToAddress                   string `json:"to_address"`
	Amount                      string `json:"amount"`
	SmartContractType           string `json:"smart_contract_type"`
	StorageHost                 string `json:"storage_host"`
	SmartContractInput          string `json:"smart_contract_input"`
	SmartContractFeeType        int    `json:"smart_contract_fee_type"`
	SmartContractCommissionSign string `json:"smart_contract_commission_sign"`
}

type GetAccountInfoRequest struct {
	Address string `json:"address"`
}

//Wallet Request
type GetWalletPagination struct {
	Offset int `json:"offset"`
	Limit int `json:"limit"`
}

type GetWalletAtAddressRequest struct {
	Address string `json:"address"`
}
type InsertWalletRequest struct {
	Address string `db:"address" json:"address"`
	Name     string    `db:"name" json:"name"`
	PendingBalance string `db:"pendingBalance" json:"pendingBalance"`
	Balance string `db:"balance" json:"balance"`
	TotalBalance string `db:"totalBalance" json:"totalBalance"`
	Bg string `db:"bg" json:"bg"`
	Color string `db:"color" json:"color"`
	IdSymbol string `db:"idSymbol" json:"idSymbol"`
	Pattern string `db:"pattern" json:"pattern"`
	Position int`db:"position" json:"position"`
}

//Transaction Request
type GetTransactionRequest struct {
	Hash string `json:"hash"`
}
type GetTransactionByAddressRequest struct {
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Address string `json:"address"`
	Status int `json:"status"`
}
type GetTransactionByHashRequest struct {
	Hash string `json:"hash"`
}
type GetTransactionPagination struct {
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Address string `json:"address"`
}

//Dapp Request
type GetDappPaginationRequest struct {
	Offset int `json:"offset"`
	Limit int `json:"limit"`
}




type GenerateCommissionSignRequest struct {
	PrivateKey string `json:"private_key"`
	Address    string `json:"address"`
}




