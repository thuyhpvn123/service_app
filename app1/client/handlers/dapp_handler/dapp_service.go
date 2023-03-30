package dapphandler

import (
	// "fmt"

	"encoding/json"
	"fmt"
	// "log"

	"github.com/jmoiron/sqlx"

	"gitlab.com/meta-node/client/models"
	"gitlab.com/meta-node/client/utils"
	"gitlab.com/meta-node/meta-node/pkg/logger"
)

// DappService struct
type DappService struct {
	db *sqlx.DB
}

// newDappService return new DappService object.
func newDappService(db *sqlx.DB) *DappService {
	return &DappService{db}
}


// getDappPagination return Dapps by Pagination.
func (ts *DappService) getDappByName(name string) []DappModel {
	dapp := []DappModel{}

	err := ts.db.Select(&dapp, "SELECT * FROM decentralizedApplicationTB WHERE name=? ", name)
	if err != nil {
		panic(err)
	}
	

	return dapp
}
func (ts *DappService) getDappByBundleId(bundleId string)*utils.ResultTransformer  {
	dapp := []DappModel{}

	err := ts.db.Select(&dapp, "SELECT * FROM decentralizedApplicationTB WHERE bundleId=? ORDER BY id DESC", bundleId)
	if err != nil {
		panic(err)
	}
	header := models.Header{ Success:true, Data: dapp}
	result := utils.NewResultTransformer(header)

	return result
}
func (ts *DappService) getAllDAppsForBrowser()([]DappModel,error){
	dapp := []DappModel{}

	err := ts.db.Select(&dapp, "SELECT * FROM decentralizedApplicationTB WHERE type = 1")
	if err != nil {
		panic(err)
	}
	// header := models.Header{ Success:true, Data: dapp}
	// result := utils.NewResultTransformer(header)

	return dapp,err
}

func (ts *DappService) getAllDappNoGroup() *utils.ResultTransformer {
	dapp := []DappModel{}
	page:=[]Body{}
	err := ts.db.Select(&dapp, "SELECT * FROM decentralizedApplicationTB WHERE type = 1 AND isInstalled = 1 AND isShowInApp = 1 AND groupId = 0 ORDER BY page, position",)
	
	if err != nil {
		panic(err)
	}
	chunkSize:=5
	// header := models.Header{Success:true,  Data: dapp}
	var chunks [][]DappModel
	if len(dapp)>chunkSize{
		for i := 0; i < len(dapp); i += chunkSize {
			end := i + chunkSize
	
			// necessary check to avoid slicing beyond
			// slice capacity
			if end > len(dapp) {
				end = len(dapp)
			}
	
			chunks = append(chunks, dapp[i:end])
		}
	
		for i:=0;i <=len(dapp)/5; i++{
			// limit:=5
			var item [][]DappModel
			// for i:=0;i<len(dapp);i+=limit{
			// 	batch := app[i:min(i+limit,len(dapp))]
			// }
			item[i]=chunks[i]
			// body:= Body{}
			body:= Body{
				Page: i,
				Name:"page"+string(i),
				DApps: item[i],
			}
			page=append(page,body)
	
		}

	}else{
		body:= Body{
			Page: 0,
			Name:"page 0",
			DApps: dapp,
		}
		page=append(page,body)
}

	header := models.Header{Success:true,  Data: page}

	result := utils.NewResultTransformer(header)

	return result
}
// GetMaxPositionWallet return maxPosition when create-wallet 
func (ts *DappService) getMaxPositionWallet() int {

	var count int

	err := ts.db.Get(&count, "SELECT position FROM decentralizedApplicationTB ORDER BY position DESC LIMIT 1")
	if err != nil {
		logger.Error(fmt.Sprintf("error when getMaxPositionWallet %", err))
		// panic(fmt.Sprintf("error when getMaxPositionWallet %v", err))
		return 0
	}


	return count
}

func (ts *DappService) getAllDAppsByGroupIdAndPage(groupId int64,page int)[]DappModel  {
	dapps := []DappModel{}

	err := ts.db.Select(&dapps, "SELECT * FROM decentralizedApplicationTB WHERE groupId=? AND page=? AND ((type = 1 AND isShowInApp = 1) OR (type = 2 AND status = 1) OR type = 3 OR type = 4) ORDER BY position", groupId,page)
	if err != nil {
		panic(err)
	}
	return dapps
}
func (ts *DappService) getAllDAppsByGroupId(groupId int64)([]DappModel,error)  {
	dapps := []DappModel{}

	err := ts.db.Select(&dapps, "SELECT * FROM decentralizedApplicationTB WHERE groupId=? AND ((type = 1 AND isShowInApp = 1) OR (type = 2 AND status = 1) OR type = 3 OR type = 4) ORDER BY page, position", groupId)
	if err != nil {
		panic(err)
	}
	return dapps,err
}
func (ts *DappService) getAllDApps()[]DappModel  {
	dapps := []DappModel{}

	err := ts.db.Select(&dapps, "SELECT * FROM decentralizedApplicationTB WHERE ((type = 1 AND isShowInApp = 1) OR type = 4) ORDER BY page, position")
	if err != nil {
		panic(err)
	}
	return dapps
}

func (ts *DappService) updateDAppPosition(position int,id int64)error{
	_, err := ts.db.Exec("UPDATE decentralizedApplicationTB SET position=? WHERE id=?",position,id)
		
	if err != nil {
		logger.Error(fmt.Sprintf("error when updateGroupDAppPosition %", err))
		panic(fmt.Sprintf("error when updateDAppPosition %v", err))
		return err
	}
	fmt.Println("updateDAppPosition in database successed")
	return nil
}
func (ts *DappService) updateDAppPageAndPosition(position int,id int64)error{
	_, err := ts.db.Exec("UPDATE decentralizedApplicationTB SET position=? WHERE id=?",position,id)
		
	if err != nil {
		logger.Error(fmt.Sprintf("error when updateGroupDAppPosition %", err))
		panic(fmt.Sprintf("error when updateDAppPosition %v", err))
		return err
	}
	fmt.Println("updateDAppPosition in database successed")
	return nil
}
func (ts *DappService) updateDAppPage(page int,id int64)error{
	_, err := ts.db.Exec("UPDATE decentralizedApplicationTB SET page=? WHERE id=?",page,id)
		
	if err != nil {
		logger.Error(fmt.Sprintf("error when updateDAppPage %", err))
		panic(fmt.Sprintf("error when updateDAppPage %v", err))
		return err
	}
	fmt.Println("updateDAppPage in database successed")
	return nil
}

func (ts *DappService) isExistDApp(bundleId string) bool  {
	dapps := []DappModel{}

	err := ts.db.Select(&dapps,"SELECT EXISTS(SELECT id FROM decentralizedApplicationTB WHERE bundleId =? AND ((type = 1 AND isShowInApp = 1) OR type = 3 OR type = 4))",bundleId)
	if err != nil {
		panic(err)
	}
	if len(dapps)==0{
		return true
	}else{
		return false
	}

}
func (ts *DappService) isExistSmartContract(bundleId string) bool  {
	dapps := []DappModel{}

	err := ts.db.Select(&dapps,"SELECT EXISTS(SELECT id FROM decentralizedApplicationTB WHERE bundleId =? AND type = 2)",bundleId)
	if err != nil {
		panic(err)
	}
	if len(dapps)==0{
		return true
	}else{
		return false
	}

}
// func (ts *DappService) insertDapp(author string, name string, hash string, sign string, version string, logo string, pathStorage string,time int, 
// 	totalWallet int, totalTransaction int,size string,bundleId string,orientation string,urlRoot string,urlLoadingScreen string,urlLauchScreen string,isInjectJs int,urlWeb string,isLocal int,fullScreen int,
// 	statusBar string,groupId int,isShowInApp int,page int,position int,isInstalled int,abiData string,binData string,status int,typeT int) int {
func (ts *DappService) insertDapp(dapp *DappModel)int{
_, err := ts.db.NamedExec("INSERT INTO decentralizedApplicationTB(author, name , hash 	, sign , version ,image,  pathStorage ,time , totalWallet , totalTransaction ,size ,bundleId ,orientation ,urlWeb,isLocal ,fullScreen ,statusBar ,groupId ,isShowInApp ,page ,position ,positionObj,isInstalled ,abiData ,binData ,status ,typeT ) values (:author, :name , :hash 	, :sign , :version  , :pathStorage ,:time , :totalWallet , :totalTransaction ,:size ,:bundleId ,:orientation ,:urlWeb,:isLocal ,:fullScreen ,:statusBar ,:groupId ,:isShowInApp ,:page ,:position ,:isInstalled ,:abiData ,:binData ,:status ,:type)",
map[string]interface{}{
		"name":dapp.Name ,
		"author":dapp.Author  ,
		"hash":dapp.Hash ,
		"sign":dapp.Sign,
		"version":dapp.Version,
		"image":dapp.Image,
		"pathStorage":dapp.PathStorage,
		"time":dapp.Time ,
		"totalWallet":dapp.TotalWallet ,
		"totalTransaction":dapp.TotalTransaction,
		"size":dapp.Size,
		"bundleId":dapp.BundleId ,
		"orientation":dapp.Orientation,
		"urlWeb":dapp.UrlWeb,
		"isLocal":dapp.IsLocal,
		"fullScreen":dapp.FullScreen,
		"statusBar":dapp.StatusBar,
		"groupId":dapp.GroupId,
		"isShowInApp":dapp.IsShowInApp,
		"page":dapp.Page,
		"position":dapp.Position,
		"positionObj":dapp.PositionObj,
		"isInstalled":dapp.IsInstalled,
		"abiData":dapp.AbiData,
		"binData":dapp.BinData,
		"status":dapp.Status,
		"type":dapp.Type,
	
	})
	
	if err != nil {
		logger.Error(fmt.Sprintf("error when insertDapp %", err))
		panic(fmt.Sprintf("error when insertDapp %v", err))
		return -1
	}

	fmt.Println("Insert Dapp in database successed")
	return 1
}


// func (ts *DappService) updateDapp(author string, name string, hash string, sign string, version string, logo string, pathStorage string,time int, 
// 	totalWallet int, totalTransaction int,size string,orientation string,urlRoot string,urlLoadingScreen string,urlLauchScreen string,isInjectJs int,urlWeb string,isLocal int,fullScreen int,
// 	statusBar string,groupId int,isShowInApp int,page int,position int,isInstalled int,abiData string,binData string,status int,typeT int,bundleId string) int {
func (ts *DappService) updateDapp(dapp *DappModel)int{
_, err := ts.db.Exec("UPDATE decentralizedApplicationTB SET author =?, name =?, hash=? , sign=? , version =?, image=?, pathStorage=? ,time =?, totalWallet =?, totalTransaction =?,size=? ,bundleId=? ,orientation=? ,urlWeb=?,isLocal =?,fullScreen=? ,statusBar=? ,groupId =?,isShowInApp =?,page =?,position =?,positionObj =?,isInstalled =?,abiData =?,binData =?,status =?,typeT=? WHERE bundleId =?",dapp.Author, dapp.Name , dapp.Hash , dapp.Sign , dapp.Version , dapp.Image, dapp.PathStorage ,dapp.Time , dapp.TotalWallet , dapp.TotalTransaction ,dapp.Size ,dapp.BundleId ,dapp.Orientation  ,dapp.UrlWeb,dapp.IsLocal ,dapp.FullScreen ,dapp.StatusBar ,dapp.GroupId ,dapp.IsShowInApp ,dapp.Page ,dapp.Position ,dapp.PositionObj,dapp.IsInstalled ,dapp.AbiData ,dapp.BinData ,dapp.Status ,dapp.Type, dapp.BundleId)
	
	if err != nil {
		logger.Error(fmt.Sprintf("error when updateDapp %", err))
		panic(fmt.Sprintf("error when updateDapp %v", err))
		return 0
	}
	fmt.Println("Update Dapp in database successed")
	return 1
}
func (ts *DappService) deleteDAppAndSmartContract(id int)int{
	_, err := ts.db.Exec("DELETE FROM decentralizedApplicationTB WHERE id=?",id)
		
	if err != nil {
		logger.Error(fmt.Sprintf("error when deleteDapp %", err))
		panic(fmt.Sprintf("error when deleteDapp %v", err))
		return 0
	}
	fmt.Println("DeleteDapp Dapp in database successed")
	return 1
}
	
// updateSQL := fmt.Sprintf("UPDATE decentralizedApplicationTB SET name = '%s', author = '%s', hash = '%s', sign = '%s', version = '%s', logo = '%s', pathStorage = '%s', time = %d, totalWallet = %d, totalTransaction = %d, size = %s, orientation = '%s', urlRoot = '%s', urlLoadingScreen = '%s', urlLauchScreen = '%s', groupId = %d WHERE bundleId = '%s'",
//     data["name"], data["author"], data["hash"], data["sign"], data["version"], data["logo"], data["pathStorage"], data["time"].(int), data["totalWallet"].(int), data["totalTransaction"].(int), data["size"], data["orientation"], data["urlRoot"], data["urlLoadingScreen"], data["urlLauchScreen"], data["groupId"].(int), data["id"])

//smart-contract in the same table
	
func (ts *DappService) getSmartContractByAddress(address string) map[string]interface{}  {

	sc := DappModelShort{}

	err := ts.db.Get(&sc, "SELECT * FROM decentralizedApplicationTB WHERE bundleId=? AND type = 2",address)
	if err != nil {
		panic(err)
	}
	var dappMap  map[string]interface{}
    bDapp, _ := json.Marshal(sc)
    json.Unmarshal(bDapp, &dappMap)

	return dappMap
}
func (ts *DappService) getLastPage() int  {

	dapp := DappModel{}

	err := ts.db.Get(&dapp, "SELECT page FROM decentralizedApplicationTB ORDER BY page DESC LIMIT 1")
	if err != nil {
		logger.Error(fmt.Sprintf("error when getLastPage%", err))
		// panic(fmt.Sprintf("error when getLastPage %v", err))
		return 0
	// panic(err)
	}

	return dapp.Page
}
func (ts *DappService) insertSmartContract(dapp *DappModel)int{
	_, err := ts.db.NamedExec("INSERT INTO decentralizedApplicationTB(author, name , hash , sign , version ,image,  pathStorage ,time , totalWallet , totalTransaction ,size ,bundleId ,orientation ,urlWeb,isLocal ,fullScreen ,statusBar ,groupId ,isShowInApp ,page ,position ,positionObj,isInstalled ,abiData ,binData ,status ,type ) values (:author, :name , :hash 	, :sign , :version  ,:image, :pathStorage ,:time , :totalWallet , :totalTransaction ,:size ,:bundleId ,:orientation ,:urlWeb,:isLocal ,:fullScreen ,:statusBar ,:groupId ,:isShowInApp ,:page ,:position ,:positionObj,:isInstalled ,:abiData ,:binData ,:status ,:type)",														
	map[string]interface{}{
			"name":dapp.Name ,
			"author":""  ,
			"hash":"" ,
			"sign":"",
			"version":"",
			"image":dapp.Image,
			"pathStorage":"",
			"time":"" ,
			"totalWallet":"" ,
			"totalTransaction":"",
			"size":"",
			"bundleId":dapp.BundleId ,
			"orientation":"",
			"urlWeb":"",
			"isLocal":"",
			"fullScreen":"",
			"statusBar":"",
			"groupId":"",
			"isShowInApp":"",
			"page":dapp.Page,
			"position":dapp.Position,
			"positionObj":dapp.PositionObj,
			"isInstalled":dapp.IsInstalled,
			"abiData":dapp.AbiData,
			"binData":dapp.BinData,
			"status":dapp.Status,
			"type":dapp.Type,
		
		})

		if err != nil {
			logger.Error(fmt.Sprintf("error when insertSmartContract %", err))
			panic(fmt.Sprintf("error when insertSmartContract %v", err))
			return -1
		}
	
		fmt.Println("Insert smart contract in database successed")
		return 1
	}
//updateSmartContractStatusByAddress update status of smart contract type 2  by address
func (ts *DappService) updateSmartContractStatusByAddress(status int, address string)error{
	_, err := ts.db.Exec("UPDATE decentralizedApplicationTB SET status =? WHERE type = 2 AND bundleId =?",status,address)
		
		if err != nil {
			logger.Error(fmt.Sprintf("error when updateSmartContractStatusByAddress %", err))
			panic(fmt.Sprintf("error when updateSmartContractStatusByAddress %v", err))
			return err
		}
		fmt.Println("Update SmartContract Status in database successed")
		return nil
}
func (ts *DappService) deleteDAppTable()error{
	_, err := ts.db.Exec("DELETE FROM decentralizedApplicationTB")
		
	if err != nil {
		logger.Error(fmt.Sprintf("error when deleteDAppTable %", err))
		panic(fmt.Sprintf("error when deleteDAppTable %v", err))
		return err
	}
	fmt.Println("deleteDAppTable in database successed")
	return nil
}
