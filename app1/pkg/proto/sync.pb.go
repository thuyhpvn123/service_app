// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: sync.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SyncAccoutStates struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StorageData map[string][]byte `protobuf:"bytes,1,rep,name=StorageData,proto3" json:"StorageData,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Finished    bool              `protobuf:"varint,2,opt,name=Finished,proto3" json:"Finished,omitempty"`
}

func (x *SyncAccoutStates) Reset() {
	*x = SyncAccoutStates{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncAccoutStates) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncAccoutStates) ProtoMessage() {}

func (x *SyncAccoutStates) ProtoReflect() protoreflect.Message {
	mi := &file_sync_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncAccoutStates.ProtoReflect.Descriptor instead.
func (*SyncAccoutStates) Descriptor() ([]byte, []int) {
	return file_sync_proto_rawDescGZIP(), []int{0}
}

func (x *SyncAccoutStates) GetStorageData() map[string][]byte {
	if x != nil {
		return x.StorageData
	}
	return nil
}

func (x *SyncAccoutStates) GetFinished() bool {
	if x != nil {
		return x.Finished
	}
	return false
}

type SyncNodeConsensusConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PacksPerEntry           uint64 `protobuf:"varint,1,opt,name=PacksPerEntry,proto3" json:"PacksPerEntry,omitempty"`
	EntriesPerSlot          uint64 `protobuf:"varint,2,opt,name=EntriesPerSlot,proto3" json:"EntriesPerSlot,omitempty"`
	EntriesPerSecond        uint64 `protobuf:"varint,3,opt,name=EntriesPerSecond,proto3" json:"EntriesPerSecond,omitempty"`
	HashesPerEntry          uint64 `protobuf:"varint,4,opt,name=HashesPerEntry,proto3" json:"HashesPerEntry,omitempty"`
	ValidatorMinStakeAmount []byte `protobuf:"bytes,5,opt,name=ValidatorMinStakeAmount,proto3" json:"ValidatorMinStakeAmount,omitempty"`
}

func (x *SyncNodeConsensusConfig) Reset() {
	*x = SyncNodeConsensusConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sync_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncNodeConsensusConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncNodeConsensusConfig) ProtoMessage() {}

func (x *SyncNodeConsensusConfig) ProtoReflect() protoreflect.Message {
	mi := &file_sync_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncNodeConsensusConfig.ProtoReflect.Descriptor instead.
func (*SyncNodeConsensusConfig) Descriptor() ([]byte, []int) {
	return file_sync_proto_rawDescGZIP(), []int{1}
}

func (x *SyncNodeConsensusConfig) GetPacksPerEntry() uint64 {
	if x != nil {
		return x.PacksPerEntry
	}
	return 0
}

func (x *SyncNodeConsensusConfig) GetEntriesPerSlot() uint64 {
	if x != nil {
		return x.EntriesPerSlot
	}
	return 0
}

func (x *SyncNodeConsensusConfig) GetEntriesPerSecond() uint64 {
	if x != nil {
		return x.EntriesPerSecond
	}
	return 0
}

func (x *SyncNodeConsensusConfig) GetHashesPerEntry() uint64 {
	if x != nil {
		return x.HashesPerEntry
	}
	return 0
}

func (x *SyncNodeConsensusConfig) GetValidatorMinStakeAmount() []byte {
	if x != nil {
		return x.ValidatorMinStakeAmount
	}
	return nil
}

var File_sync_proto protoreflect.FileDescriptor

var file_sync_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x79,
	0x6e, 0x63, 0x22, 0xb9, 0x01, 0x0a, 0x10, 0x53, 0x79, 0x6e, 0x63, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12, 0x49, 0x0a, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x73,
	0x79, 0x6e, 0x63, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x74, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x73, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x1a, 0x3e,
	0x0a, 0x10, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xf5,
	0x01, 0x0a, 0x17, 0x53, 0x79, 0x6e, 0x63, 0x4e, 0x6f, 0x64, 0x65, 0x43, 0x6f, 0x6e, 0x73, 0x65,
	0x6e, 0x73, 0x75, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x50, 0x61,
	0x63, 0x6b, 0x73, 0x50, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0d, 0x50, 0x61, 0x63, 0x6b, 0x73, 0x50, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x26, 0x0a, 0x0e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53, 0x6c,
	0x6f, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65,
	0x73, 0x50, 0x65, 0x72, 0x53, 0x6c, 0x6f, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x45, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x10, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x50, 0x65, 0x72, 0x53, 0x65,
	0x63, 0x6f, 0x6e, 0x64, 0x12, 0x26, 0x0a, 0x0e, 0x48, 0x61, 0x73, 0x68, 0x65, 0x73, 0x50, 0x65,
	0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0e, 0x48, 0x61,
	0x73, 0x68, 0x65, 0x73, 0x50, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x38, 0x0a, 0x17,
	0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x6b,
	0x65, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x17, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x4d, 0x69, 0x6e, 0x53, 0x74, 0x61, 0x6b, 0x65,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x42, 0x26, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x69,
	0x61, 0x69, 0x76, 0x6e, 0x2e, 0x66, 0x69, 0x6e, 0x73, 0x64, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sync_proto_rawDescOnce sync.Once
	file_sync_proto_rawDescData = file_sync_proto_rawDesc
)

func file_sync_proto_rawDescGZIP() []byte {
	file_sync_proto_rawDescOnce.Do(func() {
		file_sync_proto_rawDescData = protoimpl.X.CompressGZIP(file_sync_proto_rawDescData)
	})
	return file_sync_proto_rawDescData
}

var file_sync_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_sync_proto_goTypes = []interface{}{
	(*SyncAccoutStates)(nil),        // 0: sync.SyncAccoutStates
	(*SyncNodeConsensusConfig)(nil), // 1: sync.SyncNodeConsensusConfig
	nil,                             // 2: sync.SyncAccoutStates.StorageDataEntry
}
var file_sync_proto_depIdxs = []int32{
	2, // 0: sync.SyncAccoutStates.StorageData:type_name -> sync.SyncAccoutStates.StorageDataEntry
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_sync_proto_init() }
func file_sync_proto_init() {
	if File_sync_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sync_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncAccoutStates); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_sync_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncNodeConsensusConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sync_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_sync_proto_goTypes,
		DependencyIndexes: file_sync_proto_depIdxs,
		MessageInfos:      file_sync_proto_msgTypes,
	}.Build()
	File_sync_proto = out.File
	file_sync_proto_rawDesc = nil
	file_sync_proto_goTypes = nil
	file_sync_proto_depIdxs = nil
}
