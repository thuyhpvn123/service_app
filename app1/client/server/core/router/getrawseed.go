package router

import (
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"gitlab.com/meta-node/client/models"
	"gitlab.com/meta-node/client/utils"
	blst "gitlab.com/meta-node/meta-node/pkg/bls/blst/bindings/go"


	// blst "gitlab.com/meta-node/core/crypto/blst/bindings/go"
)

func randBLSTSecretKey(t [32]byte) *blstSecretKey {
	secretKey := blst.KeyGen(t[:])
	return secretKey
}

func blsGenPriKey(rawSeedHash []byte) ([]byte, []byte) {

	msgPriHex := rawSeedHash
	var msgPri [32]byte
	_ = copy(msgPri[:], msgPriHex)
	sec := randBLSTSecretKey(msgPri)
	pub := new(blstPublicKey).From(sec)
	return sec.Serialize(), pub.Compress()
}
func randomStringUniqueFromListWithLength(data []string, length int) []string {
	valueList := make([]string, 0, length)
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(valueList) < length {
		value := randomStringNotExistInArray(random, data, valueList)
		valueList = append(valueList, value)
	}
	return valueList
}

func randomStringNotExistInArray(random *rand.Rand, listData, listChecked []string) string {
	randomValue := listData[random.Intn(len(listData))]

	valueList := make([]interface{}, len(listChecked))
	for i, v := range listChecked {
		valueList[i] = v
	}
	indexExist := indexOf(valueList, randomValue)
	if indexExist == -1 {
		return randomValue
	}
	return randomStringNotExistInArray(random, listData, listChecked)
}

func indexOf(slice []interface{}, elem interface{}) int {
	for i, v := range slice {
		if v == elem {
			return i
		}
	}
	return -1
}

func (caller *CallData) GetRawSeed(call map[string]interface{}) *Result {
	// if !ok {
	//     // Return an error if the language field is not a string
	//     result.Success(map[string]interface{}{
	//         "success": false,
	//         "message": "language field must be a string",
	//     })
	//     return
	// }
	value := call["value"].(map[string]interface{})
	language := value["language"].(string)
	file := filepath.Join("./public", language+".txt")
	if _, err := os.Stat(file); os.IsNotExist(err) {
		getDownloadThread(language)
	}

	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	b := string(bytes)
	var listString []string
	//tạm bỏ qua indonesia
	// if language != "indonesia" {
	//     listString = strings.Split(b, "\n")
	// } else {
	//     listString = []string{b}
	// }
	listString = strings.Split(b, "\n")

	arrString := randomStringUniqueFromListWithLength(listString, 24)
	rawSeed = arrString
	fmt.Println("rawSeed:", rawSeed)
	rawSeedArray := make([]map[string]interface{}, len(arrString))
	for i, value := range arrString {
		rawSeedArray[i] = map[string]interface{}{
			"name":  value,
			"index": i,
		}
	}
	result := &Result{
		Success: true,
		Data:    rawSeedArray,
	}
	// fmt.Println("rawSeedArray:", result)
	header := models.Header{Success: true, Data: rawSeedArray}
	kq := utils.NewResultTransformer(header)

	go caller.sentToClient("desktop", "get-raw-seed", false, kq)

	return result
}
func getDownloadThread(language string) {
	var urlString string
	switch language {
	case "vietnamese":
		urlString = "https://gitlab.com/QuangFiIt/upload_app/-/raw/main/json/vietnamese.txt"
	case "english":
		urlString = "https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/english.txt"
	case "japanese":
		urlString = "https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/japanese.txt"
	case "korean":
		urlString = "https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/korean.txt"
	case "spainish":
		urlString = "https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/spanish.txt"
	case "chinese-simplified":
		urlString = "https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/chinese_simplified.txt"
	case "chinese-tranditional":
		urlString = "https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/chinese_traditional.txt"
	case "french":
		urlString = "https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/french.txt"
	case "italia":
		urlString = "https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/italian.txt"
	case "czech":
		urlString = "https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/czech.txt"
	case "portuguese":
		urlString = "https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/portuguese.txt"
	// case "indonesia":
	// 	urlString = "https://gitlab.com/QuangFiIt/upload_app/-/raw/main/json/indonesia.json?inline=false"
	default:
		urlString = "https://raw.githubusercontent.com/bitcoin/bips/master/bip-0039/" + language + ".txt"
	}
	// Put content on file
	url, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	resp, err := http.Get(url.String())
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Create blank file
	filePath := filepath.Join("./public", language+".txt")
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	size, err := io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Downloaded a file %s with size %d \n", language+".txt", size)

}
