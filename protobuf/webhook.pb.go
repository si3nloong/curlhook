// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: webhook.proto

package protobuf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type HealthCheckResponse_ServingStatus int32

const (
	HealthCheckResponse_UNKNOWN         HealthCheckResponse_ServingStatus = 0
	HealthCheckResponse_SERVING         HealthCheckResponse_ServingStatus = 1
	HealthCheckResponse_NOT_SERVING     HealthCheckResponse_ServingStatus = 2
	HealthCheckResponse_SERVICE_UNKNOWN HealthCheckResponse_ServingStatus = 3 // Used only by the Watch method.
)

// Enum value maps for HealthCheckResponse_ServingStatus.
var (
	HealthCheckResponse_ServingStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SERVING",
		2: "NOT_SERVING",
		3: "SERVICE_UNKNOWN",
	}
	HealthCheckResponse_ServingStatus_value = map[string]int32{
		"UNKNOWN":         0,
		"SERVING":         1,
		"NOT_SERVING":     2,
		"SERVICE_UNKNOWN": 3,
	}
)

func (x HealthCheckResponse_ServingStatus) Enum() *HealthCheckResponse_ServingStatus {
	p := new(HealthCheckResponse_ServingStatus)
	*p = x
	return p
}

func (x HealthCheckResponse_ServingStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HealthCheckResponse_ServingStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_webhook_proto_enumTypes[0].Descriptor()
}

func (HealthCheckResponse_ServingStatus) Type() protoreflect.EnumType {
	return &file_webhook_proto_enumTypes[0]
}

func (x HealthCheckResponse_ServingStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HealthCheckResponse_ServingStatus.Descriptor instead.
func (HealthCheckResponse_ServingStatus) EnumDescriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{1, 0}
}

type SendWebhookRequest_HttpMethod int32

const (
	SendWebhookRequest_GET  SendWebhookRequest_HttpMethod = 0
	SendWebhookRequest_POST SendWebhookRequest_HttpMethod = 1
)

// Enum value maps for SendWebhookRequest_HttpMethod.
var (
	SendWebhookRequest_HttpMethod_name = map[int32]string{
		0: "GET",
		1: "POST",
	}
	SendWebhookRequest_HttpMethod_value = map[string]int32{
		"GET":  0,
		"POST": 1,
	}
)

func (x SendWebhookRequest_HttpMethod) Enum() *SendWebhookRequest_HttpMethod {
	p := new(SendWebhookRequest_HttpMethod)
	*p = x
	return p
}

func (x SendWebhookRequest_HttpMethod) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SendWebhookRequest_HttpMethod) Descriptor() protoreflect.EnumDescriptor {
	return file_webhook_proto_enumTypes[1].Descriptor()
}

func (SendWebhookRequest_HttpMethod) Type() protoreflect.EnumType {
	return &file_webhook_proto_enumTypes[1]
}

func (x SendWebhookRequest_HttpMethod) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SendWebhookRequest_HttpMethod.Descriptor instead.
func (SendWebhookRequest_HttpMethod) EnumDescriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{6, 0}
}

type SendWebhookRequest_RetryMechanism int32

const (
	SendWebhookRequest_BACKOFF SendWebhookRequest_RetryMechanism = 0
	SendWebhookRequest_LINEAR  SendWebhookRequest_RetryMechanism = 1
)

// Enum value maps for SendWebhookRequest_RetryMechanism.
var (
	SendWebhookRequest_RetryMechanism_name = map[int32]string{
		0: "BACKOFF",
		1: "LINEAR",
	}
	SendWebhookRequest_RetryMechanism_value = map[string]int32{
		"BACKOFF": 0,
		"LINEAR":  1,
	}
)

func (x SendWebhookRequest_RetryMechanism) Enum() *SendWebhookRequest_RetryMechanism {
	p := new(SendWebhookRequest_RetryMechanism)
	*p = x
	return p
}

func (x SendWebhookRequest_RetryMechanism) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SendWebhookRequest_RetryMechanism) Descriptor() protoreflect.EnumDescriptor {
	return file_webhook_proto_enumTypes[2].Descriptor()
}

func (SendWebhookRequest_RetryMechanism) Type() protoreflect.EnumType {
	return &file_webhook_proto_enumTypes[2]
}

func (x SendWebhookRequest_RetryMechanism) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SendWebhookRequest_RetryMechanism.Descriptor instead.
func (SendWebhookRequest_RetryMechanism) EnumDescriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{6, 1}
}

type HealthCheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
}

func (x *HealthCheckRequest) Reset() {
	*x = HealthCheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckRequest) ProtoMessage() {}

func (x *HealthCheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckRequest.ProtoReflect.Descriptor instead.
func (*HealthCheckRequest) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{0}
}

func (x *HealthCheckRequest) GetService() string {
	if x != nil {
		return x.Service
	}
	return ""
}

type HealthCheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status HealthCheckResponse_ServingStatus `protobuf:"varint,1,opt,name=status,proto3,enum=protobuf.HealthCheckResponse_ServingStatus" json:"status,omitempty"`
}

func (x *HealthCheckResponse) Reset() {
	*x = HealthCheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HealthCheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HealthCheckResponse) ProtoMessage() {}

func (x *HealthCheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HealthCheckResponse.ProtoReflect.Descriptor instead.
func (*HealthCheckResponse) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{1}
}

func (x *HealthCheckResponse) GetStatus() HealthCheckResponse_ServingStatus {
	if x != nil {
		return x.Status
	}
	return HealthCheckResponse_UNKNOWN
}

type ListWebhooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"omitempty,required"
	PageToken string `protobuf:"bytes,1,opt,name=page_token,json=pageToken,proto3" json:"page_token,omitempty" validate:"omitempty,required"`
	// @gotags: validate:"omitempty,max=100"
	PageSize uint32 `protobuf:"varint,2,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty" validate:"omitempty,max=100"`
}

func (x *ListWebhooksRequest) Reset() {
	*x = ListWebhooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWebhooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWebhooksRequest) ProtoMessage() {}

func (x *ListWebhooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWebhooksRequest.ProtoReflect.Descriptor instead.
func (*ListWebhooksRequest) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{2}
}

func (x *ListWebhooksRequest) GetPageToken() string {
	if x != nil {
		return x.PageToken
	}
	return ""
}

func (x *ListWebhooksRequest) GetPageSize() uint32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListWebhooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Webhooks      []*Webhook `protobuf:"bytes,1,rep,name=webhooks,proto3" json:"webhooks,omitempty"`
	NextPageToken string     `protobuf:"bytes,2,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListWebhooksResponse) Reset() {
	*x = ListWebhooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListWebhooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListWebhooksResponse) ProtoMessage() {}

func (x *ListWebhooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListWebhooksResponse.ProtoReflect.Descriptor instead.
func (*ListWebhooksResponse) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{3}
}

func (x *ListWebhooksResponse) GetWebhooks() []*Webhook {
	if x != nil {
		return x.Webhooks
	}
	return nil
}

func (x *ListWebhooksResponse) GetNextPageToken() string {
	if x != nil {
		return x.NextPageToken
	}
	return ""
}

type GetWebhookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @gotags: validate:"required"
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" validate:"required"`
}

func (x *GetWebhookRequest) Reset() {
	*x = GetWebhookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWebhookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWebhookRequest) ProtoMessage() {}

func (x *GetWebhookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWebhookRequest.ProtoReflect.Descriptor instead.
func (*GetWebhookRequest) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{4}
}

func (x *GetWebhookRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetWebhookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Webhook *Webhook `protobuf:"bytes,1,opt,name=webhook,proto3" json:"webhook,omitempty"`
}

func (x *GetWebhookResponse) Reset() {
	*x = GetWebhookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetWebhookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetWebhookResponse) ProtoMessage() {}

func (x *GetWebhookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetWebhookResponse.ProtoReflect.Descriptor instead.
func (*GetWebhookResponse) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{5}
}

func (x *GetWebhookResponse) GetWebhook() *Webhook {
	if x != nil {
		return x.Webhook
	}
	return nil
}

type SendWebhookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Method SendWebhookRequest_HttpMethod `protobuf:"varint,1,opt,name=method,proto3,enum=protobuf.SendWebhookRequest_HttpMethod" json:"method,omitempty"`
	// @gotags: validate:"required,url,max=1000"
	Url     string            `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty" validate:"required,url,max=1000"`
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// @gotags: validate:"max=2048"
	Body string `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty" validate:"max=2048"`
	// @gotags: validate:"lte=10"
	Retry          uint32                            `protobuf:"varint,5,opt,name=retry,proto3" json:"retry,omitempty" validate:"lte=10"`
	RetryMechanism SendWebhookRequest_RetryMechanism `protobuf:"varint,6,opt,name=retry_mechanism,json=retryMechanism,proto3,enum=protobuf.SendWebhookRequest_RetryMechanism" json:"retry_mechanism,omitempty"`
	Timeout        uint32                            `protobuf:"varint,7,opt,name=timeout,proto3" json:"timeout,omitempty"`
	// @gotags: validate:"lte=3"
	Concurrent uint32 `protobuf:"varint,8,opt,name=concurrent,proto3" json:"concurrent,omitempty" validate:"lte=3"`
}

func (x *SendWebhookRequest) Reset() {
	*x = SendWebhookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendWebhookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendWebhookRequest) ProtoMessage() {}

func (x *SendWebhookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendWebhookRequest.ProtoReflect.Descriptor instead.
func (*SendWebhookRequest) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{6}
}

func (x *SendWebhookRequest) GetMethod() SendWebhookRequest_HttpMethod {
	if x != nil {
		return x.Method
	}
	return SendWebhookRequest_GET
}

func (x *SendWebhookRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *SendWebhookRequest) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *SendWebhookRequest) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *SendWebhookRequest) GetRetry() uint32 {
	if x != nil {
		return x.Retry
	}
	return 0
}

func (x *SendWebhookRequest) GetRetryMechanism() SendWebhookRequest_RetryMechanism {
	if x != nil {
		return x.RetryMechanism
	}
	return SendWebhookRequest_BACKOFF
}

func (x *SendWebhookRequest) GetTimeout() uint32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

func (x *SendWebhookRequest) GetConcurrent() uint32 {
	if x != nil {
		return x.Concurrent
	}
	return 0
}

type SendWebhookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
}

func (x *SendWebhookResponse) Reset() {
	*x = SendWebhookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendWebhookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendWebhookResponse) ProtoMessage() {}

func (x *SendWebhookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendWebhookResponse.ProtoReflect.Descriptor instead.
func (*SendWebhookResponse) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{7}
}

func (x *SendWebhookResponse) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

type Webhook struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Method    string                 `protobuf:"bytes,2,opt,name=method,proto3" json:"method,omitempty"`
	Body      string                 `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	Retries   uint32                 `protobuf:"varint,4,opt,name=retries,proto3" json:"retries,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Webhook) Reset() {
	*x = Webhook{}
	if protoimpl.UnsafeEnabled {
		mi := &file_webhook_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Webhook) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Webhook) ProtoMessage() {}

func (x *Webhook) ProtoReflect() protoreflect.Message {
	mi := &file_webhook_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Webhook.ProtoReflect.Descriptor instead.
func (*Webhook) Descriptor() ([]byte, []int) {
	return file_webhook_proto_rawDescGZIP(), []int{8}
}

func (x *Webhook) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Webhook) GetMethod() string {
	if x != nil {
		return x.Method
	}
	return ""
}

func (x *Webhook) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Webhook) GetRetries() uint32 {
	if x != nil {
		return x.Retries
	}
	return 0
}

func (x *Webhook) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Webhook) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_webhook_proto protoreflect.FileDescriptor

var file_webhook_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x12, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0xab, 0x01, 0x0a, 0x13, 0x48,
	0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x43, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4f, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x6e, 0x67, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e, 0x47,
	0x10, 0x01, 0x12, 0x0f, 0x0a, 0x0b, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x4e,
	0x47, 0x10, 0x02, 0x12, 0x13, 0x0a, 0x0f, 0x53, 0x45, 0x52, 0x56, 0x49, 0x43, 0x45, 0x5f, 0x55,
	0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x03, 0x22, 0x51, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1b,
	0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0d, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x6d, 0x0a, 0x14, 0x4c,
	0x69, 0x73, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x2d, 0x0a, 0x08, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x08, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6e, 0x65, 0x78,
	0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65,
	0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x41, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x07, 0x77, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x22, 0xee, 0x03, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x57, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x3f, 0x0a, 0x06, 0x6d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x74, 0x74, 0x70, 0x4d, 0x65, 0x74, 0x68,
	0x6f, 0x64, 0x52, 0x06, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x43, 0x0a, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x57, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x48, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72,
	0x73, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x74, 0x72, 0x79, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x72, 0x65, 0x74, 0x72, 0x79, 0x12, 0x54, 0x0a, 0x0f, 0x72,
	0x65, 0x74, 0x72, 0x79, 0x5f, 0x6d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x53, 0x65, 0x6e, 0x64, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x2e, 0x52, 0x65, 0x74, 0x72, 0x79, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73,
	0x6d, 0x52, 0x0e, 0x72, 0x65, 0x74, 0x72, 0x79, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73,
	0x6d, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x63,
	0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x0a, 0x63, 0x6f, 0x6e, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x1a, 0x3a, 0x0a, 0x0c, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x1f, 0x0a, 0x0a, 0x48, 0x74, 0x74, 0x70, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x45, 0x54, 0x10, 0x00, 0x12, 0x08,
	0x0a, 0x04, 0x50, 0x4f, 0x53, 0x54, 0x10, 0x01, 0x22, 0x29, 0x0a, 0x0e, 0x52, 0x65, 0x74, 0x72,
	0x79, 0x4d, 0x65, 0x63, 0x68, 0x61, 0x6e, 0x69, 0x73, 0x6d, 0x12, 0x0b, 0x0a, 0x07, 0x42, 0x41,
	0x43, 0x4b, 0x4f, 0x46, 0x46, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x4c, 0x49, 0x4e, 0x45, 0x41,
	0x52, 0x10, 0x01, 0x22, 0x2b, 0x0a, 0x13, 0x53, 0x65, 0x6e, 0x64, 0x57, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x22, 0xd5, 0x01, 0x0a, 0x07, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x07, 0x72, 0x65, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x39, 0x0a,
	0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x75,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0x82, 0x03, 0x0a, 0x0e, 0x57, 0x65, 0x62,
	0x68, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x05, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x48, 0x65,
	0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x46, 0x0a, 0x05, 0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63,
	0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x48, 0x65, 0x61, 0x6c, 0x74, 0x68, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x4d, 0x0a, 0x0c, 0x4c, 0x69, 0x73,
	0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x1d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x57,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x47, 0x65, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x47,
	0x65, 0x74, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x53, 0x65, 0x6e, 0x64, 0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b,
	0x12, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x57, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x57, 0x65,
	0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a,
	0x09, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_webhook_proto_rawDescOnce sync.Once
	file_webhook_proto_rawDescData = file_webhook_proto_rawDesc
)

func file_webhook_proto_rawDescGZIP() []byte {
	file_webhook_proto_rawDescOnce.Do(func() {
		file_webhook_proto_rawDescData = protoimpl.X.CompressGZIP(file_webhook_proto_rawDescData)
	})
	return file_webhook_proto_rawDescData
}

var file_webhook_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_webhook_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_webhook_proto_goTypes = []interface{}{
	(HealthCheckResponse_ServingStatus)(0), // 0: protobuf.HealthCheckResponse.ServingStatus
	(SendWebhookRequest_HttpMethod)(0),     // 1: protobuf.SendWebhookRequest.HttpMethod
	(SendWebhookRequest_RetryMechanism)(0), // 2: protobuf.SendWebhookRequest.RetryMechanism
	(*HealthCheckRequest)(nil),             // 3: protobuf.HealthCheckRequest
	(*HealthCheckResponse)(nil),            // 4: protobuf.HealthCheckResponse
	(*ListWebhooksRequest)(nil),            // 5: protobuf.ListWebhooksRequest
	(*ListWebhooksResponse)(nil),           // 6: protobuf.ListWebhooksResponse
	(*GetWebhookRequest)(nil),              // 7: protobuf.GetWebhookRequest
	(*GetWebhookResponse)(nil),             // 8: protobuf.GetWebhookResponse
	(*SendWebhookRequest)(nil),             // 9: protobuf.SendWebhookRequest
	(*SendWebhookResponse)(nil),            // 10: protobuf.SendWebhookResponse
	(*Webhook)(nil),                        // 11: protobuf.Webhook
	nil,                                    // 12: protobuf.SendWebhookRequest.HeadersEntry
	(*timestamppb.Timestamp)(nil),          // 13: google.protobuf.Timestamp
}
var file_webhook_proto_depIdxs = []int32{
	0,  // 0: protobuf.HealthCheckResponse.status:type_name -> protobuf.HealthCheckResponse.ServingStatus
	11, // 1: protobuf.ListWebhooksResponse.webhooks:type_name -> protobuf.Webhook
	11, // 2: protobuf.GetWebhookResponse.webhook:type_name -> protobuf.Webhook
	1,  // 3: protobuf.SendWebhookRequest.method:type_name -> protobuf.SendWebhookRequest.HttpMethod
	12, // 4: protobuf.SendWebhookRequest.headers:type_name -> protobuf.SendWebhookRequest.HeadersEntry
	2,  // 5: protobuf.SendWebhookRequest.retry_mechanism:type_name -> protobuf.SendWebhookRequest.RetryMechanism
	13, // 6: protobuf.Webhook.created_at:type_name -> google.protobuf.Timestamp
	13, // 7: protobuf.Webhook.updated_at:type_name -> google.protobuf.Timestamp
	3,  // 8: protobuf.WebhookService.Check:input_type -> protobuf.HealthCheckRequest
	3,  // 9: protobuf.WebhookService.Watch:input_type -> protobuf.HealthCheckRequest
	5,  // 10: protobuf.WebhookService.ListWebhooks:input_type -> protobuf.ListWebhooksRequest
	7,  // 11: protobuf.WebhookService.GetWebhook:input_type -> protobuf.GetWebhookRequest
	9,  // 12: protobuf.WebhookService.SendWebhook:input_type -> protobuf.SendWebhookRequest
	4,  // 13: protobuf.WebhookService.Check:output_type -> protobuf.HealthCheckResponse
	4,  // 14: protobuf.WebhookService.Watch:output_type -> protobuf.HealthCheckResponse
	6,  // 15: protobuf.WebhookService.ListWebhooks:output_type -> protobuf.ListWebhooksResponse
	8,  // 16: protobuf.WebhookService.GetWebhook:output_type -> protobuf.GetWebhookResponse
	10, // 17: protobuf.WebhookService.SendWebhook:output_type -> protobuf.SendWebhookResponse
	13, // [13:18] is the sub-list for method output_type
	8,  // [8:13] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_webhook_proto_init() }
func file_webhook_proto_init() {
	if File_webhook_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_webhook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckRequest); i {
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
		file_webhook_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HealthCheckResponse); i {
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
		file_webhook_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWebhooksRequest); i {
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
		file_webhook_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListWebhooksResponse); i {
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
		file_webhook_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWebhookRequest); i {
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
		file_webhook_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetWebhookResponse); i {
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
		file_webhook_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendWebhookRequest); i {
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
		file_webhook_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendWebhookResponse); i {
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
		file_webhook_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Webhook); i {
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
			RawDescriptor: file_webhook_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_webhook_proto_goTypes,
		DependencyIndexes: file_webhook_proto_depIdxs,
		EnumInfos:         file_webhook_proto_enumTypes,
		MessageInfos:      file_webhook_proto_msgTypes,
	}.Build()
	File_webhook_proto = out.File
	file_webhook_proto_rawDesc = nil
	file_webhook_proto_goTypes = nil
	file_webhook_proto_depIdxs = nil
}
