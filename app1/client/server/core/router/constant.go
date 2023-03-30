package router

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	// . "github.com/ethereum/go-ethereum/accounts/abi"
)

type Account struct {
	Address string
	Private string
}

const (
	// HOST        = "61.28.238.235:3011"
	// STORAGEHOST = "61.28.238.235:3051"
	HOST        = "34.138.137.194:3011"
	STORAGEHOST = "35.196.167.172:3051"
)

var PORT int
var contracts = [...]Contract{
	{Name: "token0", Address: "A913eFF3367c1a48E307Df1D23aAfF44AA8646e8"},
}

func GetPORT() int {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Some error occured. Err: %s", err)
	}
	// log.Info("PORT: ", os.Getenv("PORT"))
	PORT, err = strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		panic(err)
	}
	return PORT
}

type Contract struct {
	Name    string
	Address string
}
