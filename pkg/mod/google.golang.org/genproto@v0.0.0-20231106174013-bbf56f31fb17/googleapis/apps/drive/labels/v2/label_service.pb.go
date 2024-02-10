// Copyright 2022 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: google/apps/drive/labels/v2/label_service.proto

package labels

import (
	context "context"
	reflect "reflect"

	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_google_apps_drive_labels_v2_label_service_proto protoreflect.FileDescriptor

var file_google_apps_drive_labels_v2_label_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x61,
	0x62, 0x65, 0x6c, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64,
	0x72, 0x69, 0x76, 0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x1a, 0x1c,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f,
	0x76, 0x32, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2a,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76,
	0x65, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x76, 0x32, 0x2f, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0xb4, 0x02, 0x0a, 0x0c, 0x4c,
	0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x81, 0x01, 0x0a, 0x0a,
	0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x2e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x12, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0c, 0x12, 0x0a, 0x2f, 0x76, 0x32, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12,
	0x80, 0x01, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x2c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x2e, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x2e, 0x76, 0x32, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x22, 0x22,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f, 0x76, 0x32, 0x2f, 0x7b, 0x6e, 0x61, 0x6d,
	0x65, 0x3d, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x2f, 0x2a, 0x7d, 0xda, 0x41, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x1a, 0x1d, 0xca, 0x41, 0x1a, 0x64, 0x72, 0x69, 0x76, 0x65, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73, 0x2e, 0x63, 0x6f,
	0x6d, 0x42, 0x80, 0x01, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x61, 0x70, 0x70, 0x73, 0x2e, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2e, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x2e, 0x76, 0x32, 0x42, 0x11, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x41, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69,
	0x73, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x2f, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x2f, 0x76, 0x32, 0x3b, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0xa2, 0x02, 0x04,
	0x44, 0x4c, 0x42, 0x4c, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_google_apps_drive_labels_v2_label_service_proto_goTypes = []interface{}{
	(*ListLabelsRequest)(nil),  // 0: google.apps.drive.labels.v2.ListLabelsRequest
	(*GetLabelRequest)(nil),    // 1: google.apps.drive.labels.v2.GetLabelRequest
	(*ListLabelsResponse)(nil), // 2: google.apps.drive.labels.v2.ListLabelsResponse
	(*Label)(nil),              // 3: google.apps.drive.labels.v2.Label
}
var file_google_apps_drive_labels_v2_label_service_proto_depIdxs = []int32{
	0, // 0: google.apps.drive.labels.v2.LabelService.ListLabels:input_type -> google.apps.drive.labels.v2.ListLabelsRequest
	1, // 1: google.apps.drive.labels.v2.LabelService.GetLabel:input_type -> google.apps.drive.labels.v2.GetLabelRequest
	2, // 2: google.apps.drive.labels.v2.LabelService.ListLabels:output_type -> google.apps.drive.labels.v2.ListLabelsResponse
	3, // 3: google.apps.drive.labels.v2.LabelService.GetLabel:output_type -> google.apps.drive.labels.v2.Label
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_google_apps_drive_labels_v2_label_service_proto_init() }
func file_google_apps_drive_labels_v2_label_service_proto_init() {
	if File_google_apps_drive_labels_v2_label_service_proto != nil {
		return
	}
	file_google_apps_drive_labels_v2_label_proto_init()
	file_google_apps_drive_labels_v2_requests_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_google_apps_drive_labels_v2_label_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_google_apps_drive_labels_v2_label_service_proto_goTypes,
		DependencyIndexes: file_google_apps_drive_labels_v2_label_service_proto_depIdxs,
	}.Build()
	File_google_apps_drive_labels_v2_label_service_proto = out.File
	file_google_apps_drive_labels_v2_label_service_proto_rawDesc = nil
	file_google_apps_drive_labels_v2_label_service_proto_goTypes = nil
	file_google_apps_drive_labels_v2_label_service_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// LabelServiceClient is the client API for LabelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LabelServiceClient interface {
	// List labels.
	ListLabels(ctx context.Context, in *ListLabelsRequest, opts ...grpc.CallOption) (*ListLabelsResponse, error)
	// Get a label by its resource name.
	// Resource name may be any of:
	//
	//   - `labels/{id}` - See `labels/{id}@latest`
	//   - `labels/{id}@latest` - Gets the latest revision of the label.
	//   - `labels/{id}@published` - Gets the current published revision of the
	//     label.
	//   - `labels/{id}@{revision_id}` - Gets the label at the specified revision
	//     ID.
	GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*Label, error)
}

type labelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLabelServiceClient(cc grpc.ClientConnInterface) LabelServiceClient {
	return &labelServiceClient{cc}
}

func (c *labelServiceClient) ListLabels(ctx context.Context, in *ListLabelsRequest, opts ...grpc.CallOption) (*ListLabelsResponse, error) {
	out := new(ListLabelsResponse)
	err := c.cc.Invoke(ctx, "/google.apps.drive.labels.v2.LabelService/ListLabels", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *labelServiceClient) GetLabel(ctx context.Context, in *GetLabelRequest, opts ...grpc.CallOption) (*Label, error) {
	out := new(Label)
	err := c.cc.Invoke(ctx, "/google.apps.drive.labels.v2.LabelService/GetLabel", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LabelServiceServer is the server API for LabelService service.
type LabelServiceServer interface {
	// List labels.
	ListLabels(context.Context, *ListLabelsRequest) (*ListLabelsResponse, error)
	// Get a label by its resource name.
	// Resource name may be any of:
	//
	//   - `labels/{id}` - See `labels/{id}@latest`
	//   - `labels/{id}@latest` - Gets the latest revision of the label.
	//   - `labels/{id}@published` - Gets the current published revision of the
	//     label.
	//   - `labels/{id}@{revision_id}` - Gets the label at the specified revision
	//     ID.
	GetLabel(context.Context, *GetLabelRequest) (*Label, error)
}

// UnimplementedLabelServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLabelServiceServer struct {
}

func (*UnimplementedLabelServiceServer) ListLabels(context.Context, *ListLabelsRequest) (*ListLabelsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLabels not implemented")
}
func (*UnimplementedLabelServiceServer) GetLabel(context.Context, *GetLabelRequest) (*Label, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLabel not implemented")
}

func RegisterLabelServiceServer(s *grpc.Server, srv LabelServiceServer) {
	s.RegisterService(&_LabelService_serviceDesc, srv)
}

func _LabelService_ListLabels_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLabelsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).ListLabels(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.apps.drive.labels.v2.LabelService/ListLabels",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).ListLabels(ctx, req.(*ListLabelsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LabelService_GetLabel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLabelRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LabelServiceServer).GetLabel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.apps.drive.labels.v2.LabelService/GetLabel",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LabelServiceServer).GetLabel(ctx, req.(*GetLabelRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _LabelService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.apps.drive.labels.v2.LabelService",
	HandlerType: (*LabelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLabels",
			Handler:    _LabelService_ListLabels_Handler,
		},
		{
			MethodName: "GetLabel",
			Handler:    _LabelService_GetLabel_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/apps/drive/labels/v2/label_service.proto",
}
