// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: api.proto

package grpcapi

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

type EvaluateCargoStateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId   int32 `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	CargoState int32 `protobuf:"varint,2,opt,name=cargo_state,json=cargoState,proto3" json:"cargo_state,omitempty"`
}

func (x *EvaluateCargoStateRequest) Reset() {
	*x = EvaluateCargoStateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluateCargoStateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluateCargoStateRequest) ProtoMessage() {}

func (x *EvaluateCargoStateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluateCargoStateRequest.ProtoReflect.Descriptor instead.
func (*EvaluateCargoStateRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{0}
}

func (x *EvaluateCargoStateRequest) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

func (x *EvaluateCargoStateRequest) GetCargoState() int32 {
	if x != nil {
		return x.CargoState
	}
	return 0
}

type EvaluateDriverServiceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId      int32 `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	DriverService int32 `protobuf:"varint,2,opt,name=driver_service,json=driverService,proto3" json:"driver_service,omitempty"`
}

func (x *EvaluateDriverServiceRequest) Reset() {
	*x = EvaluateDriverServiceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluateDriverServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluateDriverServiceRequest) ProtoMessage() {}

func (x *EvaluateDriverServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluateDriverServiceRequest.ProtoReflect.Descriptor instead.
func (*EvaluateDriverServiceRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{1}
}

func (x *EvaluateDriverServiceRequest) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

func (x *EvaluateDriverServiceRequest) GetDriverService() int32 {
	if x != nil {
		return x.DriverService
	}
	return 0
}

type EvaluateDeliverySpeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId      int32 `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
	DeliverySpeed int32 `protobuf:"varint,2,opt,name=delivery_speed,json=deliverySpeed,proto3" json:"delivery_speed,omitempty"`
}

func (x *EvaluateDeliverySpeedRequest) Reset() {
	*x = EvaluateDeliverySpeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluateDeliverySpeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluateDeliverySpeedRequest) ProtoMessage() {}

func (x *EvaluateDeliverySpeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluateDeliverySpeedRequest.ProtoReflect.Descriptor instead.
func (*EvaluateDeliverySpeedRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{2}
}

func (x *EvaluateDeliverySpeedRequest) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

func (x *EvaluateDeliverySpeedRequest) GetDeliverySpeed() int32 {
	if x != nil {
		return x.DeliverySpeed
	}
	return 0
}

type EvaluateCargoStateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EvaluateCargoStateResponse) Reset() {
	*x = EvaluateCargoStateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluateCargoStateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluateCargoStateResponse) ProtoMessage() {}

func (x *EvaluateCargoStateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluateCargoStateResponse.ProtoReflect.Descriptor instead.
func (*EvaluateCargoStateResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{3}
}

func (x *EvaluateCargoStateResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EvaluateDriverServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EvaluateDriverServiceResponse) Reset() {
	*x = EvaluateDriverServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluateDriverServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluateDriverServiceResponse) ProtoMessage() {}

func (x *EvaluateDriverServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluateDriverServiceResponse.ProtoReflect.Descriptor instead.
func (*EvaluateDriverServiceResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{4}
}

func (x *EvaluateDriverServiceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type EvaluateDeliverySpeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *EvaluateDeliverySpeedResponse) Reset() {
	*x = EvaluateDeliverySpeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EvaluateDeliverySpeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EvaluateDeliverySpeedResponse) ProtoMessage() {}

func (x *EvaluateDeliverySpeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EvaluateDeliverySpeedResponse.ProtoReflect.Descriptor instead.
func (*EvaluateDeliverySpeedResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{5}
}

func (x *EvaluateDeliverySpeedResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type DriverReviewsHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DriverId int32 `protobuf:"varint,1,opt,name=driver_id,json=driverId,proto3" json:"driver_id,omitempty"`
}

func (x *DriverReviewsHistoryRequest) Reset() {
	*x = DriverReviewsHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverReviewsHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverReviewsHistoryRequest) ProtoMessage() {}

func (x *DriverReviewsHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverReviewsHistoryRequest.ProtoReflect.Descriptor instead.
func (*DriverReviewsHistoryRequest) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{6}
}

func (x *DriverReviewsHistoryRequest) GetDriverId() int32 {
	if x != nil {
		return x.DriverId
	}
	return 0
}

type DriverReviewsHistoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CargoStates    []int32 `protobuf:"varint,1,rep,packed,name=cargo_states,json=cargoStates,proto3" json:"cargo_states,omitempty"`
	DriverServices []int32 `protobuf:"varint,2,rep,packed,name=driver_services,json=driverServices,proto3" json:"driver_services,omitempty"`
	DeliverySpeeds []int32 `protobuf:"varint,3,rep,packed,name=delivery_speeds,json=deliverySpeeds,proto3" json:"delivery_speeds,omitempty"`
}

func (x *DriverReviewsHistoryResponse) Reset() {
	*x = DriverReviewsHistoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DriverReviewsHistoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DriverReviewsHistoryResponse) ProtoMessage() {}

func (x *DriverReviewsHistoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DriverReviewsHistoryResponse.ProtoReflect.Descriptor instead.
func (*DriverReviewsHistoryResponse) Descriptor() ([]byte, []int) {
	return file_api_proto_rawDescGZIP(), []int{7}
}

func (x *DriverReviewsHistoryResponse) GetCargoStates() []int32 {
	if x != nil {
		return x.CargoStates
	}
	return nil
}

func (x *DriverReviewsHistoryResponse) GetDriverServices() []int32 {
	if x != nil {
		return x.DriverServices
	}
	return nil
}

func (x *DriverReviewsHistoryResponse) GetDeliverySpeeds() []int32 {
	if x != nil {
		return x.DeliverySpeeds
	}
	return nil
}

var File_api_proto protoreflect.FileDescriptor

var file_api_proto_rawDesc = []byte{
	0x0a, 0x09, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x74, 0x61, 0x78,
	0x69, 0x22, 0x59, 0x0a, 0x19, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x67, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x63,
	0x61, 0x72, 0x67, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x62, 0x0a, 0x1c,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0d, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x62, 0x0a, 0x1c, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x25, 0x0a,
	0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53,
	0x70, 0x65, 0x65, 0x64, 0x22, 0x36, 0x0a, 0x1a, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x67, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x39, 0x0a, 0x1d,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x39, 0x0a, 0x1d, 0x45, 0x76, 0x61, 0x6c, 0x75,
	0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x70, 0x65, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x3a, 0x0a, 0x1b, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64, 0x22, 0x93,
	0x01, 0x0a, 0x1c, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x21, 0x0a, 0x0c, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0b, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x05, 0x52, 0x0e, 0x64, 0x72, 0x69,
	0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x27, 0x0a, 0x0f, 0x64,
	0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x5f, 0x73, 0x70, 0x65, 0x65, 0x64, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x05, 0x52, 0x0e, 0x64, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79, 0x53, 0x70,
	0x65, 0x65, 0x64, 0x73, 0x32, 0x89, 0x03, 0x0a, 0x0b, 0x54, 0x61, 0x78, 0x69, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x57, 0x0a, 0x12, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x67, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1f, 0x2e, 0x74, 0x61, 0x78,
	0x69, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x67, 0x6f, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x61,
	0x78, 0x69, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x67, 0x6f,
	0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a,
	0x15, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x22, 0x2e, 0x74, 0x61, 0x78, 0x69, 0x2e, 0x45, 0x76,
	0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x74, 0x61, 0x78,
	0x69, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x60, 0x0a, 0x15, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x53, 0x70, 0x65, 0x65, 0x64, 0x12, 0x22, 0x2e, 0x74, 0x61, 0x78, 0x69, 0x2e,
	0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69, 0x76, 0x65, 0x72, 0x79,
	0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x74,
	0x61, 0x78, 0x69, 0x2e, 0x45, 0x76, 0x61, 0x6c, 0x75, 0x61, 0x74, 0x65, 0x44, 0x65, 0x6c, 0x69,
	0x76, 0x65, 0x72, 0x79, 0x53, 0x70, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x5d, 0x0a, 0x14, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x76, 0x69, 0x65,
	0x77, 0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x21, 0x2e, 0x74, 0x61, 0x78, 0x69,
	0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77, 0x73, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74,
	0x61, 0x78, 0x69, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x65, 0x76, 0x69, 0x65, 0x77,
	0x73, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76,
	0x69, 0x63, 0x75, 0x61, 0x6e, 0x69, 0x2f, 0x67, 0x6f, 0x5f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65,
	0x2f, 0x67, 0x6f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x31, 0x37, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_proto_rawDescOnce sync.Once
	file_api_proto_rawDescData = file_api_proto_rawDesc
)

func file_api_proto_rawDescGZIP() []byte {
	file_api_proto_rawDescOnce.Do(func() {
		file_api_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_proto_rawDescData)
	})
	return file_api_proto_rawDescData
}

var file_api_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_api_proto_goTypes = []any{
	(*EvaluateCargoStateRequest)(nil),     // 0: taxi.EvaluateCargoStateRequest
	(*EvaluateDriverServiceRequest)(nil),  // 1: taxi.EvaluateDriverServiceRequest
	(*EvaluateDeliverySpeedRequest)(nil),  // 2: taxi.EvaluateDeliverySpeedRequest
	(*EvaluateCargoStateResponse)(nil),    // 3: taxi.EvaluateCargoStateResponse
	(*EvaluateDriverServiceResponse)(nil), // 4: taxi.EvaluateDriverServiceResponse
	(*EvaluateDeliverySpeedResponse)(nil), // 5: taxi.EvaluateDeliverySpeedResponse
	(*DriverReviewsHistoryRequest)(nil),   // 6: taxi.DriverReviewsHistoryRequest
	(*DriverReviewsHistoryResponse)(nil),  // 7: taxi.DriverReviewsHistoryResponse
}
var file_api_proto_depIdxs = []int32{
	0, // 0: taxi.TaxiService.EvaluateCargoState:input_type -> taxi.EvaluateCargoStateRequest
	1, // 1: taxi.TaxiService.EvaluateDriverService:input_type -> taxi.EvaluateDriverServiceRequest
	2, // 2: taxi.TaxiService.EvaluateDeliverySpeed:input_type -> taxi.EvaluateDeliverySpeedRequest
	6, // 3: taxi.TaxiService.DriverReviewsHistory:input_type -> taxi.DriverReviewsHistoryRequest
	3, // 4: taxi.TaxiService.EvaluateCargoState:output_type -> taxi.EvaluateCargoStateResponse
	4, // 5: taxi.TaxiService.EvaluateDriverService:output_type -> taxi.EvaluateDriverServiceResponse
	5, // 6: taxi.TaxiService.EvaluateDeliverySpeed:output_type -> taxi.EvaluateDeliverySpeedResponse
	7, // 7: taxi.TaxiService.DriverReviewsHistory:output_type -> taxi.DriverReviewsHistoryResponse
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_init() }
func file_api_proto_init() {
	if File_api_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*EvaluateCargoStateRequest); i {
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
		file_api_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*EvaluateDriverServiceRequest); i {
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
		file_api_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*EvaluateDeliverySpeedRequest); i {
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
		file_api_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*EvaluateCargoStateResponse); i {
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
		file_api_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*EvaluateDriverServiceResponse); i {
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
		file_api_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*EvaluateDeliverySpeedResponse); i {
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
		file_api_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DriverReviewsHistoryRequest); i {
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
		file_api_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*DriverReviewsHistoryResponse); i {
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
			RawDescriptor: file_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_goTypes,
		DependencyIndexes: file_api_proto_depIdxs,
		MessageInfos:      file_api_proto_msgTypes,
	}.Build()
	File_api_proto = out.File
	file_api_proto_rawDesc = nil
	file_api_proto_goTypes = nil
	file_api_proto_depIdxs = nil
}
