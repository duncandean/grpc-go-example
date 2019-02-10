// Code generated by protoc-gen-go. DO NOT EDIT.
// source: todo.proto

package todo

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Task struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	Done                 bool     `protobuf:"varint,2,opt,name=done,proto3" json:"done,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Task) Reset()         { *m = Task{} }
func (m *Task) String() string { return proto.CompactTextString(m) }
func (*Task) ProtoMessage()    {}
func (*Task) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{0}
}

func (m *Task) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Task.Unmarshal(m, b)
}
func (m *Task) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Task.Marshal(b, m, deterministic)
}
func (m *Task) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Task.Merge(m, src)
}
func (m *Task) XXX_Size() int {
	return xxx_messageInfo_Task.Size(m)
}
func (m *Task) XXX_DiscardUnknown() {
	xxx_messageInfo_Task.DiscardUnknown(m)
}

var xxx_messageInfo_Task proto.InternalMessageInfo

func (m *Task) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Task) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

type TaskList struct {
	Tasks                []*Task  `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskList) Reset()         { *m = TaskList{} }
func (m *TaskList) String() string { return proto.CompactTextString(m) }
func (*TaskList) ProtoMessage()    {}
func (*TaskList) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{1}
}

func (m *TaskList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskList.Unmarshal(m, b)
}
func (m *TaskList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskList.Marshal(b, m, deterministic)
}
func (m *TaskList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskList.Merge(m, src)
}
func (m *TaskList) XXX_Size() int {
	return xxx_messageInfo_TaskList.Size(m)
}
func (m *TaskList) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskList.DiscardUnknown(m)
}

var xxx_messageInfo_TaskList proto.InternalMessageInfo

func (m *TaskList) GetTasks() []*Task {
	if m != nil {
		return m.Tasks
	}
	return nil
}

type Text struct {
	Text                 string   `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Text) Reset()         { *m = Text{} }
func (m *Text) String() string { return proto.CompactTextString(m) }
func (*Text) ProtoMessage()    {}
func (*Text) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{2}
}

func (m *Text) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Text.Unmarshal(m, b)
}
func (m *Text) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Text.Marshal(b, m, deterministic)
}
func (m *Text) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Text.Merge(m, src)
}
func (m *Text) XXX_Size() int {
	return xxx_messageInfo_Text.Size(m)
}
func (m *Text) XXX_DiscardUnknown() {
	xxx_messageInfo_Text.DiscardUnknown(m)
}

var xxx_messageInfo_Text proto.InternalMessageInfo

func (m *Text) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

type Void struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Void) Reset()         { *m = Void{} }
func (m *Void) String() string { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()    {}
func (*Void) Descriptor() ([]byte, []int) {
	return fileDescriptor_0e4b95d0c4e09639, []int{3}
}

func (m *Void) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Void.Unmarshal(m, b)
}
func (m *Void) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Void.Marshal(b, m, deterministic)
}
func (m *Void) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Void.Merge(m, src)
}
func (m *Void) XXX_Size() int {
	return xxx_messageInfo_Void.Size(m)
}
func (m *Void) XXX_DiscardUnknown() {
	xxx_messageInfo_Void.DiscardUnknown(m)
}

var xxx_messageInfo_Void proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Task)(nil), "todo.Task")
	proto.RegisterType((*TaskList)(nil), "todo.TaskList")
	proto.RegisterType((*Text)(nil), "todo.Text")
	proto.RegisterType((*Void)(nil), "todo.Void")
}

func init() { proto.RegisterFile("todo.proto", fileDescriptor_0e4b95d0c4e09639) }

var fileDescriptor_0e4b95d0c4e09639 = []byte{
	// 178 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x8f, 0xb1, 0xaa, 0xc2, 0x40,
	0x10, 0x45, 0xb3, 0x2f, 0x9b, 0x90, 0x77, 0x05, 0x8b, 0xad, 0x42, 0x1a, 0x97, 0xc5, 0x22, 0x85,
	0xa4, 0x88, 0x5f, 0x60, 0x2f, 0x16, 0x21, 0xd8, 0x47, 0x36, 0x45, 0x08, 0x38, 0xe2, 0x4e, 0xe1,
	0xe7, 0xcb, 0xac, 0x11, 0xc5, 0xee, 0x32, 0xe7, 0xce, 0x19, 0x06, 0x60, 0xf2, 0xd4, 0xdc, 0xee,
	0xc4, 0x64, 0xb4, 0x64, 0xd7, 0x40, 0xf7, 0x43, 0x98, 0x8d, 0x81, 0xe6, 0xf1, 0xc1, 0xa5, 0xb2,
	0xaa, 0xfe, 0xef, 0x62, 0x96, 0x99, 0xa7, 0xeb, 0x58, 0xfe, 0x59, 0x55, 0x17, 0x5d, 0xcc, 0x6e,
	0x87, 0x42, 0xfa, 0xc7, 0x29, 0xb0, 0xb1, 0xc8, 0x78, 0x08, 0x73, 0x28, 0x95, 0x4d, 0xeb, 0x55,
	0x8b, 0x26, 0xda, 0x05, 0x77, 0x2f, 0xe0, 0x2a, 0xe8, 0x7e, 0x31, 0xfd, 0xda, 0x5d, 0x0e, 0x7d,
	0xa6, 0xc9, 0xb7, 0x27, 0x64, 0xb2, 0x12, 0xcc, 0x16, 0x3a, 0x6a, 0x17, 0x8f, 0xc0, 0x6a, 0xfd,
	0x71, 0x0a, 0x73, 0x89, 0xd9, 0x20, 0x3d, 0x78, 0xff, 0x2e, 0x89, 0xbd, 0xfa, 0x3a, 0xec, 0x92,
	0x4b, 0x1e, 0xdf, 0xdb, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xd0, 0xad, 0xd6, 0x8c, 0xec, 0x00,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TasksClient is the client API for Tasks service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TasksClient interface {
	List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TaskList, error)
	Add(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Task, error)
}

type tasksClient struct {
	cc *grpc.ClientConn
}

func NewTasksClient(cc *grpc.ClientConn) TasksClient {
	return &tasksClient{cc}
}

func (c *tasksClient) List(ctx context.Context, in *Void, opts ...grpc.CallOption) (*TaskList, error) {
	out := new(TaskList)
	err := c.cc.Invoke(ctx, "/todo.Tasks/List", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tasksClient) Add(ctx context.Context, in *Text, opts ...grpc.CallOption) (*Task, error) {
	out := new(Task)
	err := c.cc.Invoke(ctx, "/todo.Tasks/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TasksServer is the server API for Tasks service.
type TasksServer interface {
	List(context.Context, *Void) (*TaskList, error)
	Add(context.Context, *Text) (*Task, error)
}

func RegisterTasksServer(s *grpc.Server, srv TasksServer) {
	s.RegisterService(&_Tasks_serviceDesc, srv)
}

func _Tasks_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Tasks/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).List(ctx, req.(*Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _Tasks_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Text)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TasksServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/todo.Tasks/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TasksServer).Add(ctx, req.(*Text))
	}
	return interceptor(ctx, in, info, handler)
}

var _Tasks_serviceDesc = grpc.ServiceDesc{
	ServiceName: "todo.Tasks",
	HandlerType: (*TasksServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "List",
			Handler:    _Tasks_List_Handler,
		},
		{
			MethodName: "Add",
			Handler:    _Tasks_Add_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "todo.proto",
}
