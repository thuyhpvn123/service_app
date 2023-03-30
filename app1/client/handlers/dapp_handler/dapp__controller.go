package dapphandler

import (
	// "encoding/json"
	// "fmt"
	// "log"
	// "log"
	// "net/http"

	// "log"
	// "gitlab.com/meta-node/client/models"

	"github.com/jmoiron/sqlx"
	// "gitlab.com/meta-node/client/server/core/request"
	"gitlab.com/meta-node/client/utils"
)

// DappController struct
type DappController struct {
	service *DappService
}

// NewDappController return new DappController object.
func NewDappController(db *sqlx.DB) *DappController {
	return &DappController{newDappService(db)}
}


// getDappByAddress return Dapps by Address, limit, offset, status.
func (tc *DappController) GetDappByName(name string) []DappModel{
	dapp:= tc.service.getDappByName(name)
	return dapp
}
//get-d-app-by-bundle-id
func (tc *DappController) GetDappByBundleId(Id string)*utils.ResultTransformer  {
	result:= tc.service.getDappByBundleId(Id)
	// 	if err != nil {
	// 	log.Fatal()
	// }

	return result
}
func (tc *DappController) GetAllDAppsForBrowser() ([]DappModel,error){
	dapp,err:= tc.service.getAllDAppsForBrowser()

	return dapp,err
}

func (tc *DappController) GetAllDappNoGroup() *utils.ResultTransformer {
	dapp:= tc.service.getAllDappNoGroup()

	return dapp
}
func (tc *DappController) GetAllDApps() []DappModel {
	dapp:= tc.service.getAllDApps()

	return dapp
}

// GetMaxPositionWallet return maxPosition when create-wallet 
func (tc *DappController) GetMaxPositionWallet()int  {
	count := tc.service.getMaxPositionWallet()
	return count
}

func (tc *DappController) GetAllDAppsByGroupIdAndPage(groupId int64,page int) []DappModel{
	dapp:= tc.service.getAllDAppsByGroupIdAndPage(groupId ,page )
	return dapp
}
func (tc *DappController) UpdateDAppPage(page int,id int64) error{
	err:= tc.service.updateDAppPage(page ,id )
	return err
}

func (tc *DappController) GetAllDAppsByGroupId(groupId int64) ([]DappModel,error){
	dapp,err:= tc.service.getAllDAppsByGroupId(groupId )
	return dapp,err
}

func (tc *DappController) IsExistDApp(bundleId string) bool{
	kq:= tc.service.isExistDApp(bundleId)
	return kq
}
func (tc *DappController) IsExistSmartContract(bundleId string) bool{
	kq:= tc.service.isExistSmartContract(bundleId)
	return kq
}
// func (tc *DappController) InsertDapp(author string, name string, hash string, sign string, version string, logo string, pathStorage string,time int, 
// 	totalWallet int, totalTransaction int,size ,bundleId string,orientation string,urlRoot string,urlLoadingScreen string,urlLauchScreen string,isInjectJs int,urlWeb string,isLocal int,fullScreen int,
// 	statusBar string,groupId int,isShowInApp int,page int,position int,isInstalled int,abiData string,binData string,status int,typeT int) int {

// 	kq := tc.service.insertDapp(author, name , hash , sign , version , logo , pathStorage ,time , totalWallet , totalTransaction ,size ,bundleId ,orientation ,urlRoot ,urlLoadingScreen ,urlLauchScreen ,isInjectJs ,urlWeb,isLocal ,fullScreen ,statusBar ,groupId ,isShowInApp ,page ,position ,isInstalled ,abiData ,binData ,status ,typeT )
//    return kq
// }

func (tc *DappController) GetLastPage() int {

	kq := tc.service.getLastPage()
	return kq
}

func (tc *DappController) InsertDapp1(dapp *DappModel) int {

	kq := tc.service.insertDapp(dapp)
	return kq
}

// func (tc *DappController) UpdateDapp(author string, name string, hash string, sign string, version string, logo string, pathStorage string,time int, 
// 	totalWallet int, totalTransaction int,size ,orientation string,urlRoot string,urlLoadingScreen string,urlLauchScreen string,isInjectJs int,urlWeb string,isLocal int,fullScreen int,
// 	statusBar string,groupId int,isShowInApp int,page int,position int,isInstalled int,abiData string,binData string,status int,typeT int,bundleId string) int {

// 	kq := tc.service.updateDapp(author, name , hash , sign , version , logo , pathStorage ,time , totalWallet , totalTransaction ,size ,orientation ,urlRoot ,urlLoadingScreen ,urlLauchScreen ,isInjectJs ,urlWeb,isLocal ,fullScreen ,statusBar ,groupId ,isShowInApp ,page ,position ,isInstalled ,abiData ,binData ,status ,typeT ,bundleId)
//    return kq
// }
func (tc *DappController) UpdateDapp(dapp *DappModel) int {

	kq := tc.service.updateDapp(dapp)
	return kq
}
func (tc *DappController) UpdateDAppPosition(position int,id int64) error {

	err := tc.service.updateDAppPosition(position, id)
	return err
}

func (tc *DappController) DeleteDAppAndSmartContract(id int) int {

	kq := tc.service.deleteDAppAndSmartContract(id)
	return kq
}

func (tc *DappController) GetSmartContractByAddress(address string) map[string]interface{} {

	kq := tc.service.getSmartContractByAddress(address)
	return kq
}
func (tc *DappController) InsertSmartContract(dapp *DappModel) int {

	kq := tc.service.insertSmartContract(dapp)
	return kq
}
func (tc *DappController) UpdateSmartContractStatusByAddress(status int, address string) error{

	kq := tc.service.updateSmartContractStatusByAddress(status,address)
	return kq
}
func (tc *DappController) DeleteDAppTable()error {

	kq := tc.service.deleteDAppTable()
	return kq
}
