package wallethandler

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"gitlab.com/meta-node/client/models"
	"gitlab.com/meta-node/client/utils"
	"gitlab.com/meta-node/meta-node/pkg/logger"

)

// WalletService struct
type WalletService struct {
	db *sqlx.DB
}

// newWalletService return new WalletService object.
func newWalletService(db *sqlx.DB) *WalletService {
	return &WalletService{db}
}

// getWallets return all Wallets. //get-all-wallet
func (ts *WalletService) getAllWallets() []WalletModel {

	wallets := []WalletModel{}

	err := ts.db.Select(&wallets, "SELECT * FROM walletTB ORDER BY position")
	if err != nil {
		logger.Error(fmt.Sprintf("error when getAllWallets %", err))
	}

	// header := models.Header{ Success:true, Data: wallets}

	// result := utils.NewResultTransformer(header)

	return wallets
}
// getWalletPagination return Wallets by Pagination.
func (ts *WalletService) getWalletPagination(offset int,limit int) *utils.ResultTransformer  {
	wallets := []WalletModel{}
	err := ts.db.Select(&wallets, "SELECT * FROM walletTB ORDER BY id ASC LIMIT ? OFFSET ? ;",limit,offset)
	if err != nil {
		logger.Error(fmt.Sprintf("error when getWalletPagination %", err))
	}
	
	header := models.Header{ Success:true, Data: wallets}

	result := utils.NewResultTransformer(header)

	return result
}
// getWalletAtAddress return all Wallets.
func (ts *WalletService) getWalletByAddress(address string) (WalletModelShort,error)  {

	wallet := WalletModel{}

	err := ts.db.Get(&wallet, "SELECT * FROM walletTB WHERE address = ?",address)
	if err != nil {
		logger.Error(fmt.Sprintf("error when getWalletByAddress %", err))
		return WalletModelShort{},err

	}
	walletshort := WalletModelShort{
		Address:wallet.Address,
		Balance:wallet.Balance,
		PendingBalance:wallet.PendingBalance,
	}

	// header := models.Header{Success:true,  Data: walletshort}

	// result := utils.NewResultTransformer(header)

	return walletshort,nil
}
func (ts *WalletService) getWalletByAddress1(address string) (WalletModel,error)  {

	wallet := WalletModel{}

	err := ts.db.Get(&wallet, "SELECT * FROM walletTB WHERE address = ?",address)
	fmt.Println("getWalletByAddress1")
	if err != nil {
		logger.Error(fmt.Sprintf("error when getWalletByAddress %v", err))
		return WalletModel{},err

	}

	return wallet,nil
}

// insertWalletAtAddress return all Wallets.
func (ts *WalletService) insertWallet( wallet *WalletModel) error {

	_, err := ts.db.NamedExec("INSERT INTO walletTB(address, name, pendingBalance, balance, totalBalance, bg, color, idSymbol, pattern, position) values (:address, :name, :pendingBalance, :balance, :totalBalance, :bg, :color, :idSymbol, :pattern, :position)",
	map[string]interface{}{
		"address":wallet.Address,
	 	"name":wallet.Name,
		"pendingBalance":wallet.PendingBalance,
		"balance":wallet.Balance, 
		"totalBalance":wallet.TotalBalance, 
		"bg":wallet.Bg, 
		"color":wallet.Color, 
		"idSymbol":wallet.IdSymbol, 
		"pattern":wallet.Pattern, 
		"position":wallet.Position,
	})
	if err != nil {
		return err
	}
	fmt.Println("Insert Wallet successed")
	return nil
}
// countWalletByName return number of wallet by name.
func (ts *WalletService) countWalletByName(name string) int {

	var count int

	err := ts.db.Get(&count, "SELECT COUNT(*) FROM walletTB WHERE name LIKE :name || '%'", name)
	if err != nil {
		panic(err)
	}
	return count
}
// countWalletByName return last Transaction .
func (ts *WalletService) countWalletTable() int {

	var count int

	err := ts.db.Get(&count, "SELECT COUNT(*) FROM walletTB")
	if err != nil {
		panic(err)
	}
	return count
}
//  UpdateBalanceAtAddress update status of Transaction by hash .
func (ts *WalletService) updateBalanceAtAddress(balance string, pendingBalance string, total string, address string) error {
	// updateSQL := fmt.Sprintf()
	_,err := ts.db.Exec("UPDATE walletTB SET balance=?, pendingBalance =?, totalBalance =? where address =?", balance, pendingBalance,total,address)
	if err != nil {
		logger.Error(fmt.Sprintf("error when updateBalanceAtAddress %", err))

		// return err
	}

	return nil
}
func (ts *WalletService) editWalletUI(bg string, color string, idSymbol int,pattern string, address string) error {
	// updateSQL := fmt.Sprintf()
	_,err := ts.db.Exec("UPDATE walletTB SET bg=?, color=?, idSymbol=?, pattern=? WHERE address=?", bg, color,idSymbol,pattern,address)
	if err != nil {
		logger.Error(fmt.Sprintf("error when editWalletUI %", err))

		// return err
	}

	return nil
}

func (ts *WalletService) updateWalletPosition(position int, id int) error {
	_,err := ts.db.Exec("UPDATE walletTB SET position=? WHERE id=?", position, id)
	if err != nil {
		logger.Error(fmt.Sprintf("error when updateBalanceAtAddress %", err))

		return err
	}

	return nil
}
func (ts *WalletService) isExistWallet(address string) bool {
	_,err := ts.db.Exec("SELECT EXISTS(SELECT address FROM walletTB WHERE address =?)", address)
	if err != nil {
		logger.Error(fmt.Sprintf("error when check isExistWallet %", err))

		return false
	}

	return true
}
func (ts *WalletService) deleteAllWallets()error{
	_, err := ts.db.Exec("DELETE FROM walletTB")
		
	if err != nil {
		logger.Error(fmt.Sprintf("error when deleteAllWallets %", err))
		panic(fmt.Sprintf("error when deleteAllWallets %v", err))
		return err
	}
	fmt.Println("deleteAllWallets in database successed")
	return nil
}


