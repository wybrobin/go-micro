// Code generated by protoc-gen-go. DO NOT EDIT.
// source: router.proto

package go_micro_router

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// AdvertType defines the type of advert
type AdvertType int32

const (
	AdvertType_AdvertAnnounce AdvertType = 0
	AdvertType_AdvertUpdate   AdvertType = 1
)

var AdvertType_name = map[int32]string{
	0: "AdvertAnnounce",
	1: "AdvertUpdate",
}

var AdvertType_value = map[string]int32{
	"AdvertAnnounce": 0,
	"AdvertUpdate":   1,
}

func (x AdvertType) String() string {
	return proto.EnumName(AdvertType_name, int32(x))
}

func (AdvertType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{0}
}

// EventType defines the type of event
type EventType int32

const (
	EventType_Create EventType = 0
	EventType_Delete EventType = 1
	EventType_Update EventType = 2
)

var EventType_name = map[int32]string{
	0: "Create",
	1: "Delete",
	2: "Update",
}

var EventType_value = map[string]int32{
	"Create": 0,
	"Delete": 1,
	"Update": 2,
}

func (x EventType) String() string {
	return proto.EnumName(EventType_name, int32(x))
}

func (EventType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{1}
}

// Empty request
type Request struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

// Empty response
type Response struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

// ListResponse is returned by List
type ListResponse struct {
	Routes               []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{2}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

// LookupRequest is made to Lookup
type LookupRequest struct {
	Query                *Query   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupRequest) Reset()         { *m = LookupRequest{} }
func (m *LookupRequest) String() string { return proto.CompactTextString(m) }
func (*LookupRequest) ProtoMessage()    {}
func (*LookupRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{3}
}

func (m *LookupRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupRequest.Unmarshal(m, b)
}
func (m *LookupRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupRequest.Marshal(b, m, deterministic)
}
func (m *LookupRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupRequest.Merge(m, src)
}
func (m *LookupRequest) XXX_Size() int {
	return xxx_messageInfo_LookupRequest.Size(m)
}
func (m *LookupRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LookupRequest proto.InternalMessageInfo

func (m *LookupRequest) GetQuery() *Query {
	if m != nil {
		return m.Query
	}
	return nil
}

// LookupResponse is returned by Lookup
type LookupResponse struct {
	Routes               []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LookupResponse) Reset()         { *m = LookupResponse{} }
func (m *LookupResponse) String() string { return proto.CompactTextString(m) }
func (*LookupResponse) ProtoMessage()    {}
func (*LookupResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{4}
}

func (m *LookupResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LookupResponse.Unmarshal(m, b)
}
func (m *LookupResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LookupResponse.Marshal(b, m, deterministic)
}
func (m *LookupResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LookupResponse.Merge(m, src)
}
func (m *LookupResponse) XXX_Size() int {
	return xxx_messageInfo_LookupResponse.Size(m)
}
func (m *LookupResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LookupResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LookupResponse proto.InternalMessageInfo

func (m *LookupResponse) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

// QueryRequest queries Table for Routes
type QueryRequest struct {
	Query                *Query   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryRequest) Reset()         { *m = QueryRequest{} }
func (m *QueryRequest) String() string { return proto.CompactTextString(m) }
func (*QueryRequest) ProtoMessage()    {}
func (*QueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{5}
}

func (m *QueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryRequest.Unmarshal(m, b)
}
func (m *QueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryRequest.Marshal(b, m, deterministic)
}
func (m *QueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryRequest.Merge(m, src)
}
func (m *QueryRequest) XXX_Size() int {
	return xxx_messageInfo_QueryRequest.Size(m)
}
func (m *QueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryRequest proto.InternalMessageInfo

func (m *QueryRequest) GetQuery() *Query {
	if m != nil {
		return m.Query
	}
	return nil
}

// QueryResponse is returned by Query
type QueryResponse struct {
	Routes               []*Route `protobuf:"bytes,1,rep,name=routes,proto3" json:"routes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryResponse) Reset()         { *m = QueryResponse{} }
func (m *QueryResponse) String() string { return proto.CompactTextString(m) }
func (*QueryResponse) ProtoMessage()    {}
func (*QueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{6}
}

func (m *QueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryResponse.Unmarshal(m, b)
}
func (m *QueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryResponse.Marshal(b, m, deterministic)
}
func (m *QueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryResponse.Merge(m, src)
}
func (m *QueryResponse) XXX_Size() int {
	return xxx_messageInfo_QueryResponse.Size(m)
}
func (m *QueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryResponse proto.InternalMessageInfo

func (m *QueryResponse) GetRoutes() []*Route {
	if m != nil {
		return m.Routes
	}
	return nil
}

// WatchRequest is made to Watch Router
type WatchRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *WatchRequest) Reset()         { *m = WatchRequest{} }
func (m *WatchRequest) String() string { return proto.CompactTextString(m) }
func (*WatchRequest) ProtoMessage()    {}
func (*WatchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{7}
}

func (m *WatchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WatchRequest.Unmarshal(m, b)
}
func (m *WatchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WatchRequest.Marshal(b, m, deterministic)
}
func (m *WatchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WatchRequest.Merge(m, src)
}
func (m *WatchRequest) XXX_Size() int {
	return xxx_messageInfo_WatchRequest.Size(m)
}
func (m *WatchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WatchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WatchRequest proto.InternalMessageInfo

// Advert is router advertsement streamed by Watch
type Advert struct {
	// id of the advertising router
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// type of advertisement
	Type AdvertType `protobuf:"varint,2,opt,name=type,proto3,enum=go.micro.router.AdvertType" json:"type,omitempty"`
	// unix timestamp of the advertisement
	Timestamp int64 `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// TTL of the Advert
	Ttl int64 `protobuf:"varint,4,opt,name=ttl,proto3" json:"ttl,omitempty"`
	// events is a list of advertised events
	Events               []*Event `protobuf:"bytes,5,rep,name=events,proto3" json:"events,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Advert) Reset()         { *m = Advert{} }
func (m *Advert) String() string { return proto.CompactTextString(m) }
func (*Advert) ProtoMessage()    {}
func (*Advert) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{8}
}

func (m *Advert) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Advert.Unmarshal(m, b)
}
func (m *Advert) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Advert.Marshal(b, m, deterministic)
}
func (m *Advert) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Advert.Merge(m, src)
}
func (m *Advert) XXX_Size() int {
	return xxx_messageInfo_Advert.Size(m)
}
func (m *Advert) XXX_DiscardUnknown() {
	xxx_messageInfo_Advert.DiscardUnknown(m)
}

var xxx_messageInfo_Advert proto.InternalMessageInfo

func (m *Advert) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Advert) GetType() AdvertType {
	if m != nil {
		return m.Type
	}
	return AdvertType_AdvertAnnounce
}

func (m *Advert) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Advert) GetTtl() int64 {
	if m != nil {
		return m.Ttl
	}
	return 0
}

func (m *Advert) GetEvents() []*Event {
	if m != nil {
		return m.Events
	}
	return nil
}

// ProcessResponse is returned by Process
type ProcessResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProcessResponse) Reset()         { *m = ProcessResponse{} }
func (m *ProcessResponse) String() string { return proto.CompactTextString(m) }
func (*ProcessResponse) ProtoMessage()    {}
func (*ProcessResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{9}
}

func (m *ProcessResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProcessResponse.Unmarshal(m, b)
}
func (m *ProcessResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProcessResponse.Marshal(b, m, deterministic)
}
func (m *ProcessResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProcessResponse.Merge(m, src)
}
func (m *ProcessResponse) XXX_Size() int {
	return xxx_messageInfo_ProcessResponse.Size(m)
}
func (m *ProcessResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProcessResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProcessResponse proto.InternalMessageInfo

// CreateResponse is returned by Create
type CreateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateResponse) Reset()         { *m = CreateResponse{} }
func (m *CreateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()    {}
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{10}
}

func (m *CreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateResponse.Unmarshal(m, b)
}
func (m *CreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateResponse.Marshal(b, m, deterministic)
}
func (m *CreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateResponse.Merge(m, src)
}
func (m *CreateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateResponse.Size(m)
}
func (m *CreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateResponse proto.InternalMessageInfo

// DeleteResponse is returned by Delete
type DeleteResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeleteResponse) Reset()         { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()    {}
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{11}
}

func (m *DeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeleteResponse.Unmarshal(m, b)
}
func (m *DeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeleteResponse.Marshal(b, m, deterministic)
}
func (m *DeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeleteResponse.Merge(m, src)
}
func (m *DeleteResponse) XXX_Size() int {
	return xxx_messageInfo_DeleteResponse.Size(m)
}
func (m *DeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeleteResponse proto.InternalMessageInfo

// UpdateResponse is returned by Update
type UpdateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UpdateResponse) Reset()         { *m = UpdateResponse{} }
func (m *UpdateResponse) String() string { return proto.CompactTextString(m) }
func (*UpdateResponse) ProtoMessage()    {}
func (*UpdateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{12}
}

func (m *UpdateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateResponse.Unmarshal(m, b)
}
func (m *UpdateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateResponse.Marshal(b, m, deterministic)
}
func (m *UpdateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateResponse.Merge(m, src)
}
func (m *UpdateResponse) XXX_Size() int {
	return xxx_messageInfo_UpdateResponse.Size(m)
}
func (m *UpdateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateResponse proto.InternalMessageInfo

// Event is routing table event
type Event struct {
	// type of event
	Type EventType `protobuf:"varint,1,opt,name=type,proto3,enum=go.micro.router.EventType" json:"type,omitempty"`
	// unix timestamp of event
	Timestamp int64 `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// service route
	Route                *Route   `protobuf:"bytes,3,opt,name=route,proto3" json:"route,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{13}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetType() EventType {
	if m != nil {
		return m.Type
	}
	return EventType_Create
}

func (m *Event) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Event) GetRoute() *Route {
	if m != nil {
		return m.Route
	}
	return nil
}

// Query is passed in a LookupRequest
type Query struct {
	// service to lookup
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// gateway to lookup
	Gateway string `protobuf:"bytes,2,opt,name=gateway,proto3" json:"gateway,omitempty"`
	// network to lookup
	Network              string   `protobuf:"bytes,3,opt,name=network,proto3" json:"network,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Query) Reset()         { *m = Query{} }
func (m *Query) String() string { return proto.CompactTextString(m) }
func (*Query) ProtoMessage()    {}
func (*Query) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{14}
}

func (m *Query) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Query.Unmarshal(m, b)
}
func (m *Query) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Query.Marshal(b, m, deterministic)
}
func (m *Query) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Query.Merge(m, src)
}
func (m *Query) XXX_Size() int {
	return xxx_messageInfo_Query.Size(m)
}
func (m *Query) XXX_DiscardUnknown() {
	xxx_messageInfo_Query.DiscardUnknown(m)
}

var xxx_messageInfo_Query proto.InternalMessageInfo

func (m *Query) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Query) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *Query) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

// Route is a service route
type Route struct {
	// service for the route
	Service string `protobuf:"bytes,1,opt,name=service,proto3" json:"service,omitempty"`
	// the address that advertise this route
	Address string `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	// gateway as the next hop
	Gateway string `protobuf:"bytes,3,opt,name=gateway,proto3" json:"gateway,omitempty"`
	// the network for this destination
	Network string `protobuf:"bytes,4,opt,name=network,proto3" json:"network,omitempty"`
	// router if the router id
	Router string `protobuf:"bytes,5,opt,name=router,proto3" json:"router,omitempty"`
	// the network link
	Link string `protobuf:"bytes,6,opt,name=link,proto3" json:"link,omitempty"`
	// the metric / score of this route
	Metric               int64    `protobuf:"varint,7,opt,name=metric,proto3" json:"metric,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Route) Reset()         { *m = Route{} }
func (m *Route) String() string { return proto.CompactTextString(m) }
func (*Route) ProtoMessage()    {}
func (*Route) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{15}
}

func (m *Route) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Route.Unmarshal(m, b)
}
func (m *Route) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Route.Marshal(b, m, deterministic)
}
func (m *Route) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Route.Merge(m, src)
}
func (m *Route) XXX_Size() int {
	return xxx_messageInfo_Route.Size(m)
}
func (m *Route) XXX_DiscardUnknown() {
	xxx_messageInfo_Route.DiscardUnknown(m)
}

var xxx_messageInfo_Route proto.InternalMessageInfo

func (m *Route) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Route) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Route) GetGateway() string {
	if m != nil {
		return m.Gateway
	}
	return ""
}

func (m *Route) GetNetwork() string {
	if m != nil {
		return m.Network
	}
	return ""
}

func (m *Route) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *Route) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

func (m *Route) GetMetric() int64 {
	if m != nil {
		return m.Metric
	}
	return 0
}

type Status struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Status) Reset()         { *m = Status{} }
func (m *Status) String() string { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()    {}
func (*Status) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{16}
}

func (m *Status) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Status.Unmarshal(m, b)
}
func (m *Status) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Status.Marshal(b, m, deterministic)
}
func (m *Status) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Status.Merge(m, src)
}
func (m *Status) XXX_Size() int {
	return xxx_messageInfo_Status.Size(m)
}
func (m *Status) XXX_DiscardUnknown() {
	xxx_messageInfo_Status.DiscardUnknown(m)
}

var xxx_messageInfo_Status proto.InternalMessageInfo

func (m *Status) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Status) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

type StatusResponse struct {
	Status               *Status  `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_367072455c71aedc, []int{17}
}

func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetStatus() *Status {
	if m != nil {
		return m.Status
	}
	return nil
}

func init() {
	proto.RegisterEnum("go.micro.router.AdvertType", AdvertType_name, AdvertType_value)
	proto.RegisterEnum("go.micro.router.EventType", EventType_name, EventType_value)
	proto.RegisterType((*Request)(nil), "go.micro.router.Request")
	proto.RegisterType((*Response)(nil), "go.micro.router.Response")
	proto.RegisterType((*ListResponse)(nil), "go.micro.router.ListResponse")
	proto.RegisterType((*LookupRequest)(nil), "go.micro.router.LookupRequest")
	proto.RegisterType((*LookupResponse)(nil), "go.micro.router.LookupResponse")
	proto.RegisterType((*QueryRequest)(nil), "go.micro.router.QueryRequest")
	proto.RegisterType((*QueryResponse)(nil), "go.micro.router.QueryResponse")
	proto.RegisterType((*WatchRequest)(nil), "go.micro.router.WatchRequest")
	proto.RegisterType((*Advert)(nil), "go.micro.router.Advert")
	proto.RegisterType((*ProcessResponse)(nil), "go.micro.router.ProcessResponse")
	proto.RegisterType((*CreateResponse)(nil), "go.micro.router.CreateResponse")
	proto.RegisterType((*DeleteResponse)(nil), "go.micro.router.DeleteResponse")
	proto.RegisterType((*UpdateResponse)(nil), "go.micro.router.UpdateResponse")
	proto.RegisterType((*Event)(nil), "go.micro.router.Event")
	proto.RegisterType((*Query)(nil), "go.micro.router.Query")
	proto.RegisterType((*Route)(nil), "go.micro.router.Route")
	proto.RegisterType((*Status)(nil), "go.micro.router.Status")
	proto.RegisterType((*StatusResponse)(nil), "go.micro.router.StatusResponse")
}

func init() { proto.RegisterFile("router.proto", fileDescriptor_367072455c71aedc) }

var fileDescriptor_367072455c71aedc = []byte{
	// 693 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x4f, 0xdb, 0x4a,
	0x10, 0xb7, 0x93, 0xd8, 0x79, 0x99, 0x17, 0x8c, 0xdf, 0xe8, 0x09, 0xac, 0xb4, 0x40, 0xe4, 0x13,
	0x42, 0xc8, 0x54, 0xe9, 0xb5, 0xff, 0x02, 0xa5, 0xaa, 0x54, 0x0e, 0xad, 0x0b, 0xea, 0xd9, 0xd8,
	0x23, 0x6a, 0x91, 0xd8, 0x66, 0x77, 0x03, 0xca, 0xb9, 0x9f, 0xa6, 0xe7, 0x7e, 0xa4, 0x5e, 0xfb,
	0x21, 0x2a, 0xef, 0xae, 0x43, 0x88, 0x31, 0x12, 0x9c, 0xbc, 0xf3, 0xef, 0x37, 0xb3, 0x3b, 0xbf,
	0x19, 0x43, 0x9f, 0xe5, 0x33, 0x41, 0x2c, 0x28, 0x58, 0x2e, 0x72, 0x5c, 0xbf, 0xc8, 0x83, 0x69,
	0x1a, 0xb3, 0x3c, 0x50, 0x6a, 0xbf, 0x07, 0xdd, 0x90, 0xae, 0x66, 0xc4, 0x85, 0x0f, 0xf0, 0x4f,
	0x48, 0xbc, 0xc8, 0x33, 0x4e, 0xfe, 0x1b, 0xe8, 0x9f, 0xa4, 0x5c, 0x54, 0x32, 0x06, 0x60, 0xcb,
	0x00, 0xee, 0x99, 0xc3, 0xf6, 0xee, 0xbf, 0xa3, 0x8d, 0x60, 0x05, 0x28, 0x08, 0xcb, 0x4f, 0xa8,
	0xbd, 0xfc, 0xd7, 0xb0, 0x76, 0x92, 0xe7, 0x97, 0xb3, 0x42, 0x83, 0xe3, 0x3e, 0x58, 0x57, 0x33,
	0x62, 0x73, 0xcf, 0x1c, 0x9a, 0xf7, 0xc6, 0x7f, 0x29, 0xad, 0xa1, 0x72, 0xf2, 0xdf, 0x81, 0x53,
	0x85, 0x3f, 0xb1, 0x80, 0x57, 0xd0, 0x57, 0x88, 0x4f, 0xca, 0xff, 0x16, 0xd6, 0x74, 0xf4, 0x13,
	0xd3, 0x3b, 0xd0, 0xff, 0x16, 0x89, 0xf8, 0x7b, 0xf5, 0xb6, 0x3f, 0x4d, 0xb0, 0xc7, 0xc9, 0x35,
	0x31, 0x81, 0x0e, 0xb4, 0xd2, 0x44, 0x96, 0xd1, 0x0b, 0x5b, 0x69, 0x82, 0x07, 0xd0, 0x11, 0xf3,
	0x82, 0xbc, 0xd6, 0xd0, 0xdc, 0x75, 0x46, 0xcf, 0x6a, 0xc0, 0x2a, 0xec, 0x74, 0x5e, 0x50, 0x28,
	0x1d, 0xf1, 0x39, 0xf4, 0x44, 0x3a, 0x25, 0x2e, 0xa2, 0x69, 0xe1, 0xb5, 0x87, 0xe6, 0x6e, 0x3b,
	0xbc, 0x55, 0xa0, 0x0b, 0x6d, 0x21, 0x26, 0x5e, 0x47, 0xea, 0xcb, 0x63, 0x59, 0x3b, 0x5d, 0x53,
	0x26, 0xb8, 0x67, 0x35, 0xd4, 0x7e, 0x5c, 0x9a, 0x43, 0xed, 0xe5, 0xff, 0x07, 0xeb, 0x9f, 0x59,
	0x1e, 0x13, 0xe7, 0x0b, 0x3a, 0xb8, 0xe0, 0x1c, 0x31, 0x8a, 0x04, 0x2d, 0x6b, 0xde, 0xd3, 0x84,
	0xee, 0x6a, 0xce, 0x8a, 0x64, 0xd9, 0xe7, 0x87, 0x09, 0x96, 0x84, 0xc6, 0x40, 0xdf, 0xd1, 0x94,
	0x77, 0x1c, 0xdc, 0x5f, 0x40, 0xd3, 0x15, 0x5b, 0xab, 0x57, 0xdc, 0x07, 0x4b, 0xc6, 0xc9, 0xcb,
	0x37, 0xf7, 0x42, 0x39, 0xf9, 0x67, 0x60, 0xc9, 0x5e, 0xa2, 0x07, 0x5d, 0x4e, 0xec, 0x3a, 0x8d,
	0x49, 0xbf, 0x7e, 0x25, 0x96, 0x96, 0x8b, 0x48, 0xd0, 0x4d, 0x34, 0x97, 0xc9, 0x7a, 0x61, 0x25,
	0x96, 0x96, 0x8c, 0xc4, 0x4d, 0xce, 0x2e, 0x65, 0xb2, 0x5e, 0x58, 0x89, 0xfe, 0x2f, 0x13, 0x2c,
	0x99, 0xe7, 0x61, 0xdc, 0x28, 0x49, 0x18, 0x71, 0x5e, 0xe1, 0x6a, 0x71, 0x39, 0x63, 0xbb, 0x31,
	0x63, 0xe7, 0x4e, 0x46, 0xdc, 0xd0, 0x1c, 0x64, 0x9e, 0x25, 0x0d, 0x5a, 0x42, 0x84, 0xce, 0x24,
	0xcd, 0x2e, 0x3d, 0x5b, 0x6a, 0xe5, 0xb9, 0xf4, 0x9d, 0x92, 0x60, 0x69, 0xec, 0x75, 0xe5, 0xeb,
	0x69, 0xc9, 0x1f, 0x81, 0xfd, 0x55, 0x44, 0x62, 0xc6, 0xcb, 0xa8, 0x38, 0x4f, 0xaa, 0x92, 0xe5,
	0x19, 0xff, 0x07, 0x8b, 0x18, 0xcb, 0x99, 0xae, 0x56, 0x09, 0xfe, 0x18, 0x1c, 0x15, 0xb3, 0x98,
	0x86, 0x03, 0xb0, 0xb9, 0xd4, 0xe8, 0x69, 0xda, 0xac, 0x75, 0x40, 0x07, 0x68, 0xb7, 0xbd, 0x11,
	0xc0, 0x2d, 0x8d, 0x11, 0xc1, 0x51, 0xd2, 0x38, 0xcb, 0xf2, 0x59, 0x16, 0x93, 0x6b, 0xa0, 0x0b,
	0x7d, 0xa5, 0x53, 0x1c, 0x72, 0xcd, 0xbd, 0x03, 0xe8, 0x2d, 0x68, 0x81, 0x00, 0xb6, 0x22, 0xa0,
	0x6b, 0x94, 0x67, 0x45, 0x3d, 0xd7, 0x2c, 0xcf, 0x3a, 0xa0, 0x35, 0xfa, 0xd3, 0x02, 0x3b, 0x54,
	0x4f, 0xf2, 0x09, 0x6c, 0xb5, 0x3f, 0x70, 0xbb, 0x56, 0xda, 0x9d, 0xbd, 0x34, 0xd8, 0x69, 0xb4,
	0x6b, 0x12, 0x1b, 0x78, 0x08, 0x96, 0x9c, 0x65, 0xdc, 0xaa, 0xf9, 0x2e, 0xcf, 0xf8, 0xa0, 0x61,
	0xae, 0x7c, 0xe3, 0x85, 0x89, 0x87, 0xd0, 0x53, 0xd7, 0x4b, 0x39, 0xa1, 0x57, 0x27, 0xac, 0x86,
	0xd8, 0x6c, 0x98, 0x7e, 0x89, 0xf1, 0x01, 0xba, 0x7a, 0x2e, 0xb1, 0xc9, 0x6f, 0x30, 0xac, 0x19,
	0x56, 0x47, 0xd9, 0xc0, 0xe3, 0x05, 0x07, 0x9a, 0x0b, 0xd9, 0x69, 0xea, 0xe8, 0x02, 0x66, 0xf4,
	0xbb, 0x05, 0xd6, 0x69, 0x74, 0x3e, 0x21, 0x3c, 0xaa, 0x9a, 0x83, 0x0d, 0xa3, 0x78, 0x0f, 0xdc,
	0xca, 0x3a, 0x31, 0x4a, 0x10, 0xd5, 0xd5, 0x47, 0x80, 0xac, 0x6c, 0x20, 0x09, 0xa2, 0xe8, 0xf0,
	0x08, 0x90, 0x95, 0xa5, 0x65, 0xe0, 0x18, 0x3a, 0xe5, 0xbf, 0xef, 0x81, 0xd7, 0xa9, 0x13, 0x61,
	0xf9, 0x67, 0xe9, 0x1b, 0xf8, 0xb1, 0xda, 0x39, 0x5b, 0x0d, 0xff, 0x19, 0x0d, 0xb4, 0xdd, 0x64,
	0xae, 0x90, 0xce, 0x6d, 0xf9, 0xdf, 0x7e, 0xf9, 0x37, 0x00, 0x00, 0xff, 0xff, 0x86, 0x75, 0x28,
	0x0b, 0xc7, 0x07, 0x00, 0x00,
}
