package pos

import (
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	pb "gitlab.com/meta-node/meta-node/pkg/proto"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	errInvalidUnstakeAmount = errors.New("invalid unstake amount")
)

type StakePool struct {
	stakedAmount              map[common.Address]*uint256.Int
	stakerConnectionAddresses map[common.Address]string
	totalStakedAmount         *uint256.Int
}

func NewStakePool() *StakePool {
	return &StakePool{
		stakedAmount:              make(map[common.Address]*uint256.Int),
		stakerConnectionAddresses: make(map[common.Address]string),
		totalStakedAmount:         uint256.NewInt(0),
	}
}

func NewStakePoolFromProto(protoSp *pb.PosStakePool) *StakePool {
	totalStakedAmount := uint256.NewInt(0)
	stakedAmount := make(map[common.Address]*uint256.Int, len(protoSp.StakedAmount))
	stakerConnectionAddresses := make(map[common.Address]string, len(protoSp.StakerConnectionAddresses))
	for a, v := range protoSp.StakedAmount {
		stakedAmount[common.HexToAddress(a)] = uint256.NewInt(0).SetBytes(v)
		totalStakedAmount = totalStakedAmount.Add(totalStakedAmount, stakedAmount[common.HexToAddress(a)])
	}
	for a, v := range protoSp.StakerConnectionAddresses {
		stakerConnectionAddresses[common.HexToAddress(a)] = v
	}
	return &StakePool{
		stakedAmount:              stakedAmount,
		totalStakedAmount:         totalStakedAmount,
		stakerConnectionAddresses: stakerConnectionAddresses,
	}
}

func Unmarshal(b []byte) (*StakePool, error) {
	protoSp := &pb.PosStakePool{}
	err := proto.Unmarshal(b, protoSp)
	if err != nil {
		return nil, err
	}
	return NewStakePoolFromProto(protoSp), nil
}

func (p *StakePool) Stake(address common.Address, amount *uint256.Int) {
	if p.stakedAmount[address] == nil {
		p.stakedAmount[address] = uint256.NewInt(0)
	}
	p.stakedAmount[address] = uint256.NewInt(0).Add(p.stakedAmount[address], amount)
	p.totalStakedAmount = p.totalStakedAmount.Add(p.totalStakedAmount, amount)
}

func (p *StakePool) UpdateConnectionAddress(address common.Address, connectionAddress string) {
	p.stakerConnectionAddresses[address] = connectionAddress
}

func (p *StakePool) Unstake(address common.Address, amount *uint256.Int) error {
	if amount.Gt(p.stakedAmount[address]) {
		return errInvalidUnstakeAmount
	}
	p.stakedAmount[address] = p.stakedAmount[address].Sub(p.stakedAmount[address], amount)
	p.totalStakedAmount = p.totalStakedAmount.Sub(p.totalStakedAmount, amount)
	if p.stakedAmount[address].Eq(uint256.NewInt(0)) {
		delete(p.stakedAmount, address)
		delete(p.stakerConnectionAddresses, address)
	}
	return nil
}

func (p *StakePool) GetStakedAmounts() map[common.Address]*uint256.Int {
	return p.stakedAmount
}

func (p *StakePool) GetStakedAmount(address common.Address) *uint256.Int {
	if p.stakedAmount[address] == nil {
		return uint256.NewInt(0)
	}
	return p.stakedAmount[address]
}

func (p *StakePool) GetTotalStakedAmount() *uint256.Int {
	return p.totalStakedAmount
}

func (p *StakePool) GetStakerConnectionAddress(address common.Address) string {
	return p.stakerConnectionAddresses[address]
}

func (p *StakePool) Copy() *StakePool {
	copy := NewStakePool()
	for k, v := range p.stakedAmount {
		copy.stakedAmount[k] = v
	}
	for k, v := range p.stakerConnectionAddresses {
		copy.stakerConnectionAddresses[k] = v
	}
	copy.totalStakedAmount = p.totalStakedAmount.Clone()
	return copy
}

func (p *StakePool) GetStakedAmountsAboveThreshHold(threshHold *uint256.Int) map[common.Address]*uint256.Int {
	rs := map[common.Address]*uint256.Int{}
	for k, v := range p.stakedAmount {
		if v.Gt(threshHold) || v.Eq(threshHold) {
			rs[k] = v.Clone()
		}
	}
	return rs
}

func (p *StakePool) GetStakerConnectionAddressAboveThreshHold(threshHold *uint256.Int) map[common.Address]string {
	rs := map[common.Address]string{}
	for k, v := range p.stakedAmount {
		if v.Gt(threshHold) || v.Eq(threshHold) {
			rs[k] = p.stakerConnectionAddresses[k]
		}
	}
	return rs
}

func (p *StakePool) GetProto() protoreflect.ProtoMessage {
	protoSp := &pb.PosStakePool{
		StakedAmount:              make(map[string][]byte, len(p.stakedAmount)),
		StakerConnectionAddresses: make(map[string]string, len(p.stakerConnectionAddresses)),
	}
	for a, v := range p.stakedAmount {
		protoSp.StakedAmount[hex.EncodeToString(a.Bytes())] = v.Bytes()
	}
	for a, v := range p.stakerConnectionAddresses {
		protoSp.StakerConnectionAddresses[hex.EncodeToString(a.Bytes())] = v
	}
	return protoSp
}

func (p *StakePool) Marshal() ([]byte, error) {
	protoSp := p.GetProto()
	return proto.Marshal(protoSp)
}
