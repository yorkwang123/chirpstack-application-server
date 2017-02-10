// Code generated by protoc-gen-go.
// source: node.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateNodeRequest struct {
	// Name of the application to which the node must be added.
	ApplicationName string `protobuf:"bytes,13,opt,name=applicationName" json:"applicationName,omitempty"`
	// Hex encoded DevEUI.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	// Hex encoded AppEUI.
	AppEUI string `protobuf:"bytes,2,opt,name=appEUI" json:"appEUI,omitempty"`
	// Hex encoded AppKey.
	AppKey string `protobuf:"bytes,3,opt,name=appKey" json:"appKey,omitempty"`
	// RX delay.
	RxDelay uint32 `protobuf:"varint,4,opt,name=rxDelay" json:"rxDelay,omitempty"`
	// RX1 data-rate offset.
	Rx1DROffset uint32 `protobuf:"varint,5,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	// Channel-list ID used for CFlist (see LoRaWAN regional parameters if this applies to your region).
	ChannelListID int64 `protobuf:"varint,6,opt,name=channelListID" json:"channelListID,omitempty"`
	// RX window to use.
	RxWindow RXWindow `protobuf:"varint,7,opt,name=rxWindow,enum=api.RXWindow" json:"rxWindow,omitempty"`
	// Data-rate to use for RX2.
	Rx2DR uint32 `protobuf:"varint,8,opt,name=rx2DR" json:"rx2DR,omitempty"`
	// Name of the node (if left blank, it will be set to the DevEUI).
	Name string `protobuf:"bytes,9,opt,name=name" json:"name,omitempty"`
	// Relax frame-counter mode is enabled.
	RelaxFCnt bool `protobuf:"varint,10,opt,name=relaxFCnt" json:"relaxFCnt,omitempty"`
	// Interval (in frames) in which the ADR engine may adapt the data-rate of the node (0 = disabled).
	AdrInterval uint32 `protobuf:"varint,11,opt,name=adrInterval" json:"adrInterval,omitempty"`
	// Installation-margin to use for ADR calculation.
	InstallationMargin float64 `protobuf:"fixed64,12,opt,name=installationMargin" json:"installationMargin,omitempty"`
	// Description of the node.
	Description string `protobuf:"bytes,14,opt,name=description" json:"description,omitempty"`
	// Node is activated by ABP.
	IsABP bool `protobuf:"varint,15,opt,name=isABP" json:"isABP,omitempty"`
}

func (m *CreateNodeRequest) Reset()                    { *m = CreateNodeRequest{} }
func (m *CreateNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateNodeRequest) ProtoMessage()               {}
func (*CreateNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateNodeRequest) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *CreateNodeRequest) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *CreateNodeRequest) GetAppEUI() string {
	if m != nil {
		return m.AppEUI
	}
	return ""
}

func (m *CreateNodeRequest) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *CreateNodeRequest) GetRxDelay() uint32 {
	if m != nil {
		return m.RxDelay
	}
	return 0
}

func (m *CreateNodeRequest) GetRx1DROffset() uint32 {
	if m != nil {
		return m.Rx1DROffset
	}
	return 0
}

func (m *CreateNodeRequest) GetChannelListID() int64 {
	if m != nil {
		return m.ChannelListID
	}
	return 0
}

func (m *CreateNodeRequest) GetRxWindow() RXWindow {
	if m != nil {
		return m.RxWindow
	}
	return RXWindow_RX1
}

func (m *CreateNodeRequest) GetRx2DR() uint32 {
	if m != nil {
		return m.Rx2DR
	}
	return 0
}

func (m *CreateNodeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CreateNodeRequest) GetRelaxFCnt() bool {
	if m != nil {
		return m.RelaxFCnt
	}
	return false
}

func (m *CreateNodeRequest) GetAdrInterval() uint32 {
	if m != nil {
		return m.AdrInterval
	}
	return 0
}

func (m *CreateNodeRequest) GetInstallationMargin() float64 {
	if m != nil {
		return m.InstallationMargin
	}
	return 0
}

func (m *CreateNodeRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *CreateNodeRequest) GetIsABP() bool {
	if m != nil {
		return m.IsABP
	}
	return false
}

type CreateNodeResponse struct {
}

func (m *CreateNodeResponse) Reset()                    { *m = CreateNodeResponse{} }
func (m *CreateNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateNodeResponse) ProtoMessage()               {}
func (*CreateNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type GetNodeRequest struct {
	// Name of the application owning the node.
	ApplicationName string `protobuf:"bytes,2,opt,name=applicationName" json:"applicationName,omitempty"`
	// Name of the node.
	NodeName string `protobuf:"bytes,1,opt,name=nodeName" json:"nodeName,omitempty"`
}

func (m *GetNodeRequest) Reset()                    { *m = GetNodeRequest{} }
func (m *GetNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*GetNodeRequest) ProtoMessage()               {}
func (*GetNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *GetNodeRequest) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *GetNodeRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type GetNodeResponse struct {
	// Hex encoded DevEUI.
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	// Hex encoded AppEUI.
	AppEUI string `protobuf:"bytes,2,opt,name=appEUI" json:"appEUI,omitempty"`
	// Hex encoded AppKey.
	AppKey string `protobuf:"bytes,3,opt,name=appKey" json:"appKey,omitempty"`
	// RX delay.
	RxDelay uint32 `protobuf:"varint,4,opt,name=rxDelay" json:"rxDelay,omitempty"`
	// RX1 data-rate offset.
	Rx1DROffset uint32 `protobuf:"varint,5,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	// Channel-list ID used for CFlist (see LoRaWAN regional parameters if this applies to your region).
	ChannelListID int64 `protobuf:"varint,6,opt,name=channelListID" json:"channelListID,omitempty"`
	// RX window to use.
	RxWindow RXWindow `protobuf:"varint,7,opt,name=rxWindow,enum=api.RXWindow" json:"rxWindow,omitempty"`
	// Data-rate to use for RX2.
	Rx2DR uint32 `protobuf:"varint,8,opt,name=rx2DR" json:"rx2DR,omitempty"`
	// Name of the node.
	Name string `protobuf:"bytes,9,opt,name=name" json:"name,omitempty"`
	// Relax frame-counter mode is enabled.
	RelaxFCnt bool `protobuf:"varint,10,opt,name=relaxFCnt" json:"relaxFCnt,omitempty"`
	// Interval (in frames) in which the ADR engine may adapt the data-rate of the node (0 = disabled).
	AdrInterval uint32 `protobuf:"varint,11,opt,name=adrInterval" json:"adrInterval,omitempty"`
	// Installation-margin to use for ADR calculation.
	InstallationMargin float64 `protobuf:"fixed64,12,opt,name=installationMargin" json:"installationMargin,omitempty"`
	// Description of the node.
	Description string `protobuf:"bytes,13,opt,name=description" json:"description,omitempty"`
	// Node is activated by ABP.
	IsABP bool `protobuf:"varint,14,opt,name=isABP" json:"isABP,omitempty"`
}

func (m *GetNodeResponse) Reset()                    { *m = GetNodeResponse{} }
func (m *GetNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*GetNodeResponse) ProtoMessage()               {}
func (*GetNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *GetNodeResponse) GetDevEUI() string {
	if m != nil {
		return m.DevEUI
	}
	return ""
}

func (m *GetNodeResponse) GetAppEUI() string {
	if m != nil {
		return m.AppEUI
	}
	return ""
}

func (m *GetNodeResponse) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *GetNodeResponse) GetRxDelay() uint32 {
	if m != nil {
		return m.RxDelay
	}
	return 0
}

func (m *GetNodeResponse) GetRx1DROffset() uint32 {
	if m != nil {
		return m.Rx1DROffset
	}
	return 0
}

func (m *GetNodeResponse) GetChannelListID() int64 {
	if m != nil {
		return m.ChannelListID
	}
	return 0
}

func (m *GetNodeResponse) GetRxWindow() RXWindow {
	if m != nil {
		return m.RxWindow
	}
	return RXWindow_RX1
}

func (m *GetNodeResponse) GetRx2DR() uint32 {
	if m != nil {
		return m.Rx2DR
	}
	return 0
}

func (m *GetNodeResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GetNodeResponse) GetRelaxFCnt() bool {
	if m != nil {
		return m.RelaxFCnt
	}
	return false
}

func (m *GetNodeResponse) GetAdrInterval() uint32 {
	if m != nil {
		return m.AdrInterval
	}
	return 0
}

func (m *GetNodeResponse) GetInstallationMargin() float64 {
	if m != nil {
		return m.InstallationMargin
	}
	return 0
}

func (m *GetNodeResponse) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *GetNodeResponse) GetIsABP() bool {
	if m != nil {
		return m.IsABP
	}
	return false
}

type DeleteNodeRequest struct {
	// Name of the application owning the node.
	ApplicationName string `protobuf:"bytes,2,opt,name=applicationName" json:"applicationName,omitempty"`
	// Name of the node.
	NodeName string `protobuf:"bytes,1,opt,name=nodeName" json:"nodeName,omitempty"`
}

func (m *DeleteNodeRequest) Reset()                    { *m = DeleteNodeRequest{} }
func (m *DeleteNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeRequest) ProtoMessage()               {}
func (*DeleteNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *DeleteNodeRequest) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *DeleteNodeRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

type DeleteNodeResponse struct {
}

func (m *DeleteNodeResponse) Reset()                    { *m = DeleteNodeResponse{} }
func (m *DeleteNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteNodeResponse) ProtoMessage()               {}
func (*DeleteNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type ListNodeRequest struct {
	// Name of the application owning the node.
	ApplicationName string `protobuf:"bytes,3,opt,name=applicationName" json:"applicationName,omitempty"`
	// Max number of nodes to return in the result-set.
	Limit int64 `protobuf:"varint,1,opt,name=limit" json:"limit,omitempty"`
	// Offset of the result-set (for pagination).
	Offset int64 `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
}

func (m *ListNodeRequest) Reset()                    { *m = ListNodeRequest{} }
func (m *ListNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*ListNodeRequest) ProtoMessage()               {}
func (*ListNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ListNodeRequest) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *ListNodeRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *ListNodeRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

type ListNodeResponse struct {
	// Total number of nodes available within the result-set.
	TotalCount int64 `protobuf:"varint,1,opt,name=totalCount" json:"totalCount,omitempty"`
	// Nodes within this result-set.
	Result []*GetNodeResponse `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
}

func (m *ListNodeResponse) Reset()                    { *m = ListNodeResponse{} }
func (m *ListNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*ListNodeResponse) ProtoMessage()               {}
func (*ListNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

func (m *ListNodeResponse) GetTotalCount() int64 {
	if m != nil {
		return m.TotalCount
	}
	return 0
}

func (m *ListNodeResponse) GetResult() []*GetNodeResponse {
	if m != nil {
		return m.Result
	}
	return nil
}

type UpdateNodeRequest struct {
	// Name of the application owning the node.
	ApplicationName string `protobuf:"bytes,13,opt,name=applicationName" json:"applicationName,omitempty"`
	// Name of the node to update.
	NodeName string `protobuf:"bytes,14,opt,name=nodeName" json:"nodeName,omitempty"`
	// Hex encoded AppEUI.
	AppEUI string `protobuf:"bytes,2,opt,name=appEUI" json:"appEUI,omitempty"`
	// Hex encoded AppKey.
	AppKey string `protobuf:"bytes,3,opt,name=appKey" json:"appKey,omitempty"`
	// RX delay.
	RxDelay uint32 `protobuf:"varint,4,opt,name=rxDelay" json:"rxDelay,omitempty"`
	// RX1 data-rate offset.
	Rx1DROffset uint32 `protobuf:"varint,5,opt,name=rx1DROffset" json:"rx1DROffset,omitempty"`
	// Channel-list ID used for CFlist (see LoRaWAN regional parameters if this applies to your region).
	ChannelListID int64 `protobuf:"varint,6,opt,name=channelListID" json:"channelListID,omitempty"`
	// RX window to use.
	RxWindow RXWindow `protobuf:"varint,7,opt,name=rxWindow,enum=api.RXWindow" json:"rxWindow,omitempty"`
	// Data-rate to use for RX2.
	Rx2DR uint32 `protobuf:"varint,8,opt,name=rx2DR" json:"rx2DR,omitempty"`
	// Name of the node (note that renaming the node affects its api endpoint)
	Name string `protobuf:"bytes,9,opt,name=name" json:"name,omitempty"`
	// Relax frame-counter mode is enabled.
	RelaxFCnt bool `protobuf:"varint,10,opt,name=relaxFCnt" json:"relaxFCnt,omitempty"`
	// Interval (in frames) in which the ADR engine may adapt the data-rate of the node (0 = disabled).
	AdrInterval uint32 `protobuf:"varint,11,opt,name=adrInterval" json:"adrInterval,omitempty"`
	// Installation-margin to use for ADR calculation.
	InstallationMargin float64 `protobuf:"fixed64,12,opt,name=installationMargin" json:"installationMargin,omitempty"`
	// Description of the node.
	Description string `protobuf:"bytes,15,opt,name=description" json:"description,omitempty"`
	// Node is activated by ABP.
	IsABP bool `protobuf:"varint,16,opt,name=isABP" json:"isABP,omitempty"`
}

func (m *UpdateNodeRequest) Reset()                    { *m = UpdateNodeRequest{} }
func (m *UpdateNodeRequest) String() string            { return proto.CompactTextString(m) }
func (*UpdateNodeRequest) ProtoMessage()               {}
func (*UpdateNodeRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

func (m *UpdateNodeRequest) GetApplicationName() string {
	if m != nil {
		return m.ApplicationName
	}
	return ""
}

func (m *UpdateNodeRequest) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

func (m *UpdateNodeRequest) GetAppEUI() string {
	if m != nil {
		return m.AppEUI
	}
	return ""
}

func (m *UpdateNodeRequest) GetAppKey() string {
	if m != nil {
		return m.AppKey
	}
	return ""
}

func (m *UpdateNodeRequest) GetRxDelay() uint32 {
	if m != nil {
		return m.RxDelay
	}
	return 0
}

func (m *UpdateNodeRequest) GetRx1DROffset() uint32 {
	if m != nil {
		return m.Rx1DROffset
	}
	return 0
}

func (m *UpdateNodeRequest) GetChannelListID() int64 {
	if m != nil {
		return m.ChannelListID
	}
	return 0
}

func (m *UpdateNodeRequest) GetRxWindow() RXWindow {
	if m != nil {
		return m.RxWindow
	}
	return RXWindow_RX1
}

func (m *UpdateNodeRequest) GetRx2DR() uint32 {
	if m != nil {
		return m.Rx2DR
	}
	return 0
}

func (m *UpdateNodeRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *UpdateNodeRequest) GetRelaxFCnt() bool {
	if m != nil {
		return m.RelaxFCnt
	}
	return false
}

func (m *UpdateNodeRequest) GetAdrInterval() uint32 {
	if m != nil {
		return m.AdrInterval
	}
	return 0
}

func (m *UpdateNodeRequest) GetInstallationMargin() float64 {
	if m != nil {
		return m.InstallationMargin
	}
	return 0
}

func (m *UpdateNodeRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *UpdateNodeRequest) GetIsABP() bool {
	if m != nil {
		return m.IsABP
	}
	return false
}

type UpdateNodeResponse struct {
}

func (m *UpdateNodeResponse) Reset()                    { *m = UpdateNodeResponse{} }
func (m *UpdateNodeResponse) String() string            { return proto.CompactTextString(m) }
func (*UpdateNodeResponse) ProtoMessage()               {}
func (*UpdateNodeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

func init() {
	proto.RegisterType((*CreateNodeRequest)(nil), "api.CreateNodeRequest")
	proto.RegisterType((*CreateNodeResponse)(nil), "api.CreateNodeResponse")
	proto.RegisterType((*GetNodeRequest)(nil), "api.GetNodeRequest")
	proto.RegisterType((*GetNodeResponse)(nil), "api.GetNodeResponse")
	proto.RegisterType((*DeleteNodeRequest)(nil), "api.DeleteNodeRequest")
	proto.RegisterType((*DeleteNodeResponse)(nil), "api.DeleteNodeResponse")
	proto.RegisterType((*ListNodeRequest)(nil), "api.ListNodeRequest")
	proto.RegisterType((*ListNodeResponse)(nil), "api.ListNodeResponse")
	proto.RegisterType((*UpdateNodeRequest)(nil), "api.UpdateNodeRequest")
	proto.RegisterType((*UpdateNodeResponse)(nil), "api.UpdateNodeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Node service

type NodeClient interface {
	// Create creates the given node.
	Create(ctx context.Context, in *CreateNodeRequest, opts ...grpc.CallOption) (*CreateNodeResponse, error)
	// Get returns the node for the requested name.
	Get(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error)
	// Delete deletes the node matching the given name.
	Delete(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*DeleteNodeResponse, error)
	// List lists the nodes.
	List(ctx context.Context, in *ListNodeRequest, opts ...grpc.CallOption) (*ListNodeResponse, error)
	// Update updates the node matching the given name.
	Update(ctx context.Context, in *UpdateNodeRequest, opts ...grpc.CallOption) (*UpdateNodeResponse, error)
}

type nodeClient struct {
	cc *grpc.ClientConn
}

func NewNodeClient(cc *grpc.ClientConn) NodeClient {
	return &nodeClient{cc}
}

func (c *nodeClient) Create(ctx context.Context, in *CreateNodeRequest, opts ...grpc.CallOption) (*CreateNodeResponse, error) {
	out := new(CreateNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Create", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Get(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error) {
	out := new(GetNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Delete(ctx context.Context, in *DeleteNodeRequest, opts ...grpc.CallOption) (*DeleteNodeResponse, error) {
	out := new(DeleteNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) List(ctx context.Context, in *ListNodeRequest, opts ...grpc.CallOption) (*ListNodeResponse, error) {
	out := new(ListNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nodeClient) Update(ctx context.Context, in *UpdateNodeRequest, opts ...grpc.CallOption) (*UpdateNodeResponse, error) {
	out := new(UpdateNodeResponse)
	err := grpc.Invoke(ctx, "/api.Node/Update", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Node service

type NodeServer interface {
	// Create creates the given node.
	Create(context.Context, *CreateNodeRequest) (*CreateNodeResponse, error)
	// Get returns the node for the requested name.
	Get(context.Context, *GetNodeRequest) (*GetNodeResponse, error)
	// Delete deletes the node matching the given name.
	Delete(context.Context, *DeleteNodeRequest) (*DeleteNodeResponse, error)
	// List lists the nodes.
	List(context.Context, *ListNodeRequest) (*ListNodeResponse, error)
	// Update updates the node matching the given name.
	Update(context.Context, *UpdateNodeRequest) (*UpdateNodeResponse, error)
}

func RegisterNodeServer(s *grpc.Server, srv NodeServer) {
	s.RegisterService(&_Node_serviceDesc, srv)
}

func _Node_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Create(ctx, req.(*CreateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Get(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Delete(ctx, req.(*DeleteNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).List(ctx, req.(*ListNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Node_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NodeServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Node/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NodeServer).Update(ctx, req.(*UpdateNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Node_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Node",
	HandlerType: (*NodeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _Node_Create_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Node_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Node_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _Node_List_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _Node_Update_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "node.proto",
}

func init() { proto.RegisterFile("node.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 704 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xec, 0x56, 0xdd, 0x4e, 0x13, 0x41,
	0x14, 0xce, 0x76, 0xcb, 0x52, 0x0e, 0xb4, 0x85, 0xb1, 0xe2, 0xa4, 0x21, 0x66, 0xb3, 0xf1, 0x62,
	0x41, 0xd3, 0x86, 0x4a, 0x4c, 0x24, 0x26, 0x46, 0xa9, 0x12, 0x82, 0xa2, 0xd9, 0x04, 0x7f, 0xee,
	0x1c, 0xbb, 0x03, 0x4e, 0xb2, 0x9d, 0x59, 0x77, 0x07, 0x28, 0x12, 0x6e, 0x7c, 0x05, 0x5f, 0xc6,
	0x07, 0xf0, 0x0d, 0x7c, 0x05, 0x2e, 0x7c, 0x0c, 0x33, 0x33, 0x4b, 0x59, 0xda, 0x86, 0x92, 0x5e,
	0xe8, 0x0d, 0x77, 0x3d, 0xdf, 0x4c, 0xcf, 0x77, 0x7e, 0xbe, 0x73, 0x66, 0x01, 0xb8, 0x08, 0x69,
	0x23, 0x4e, 0x84, 0x14, 0xc8, 0x26, 0x31, 0xab, 0x2f, 0xed, 0x0b, 0xb1, 0x1f, 0xd1, 0x26, 0x89,
	0x59, 0x93, 0x70, 0x2e, 0x24, 0x91, 0x4c, 0xf0, 0xd4, 0x5c, 0xa9, 0xcf, 0x75, 0x44, 0xb7, 0x2b,
	0xb8, 0xb1, 0xbc, 0x33, 0x1b, 0x16, 0x36, 0x12, 0x4a, 0x24, 0xdd, 0x11, 0x21, 0x0d, 0xe8, 0xd7,
	0x03, 0x9a, 0x4a, 0xe4, 0x43, 0x95, 0xc4, 0x71, 0xc4, 0x3a, 0xfa, 0x9f, 0x3b, 0xa4, 0x4b, 0x71,
	0xd9, 0xb5, 0xfc, 0x99, 0x60, 0x10, 0x46, 0x8b, 0xe0, 0x84, 0xf4, 0xf0, 0xc5, 0xee, 0x16, 0xb6,
	0xf4, 0x85, 0xcc, 0x52, 0x38, 0x89, 0x63, 0x85, 0x17, 0x0c, 0x6e, 0xac, 0x0c, 0xdf, 0xa6, 0xc7,
	0xd8, 0xee, 0xe3, 0xdb, 0xf4, 0x18, 0x61, 0x98, 0x4e, 0x7a, 0x6d, 0x1a, 0x91, 0x63, 0x5c, 0x74,
	0x2d, 0xbf, 0x1c, 0x9c, 0x9b, 0xc8, 0x85, 0xd9, 0xa4, 0xb7, 0xda, 0x0e, 0xde, 0xec, 0xed, 0xa5,
	0x54, 0xe2, 0x29, 0x7d, 0x9a, 0x87, 0xd0, 0x3d, 0x28, 0x77, 0xbe, 0x10, 0xce, 0x69, 0xf4, 0x8a,
	0xa5, 0x72, 0xab, 0x8d, 0x1d, 0xd7, 0xf2, 0xed, 0xe0, 0x32, 0x88, 0x96, 0xa1, 0x94, 0xf4, 0xde,
	0x33, 0x1e, 0x8a, 0x23, 0x3c, 0xed, 0x5a, 0x7e, 0xa5, 0x55, 0x6e, 0x90, 0x98, 0x35, 0x82, 0x0f,
	0x06, 0x0c, 0xfa, 0xc7, 0xa8, 0x06, 0x53, 0x49, 0xaf, 0xd5, 0x0e, 0x70, 0x49, 0x93, 0x19, 0x03,
	0x21, 0x28, 0x72, 0x55, 0x89, 0x19, 0x1d, 0xb8, 0xfe, 0x8d, 0x96, 0x60, 0x26, 0xa1, 0x11, 0xe9,
	0xbd, 0xdc, 0xe0, 0x12, 0x83, 0x6b, 0xf9, 0xa5, 0xe0, 0x02, 0x50, 0xa1, 0x93, 0x30, 0xd9, 0xe2,
	0x92, 0x26, 0x87, 0x24, 0xc2, 0xb3, 0x26, 0xf4, 0x1c, 0x84, 0x1a, 0x80, 0x18, 0x4f, 0x25, 0x89,
	0x22, 0x5d, 0xd2, 0xd7, 0x24, 0xd9, 0x67, 0x1c, 0xcf, 0xb9, 0x96, 0x6f, 0x05, 0x23, 0x4e, 0x94,
	0xc7, 0x90, 0xa6, 0x9d, 0x84, 0xc5, 0x0a, 0xc4, 0x15, 0x1d, 0x4a, 0x1e, 0x52, 0xb1, 0xb3, 0xf4,
	0xd9, 0xf3, 0xb7, 0xb8, 0xaa, 0xa3, 0x31, 0x86, 0x57, 0x03, 0x94, 0xef, 0x72, 0x1a, 0x0b, 0x9e,
	0x52, 0xef, 0x1d, 0x54, 0x36, 0xa9, 0x1c, 0xd3, 0xf8, 0xc2, 0xe8, 0xc6, 0xd7, 0xa1, 0xa4, 0x74,
	0xa7, 0xaf, 0x98, 0xd6, 0xf7, 0x6d, 0xef, 0xa7, 0x0d, 0xd5, 0xbe, 0x63, 0xc3, 0x75, 0x23, 0x94,
	0xff, 0x2a, 0x94, 0xf2, 0x15, 0x42, 0xa9, 0xe4, 0x85, 0xf2, 0x11, 0x16, 0xda, 0x34, 0xa2, 0x63,
	0xd7, 0xc1, 0x04, 0xaa, 0xa8, 0x01, 0xca, 0xbb, 0xce, 0x34, 0xc8, 0xa0, 0xaa, 0xea, 0x3e, 0x86,
	0xce, 0x1e, 0x4d, 0x57, 0x83, 0xa9, 0x88, 0x75, 0x99, 0xd4, 0x5c, 0x76, 0x60, 0x0c, 0x25, 0x1d,
	0x61, 0x34, 0x50, 0xd0, 0x70, 0x66, 0x79, 0x9f, 0x60, 0xfe, 0x82, 0x2a, 0x93, 0xe5, 0x5d, 0x00,
	0x29, 0x24, 0x89, 0x36, 0xc4, 0x01, 0x3f, 0x77, 0x93, 0x43, 0xd0, 0x03, 0x70, 0x12, 0x9a, 0x1e,
	0x44, 0xca, 0x97, 0xed, 0xcf, 0xb6, 0x6a, 0x5a, 0x0a, 0x03, 0xe2, 0x0e, 0xb2, 0x3b, 0xde, 0x1f,
	0x1b, 0x16, 0x76, 0xe3, 0x70, 0xe2, 0x6d, 0x9a, 0x2f, 0x5f, 0xe5, 0x72, 0xf9, 0x6e, 0x06, 0xe5,
	0x9f, 0x0c, 0x4a, 0xf5, 0x8a, 0x41, 0x99, 0x1f, 0xd8, 0xa8, 0xf9, 0x4e, 0x1b, 0x21, 0xb4, 0x7e,
	0x15, 0xa1, 0xa8, 0x00, 0x24, 0xc0, 0x31, 0x0b, 0x17, 0x2d, 0xea, 0x9a, 0x0c, 0xbd, 0xb1, 0xf5,
	0x3b, 0x43, 0x78, 0x36, 0x11, 0x6b, 0xdf, 0x7f, 0x9f, 0xfd, 0x28, 0x34, 0xbc, 0x65, 0xf3, 0x80,
	0x5f, 0x48, 0x24, 0x6d, 0x9e, 0x0c, 0x08, 0xe6, 0xb4, 0xa9, 0xe4, 0x91, 0xae, 0x5b, 0x2b, 0x88,
	0x83, 0xbd, 0x49, 0x25, 0xba, 0x75, 0x59, 0x9f, 0x86, 0x6a, 0xa4, 0x68, 0xbd, 0x27, 0x9a, 0xe7,
	0x11, 0x5a, 0xbb, 0x36, 0x4f, 0xf3, 0xe4, 0x5c, 0x8d, 0xa7, 0xe8, 0x08, 0x1c, 0x33, 0xcd, 0x59,
	0x82, 0x43, 0x5b, 0x23, 0x4b, 0x70, 0xc4, 0xc8, 0x67, 0xc4, 0x2b, 0x93, 0x11, 0xef, 0x41, 0x51,
	0xe9, 0x0f, 0x99, 0xa4, 0x06, 0x76, 0x47, 0xfd, 0xf6, 0x00, 0x9a, 0x51, 0xae, 0x6a, 0xca, 0xfb,
	0xe8, 0xfa, 0x35, 0x45, 0xdf, 0xc0, 0x31, 0x0d, 0xce, 0x12, 0x1c, 0x9a, 0xeb, 0x2c, 0xc1, 0x61,
	0x15, 0x78, 0x4f, 0x35, 0xdb, 0xe3, 0xfa, 0x44, 0x09, 0xae, 0x5b, 0x2b, 0x9f, 0x1d, 0xfd, 0x71,
	0xf6, 0xf0, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf8, 0x24, 0xdb, 0xfd, 0xdb, 0x09, 0x00, 0x00,
}
