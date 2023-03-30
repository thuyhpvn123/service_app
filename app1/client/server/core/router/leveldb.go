package router

import (
	// "fmt"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)
// type Result struct {
// 	Success bool      `json:"success"`
// 	Data  interface{} `json:"data"`
// }

// }
// func main(){
// 	db, err := leveldb.OpenFile("./db", nil)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()
// 	call :=map[string]interface{}{
// 		"key":"0x3E65Ce8502fb2a85F061B3fD8256D61Cc8C9d440",
// 		"data":"thuy",
// 	}
// 	// call1 :=map[string]interface{}{
// 	// 	"key":"0x3E65Ce8502fb2a85F061B3fD8256D61Cc8C9d441",
// 	// }

// 	result:=writeValueStorage(call,db)
// 	fmt.Println("write data",result)
// 	result1:=readValueStorage(call,db)
// 	fmt.Println("read data",result1)
// 	result2:=deleteKeyStorage(call,db)
// 	fmt.Println("delete data",result2)
// 	result3:=readValueStorage(call,db)
// 	fmt.Println("read data",result3)

// }

func WriteValueStorage(call map[string]interface{},db *leveldb.DB) map[string]interface{} {
	key, ok := call["key"].(string)
	data, ok := call["data"].(string)
	if !ok || key == "" || data == "" {
	
	return map[string]interface{}{
	"error": "EWVS-001",
	"message": "Key and data is required",
	}
	}

	// securyDb.Write(key, data)
	err := db.Put([]byte(key), []byte(data), nil)
	if err!=nil{
		log.Fatal()
	}
	return map[string]interface{}{
	"success": true,
	}
}
 func ReadValueStorage(call map[string]interface{},db *leveldb.DB) map[string]interface{} {
	key, ok := call["key"].(string)
	if !ok || key == "" {
		return map[string]interface{}{
			"success": false,
			"msg": "Key not found",
		}
	} else {
		value,err := db.Get([]byte(key), nil)
		if err != nil {
			return map[string]interface{}{
				"success": false,
				"msg": "",
			}
		
		} else {
			return map[string]interface{}{
				"value": value,
				"success": true,
				"msg": "",
			}
			
		}
	}
}
func DeleteKeyStorage(call map[string]interface{},db *leveldb.DB) map[string]interface{} {
	key, ok := call["key"].(string)
	if !ok || key == "" {
	return map[string]interface{}{
	"error": "EDKS-001",
	"message": "Key is required",
	}
	}
	// securyDb.Delete(key)
	err := db.Delete([]byte(key), nil)
	if err!=nil{
		log.Fatal()
	}
	return map[string]interface{}{
	"success": true,
	}
}