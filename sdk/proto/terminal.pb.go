// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.11.2
// source: terminal.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type TerminalUI struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *TerminalUI) Reset() {
	*x = TerminalUI{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terminal_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminalUI) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminalUI) ProtoMessage() {}

func (x *TerminalUI) ProtoReflect() protoreflect.Message {
	mi := &file_terminal_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminalUI.ProtoReflect.Descriptor instead.
func (*TerminalUI) Descriptor() ([]byte, []int) {
	return file_terminal_proto_rawDescGZIP(), []int{0}
}

type TerminalUI_OutputRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Lines []string `protobuf:"bytes,1,rep,name=lines,proto3" json:"lines,omitempty"`
}

func (x *TerminalUI_OutputRequest) Reset() {
	*x = TerminalUI_OutputRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terminal_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminalUI_OutputRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminalUI_OutputRequest) ProtoMessage() {}

func (x *TerminalUI_OutputRequest) ProtoReflect() protoreflect.Message {
	mi := &file_terminal_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminalUI_OutputRequest.ProtoReflect.Descriptor instead.
func (*TerminalUI_OutputRequest) Descriptor() ([]byte, []int) {
	return file_terminal_proto_rawDescGZIP(), []int{0, 0}
}

func (x *TerminalUI_OutputRequest) GetLines() []string {
	if x != nil {
		return x.Lines
	}
	return nil
}

type TerminalUI_Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Event:
	//	*TerminalUI_Event_Line_
	//	*TerminalUI_Event_Status_
	//	*TerminalUI_Event_NamedValues_
	//	*TerminalUI_Event_Raw_
	//	*TerminalUI_Event_Table_
	Event isTerminalUI_Event_Event `protobuf_oneof:"event"`
}

func (x *TerminalUI_Event) Reset() {
	*x = TerminalUI_Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terminal_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminalUI_Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminalUI_Event) ProtoMessage() {}

func (x *TerminalUI_Event) ProtoReflect() protoreflect.Message {
	mi := &file_terminal_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminalUI_Event.ProtoReflect.Descriptor instead.
func (*TerminalUI_Event) Descriptor() ([]byte, []int) {
	return file_terminal_proto_rawDescGZIP(), []int{0, 1}
}

func (m *TerminalUI_Event) GetEvent() isTerminalUI_Event_Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func (x *TerminalUI_Event) GetLine() *TerminalUI_Event_Line {
	if x, ok := x.GetEvent().(*TerminalUI_Event_Line_); ok {
		return x.Line
	}
	return nil
}

func (x *TerminalUI_Event) GetStatus() *TerminalUI_Event_Status {
	if x, ok := x.GetEvent().(*TerminalUI_Event_Status_); ok {
		return x.Status
	}
	return nil
}

func (x *TerminalUI_Event) GetNamedValues() *TerminalUI_Event_NamedValues {
	if x, ok := x.GetEvent().(*TerminalUI_Event_NamedValues_); ok {
		return x.NamedValues
	}
	return nil
}

func (x *TerminalUI_Event) GetRaw() *TerminalUI_Event_Raw {
	if x, ok := x.GetEvent().(*TerminalUI_Event_Raw_); ok {
		return x.Raw
	}
	return nil
}

func (x *TerminalUI_Event) GetTable() *TerminalUI_Event_Table {
	if x, ok := x.GetEvent().(*TerminalUI_Event_Table_); ok {
		return x.Table
	}
	return nil
}

type isTerminalUI_Event_Event interface {
	isTerminalUI_Event_Event()
}

type TerminalUI_Event_Line_ struct {
	Line *TerminalUI_Event_Line `protobuf:"bytes,1,opt,name=line,proto3,oneof"`
}

type TerminalUI_Event_Status_ struct {
	Status *TerminalUI_Event_Status `protobuf:"bytes,2,opt,name=status,proto3,oneof"`
}

type TerminalUI_Event_NamedValues_ struct {
	NamedValues *TerminalUI_Event_NamedValues `protobuf:"bytes,3,opt,name=named_values,json=namedValues,proto3,oneof"`
}

type TerminalUI_Event_Raw_ struct {
	Raw *TerminalUI_Event_Raw `protobuf:"bytes,4,opt,name=raw,proto3,oneof"`
}

type TerminalUI_Event_Table_ struct {
	Table *TerminalUI_Event_Table `protobuf:"bytes,5,opt,name=table,proto3,oneof"`
}

func (*TerminalUI_Event_Line_) isTerminalUI_Event_Event() {}

func (*TerminalUI_Event_Status_) isTerminalUI_Event_Event() {}

func (*TerminalUI_Event_NamedValues_) isTerminalUI_Event_Event() {}

func (*TerminalUI_Event_Raw_) isTerminalUI_Event_Event() {}

func (*TerminalUI_Event_Table_) isTerminalUI_Event_Event() {}

type TerminalUI_Event_Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg    string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Step   bool   `protobuf:"varint,3,opt,name=step,proto3" json:"step,omitempty"`
}

func (x *TerminalUI_Event_Status) Reset() {
	*x = TerminalUI_Event_Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terminal_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminalUI_Event_Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminalUI_Event_Status) ProtoMessage() {}

func (x *TerminalUI_Event_Status) ProtoReflect() protoreflect.Message {
	mi := &file_terminal_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminalUI_Event_Status.ProtoReflect.Descriptor instead.
func (*TerminalUI_Event_Status) Descriptor() ([]byte, []int) {
	return file_terminal_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *TerminalUI_Event_Status) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TerminalUI_Event_Status) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *TerminalUI_Event_Status) GetStep() bool {
	if x != nil {
		return x.Step
	}
	return false
}

type TerminalUI_Event_Line struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg   string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
	Style string `protobuf:"bytes,2,opt,name=style,proto3" json:"style,omitempty"`
}

func (x *TerminalUI_Event_Line) Reset() {
	*x = TerminalUI_Event_Line{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terminal_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminalUI_Event_Line) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminalUI_Event_Line) ProtoMessage() {}

func (x *TerminalUI_Event_Line) ProtoReflect() protoreflect.Message {
	mi := &file_terminal_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminalUI_Event_Line.ProtoReflect.Descriptor instead.
func (*TerminalUI_Event_Line) Descriptor() ([]byte, []int) {
	return file_terminal_proto_rawDescGZIP(), []int{0, 1, 1}
}

func (x *TerminalUI_Event_Line) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *TerminalUI_Event_Line) GetStyle() string {
	if x != nil {
		return x.Style
	}
	return ""
}

type TerminalUI_Event_Raw struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data   []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Stderr bool   `protobuf:"varint,2,opt,name=stderr,proto3" json:"stderr,omitempty"`
}

func (x *TerminalUI_Event_Raw) Reset() {
	*x = TerminalUI_Event_Raw{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terminal_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminalUI_Event_Raw) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminalUI_Event_Raw) ProtoMessage() {}

func (x *TerminalUI_Event_Raw) ProtoReflect() protoreflect.Message {
	mi := &file_terminal_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminalUI_Event_Raw.ProtoReflect.Descriptor instead.
func (*TerminalUI_Event_Raw) Descriptor() ([]byte, []int) {
	return file_terminal_proto_rawDescGZIP(), []int{0, 1, 2}
}

func (x *TerminalUI_Event_Raw) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TerminalUI_Event_Raw) GetStderr() bool {
	if x != nil {
		return x.Stderr
	}
	return false
}

type TerminalUI_Event_NamedValue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *TerminalUI_Event_NamedValue) Reset() {
	*x = TerminalUI_Event_NamedValue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terminal_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminalUI_Event_NamedValue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminalUI_Event_NamedValue) ProtoMessage() {}

func (x *TerminalUI_Event_NamedValue) ProtoReflect() protoreflect.Message {
	mi := &file_terminal_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminalUI_Event_NamedValue.ProtoReflect.Descriptor instead.
func (*TerminalUI_Event_NamedValue) Descriptor() ([]byte, []int) {
	return file_terminal_proto_rawDescGZIP(), []int{0, 1, 3}
}

func (x *TerminalUI_Event_NamedValue) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *TerminalUI_Event_NamedValue) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type TerminalUI_Event_NamedValues struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []*TerminalUI_Event_NamedValue `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *TerminalUI_Event_NamedValues) Reset() {
	*x = TerminalUI_Event_NamedValues{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terminal_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminalUI_Event_NamedValues) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminalUI_Event_NamedValues) ProtoMessage() {}

func (x *TerminalUI_Event_NamedValues) ProtoReflect() protoreflect.Message {
	mi := &file_terminal_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminalUI_Event_NamedValues.ProtoReflect.Descriptor instead.
func (*TerminalUI_Event_NamedValues) Descriptor() ([]byte, []int) {
	return file_terminal_proto_rawDescGZIP(), []int{0, 1, 4}
}

func (x *TerminalUI_Event_NamedValues) GetValues() []*TerminalUI_Event_NamedValue {
	if x != nil {
		return x.Values
	}
	return nil
}

type TerminalUI_Event_TableEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	Color string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *TerminalUI_Event_TableEntry) Reset() {
	*x = TerminalUI_Event_TableEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terminal_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminalUI_Event_TableEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminalUI_Event_TableEntry) ProtoMessage() {}

func (x *TerminalUI_Event_TableEntry) ProtoReflect() protoreflect.Message {
	mi := &file_terminal_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminalUI_Event_TableEntry.ProtoReflect.Descriptor instead.
func (*TerminalUI_Event_TableEntry) Descriptor() ([]byte, []int) {
	return file_terminal_proto_rawDescGZIP(), []int{0, 1, 5}
}

func (x *TerminalUI_Event_TableEntry) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *TerminalUI_Event_TableEntry) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type TerminalUI_Event_TableRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entries []*TerminalUI_Event_TableEntry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
}

func (x *TerminalUI_Event_TableRow) Reset() {
	*x = TerminalUI_Event_TableRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terminal_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminalUI_Event_TableRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminalUI_Event_TableRow) ProtoMessage() {}

func (x *TerminalUI_Event_TableRow) ProtoReflect() protoreflect.Message {
	mi := &file_terminal_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminalUI_Event_TableRow.ProtoReflect.Descriptor instead.
func (*TerminalUI_Event_TableRow) Descriptor() ([]byte, []int) {
	return file_terminal_proto_rawDescGZIP(), []int{0, 1, 6}
}

func (x *TerminalUI_Event_TableRow) GetEntries() []*TerminalUI_Event_TableEntry {
	if x != nil {
		return x.Entries
	}
	return nil
}

type TerminalUI_Event_Table struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Headers []string                     `protobuf:"bytes,1,rep,name=headers,proto3" json:"headers,omitempty"`
	Rows    []*TerminalUI_Event_TableRow `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows,omitempty"`
}

func (x *TerminalUI_Event_Table) Reset() {
	*x = TerminalUI_Event_Table{}
	if protoimpl.UnsafeEnabled {
		mi := &file_terminal_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TerminalUI_Event_Table) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TerminalUI_Event_Table) ProtoMessage() {}

func (x *TerminalUI_Event_Table) ProtoReflect() protoreflect.Message {
	mi := &file_terminal_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TerminalUI_Event_Table.ProtoReflect.Descriptor instead.
func (*TerminalUI_Event_Table) Descriptor() ([]byte, []int) {
	return file_terminal_proto_rawDescGZIP(), []int{0, 1, 7}
}

func (x *TerminalUI_Event_Table) GetHeaders() []string {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *TerminalUI_Event_Table) GetRows() []*TerminalUI_Event_TableRow {
	if x != nil {
		return x.Rows
	}
	return nil
}

var File_terminal_proto protoreflect.FileDescriptor

var file_terminal_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf1, 0x06, 0x0a, 0x0a, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61,
	0x6c, 0x55, 0x49, 0x1a, 0x25, 0x0a, 0x0d, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x1a, 0xbb, 0x06, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x12, 0x32, 0x0a, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x61, 0x6c, 0x55, 0x49, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x4c, 0x69, 0x6e, 0x65,
	0x48, 0x00, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x38, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x49, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x48, 0x00, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x48, 0x0a, 0x0c, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x5f, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x49, 0x2e, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x48, 0x00, 0x52,
	0x0b, 0x6e, 0x61, 0x6d, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x03,
	0x72, 0x61, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x49, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x52, 0x61, 0x77, 0x48, 0x00, 0x52, 0x03, 0x72, 0x61, 0x77, 0x12, 0x35, 0x0a,
	0x05, 0x74, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x49, 0x2e,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x48, 0x00, 0x52, 0x05, 0x74,
	0x61, 0x62, 0x6c, 0x65, 0x1a, 0x46, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x73, 0x74, 0x65, 0x70,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x04, 0x73, 0x74, 0x65, 0x70, 0x1a, 0x2e, 0x0a, 0x04,
	0x4c, 0x69, 0x6e, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x1a, 0x31, 0x0a, 0x03,
	0x52, 0x61, 0x77, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x64, 0x65, 0x72, 0x72, 0x1a,
	0x36, 0x0a, 0x0a, 0x4e, 0x61, 0x6d, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x49, 0x0a, 0x0b, 0x4e, 0x61, 0x6d, 0x65, 0x64,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x3a, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54,
	0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x49, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e,
	0x4e, 0x61, 0x6d, 0x65, 0x64, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x73, 0x1a, 0x38, 0x0a, 0x0a, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x1a, 0x48, 0x0a, 0x08,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x77, 0x12, 0x3c, 0x0a, 0x07, 0x65, 0x6e, 0x74, 0x72,
	0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x49, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x65,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x1a, 0x57, 0x0a, 0x05, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x34, 0x0a, 0x04, 0x72, 0x6f, 0x77,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x49, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x6f, 0x77, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x42,
	0x07, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x32, 0x93, 0x01, 0x0a, 0x11, 0x54, 0x65, 0x72,
	0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x49, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41,
	0x0a, 0x06, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x49, 0x2e, 0x4f, 0x75, 0x74, 0x70,
	0x75, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x12, 0x3b, 0x0a, 0x06, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x17, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x6c, 0x55, 0x49, 0x2e, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x28, 0x01, 0x42, 0x09,
	0x5a, 0x07, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_terminal_proto_rawDescOnce sync.Once
	file_terminal_proto_rawDescData = file_terminal_proto_rawDesc
)

func file_terminal_proto_rawDescGZIP() []byte {
	file_terminal_proto_rawDescOnce.Do(func() {
		file_terminal_proto_rawDescData = protoimpl.X.CompressGZIP(file_terminal_proto_rawDescData)
	})
	return file_terminal_proto_rawDescData
}

var file_terminal_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_terminal_proto_goTypes = []interface{}{
	(*TerminalUI)(nil),                   // 0: proto.TerminalUI
	(*TerminalUI_OutputRequest)(nil),     // 1: proto.TerminalUI.OutputRequest
	(*TerminalUI_Event)(nil),             // 2: proto.TerminalUI.Event
	(*TerminalUI_Event_Status)(nil),      // 3: proto.TerminalUI.Event.Status
	(*TerminalUI_Event_Line)(nil),        // 4: proto.TerminalUI.Event.Line
	(*TerminalUI_Event_Raw)(nil),         // 5: proto.TerminalUI.Event.Raw
	(*TerminalUI_Event_NamedValue)(nil),  // 6: proto.TerminalUI.Event.NamedValue
	(*TerminalUI_Event_NamedValues)(nil), // 7: proto.TerminalUI.Event.NamedValues
	(*TerminalUI_Event_TableEntry)(nil),  // 8: proto.TerminalUI.Event.TableEntry
	(*TerminalUI_Event_TableRow)(nil),    // 9: proto.TerminalUI.Event.TableRow
	(*TerminalUI_Event_Table)(nil),       // 10: proto.TerminalUI.Event.Table
	(*empty.Empty)(nil),                  // 11: google.protobuf.Empty
}
var file_terminal_proto_depIdxs = []int32{
	4,  // 0: proto.TerminalUI.Event.line:type_name -> proto.TerminalUI.Event.Line
	3,  // 1: proto.TerminalUI.Event.status:type_name -> proto.TerminalUI.Event.Status
	7,  // 2: proto.TerminalUI.Event.named_values:type_name -> proto.TerminalUI.Event.NamedValues
	5,  // 3: proto.TerminalUI.Event.raw:type_name -> proto.TerminalUI.Event.Raw
	10, // 4: proto.TerminalUI.Event.table:type_name -> proto.TerminalUI.Event.Table
	6,  // 5: proto.TerminalUI.Event.NamedValues.values:type_name -> proto.TerminalUI.Event.NamedValue
	8,  // 6: proto.TerminalUI.Event.TableRow.entries:type_name -> proto.TerminalUI.Event.TableEntry
	9,  // 7: proto.TerminalUI.Event.Table.rows:type_name -> proto.TerminalUI.Event.TableRow
	1,  // 8: proto.TerminalUIService.Output:input_type -> proto.TerminalUI.OutputRequest
	2,  // 9: proto.TerminalUIService.Events:input_type -> proto.TerminalUI.Event
	11, // 10: proto.TerminalUIService.Output:output_type -> google.protobuf.Empty
	11, // 11: proto.TerminalUIService.Events:output_type -> google.protobuf.Empty
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_terminal_proto_init() }
func file_terminal_proto_init() {
	if File_terminal_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_terminal_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminalUI); i {
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
		file_terminal_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminalUI_OutputRequest); i {
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
		file_terminal_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminalUI_Event); i {
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
		file_terminal_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminalUI_Event_Status); i {
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
		file_terminal_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminalUI_Event_Line); i {
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
		file_terminal_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminalUI_Event_Raw); i {
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
		file_terminal_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminalUI_Event_NamedValue); i {
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
		file_terminal_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminalUI_Event_NamedValues); i {
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
		file_terminal_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminalUI_Event_TableEntry); i {
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
		file_terminal_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminalUI_Event_TableRow); i {
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
		file_terminal_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TerminalUI_Event_Table); i {
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
	file_terminal_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*TerminalUI_Event_Line_)(nil),
		(*TerminalUI_Event_Status_)(nil),
		(*TerminalUI_Event_NamedValues_)(nil),
		(*TerminalUI_Event_Raw_)(nil),
		(*TerminalUI_Event_Table_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_terminal_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_terminal_proto_goTypes,
		DependencyIndexes: file_terminal_proto_depIdxs,
		MessageInfos:      file_terminal_proto_msgTypes,
	}.Build()
	File_terminal_proto = out.File
	file_terminal_proto_rawDesc = nil
	file_terminal_proto_goTypes = nil
	file_terminal_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TerminalUIServiceClient is the client API for TerminalUIService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TerminalUIServiceClient interface {
	Output(ctx context.Context, in *TerminalUI_OutputRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	Events(ctx context.Context, opts ...grpc.CallOption) (TerminalUIService_EventsClient, error)
}

type terminalUIServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTerminalUIServiceClient(cc grpc.ClientConnInterface) TerminalUIServiceClient {
	return &terminalUIServiceClient{cc}
}

func (c *terminalUIServiceClient) Output(ctx context.Context, in *TerminalUI_OutputRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/proto.TerminalUIService/Output", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *terminalUIServiceClient) Events(ctx context.Context, opts ...grpc.CallOption) (TerminalUIService_EventsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TerminalUIService_serviceDesc.Streams[0], "/proto.TerminalUIService/Events", opts...)
	if err != nil {
		return nil, err
	}
	x := &terminalUIServiceEventsClient{stream}
	return x, nil
}

type TerminalUIService_EventsClient interface {
	Send(*TerminalUI_Event) error
	CloseAndRecv() (*empty.Empty, error)
	grpc.ClientStream
}

type terminalUIServiceEventsClient struct {
	grpc.ClientStream
}

func (x *terminalUIServiceEventsClient) Send(m *TerminalUI_Event) error {
	return x.ClientStream.SendMsg(m)
}

func (x *terminalUIServiceEventsClient) CloseAndRecv() (*empty.Empty, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(empty.Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TerminalUIServiceServer is the server API for TerminalUIService service.
type TerminalUIServiceServer interface {
	Output(context.Context, *TerminalUI_OutputRequest) (*empty.Empty, error)
	Events(TerminalUIService_EventsServer) error
}

// UnimplementedTerminalUIServiceServer can be embedded to have forward compatible implementations.
type UnimplementedTerminalUIServiceServer struct {
}

func (*UnimplementedTerminalUIServiceServer) Output(context.Context, *TerminalUI_OutputRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Output not implemented")
}
func (*UnimplementedTerminalUIServiceServer) Events(TerminalUIService_EventsServer) error {
	return status.Errorf(codes.Unimplemented, "method Events not implemented")
}

func RegisterTerminalUIServiceServer(s *grpc.Server, srv TerminalUIServiceServer) {
	s.RegisterService(&_TerminalUIService_serviceDesc, srv)
}

func _TerminalUIService_Output_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TerminalUI_OutputRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TerminalUIServiceServer).Output(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.TerminalUIService/Output",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TerminalUIServiceServer).Output(ctx, req.(*TerminalUI_OutputRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TerminalUIService_Events_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TerminalUIServiceServer).Events(&terminalUIServiceEventsServer{stream})
}

type TerminalUIService_EventsServer interface {
	SendAndClose(*empty.Empty) error
	Recv() (*TerminalUI_Event, error)
	grpc.ServerStream
}

type terminalUIServiceEventsServer struct {
	grpc.ServerStream
}

func (x *terminalUIServiceEventsServer) SendAndClose(m *empty.Empty) error {
	return x.ServerStream.SendMsg(m)
}

func (x *terminalUIServiceEventsServer) Recv() (*TerminalUI_Event, error) {
	m := new(TerminalUI_Event)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TerminalUIService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.TerminalUIService",
	HandlerType: (*TerminalUIServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Output",
			Handler:    _TerminalUIService_Output_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Events",
			Handler:       _TerminalUIService_Events_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "terminal.proto",
}
