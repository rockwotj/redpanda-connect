// Copyright 2025 Redpanda Data, Inc.
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
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: runtime.proto

package runtimepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// `NullValue` is a representation of a null value.
type NullValue int32

const (
	NullValue_NULL_VALUE NullValue = 0
)

// Enum value maps for NullValue.
var (
	NullValue_name = map[int32]string{
		0: "NULL_VALUE",
	}
	NullValue_value = map[string]int32{
		"NULL_VALUE": 0,
	}
)

func (x NullValue) Enum() *NullValue {
	p := new(NullValue)
	*p = x
	return p
}

func (x NullValue) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (NullValue) Descriptor() protoreflect.EnumDescriptor {
	return file_runtime_proto_enumTypes[0].Descriptor()
}

func (NullValue) Type() protoreflect.EnumType {
	return &file_runtime_proto_enumTypes[0]
}

func (x NullValue) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use NullValue.Descriptor instead.
func (NullValue) EnumDescriptor() ([]byte, []int) {
	return file_runtime_proto_rawDescGZIP(), []int{0}
}

// `StructValue` represents a struct value which can be used to represent a
// structured data value.
type StructValue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Fields        map[string]*Value      `protobuf:"bytes,1,rep,name=fields,proto3" json:"fields,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *StructValue) Reset() {
	*x = StructValue{}
	mi := &file_runtime_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StructValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StructValue) ProtoMessage() {}

func (x *StructValue) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StructValue.ProtoReflect.Descriptor instead.
func (*StructValue) Descriptor() ([]byte, []int) {
	return file_runtime_proto_rawDescGZIP(), []int{0}
}

func (x *StructValue) GetFields() map[string]*Value {
	if x != nil {
		return x.Fields
	}
	return nil
}

// `ListValue` represents a list value which can be used to represent a list of
// values.
type ListValue struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Values        []*Value               `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListValue) Reset() {
	*x = ListValue{}
	mi := &file_runtime_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListValue) ProtoMessage() {}

func (x *ListValue) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListValue.ProtoReflect.Descriptor instead.
func (*ListValue) Descriptor() ([]byte, []int) {
	return file_runtime_proto_rawDescGZIP(), []int{1}
}

func (x *ListValue) GetValues() []*Value {
	if x != nil {
		return x.Values
	}
	return nil
}

// `Value` represents a dynamically typed value which can be used to represent
// a value passed to an agent.
type Value struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Kind:
	//
	//	*Value_NullValue
	//	*Value_StringValue
	//	*Value_IntegerValue
	//	*Value_DoubleValue
	//	*Value_BoolValue
	//	*Value_TimestampValue
	//	*Value_BytesValue
	//	*Value_StructValue
	//	*Value_ListValue
	Kind          isValue_Kind `protobuf_oneof:"kind"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Value) Reset() {
	*x = Value{}
	mi := &file_runtime_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_runtime_proto_rawDescGZIP(), []int{2}
}

func (x *Value) GetKind() isValue_Kind {
	if x != nil {
		return x.Kind
	}
	return nil
}

func (x *Value) GetNullValue() NullValue {
	if x != nil {
		if x, ok := x.Kind.(*Value_NullValue); ok {
			return x.NullValue
		}
	}
	return NullValue_NULL_VALUE
}

func (x *Value) GetStringValue() string {
	if x != nil {
		if x, ok := x.Kind.(*Value_StringValue); ok {
			return x.StringValue
		}
	}
	return ""
}

func (x *Value) GetIntegerValue() int64 {
	if x != nil {
		if x, ok := x.Kind.(*Value_IntegerValue); ok {
			return x.IntegerValue
		}
	}
	return 0
}

func (x *Value) GetDoubleValue() float64 {
	if x != nil {
		if x, ok := x.Kind.(*Value_DoubleValue); ok {
			return x.DoubleValue
		}
	}
	return 0
}

func (x *Value) GetBoolValue() bool {
	if x != nil {
		if x, ok := x.Kind.(*Value_BoolValue); ok {
			return x.BoolValue
		}
	}
	return false
}

func (x *Value) GetTimestampValue() *timestamppb.Timestamp {
	if x != nil {
		if x, ok := x.Kind.(*Value_TimestampValue); ok {
			return x.TimestampValue
		}
	}
	return nil
}

func (x *Value) GetBytesValue() []byte {
	if x != nil {
		if x, ok := x.Kind.(*Value_BytesValue); ok {
			return x.BytesValue
		}
	}
	return nil
}

func (x *Value) GetStructValue() *StructValue {
	if x != nil {
		if x, ok := x.Kind.(*Value_StructValue); ok {
			return x.StructValue
		}
	}
	return nil
}

func (x *Value) GetListValue() *ListValue {
	if x != nil {
		if x, ok := x.Kind.(*Value_ListValue); ok {
			return x.ListValue
		}
	}
	return nil
}

type isValue_Kind interface {
	isValue_Kind()
}

type Value_NullValue struct {
	NullValue NullValue `protobuf:"varint,1,opt,name=null_value,json=nullValue,proto3,enum=redpanda.runtime.v1alpha1.NullValue,oneof"`
}

type Value_StringValue struct {
	StringValue string `protobuf:"bytes,2,opt,name=string_value,json=stringValue,proto3,oneof"`
}

type Value_IntegerValue struct {
	IntegerValue int64 `protobuf:"varint,3,opt,name=integer_value,json=integerValue,proto3,oneof"`
}

type Value_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,4,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Value_BoolValue struct {
	BoolValue bool `protobuf:"varint,5,opt,name=bool_value,json=boolValue,proto3,oneof"`
}

type Value_TimestampValue struct {
	TimestampValue *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=timestamp_value,json=timestampValue,proto3,oneof"`
}

type Value_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,7,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

type Value_StructValue struct {
	StructValue *StructValue `protobuf:"bytes,8,opt,name=struct_value,json=structValue,proto3,oneof"`
}

type Value_ListValue struct {
	ListValue *ListValue `protobuf:"bytes,9,opt,name=list_value,json=listValue,proto3,oneof"`
}

func (*Value_NullValue) isValue_Kind() {}

func (*Value_StringValue) isValue_Kind() {}

func (*Value_IntegerValue) isValue_Kind() {}

func (*Value_DoubleValue) isValue_Kind() {}

func (*Value_BoolValue) isValue_Kind() {}

func (*Value_TimestampValue) isValue_Kind() {}

func (*Value_BytesValue) isValue_Kind() {}

func (*Value_StructValue) isValue_Kind() {}

func (*Value_ListValue) isValue_Kind() {}

// Message represents a piece of structured data that flows through the runtime.
type Message struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Payload:
	//
	//	*Message_Serialized
	//	*Message_Structured
	Payload       isMessage_Payload `protobuf_oneof:"payload"`
	Metadata      *StructValue      `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Message) Reset() {
	*x = Message{}
	mi := &file_runtime_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_runtime_proto_rawDescGZIP(), []int{3}
}

func (x *Message) GetPayload() isMessage_Payload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Message) GetSerialized() []byte {
	if x != nil {
		if x, ok := x.Payload.(*Message_Serialized); ok {
			return x.Serialized
		}
	}
	return nil
}

func (x *Message) GetStructured() *Value {
	if x != nil {
		if x, ok := x.Payload.(*Message_Structured); ok {
			return x.Structured
		}
	}
	return nil
}

func (x *Message) GetMetadata() *StructValue {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type isMessage_Payload interface {
	isMessage_Payload()
}

type Message_Serialized struct {
	Serialized []byte `protobuf:"bytes,1,opt,name=serialized,proto3,oneof"`
}

type Message_Structured struct {
	Structured *Value `protobuf:"bytes,2,opt,name=structured,proto3,oneof"`
}

func (*Message_Serialized) isMessage_Payload() {}

func (*Message_Structured) isMessage_Payload() {}

type TraceContext struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	TraceId       string                 `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	SpanId        string                 `protobuf:"bytes,2,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	TraceFlags    string                 `protobuf:"bytes,4,opt,name=trace_flags,json=traceFlags,proto3" json:"trace_flags,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TraceContext) Reset() {
	*x = TraceContext{}
	mi := &file_runtime_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TraceContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TraceContext) ProtoMessage() {}

func (x *TraceContext) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TraceContext.ProtoReflect.Descriptor instead.
func (*TraceContext) Descriptor() ([]byte, []int) {
	return file_runtime_proto_rawDescGZIP(), []int{4}
}

func (x *TraceContext) GetTraceId() string {
	if x != nil {
		return x.TraceId
	}
	return ""
}

func (x *TraceContext) GetSpanId() string {
	if x != nil {
		return x.SpanId
	}
	return ""
}

func (x *TraceContext) GetTraceFlags() string {
	if x != nil {
		return x.TraceFlags
	}
	return ""
}

type Trace struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Spans         []*Span                `protobuf:"bytes,1,rep,name=spans,proto3" json:"spans,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Trace) Reset() {
	*x = Trace{}
	mi := &file_runtime_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Trace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Trace) ProtoMessage() {}

func (x *Trace) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Trace.ProtoReflect.Descriptor instead.
func (*Trace) Descriptor() ([]byte, []int) {
	return file_runtime_proto_rawDescGZIP(), []int{5}
}

func (x *Trace) GetSpans() []*Span {
	if x != nil {
		return x.Spans
	}
	return nil
}

type Span struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	SpanId        string                 `protobuf:"bytes,1,opt,name=span_id,json=spanId,proto3" json:"span_id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	StartTime     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	EndTime       *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=end_time,json=endTime,proto3" json:"end_time,omitempty"`
	Attributes    map[string]*Value      `protobuf:"bytes,5,rep,name=attributes,proto3" json:"attributes,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ChildSpans    []*Span                `protobuf:"bytes,6,rep,name=child_spans,json=childSpans,proto3" json:"child_spans,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Span) Reset() {
	*x = Span{}
	mi := &file_runtime_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Span) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Span) ProtoMessage() {}

func (x *Span) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Span.ProtoReflect.Descriptor instead.
func (*Span) Descriptor() ([]byte, []int) {
	return file_runtime_proto_rawDescGZIP(), []int{6}
}

func (x *Span) GetSpanId() string {
	if x != nil {
		return x.SpanId
	}
	return ""
}

func (x *Span) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Span) GetStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.StartTime
	}
	return nil
}

func (x *Span) GetEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.EndTime
	}
	return nil
}

func (x *Span) GetAttributes() map[string]*Value {
	if x != nil {
		return x.Attributes
	}
	return nil
}

func (x *Span) GetChildSpans() []*Span {
	if x != nil {
		return x.ChildSpans
	}
	return nil
}

// InvokeAgentRequest is the request message for the `InvokeAgent` method.
type InvokeAgentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       *Message               `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	TraceContext  *TraceContext          `protobuf:"bytes,2,opt,name=trace_context,json=traceContext,proto3" json:"trace_context,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InvokeAgentRequest) Reset() {
	*x = InvokeAgentRequest{}
	mi := &file_runtime_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InvokeAgentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeAgentRequest) ProtoMessage() {}

func (x *InvokeAgentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeAgentRequest.ProtoReflect.Descriptor instead.
func (*InvokeAgentRequest) Descriptor() ([]byte, []int) {
	return file_runtime_proto_rawDescGZIP(), []int{7}
}

func (x *InvokeAgentRequest) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *InvokeAgentRequest) GetTraceContext() *TraceContext {
	if x != nil {
		return x.TraceContext
	}
	return nil
}

// InvokeAgentResponse is the response message for the `InvokeAgent` method.
type InvokeAgentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Message       *Message               `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	Trace         *Trace                 `protobuf:"bytes,2,opt,name=trace,proto3" json:"trace,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *InvokeAgentResponse) Reset() {
	*x = InvokeAgentResponse{}
	mi := &file_runtime_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *InvokeAgentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InvokeAgentResponse) ProtoMessage() {}

func (x *InvokeAgentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_runtime_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InvokeAgentResponse.ProtoReflect.Descriptor instead.
func (*InvokeAgentResponse) Descriptor() ([]byte, []int) {
	return file_runtime_proto_rawDescGZIP(), []int{8}
}

func (x *InvokeAgentResponse) GetMessage() *Message {
	if x != nil {
		return x.Message
	}
	return nil
}

func (x *InvokeAgentResponse) GetTrace() *Trace {
	if x != nil {
		return x.Trace
	}
	return nil
}

var File_runtime_proto protoreflect.FileDescriptor

const file_runtime_proto_rawDesc = "" +
	"\n" +
	"\rruntime.proto\x12\x19redpanda.runtime.v1alpha1\x1a\x1fgoogle/protobuf/timestamp.proto\"\xb6\x01\n" +
	"\vStructValue\x12J\n" +
	"\x06fields\x18\x01 \x03(\v22.redpanda.runtime.v1alpha1.StructValue.FieldsEntryR\x06fields\x1a[\n" +
	"\vFieldsEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x126\n" +
	"\x05value\x18\x02 \x01(\v2 .redpanda.runtime.v1alpha1.ValueR\x05value:\x028\x01\"E\n" +
	"\tListValue\x128\n" +
	"\x06values\x18\x01 \x03(\v2 .redpanda.runtime.v1alpha1.ValueR\x06values\"\xe6\x03\n" +
	"\x05Value\x12E\n" +
	"\n" +
	"null_value\x18\x01 \x01(\x0e2$.redpanda.runtime.v1alpha1.NullValueH\x00R\tnullValue\x12#\n" +
	"\fstring_value\x18\x02 \x01(\tH\x00R\vstringValue\x12%\n" +
	"\rinteger_value\x18\x03 \x01(\x03H\x00R\fintegerValue\x12#\n" +
	"\fdouble_value\x18\x04 \x01(\x01H\x00R\vdoubleValue\x12\x1f\n" +
	"\n" +
	"bool_value\x18\x05 \x01(\bH\x00R\tboolValue\x12E\n" +
	"\x0ftimestamp_value\x18\x06 \x01(\v2\x1a.google.protobuf.TimestampH\x00R\x0etimestampValue\x12!\n" +
	"\vbytes_value\x18\a \x01(\fH\x00R\n" +
	"bytesValue\x12K\n" +
	"\fstruct_value\x18\b \x01(\v2&.redpanda.runtime.v1alpha1.StructValueH\x00R\vstructValue\x12E\n" +
	"\n" +
	"list_value\x18\t \x01(\v2$.redpanda.runtime.v1alpha1.ListValueH\x00R\tlistValueB\x06\n" +
	"\x04kind\"\xbe\x01\n" +
	"\aMessage\x12 \n" +
	"\n" +
	"serialized\x18\x01 \x01(\fH\x00R\n" +
	"serialized\x12B\n" +
	"\n" +
	"structured\x18\x02 \x01(\v2 .redpanda.runtime.v1alpha1.ValueH\x00R\n" +
	"structured\x12B\n" +
	"\bmetadata\x18\x03 \x01(\v2&.redpanda.runtime.v1alpha1.StructValueR\bmetadataB\t\n" +
	"\apayload\"c\n" +
	"\fTraceContext\x12\x19\n" +
	"\btrace_id\x18\x01 \x01(\tR\atraceId\x12\x17\n" +
	"\aspan_id\x18\x02 \x01(\tR\x06spanId\x12\x1f\n" +
	"\vtrace_flags\x18\x04 \x01(\tR\n" +
	"traceFlags\">\n" +
	"\x05Trace\x125\n" +
	"\x05spans\x18\x01 \x03(\v2\x1f.redpanda.runtime.v1alpha1.SpanR\x05spans\"\x99\x03\n" +
	"\x04Span\x12\x17\n" +
	"\aspan_id\x18\x01 \x01(\tR\x06spanId\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x129\n" +
	"\n" +
	"start_time\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tstartTime\x125\n" +
	"\bend_time\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\aendTime\x12O\n" +
	"\n" +
	"attributes\x18\x05 \x03(\v2/.redpanda.runtime.v1alpha1.Span.AttributesEntryR\n" +
	"attributes\x12@\n" +
	"\vchild_spans\x18\x06 \x03(\v2\x1f.redpanda.runtime.v1alpha1.SpanR\n" +
	"childSpans\x1a_\n" +
	"\x0fAttributesEntry\x12\x10\n" +
	"\x03key\x18\x01 \x01(\tR\x03key\x126\n" +
	"\x05value\x18\x02 \x01(\v2 .redpanda.runtime.v1alpha1.ValueR\x05value:\x028\x01\"\xa0\x01\n" +
	"\x12InvokeAgentRequest\x12<\n" +
	"\amessage\x18\x01 \x01(\v2\".redpanda.runtime.v1alpha1.MessageR\amessage\x12L\n" +
	"\rtrace_context\x18\x02 \x01(\v2'.redpanda.runtime.v1alpha1.TraceContextR\ftraceContext\"\x8b\x01\n" +
	"\x13InvokeAgentResponse\x12<\n" +
	"\amessage\x18\x01 \x01(\v2\".redpanda.runtime.v1alpha1.MessageR\amessage\x126\n" +
	"\x05trace\x18\x02 \x01(\v2 .redpanda.runtime.v1alpha1.TraceR\x05trace*\x1b\n" +
	"\tNullValue\x12\x0e\n" +
	"\n" +
	"NULL_VALUE\x10\x002w\n" +
	"\aRuntime\x12l\n" +
	"\vInvokeAgent\x12-.redpanda.runtime.v1alpha1.InvokeAgentRequest\x1a..redpanda.runtime.v1alpha1.InvokeAgentResponseB\x1aZ\x18internal/agent/runtimepbb\x06proto3"

var (
	file_runtime_proto_rawDescOnce sync.Once
	file_runtime_proto_rawDescData []byte
)

func file_runtime_proto_rawDescGZIP() []byte {
	file_runtime_proto_rawDescOnce.Do(func() {
		file_runtime_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_runtime_proto_rawDesc), len(file_runtime_proto_rawDesc)))
	})
	return file_runtime_proto_rawDescData
}

var file_runtime_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_runtime_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_runtime_proto_goTypes = []any{
	(NullValue)(0),                // 0: redpanda.runtime.v1alpha1.NullValue
	(*StructValue)(nil),           // 1: redpanda.runtime.v1alpha1.StructValue
	(*ListValue)(nil),             // 2: redpanda.runtime.v1alpha1.ListValue
	(*Value)(nil),                 // 3: redpanda.runtime.v1alpha1.Value
	(*Message)(nil),               // 4: redpanda.runtime.v1alpha1.Message
	(*TraceContext)(nil),          // 5: redpanda.runtime.v1alpha1.TraceContext
	(*Trace)(nil),                 // 6: redpanda.runtime.v1alpha1.Trace
	(*Span)(nil),                  // 7: redpanda.runtime.v1alpha1.Span
	(*InvokeAgentRequest)(nil),    // 8: redpanda.runtime.v1alpha1.InvokeAgentRequest
	(*InvokeAgentResponse)(nil),   // 9: redpanda.runtime.v1alpha1.InvokeAgentResponse
	nil,                           // 10: redpanda.runtime.v1alpha1.StructValue.FieldsEntry
	nil,                           // 11: redpanda.runtime.v1alpha1.Span.AttributesEntry
	(*timestamppb.Timestamp)(nil), // 12: google.protobuf.Timestamp
}
var file_runtime_proto_depIdxs = []int32{
	10, // 0: redpanda.runtime.v1alpha1.StructValue.fields:type_name -> redpanda.runtime.v1alpha1.StructValue.FieldsEntry
	3,  // 1: redpanda.runtime.v1alpha1.ListValue.values:type_name -> redpanda.runtime.v1alpha1.Value
	0,  // 2: redpanda.runtime.v1alpha1.Value.null_value:type_name -> redpanda.runtime.v1alpha1.NullValue
	12, // 3: redpanda.runtime.v1alpha1.Value.timestamp_value:type_name -> google.protobuf.Timestamp
	1,  // 4: redpanda.runtime.v1alpha1.Value.struct_value:type_name -> redpanda.runtime.v1alpha1.StructValue
	2,  // 5: redpanda.runtime.v1alpha1.Value.list_value:type_name -> redpanda.runtime.v1alpha1.ListValue
	3,  // 6: redpanda.runtime.v1alpha1.Message.structured:type_name -> redpanda.runtime.v1alpha1.Value
	1,  // 7: redpanda.runtime.v1alpha1.Message.metadata:type_name -> redpanda.runtime.v1alpha1.StructValue
	7,  // 8: redpanda.runtime.v1alpha1.Trace.spans:type_name -> redpanda.runtime.v1alpha1.Span
	12, // 9: redpanda.runtime.v1alpha1.Span.start_time:type_name -> google.protobuf.Timestamp
	12, // 10: redpanda.runtime.v1alpha1.Span.end_time:type_name -> google.protobuf.Timestamp
	11, // 11: redpanda.runtime.v1alpha1.Span.attributes:type_name -> redpanda.runtime.v1alpha1.Span.AttributesEntry
	7,  // 12: redpanda.runtime.v1alpha1.Span.child_spans:type_name -> redpanda.runtime.v1alpha1.Span
	4,  // 13: redpanda.runtime.v1alpha1.InvokeAgentRequest.message:type_name -> redpanda.runtime.v1alpha1.Message
	5,  // 14: redpanda.runtime.v1alpha1.InvokeAgentRequest.trace_context:type_name -> redpanda.runtime.v1alpha1.TraceContext
	4,  // 15: redpanda.runtime.v1alpha1.InvokeAgentResponse.message:type_name -> redpanda.runtime.v1alpha1.Message
	6,  // 16: redpanda.runtime.v1alpha1.InvokeAgentResponse.trace:type_name -> redpanda.runtime.v1alpha1.Trace
	3,  // 17: redpanda.runtime.v1alpha1.StructValue.FieldsEntry.value:type_name -> redpanda.runtime.v1alpha1.Value
	3,  // 18: redpanda.runtime.v1alpha1.Span.AttributesEntry.value:type_name -> redpanda.runtime.v1alpha1.Value
	8,  // 19: redpanda.runtime.v1alpha1.Runtime.InvokeAgent:input_type -> redpanda.runtime.v1alpha1.InvokeAgentRequest
	9,  // 20: redpanda.runtime.v1alpha1.Runtime.InvokeAgent:output_type -> redpanda.runtime.v1alpha1.InvokeAgentResponse
	20, // [20:21] is the sub-list for method output_type
	19, // [19:20] is the sub-list for method input_type
	19, // [19:19] is the sub-list for extension type_name
	19, // [19:19] is the sub-list for extension extendee
	0,  // [0:19] is the sub-list for field type_name
}

func init() { file_runtime_proto_init() }
func file_runtime_proto_init() {
	if File_runtime_proto != nil {
		return
	}
	file_runtime_proto_msgTypes[2].OneofWrappers = []any{
		(*Value_NullValue)(nil),
		(*Value_StringValue)(nil),
		(*Value_IntegerValue)(nil),
		(*Value_DoubleValue)(nil),
		(*Value_BoolValue)(nil),
		(*Value_TimestampValue)(nil),
		(*Value_BytesValue)(nil),
		(*Value_StructValue)(nil),
		(*Value_ListValue)(nil),
	}
	file_runtime_proto_msgTypes[3].OneofWrappers = []any{
		(*Message_Serialized)(nil),
		(*Message_Structured)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_runtime_proto_rawDesc), len(file_runtime_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_runtime_proto_goTypes,
		DependencyIndexes: file_runtime_proto_depIdxs,
		EnumInfos:         file_runtime_proto_enumTypes,
		MessageInfos:      file_runtime_proto_msgTypes,
	}.Build()
	File_runtime_proto = out.File
	file_runtime_proto_goTypes = nil
	file_runtime_proto_depIdxs = nil
}
