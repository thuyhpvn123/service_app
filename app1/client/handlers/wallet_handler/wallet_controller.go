package wallethandler

import (
	// "encoding/json"
	// "fmt"
	"log"
	// "strconv"
	// "net/http"
	// "gitlab.com/meta-node/client/models"

	"github.com/jmoiron/sqlx"
	// "gitlab.com/meta-node/client/server/core/request"
	"gitlab.com/meta-node/client/utils"
)

// WalletController struct
type WalletController struct {
	service *WalletService
}

// NewWalletController return new WalletController object.
func NewWalletController(db *sqlx.DB) *WalletController {
	return &WalletController{newWalletService(db)}
}

// GetWallets return all Wallets.
func (tc *WalletController) GetAllWallets() []WalletModel {

	wallets := tc.service.getAllWallets()
	return wallets
}

// GetWallets return all Wallets.
func (tc *WalletController) GetWalletPagination(offset int, limit int) *utils.ResultTransformer {
	wallets := tc.service.getWalletPagination(offset, limit)
	return wallets
}

// GetWalletAtAddress return all Wallets At Address.
func (tc *WalletController) GetWalletByAddress(address string) (WalletModelShort, error) {

	wallets, err := tc.service.getWalletByAddress(address)
	return wallets, err
}
func (tc *WalletController) GetWalletByAddress1(address string) (WalletModel, error) {

	wallets, err := tc.service.getWalletByAddress1(address)
	return wallets, err
}

// InsertWallet
func (tc *WalletController) InsertWallet1(wallet *WalletModel) {

	err := tc.service.insertWallet( wallet)
	if err != nil {
		log.Fatal()
		return
	}
}

// CountWalletByName return number of wallets  by name.
func (tc *WalletController) CountWalletByName(name string) int {
	count := tc.service.countWalletByName(name)
	return count
}

// CountWalletTable return number of wallets
func (tc *WalletController) CountWalletTable() int {
	count := tc.service.countWalletTable()
	return count
}
func (tc *WalletController) UpdateBalanceAtAddress(balance string, pendingBalance string, total string, address string) error {

	err := tc.service.updateBalanceAtAddress(balance, pendingBalance, total, address)
	if err != nil {
		log.Fatal()
		return err
	}
	return nil
}
func (tc *WalletController) IsExistWallet(address string) bool {

	kq := tc.service.isExistWallet(address)
	return kq
}
func (tc *WalletController) UpdateWalletPosition(position int, id int) error {

	err := tc.service.updateWalletPosition(position, id)
	return err
}
func (tc *WalletController) EditWalletUI(bg string, color string, idSymbol int,pattern string, address string) error {

	err := tc.service.editWalletUI(bg , color , idSymbol ,pattern , address)
	return err
}

func (tc *WalletController) DeleteAllWallets()error {

	kq := tc.service.deleteAllWallets()
	return kq
}
