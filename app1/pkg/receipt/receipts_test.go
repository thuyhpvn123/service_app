package receipt

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
	pb "gitlab.com/meta-node/meta-node/pkg/proto"
)

var (
	testAddress1    = common.HexToAddress("067d82035bacafcf39258296bcbaa96ddf8678f6")
	testAddress2    = common.HexToAddress("d381f9537af2a152aa2fb2d721fe9b285f1e87f8")
	testHash        = common.HexToHash("1111111111111111111111111111111111111111111111111111111111111111")
	testAmount      = uint256.NewInt(10)
	testStatus      = pb.RECEIPT_STATUS_RETURNED
	testReturnValue = []byte{1, 2, 3, 4}
	testException   = pb.EXCEPTION_NONE
)

func TestAddReceipt(t *testing.T) {
	receipts := NewReceipts()
	receipt := NewReceipt(
		testHash,
		testAddress1,
		testAddress2,
		testAmount,
		testStatus,
		testReturnValue,
		testException,
	)
	err := receipts.AddReceipt(receipt)
	assert.Nil(t, err)
	rcpts := receipts.GetReceiptsMap()
	assert.Equal(t, rcpts[testHash], receipt)
	b, _ := receipt.Marshal()
	trieB, err := receipts.(*Receipts).trie.Get(testHash.Bytes())
	assert.Nil(t, err)
	assert.Equal(t, b, trieB)
}

func TestGetReceiptsRoot(t *testing.T) {
	receipts := NewReceipts()
	receipt := NewReceipt(
		testHash,
		testAddress1,
		testAddress2,
		testAmount,
		testStatus,
		testReturnValue,
		testException,
	)
	receipts.AddReceipt(receipt)
	receiptRoot, err := receipts.GetReceiptsRoot()
	assert.Nil(t, err)
	assert.Equal(t, common.HexToHash("0xd5494d73399638c00165a4873fa02cae9964f9ea82080735daa377ac77f38a34"), receiptRoot)
}
