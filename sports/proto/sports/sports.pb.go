// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.21.5
// source: sports/sports.proto

package sports

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

// Request for ListEvents call.
type ListEventsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Filter  *ListEventsRequestFilter `protobuf:"bytes,1,opt,name=filter,proto3" json:"filter,omitempty"`
	OrderBy string                   `protobuf:"bytes,2,opt,name=order_by,json=orderBy,proto3" json:"order_by,omitempty"`
}

func (x *ListEventsRequest) Reset() {
	*x = ListEventsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsRequest) ProtoMessage() {}

func (x *ListEventsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsRequest.ProtoReflect.Descriptor instead.
func (*ListEventsRequest) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{0}
}

func (x *ListEventsRequest) GetFilter() *ListEventsRequestFilter {
	if x != nil {
		return x.Filter
	}
	return nil
}

func (x *ListEventsRequest) GetOrderBy() string {
	if x != nil {
		return x.OrderBy
	}
	return ""
}

// Response for ListEvents call.
type ListEventsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Events []*Event `protobuf:"bytes,1,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *ListEventsResponse) Reset() {
	*x = ListEventsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsResponse) ProtoMessage() {}

func (x *ListEventsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsResponse.ProtoReflect.Descriptor instead.
func (*ListEventsResponse) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{1}
}

func (x *ListEventsResponse) GetEvents() []*Event {
	if x != nil {
		return x.Events
	}
	return nil
}

// Request for ListEvents call.
type ListEventsRequestFilter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SportId *int64 `protobuf:"varint,1,opt,name=sport_id,json=sportId,proto3,oneof" json:"sport_id,omitempty"`
	Status  string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ListEventsRequestFilter) Reset() {
	*x = ListEventsRequestFilter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListEventsRequestFilter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListEventsRequestFilter) ProtoMessage() {}

func (x *ListEventsRequestFilter) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListEventsRequestFilter.ProtoReflect.Descriptor instead.
func (*ListEventsRequestFilter) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{2}
}

func (x *ListEventsRequestFilter) GetSportId() int64 {
	if x != nil && x.SportId != nil {
		return *x.SportId
	}
	return 0
}

func (x *ListEventsRequestFilter) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

// Request for GetEvent call.
type GetEventRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetEventRequest) Reset() {
	*x = GetEventRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventRequest) ProtoMessage() {}

func (x *GetEventRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventRequest.ProtoReflect.Descriptor instead.
func (*GetEventRequest) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{3}
}

func (x *GetEventRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

// Response for GetEvent call.
type GetEventResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Event *Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
}

func (x *GetEventResponse) Reset() {
	*x = GetEventResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEventResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEventResponse) ProtoMessage() {}

func (x *GetEventResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEventResponse.ProtoReflect.Descriptor instead.
func (*GetEventResponse) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{4}
}

func (x *GetEventResponse) GetEvent() *Event {
	if x != nil {
		return x.Event
	}
	return nil
}

// A sport event resource.
type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// ID represents a unique identifier for the event.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name is the official name given to the event.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// VenueId represents a unique identifier for the venue.
	VenueId int64 `protobuf:"varint,3,opt,name=venue_id,json=venueId,proto3" json:"venue_id,omitempty"`
	// SportsId represents a unique identifier for the type of sport.
	SportId int64 `protobuf:"varint,4,opt,name=sport_id,json=sportId,proto3" json:"sport_id,omitempty"`
	// ParticipantsId represents a unique identifier for the participants of the event.
	ParticipantsId int64 `protobuf:"varint,5,opt,name=participants_id,json=participantsId,proto3" json:"participants_id,omitempty"`
	// AdvertisedStartTime is the time the event is advertised to run.
	AdvertisedStartTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=advertised_start_time,json=advertisedStartTime,proto3" json:"advertised_start_time,omitempty"`
	// AdvertisedEndTime is the time the event is advertised to end.
	AdvertisedEndTime *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=advertised_end_time,json=advertisedEndTime,proto3" json:"advertised_end_time,omitempty"`
	// Status represents whether the event is still open, ongoing or closed.
	Status string `protobuf:"bytes,8,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sports_sports_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_sports_sports_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_sports_sports_proto_rawDescGZIP(), []int{5}
}

func (x *Event) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Event) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Event) GetVenueId() int64 {
	if x != nil {
		return x.VenueId
	}
	return 0
}

func (x *Event) GetSportId() int64 {
	if x != nil {
		return x.SportId
	}
	return 0
}

func (x *Event) GetParticipantsId() int64 {
	if x != nil {
		return x.ParticipantsId
	}
	return 0
}

func (x *Event) GetAdvertisedStartTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AdvertisedStartTime
	}
	return nil
}

func (x *Event) GetAdvertisedEndTime() *timestamppb.Timestamp {
	if x != nil {
		return x.AdvertisedEndTime
	}
	return nil
}

func (x *Event) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_sports_sports_proto protoreflect.FileDescriptor

var file_sports_sports_proto_rawDesc = []byte{
	0x0a, 0x13, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2f, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x67,
	0x0a, 0x11, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x37, 0x0a, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69,
	0x6c, 0x74, 0x65, 0x72, 0x52, 0x06, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x08,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6f, 0x72, 0x64, 0x65, 0x72, 0x42, 0x79, 0x22, 0x3b, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x06, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x73, 0x22, 0x5e, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x12,
	0x1e, 0x0a, 0x08, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x48, 0x00, 0x52, 0x07, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x73, 0x70, 0x6f, 0x72,
	0x74, 0x5f, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23, 0x0a, 0x05, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x22, 0xbe, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x49, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70,
	0x61, 0x6e, 0x74, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x70,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x49, 0x64, 0x12, 0x4e, 0x0a,
	0x15, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x5f, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x13, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74,
	0x69, 0x73, 0x65, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x4a, 0x0a,
	0x13, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x64, 0x5f,
	0x74, 0x69, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x11, 0x61, 0x64, 0x76, 0x65, 0x72, 0x74, 0x69, 0x73,
	0x65, 0x64, 0x45, 0x6e, 0x64, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x32, 0x90, 0x01, 0x0a, 0x06, 0x53, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x12, 0x45, 0x0a, 0x0a,
	0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x19, 0x2e, 0x73, 0x70, 0x6f,
	0x72, 0x74, 0x73, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3f, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x17, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x73, 0x70, 0x6f, 0x72, 0x74,
	0x73, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2f, 0x73, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sports_sports_proto_rawDescOnce sync.Once
	file_sports_sports_proto_rawDescData = file_sports_sports_proto_rawDesc
)

func file_sports_sports_proto_rawDescGZIP() []byte {
	file_sports_sports_proto_rawDescOnce.Do(func() {
		file_sports_sports_proto_rawDescData = protoimpl.X.CompressGZIP(file_sports_sports_proto_rawDescData)
	})
	return file_sports_sports_proto_rawDescData
}

var file_sports_sports_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_sports_sports_proto_goTypes = []interface{}{
	(*ListEventsRequest)(nil),       // 0: sports.ListEventsRequest
	(*ListEventsResponse)(nil),      // 1: sports.ListEventsResponse
	(*ListEventsRequestFilter)(nil), // 2: sports.ListEventsRequestFilter
	(*GetEventRequest)(nil),         // 3: sports.GetEventRequest
	(*GetEventResponse)(nil),        // 4: sports.GetEventResponse
	(*Event)(nil),                   // 5: sports.Event
	(*timestamppb.Timestamp)(nil),   // 6: google.protobuf.Timestamp
}
var file_sports_sports_proto_depIdxs = []int32{
	2, // 0: sports.ListEventsRequest.filter:type_name -> sports.ListEventsRequestFilter
	5, // 1: sports.ListEventsResponse.events:type_name -> sports.Event
	5, // 2: sports.GetEventResponse.event:type_name -> sports.Event
	6, // 3: sports.Event.advertised_start_time:type_name -> google.protobuf.Timestamp
	6, // 4: sports.Event.advertised_end_time:type_name -> google.protobuf.Timestamp
	0, // 5: sports.Sports.ListEvents:input_type -> sports.ListEventsRequest
	3, // 6: sports.Sports.GetEvent:input_type -> sports.GetEventRequest
	1, // 7: sports.Sports.ListEvents:output_type -> sports.ListEventsResponse
	4, // 8: sports.Sports.GetEvent:output_type -> sports.GetEventResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_sports_sports_proto_init() }
func file_sports_sports_proto_init() {
	if File_sports_sports_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sports_sports_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsRequest); i {
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
		file_sports_sports_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsResponse); i {
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
		file_sports_sports_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListEventsRequestFilter); i {
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
		file_sports_sports_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventRequest); i {
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
		file_sports_sports_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEventResponse); i {
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
		file_sports_sports_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
	file_sports_sports_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sports_sports_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sports_sports_proto_goTypes,
		DependencyIndexes: file_sports_sports_proto_depIdxs,
		MessageInfos:      file_sports_sports_proto_msgTypes,
	}.Build()
	File_sports_sports_proto = out.File
	file_sports_sports_proto_rawDesc = nil
	file_sports_sports_proto_goTypes = nil
	file_sports_sports_proto_depIdxs = nil
}
