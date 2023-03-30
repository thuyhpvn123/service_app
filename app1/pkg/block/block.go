package block

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/holiman/uint256"
	e "gitlab.com/meta-node/meta-node/pkg/entry"
	"gitlab.com/meta-node/meta-node/pkg/logger"
	pb "gitlab.com/meta-node/meta-node/pkg/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type IBlock interface {
	// general
	Marshal() ([]byte, error)
	GetProto() protoreflect.ProtoMessage
	String() string

	// getter
	GetHash() common.Hash
	GetCount() *uint256.Int
	GetAccountStatesRoot() common.Hash
	GetLastEntryHash() common.Hash
	GetTimeStamp() int64
	IsVirtual() bool
	CalculateHash() (common.Hash, error)

	// setter
	SetTimeStamp(int64)
}

type Block struct {
	proto *pb.Block
}

func NewBlock(bProto *pb.Block) *Block {
	return &Block{
		proto: bProto,
	}
}

func Unmarshal(bytes []byte) (IBlock, error) {
	p := &pb.Block{}
	err := proto.Unmarshal(bytes, p)
	if err != nil {
		return nil, err
	}
	if proto.Equal(p, &pb.Block{}) {
		return nil, nil
	}
	return NewBlock(p), nil
}

func (b *Block) GetHash() common.Hash {
	return common.BytesToHash(b.proto.Hash)
}

func (b *Block) SetHash(hash common.Hash) {
	b.proto.Hash = hash.Bytes()
}

func (b *Block) CalculateHash() (common.Hash, error) {
	blockHashData := &pb.BlockHashData{
		Count:             b.proto.Count,
		LastEntryHash:     b.proto.Count,
		AccountStatesRoot: b.proto.AccountStatesRoot,
		ReceiptRoot:       b.proto.ReceiptRoot,
	}

	bData, err := proto.Marshal(blockHashData)
	if err != nil {
		return common.Hash{}, err
	}
	hash := crypto.Keccak256Hash(bData)
	return hash, nil
}

func (b *Block) GetCount() *uint256.Int {
	return uint256.NewInt(0).SetBytes(b.proto.Count)
}

func (b *Block) GetLastEntryHash() common.Hash {
	return common.BytesToHash(b.proto.LastEntryHash)
}

func (b *Block) GetAccountStatesRoot() common.Hash {
	return common.BytesToHash(b.proto.AccountStatesRoot)
}

func (b *Block) IsVirtual() bool {
	return b.proto.Hash != nil && (len(b.proto.AccountStatesRoot) == 0 && len(b.proto.ReceiptRoot) == 0)
}

func (b *Block) Marshal() ([]byte, error) {
	return proto.Marshal(b.proto)
}

func (b *Block) GetProto() protoreflect.ProtoMessage {
	return b.proto
}

func (b *Block) SetTimeStamp(timestamp int64) {
	b.proto.TimeStamp = timestamp
}

func CheckBlockHash(block IBlock) bool {
	correctHash, err := block.CalculateHash()
	if err != nil {
		logger.Warn(fmt.Sprintf("Error when calculate hash %v", err))
		return false
	}
	return block.GetHash() == correctHash
}

func NewVirtualBlock(
	lastBlock IBlock,
	entriesPerSlot uint64,
	hashesPerEntry uint64,
	entriesPerSecond uint64,
) IBlock {
	count := uint256.NewInt(0).AddUint64(lastBlock.GetCount(), 1)
	virtualBlockProto := &pb.Block{
		Count:             count.Bytes(),
		AccountStatesRoot: lastBlock.GetAccountStatesRoot().Bytes(),
		ReceiptRoot:       nil,
	}
	var lastEntry e.IEntry
	lastEntryHash := lastBlock.GetLastEntryHash()
	for i := uint64(0); i < entriesPerSlot; i++ {
		lastEntry = e.NewEntry(
			count,
			lastEntryHash,
			hashesPerEntry,
			nil,
		)
		lastEntryHash = lastEntry.GetHash()
	}
	virtualBlockProto.LastEntryHash = lastEntry.GetHash().Bytes()

	virtualBlock := NewBlock(virtualBlockProto)
	hash, err := virtualBlock.CalculateHash()
	if err != nil {
		logger.Warn(fmt.Sprintf("Error when hash virtual block %v", err))
	}
	virtualBlock.SetHash(hash)
	return virtualBlock
}

func (b *Block) String() string {
	str := fmt.Sprintf(`
	Hash: %v
	Count: %v
	Last entry hash: %v
	Account states root: %v
	Receipt root: %v
	Timestamp: %v
`,
		common.BytesToHash(b.proto.Hash),
		uint256.NewInt(0).SetBytes(b.proto.Count),
		common.BytesToHash(b.proto.LastEntryHash),
		common.BytesToHash(b.proto.AccountStatesRoot),
		common.BytesToHash(b.proto.ReceiptRoot),
		b.proto.TimeStamp,
	)
	return str
}

func (b *Block) GetTimeStamp() int64 {
	return b.proto.TimeStamp
}
