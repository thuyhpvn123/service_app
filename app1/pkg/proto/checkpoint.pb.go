// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: checkpoint.proto

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

type Checkpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastBlock          *Block          `protobuf:"bytes,1,opt,name=LastBlock,proto3" json:"LastBlock,omitempty"`
	ThisLeaderSchedule *LeaderSchedule `protobuf:"bytes,2,opt,name=ThisLeaderSchedule,proto3" json:"ThisLeaderSchedule,omitempty"`
	NextLeaderSchedule *LeaderSchedule `protobuf:"bytes,3,opt,name=NextLeaderSchedule,proto3" json:"NextLeaderSchedule,omitempty"`
}

func (x *Checkpoint) Reset() {
	*x = Checkpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checkpoint_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Checkpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Checkpoint) ProtoMessage() {}

func (x *Checkpoint) ProtoReflect() protoreflect.Message {
	mi := &file_checkpoint_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Checkpoint.ProtoReflect.Descriptor instead.
func (*Checkpoint) Descriptor() ([]byte, []int) {
	return file_checkpoint_proto_rawDescGZIP(), []int{0}
}

func (x *Checkpoint) GetLastBlock() *Block {
	if x != nil {
		return x.LastBlock
	}
	return nil
}

func (x *Checkpoint) GetThisLeaderSchedule() *LeaderSchedule {
	if x != nil {
		return x.ThisLeaderSchedule
	}
	return nil
}

func (x *Checkpoint) GetNextLeaderSchedule() *LeaderSchedule {
	if x != nil {
		return x.NextLeaderSchedule
	}
	return nil
}

var File_checkpoint_proto protoreflect.FileDescriptor

var file_checkpoint_proto_rawDesc = []byte{
	0x0a, 0x10, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0a, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x1a, 0x0b,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x70, 0x6f, 0x68,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc2, 0x01, 0x0a, 0x0a, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x2a, 0x0a, 0x09, 0x4c, 0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f,
	0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x09, 0x4c, 0x61, 0x73, 0x74, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x43, 0x0a, 0x12, 0x54, 0x68, 0x69, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e,
	0x70, 0x6f, 0x68, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x52, 0x12, 0x54, 0x68, 0x69, 0x73, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x43, 0x0a, 0x12, 0x4e, 0x65, 0x78, 0x74, 0x4c, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x6f, 0x68, 0x2e, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x12, 0x4e, 0x65, 0x78, 0x74, 0x4c, 0x65, 0x61,
	0x64, 0x65, 0x72, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x42, 0x2c, 0x0a, 0x22, 0x63,
	0x6f, 0x6d, 0x2e, 0x66, 0x69, 0x61, 0x69, 0x76, 0x6e, 0x2e, 0x66, 0x69, 0x6e, 0x73, 0x64, 0x6b,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x5a, 0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_checkpoint_proto_rawDescOnce sync.Once
	file_checkpoint_proto_rawDescData = file_checkpoint_proto_rawDesc
)

func file_checkpoint_proto_rawDescGZIP() []byte {
	file_checkpoint_proto_rawDescOnce.Do(func() {
		file_checkpoint_proto_rawDescData = protoimpl.X.CompressGZIP(file_checkpoint_proto_rawDescData)
	})
	return file_checkpoint_proto_rawDescData
}

var file_checkpoint_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_checkpoint_proto_goTypes = []interface{}{
	(*Checkpoint)(nil),     // 0: checkpoint.Checkpoint
	(*Block)(nil),          // 1: block.Block
	(*LeaderSchedule)(nil), // 2: poh.LeaderSchedule
}
var file_checkpoint_proto_depIdxs = []int32{
	1, // 0: checkpoint.Checkpoint.LastBlock:type_name -> block.Block
	2, // 1: checkpoint.Checkpoint.ThisLeaderSchedule:type_name -> poh.LeaderSchedule
	2, // 2: checkpoint.Checkpoint.NextLeaderSchedule:type_name -> poh.LeaderSchedule
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_checkpoint_proto_init() }
func file_checkpoint_proto_init() {
	if File_checkpoint_proto != nil {
		return
	}
	file_block_proto_init()
	file_poh_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_checkpoint_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Checkpoint); i {
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
			RawDescriptor: file_checkpoint_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_checkpoint_proto_goTypes,
		DependencyIndexes: file_checkpoint_proto_depIdxs,
		MessageInfos:      file_checkpoint_proto_msgTypes,
	}.Build()
	File_checkpoint_proto = out.File
	file_checkpoint_proto_rawDesc = nil
	file_checkpoint_proto_goTypes = nil
	file_checkpoint_proto_depIdxs = nil
}
