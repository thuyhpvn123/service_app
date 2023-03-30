package router

// import (
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"

// 	"github.com/ethereum/go-ethereum/common"
// 	"gitlab.com/meta-node/client/config"
// 	"gitlab.com/meta-node/client/network/messages"
// 	cn "gitlab.com/meta-node/meta-node/pkg/network"
// 	pb "gitlab.com/meta-node/meta-node/pkg/proto"
// 	"google.golang.org/protobuf/proto"
// )

// func readDataLastHash(hash string) *pb.Transaction {
// 	dat, err := os.ReadFile(fmt.Sprintf("./datas/%s", hash))
// 	if err != nil {
// 		log.Fatalf("Error when write data %v", err)
// 	}
// 	transaction := &pb.Transaction{}
// 	proto.Unmarshal(dat, transaction)
// 	return transaction
// }

// func sendGetAccountState1(
// 	version string,
// 	priKey []byte,
// 	pubKey []byte,
// 	parentConn *cn.Connection,
// 	address common.Address,
// 	sign []byte,
// ) {
// 	cn.SendBytes1(version, priKey,pubKey, parentConn, messages.GetAccountState, address.Bytes(),sign)
// 	fmt.Println("Sended Account Info")

// }
// func sendGetAccountState(
// 	config *config.Config,
// 	parentConn *cn.Connection,
// 	address common.Address,
// ) {
// 	cn.SendBytes(config, parentConn, messages.GetAccountState, address.Bytes())
// 	fmt.Println("Sended Account Info")

// }


// func enterAddress(message string) common.Address {
// 	address := strings.Replace(message, "\n", "", -1)
// 	address = strings.ToLower(message)
// 	return common.HexToAddress(address)
// }
