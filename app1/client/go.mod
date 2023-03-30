module gitlab.com/meta-node/client

go 1.18

// replace gitlab.com/meta-node/core => ../core

replace gitlab.com/meta-node/meta-node/pkg => ../pkg

require (
	github.com/denisbrodbeck/machineid v1.0.1
	github.com/ethereum/go-ethereum v1.11.1
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.0
	github.com/holiman/uint256 v1.2.1
	github.com/jmoiron/sqlx v1.3.5
	github.com/joho/godotenv v1.4.0
	github.com/mattn/go-sqlite3 v1.14.16
	github.com/sirupsen/logrus v1.9.0
	github.com/supranational/blst v0.3.8-0.20220526154634-513d2456b344
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
	gitlab.com/meta-node/core v0.0.0-00010101000000-000000000000
	gitlab.com/meta-node/meta-node/pkg v0.0.0-00010101000000-000000000000
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	golang.org/x/crypto v0.1.0 // indirect
	golang.org/x/sys v0.5.0 // indirect
)
