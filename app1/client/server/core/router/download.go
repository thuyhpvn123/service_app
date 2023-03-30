package router

import (
	"archive/zip"
	"reflect"
	// "container/heap"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	hdlDapp "gitlab.com/meta-node/client/handlers/dapp_handler"
	"gitlab.com/meta-node/client/models"
	"gitlab.com/meta-node/client/utils"

	"github.com/jmoiron/sqlx"
	log "github.com/sirupsen/logrus"
)

var versionJsonFile int 
type ResultDownload struct {
	// success bool`db:"success" json:"success"`
	Value interface{} `db:"value" json:"value"`
}

type App struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Hash    string `json:"hash"`
	UrlZip  string `json:"urlZip"`
	Id      string `json:"id"`
}
type DefaultJsonFile struct {
	Version string `json:"version"`
	Apps    []App  `json:"apps"`
}

type MetaApp struct {
	Version          string `json:"version"`
	UrlDefault       string `json:"urlDefault"`
	UrlLoadingScreen string `json:"urlLoadingScreen"`
	UrlZip           string `json:"urlZip"`
	Apps             []App  `json:"apps"`
	Hash			 string `json:"hash"`
}

// Download file
var (
	fileName string
	// fullURLFile string
)

func (caller *CallData) Download(bundleId string, fullURLFile string, hash string) {
	// bs, err := ioutil.ReadFile("./frontend/default_1.0.6.json")
	// if err != nil {
	// 	fmt.Println("Error: ", err)
	// 	return
	// }
	// var metaApp MetaApp
	// var fullURLFile string
	// var hash string
	// json.Unmarshal(bs, &metaApp)
	// for i, _ := range metaApp.Apps {
	// 	if metaApp.Apps[i].Name == bundleId {
	// 		hash = metaApp.Apps[i].Hash
	// 		fullURLFile = metaApp.Apps[i].UrlZip
	// 	}

	// }
	fmt.Println("fullURLFile:", fullURLFile)

	// Build fileName from fullPath
	fileURL, err := url.Parse(fullURLFile)
	if err != nil {
		log.Error(err)
	}
	fmt.Println("fileURL", fileURL)
	path := fileURL.Path
	fmt.Println("path:", path)
	segments := strings.Split(path, "/")
	fileName = segments[len(segments)-1]
	// fileName = "dapp"

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		log.Error(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fullURLFile)
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d", fileName, size)
	// hash:=
	// des := "public/"+hash+"/"+fileName
	des := "frontend/public/" + hash + "/" + bundleId
	// go caller.sentToClient("Dapp Path:", pathApp)
	
	fmt.Println("des:", des)
	err = caller.unzipSource(fileName, des)
	if err != nil {
		log.Error(err)
	}

}

//Unzip file
func (caller *CallData) unzipSource(source, destination string) error {
	fmt.Println("Unzipping file")
	// 1. Open the zip file
	reader, err := zip.OpenReader(source)
	if err != nil {
		log.Error(err)
	}
	defer reader.Close()

	// 2. Get the absolute destination path
	destination, err = filepath.Abs(destination)
	if err != nil {
		log.Error(err)
	}

	// 3. Iterate over zip files inside the archive and unzip each of them
	for _, f := range reader.File {
		err := unzipFile(f, destination)
		if err != nil {
			log.Error(err)
		}

	}
	// go caller.sentToClient("Dapp Path:", destination)

	fmt.Println("Unzip file completed")

	return nil
}

func unzipFile(f *zip.File, destination string) error {
	// 4. Check if file paths are not vulnerable to Zip Slip
	filePath := filepath.Join(destination, f.Name)
	if !strings.HasPrefix(filePath, filepath.Clean(destination)+string(os.PathSeparator)) {
		return fmt.Errorf("invalid file path: %s", filePath)
	}

	// 5. Create directory tree
	if f.FileInfo().IsDir() {
		if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
			log.Error(err)
		}
		return nil
	}

	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		log.Error(err)
	}

	// 6. Create a destination file for unzipped content
	destinationFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
	if err != nil {
		log.Error(err)
	}
	defer destinationFile.Close()

	// 7. Unzip the content of a file and copy it to the destination file
	zippedFile, err := f.Open()
	if err != nil {
		log.Error(err)
	}
	defer zippedFile.Close()

	if _, err := io.Copy(destinationFile, zippedFile); err != nil {
		log.Error(err)
	}
	return nil
}

var myClient = &http.Client{Timeout: 10 * time.Second}

func (caller *CallData) GetVersionMTN(url string) (string, string, string,string,string) {
	metaApp := new(MetaApp)
	r, err := myClient.Get(url)
	if err != nil {
		log.Error(err)
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(metaApp)

	return metaApp.Version, metaApp.UrlDefault, metaApp.UrlZip,metaApp.Hash,metaApp.UrlLoadingScreen
}
func (caller *CallData) GetVersionDefaultJson(url string) string {
	defaultJsonFile := new(DefaultJsonFile)
	r, err := myClient.Get(url)
	if err != nil {
		log.Error(err)
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(defaultJsonFile)

	return defaultJsonFile.Version
}
func GetDapp(url string, bundleId string) (string, string, string) {
	DefaultJsonFile := new(DefaultJsonFile)
	var ver string
	var urlZipDapp = ""
	var hash = ""

	r, err := myClient.Get(url)
	if err != nil {
		log.Error(err)
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(DefaultJsonFile)
	apps := DefaultJsonFile.Apps
	for i := range apps {
		if apps[i].Id == bundleId {
			ver = apps[i].Version
			urlZipDapp = apps[i].UrlZip
			hash = apps[i].Hash
			break
		}
	}
	fmt.Println("version Dapp", ver)
	return ver, urlZipDapp, hash
}
func GetAllBundleIdDapp(url string) []string {
	DefaultJsonFile := new(DefaultJsonFile)
	r, err := myClient.Get(url)
	if err != nil {
		log.Error(err)
	}
	defer r.Body.Close()
	err = json.NewDecoder(r.Body).Decode(DefaultJsonFile)
	apps := DefaultJsonFile.Apps
	var arr []string
	for i := range apps {
		if apps[i].UrlZip !=""  {
		arr=append(arr,apps[i].Id)
		}
	}
	return arr
}

func versionCheck(str string) int {
	item := strings.Split(str, ".")
	char1Int, _ := strconv.Atoi(item[0])
	char2Int, _ := strconv.Atoi(item[1])
	char3Int, _ := strconv.Atoi(item[2])
	verCheck := char1Int*1000000 + char2Int*10000 + char3Int
	return verCheck
}
func (caller *CallData) CheckVersion(bundleId string) {
	url := "https://metanode.co/json/metanode_dev_desktop.json"

	switch bundleId {
	case "Metanode":
		//get version of metanode app in file
		version:=readfile()
		//check version of metanode
		version1:=versionCheck(version)
		ver, urlDefaultMTN, urlZip ,hash,UrlLoadingScreen:= caller.GetVersionMTN(url)
		verCheck := versionCheck(ver)

		if version1 < verCheck {
			//create new file metanode.json to update
			createmetanodeJsonFile()
			header := models.Header{ Success:true, Data: UrlLoadingScreen}
			kq := utils.NewResultTransformer(header)
		
			go caller.sentToClient("desktop","download", false,kq)

			caller.Download("Metanode", urlZip,hash)
			bundleIdArr:=GetAllBundleIdDapp(urlDefaultMTN)
			fmt.Println("bundleIdArr:",bundleIdArr)
			for i :=range bundleIdArr{
				fmt.Println("urlDefaultMTN:",urlDefaultMTN)
	
				_, urlZipDapp, hash:=GetDapp(urlDefaultMTN,bundleIdArr[i])
				fmt.Println("urlZipDapp:",urlZipDapp)
	
				caller.Download(bundleIdArr[i], urlZipDapp, hash)
	
			}
	

		}
		des := "/public/" + hash + "/" + bundleId
		kq := utils.NewResultTransformer(des)

		go caller.sentToClient("desktop","metanode-path",false, kq)
		//download all default apps in metanode urlDefault file  json
		
	default:
		// get version of metanode app in file
		version:=readfile()
		//check version of metanode
		version1:=versionCheck(version)
		//check Metanode version first
		ver, urlDefaultMTN, urlZip ,hashMTN,_:= caller.GetVersionMTN(url)
		verCheck := versionCheck(ver)

		
		if version1 < verCheck {
			//create new file metanode.json to update
			createmetanodeJsonFile()

			caller.Download("Metanode", urlZip,hashMTN)
		}

		//check DefaulJsonFile
		verFileJson := caller.GetVersionDefaultJson(urlDefaultMTN)
		fmt.Println("verFileJson:", verFileJson)
		verCheck1 := versionCheck(verFileJson)
		versionJsonFile=19

		fmt.Println("versionJsonFile:", versionJsonFile)
		fmt.Println("verCheck1:", verCheck1)
		if versionJsonFile < verCheck1 {
			var lastVer int
			verDapp, urlZipDapp, hash:=GetDapp(urlDefaultMTN,bundleId)
			db, err := sqlx.Connect("sqlite3","./database/doc_2022-12-26_09-16-03.meta_findsdk.db")
			if err != nil {
				log.Fatal(err)
			}
	
			dappCtrl := hdlDapp.NewDappController(db)
	
			 result:=dappCtrl.GetDappByBundleId(bundleId)
			 
			 var v interface{} = result.Data.(*models.Header).Data

			 var out []interface{}
			 rv := reflect.ValueOf(v)
			 if rv.Kind() == reflect.Slice {
				 for i := 0; i < rv.Len(); i++ {
					 out = append(out, rv.Index(i).Interface())
				 }
			 }
			
			 if len(out)==0{
				lastVer=0
			 }else {
				data:=result.Data.(*models.Header).Data.([]hdlDapp.DappModel)[0].Version
				lastVer = versionCheck(data)
			 }
			verCheck2 := versionCheck(verDapp)
			fmt.Println("lastVer:",lastVer)
			fmt.Println("verCheck2:",verCheck2)

			if lastVer < verCheck2 {
				caller.Download(bundleId, urlZipDapp, hash)
			}
			versionJsonFile=verCheck1
		}
		// hashPath=hash
		// go caller.sentToClient("MTN Path:", des)
		_, _, hashDapp:=GetDapp(urlDefaultMTN,bundleId)

		des := "/public/" + hashDapp + "/" + bundleId
		header := models.Header{ Success:true, Data: des}
		kq := utils.NewResultTransformer(header)
	
		go caller.sentToClient("desktop","dapp-path",false ,kq)

	}		
}
func createmetanodeJsonFile()*os.File{
	url := "https://metanode.co/json/metanode_dev_desktop.json"

	// Create blank file
	des := "frontend" 
    path := filepath.Join(des, "metanode.json")
    fmt.Println(path)
    file, err := os.Create(path)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("File created successfully")
    defer file.Close()

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(url)
	if err != nil {
		log.Error(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Metanode.json created")
	return file

}
func  readfile()string {
	bs, err := ioutil.ReadFile("./frontend/metanode.json")
if err != nil {
	fmt.Println("Error: ", err)
	return "Error when read file metanode.json"
}
var mtApp App
json.Unmarshal(bs, &mtApp)
fmt.Println(mtApp.Version)
return mtApp.Version
}
