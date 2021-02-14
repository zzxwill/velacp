// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: pkg/datastore/model/cluster.proto

package model

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Cluster struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// baes64 encoded string of kubeconfig
	Kubeconfig string `protobuf:"bytes,2,opt,name=kubeconfig,proto3" json:"kubeconfig,omitempty"`
	// Unix time of the last time when the cluster is updated.
	UpdatedAt int64 `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Cluster) Reset() {
	*x = Cluster{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_datastore_model_cluster_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cluster) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cluster) ProtoMessage() {}

func (x *Cluster) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_datastore_model_cluster_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cluster.ProtoReflect.Descriptor instead.
func (*Cluster) Descriptor() ([]byte, []int) {
	return file_pkg_datastore_model_cluster_proto_rawDescGZIP(), []int{0}
}

func (x *Cluster) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Cluster) GetKubeconfig() string {
	if x != nil {
		return x.Kubeconfig
	}
	return ""
}

func (x *Cluster) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

var File_pkg_datastore_model_cluster_proto protoreflect.FileDescriptor

var file_pkg_datastore_model_cluster_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f,
	0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x76, 0x65, 0x6c, 0x61, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6e, 0x0a, 0x07,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x6b, 0x75, 0x62, 0x65, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x26, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20,
	0x00, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x2f, 0x5a, 0x2d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x61, 0x6d, 0x2d, 0x64,
	0x65, 0x76, 0x2f, 0x76, 0x65, 0x6c, 0x61, 0x63, 0x70, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_datastore_model_cluster_proto_rawDescOnce sync.Once
	file_pkg_datastore_model_cluster_proto_rawDescData = file_pkg_datastore_model_cluster_proto_rawDesc
)

func file_pkg_datastore_model_cluster_proto_rawDescGZIP() []byte {
	file_pkg_datastore_model_cluster_proto_rawDescOnce.Do(func() {
		file_pkg_datastore_model_cluster_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_datastore_model_cluster_proto_rawDescData)
	})
	return file_pkg_datastore_model_cluster_proto_rawDescData
}

var file_pkg_datastore_model_cluster_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_pkg_datastore_model_cluster_proto_goTypes = []interface{}{
	(*Cluster)(nil), // 0: vela.api.model.Cluster
}
var file_pkg_datastore_model_cluster_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_datastore_model_cluster_proto_init() }
func file_pkg_datastore_model_cluster_proto_init() {
	if File_pkg_datastore_model_cluster_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_datastore_model_cluster_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cluster); i {
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
			RawDescriptor: file_pkg_datastore_model_cluster_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pkg_datastore_model_cluster_proto_goTypes,
		DependencyIndexes: file_pkg_datastore_model_cluster_proto_depIdxs,
		MessageInfos:      file_pkg_datastore_model_cluster_proto_msgTypes,
	}.Build()
	File_pkg_datastore_model_cluster_proto = out.File
	file_pkg_datastore_model_cluster_proto_rawDesc = nil
	file_pkg_datastore_model_cluster_proto_goTypes = nil
	file_pkg_datastore_model_cluster_proto_depIdxs = nil
}