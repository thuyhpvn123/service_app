package block

import (
	"encoding/hex"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	cm "gitlab.com/meta-node/meta-node/pkg/common"
	pb "gitlab.com/meta-node/meta-node/pkg/proto"
	"gitlab.com/meta-node/meta-node/pkg/receipt"
	"gitlab.com/meta-node/meta-node/pkg/state"
	"google.golang.org/protobuf/proto"
)

type IFullBlock interface {
	Unmarshal(b []byte) error
	LoadFromProto(fbProto *pb.FullBlock)
	GetBlock() IBlock
	SetBlock(IBlock)
	GetValidatorSigns() map[cm.PublicKey]cm.Sign
	AddValidatorSign(cm.PublicKey, cm.Sign)
	SetValidatorSigns(map[cm.PublicKey]cm.Sign)
	GetAccountStateChanges() map[common.Address]state.IAccountState
	GetReceipts() receipt.IReceipts
	Marshal() ([]byte, error)
	SetAccountStates(state.IAccountStates)
	GetAccountStates() state.IAccountStates

	GetStakeChange() map[common.Address]*uint256.Int
	GetUnstakeChange() map[common.Address]*uint256.Int
	GetStakerConnectionAddresses() map[common.Address]string

	SetTimeStamp(int64)
}

type FullBlock struct {
	block    IBlock
	receipts receipt.IReceipts
	// account state data
	accountStateChanges map[common.Address]state.IAccountState
	accountStates       state.IAccountStates
	// stake data
	stakeChange               map[common.Address]*uint256.Int
	stakerConnectionAddresses map[common.Address]string
	unstakeChange             map[common.Address]*uint256.Int

	//
	validatorSigns map[cm.PublicKey]cm.Sign
}

func NewFullBlock(
	block IBlock,
	receipts receipt.IReceipts,
	validatorSigns map[cm.PublicKey]cm.Sign,
	accountStates state.IAccountStates,
	accountStateChanges map[common.Address]state.IAccountState,
	stakeChange map[common.Address]*uint256.Int,
	stakerConnectionAddresses map[common.Address]string,
	unstakeChange map[common.Address]*uint256.Int,
) IFullBlock {
	return &FullBlock{
		block:                     block,
		receipts:                  receipts,
		accountStates:             accountStates,
		accountStateChanges:       accountStateChanges,
		stakeChange:               stakeChange,
		stakerConnectionAddresses: stakerConnectionAddresses,
		unstakeChange:             unstakeChange,
	}
}

func (fb *FullBlock) Unmarshal(b []byte) error {
	fbProto := &pb.FullBlock{}
	err := proto.Unmarshal(b, fbProto)
	if err != nil {
		return err
	}
	fb.LoadFromProto(fbProto)
	return nil
}

func (fb *FullBlock) LoadFromProto(fbProto *pb.FullBlock) {
	block := NewBlock(fbProto.Block)

	receipts := receipt.NewReceipts()
	for _, v := range fbProto.Receipts {
		receipts.AddReceipt(receipt.ReceiptFromProto(v))
	}

	accountStateChanges := map[common.Address]state.IAccountState{}
	for _, v := range fbProto.AccountStateChanges {
		accountStateChanges[common.BytesToAddress(v.Address)] = state.AccountStateFromProto(v)
	}

	stakeChange := make(map[common.Address]*uint256.Int, len(fbProto.StakeChange))
	stakerConnectionAddresses := make(map[common.Address]string, len(fbProto.StakeChange))
	unstakeChange := make(map[common.Address]*uint256.Int, len(fbProto.UnstakeChange))

	for i, v := range fbProto.StakeChange {
		stakeChange[common.HexToAddress(i)] = uint256.NewInt(0).SetBytes(v)
	}

	for i, v := range fbProto.StakerConnectionAddresses {
		stakerConnectionAddresses[common.HexToAddress(i)] = v
	}

	for i, v := range fbProto.UnstakeChange {
		unstakeChange[common.HexToAddress(i)] = uint256.NewInt(0).SetBytes(v)
	}

	validatorSigns := make(map[cm.PublicKey]cm.Sign)
	for i, v := range fbProto.ValidatorSigns {
		validatorSigns[cm.PubkeyFromBytes(common.FromHex(i))] = cm.SignFromBytes(v)
	}
	*fb = FullBlock{
		block:               block,
		receipts:            receipts,
		accountStateChanges: accountStateChanges,
		validatorSigns:      validatorSigns,
		stakeChange:         stakeChange,
		unstakeChange:       unstakeChange,
	}
}

func (fb *FullBlock) GetBlock() IBlock {
	return fb.block
}

func (fb *FullBlock) SetBlock(b IBlock) {
	fb.block = b
}

func (fb *FullBlock) GetValidatorSigns() map[cm.PublicKey]cm.Sign {
	return fb.validatorSigns
}

func (fb *FullBlock) AddValidatorSign(pk cm.PublicKey, sign cm.Sign) {
	fb.validatorSigns[pk] = sign
}

func (fb *FullBlock) SetValidatorSigns(signs map[cm.PublicKey]cm.Sign) {
	fb.validatorSigns = signs
}

func (fb *FullBlock) GetAccountStateChanges() map[common.Address]state.IAccountState {
	return fb.accountStateChanges
}

func (fb *FullBlock) GetReceipts() receipt.IReceipts {
	return fb.receipts
}

func (fb *FullBlock) SetAccountStates(as state.IAccountStates) {
	fb.accountStates = as.Copy()
}

func (fb *FullBlock) GetAccountStates() state.IAccountStates {
	return fb.accountStates
}

func (fb *FullBlock) Marshal() ([]byte, error) {
	validatorSigns := make(map[string][]byte, len(fb.validatorSigns))
	for i, v := range fb.validatorSigns {
		validatorSigns[hex.EncodeToString(i.Bytes())] = v.Bytes()
	}
	var receipts []*pb.Receipt
	if fb.receipts != nil {
		receiptsMap := fb.receipts.GetReceiptsMap()
		receipts = make([]*pb.Receipt, len(receiptsMap))
		i := 0
		for _, v := range receiptsMap {
			receipts[i] = v.GetProto().(*pb.Receipt)
			i++
		}
	}

	accountStateChanges := make([]*pb.AccountState, len(fb.accountStateChanges))
	i := 0
	for _, v := range fb.accountStateChanges {
		accountStateChanges = append(accountStateChanges, v.GetProto().(*pb.AccountState))
		i++
	}
	// stake data
	stakeChange := make(map[string][]byte, len(fb.stakeChange))
	for i, v := range fb.stakeChange {
		stakeChange[hex.EncodeToString(i.Bytes())] = v.Bytes()
	}

	stakerConnectionAddresses := make(map[string]string, len(fb.stakerConnectionAddresses))
	for i, v := range fb.stakerConnectionAddresses {
		stakerConnectionAddresses[hex.EncodeToString(i.Bytes())] = v
	}

	unstakeChange := make(map[string][]byte, len(fb.unstakeChange))
	for i, v := range fb.unstakeChange {
		unstakeChange[hex.EncodeToString(i.Bytes())] = v.Bytes()
	}

	fbProto := &pb.FullBlock{
		Block:               fb.block.GetProto().(*pb.Block),
		Receipts:            receipts,
		AccountStateChanges: accountStateChanges,

		StakeChange:               stakeChange,
		StakerConnectionAddresses: stakerConnectionAddresses,
		UnstakeChange:             unstakeChange,

		ValidatorSigns: validatorSigns,
	}
	return proto.Marshal(fbProto)
}

func (fb *FullBlock) SetTimeStamp(timestamp int64) {
	fb.block.SetTimeStamp(timestamp)
}

func (fb *FullBlock) GetStakeChange() map[common.Address]*uint256.Int {
	return fb.stakeChange
}

func (fb *FullBlock) GetUnstakeChange() map[common.Address]*uint256.Int {
	return fb.unstakeChange
}

func (fb *FullBlock) GetStakerConnectionAddresses() map[common.Address]string {
	return fb.stakerConnectionAddresses
}
