// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: interview/interview.proto

package interview

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The basic message containing the primary information about a interview.
type Interview struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                     uint32     `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	InterviewerId          uint32     `protobuf:"varint,2,opt,name=interviewer_id,json=interviewerId,proto3" json:"interviewer_id,omitempty"`
	StudentId              uint32     `protobuf:"varint,3,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	InterviewDate          string     `protobuf:"bytes,4,opt,name=interview_date,json=interviewDate,proto3" json:"interview_date,omitempty"`
	Grade                  uint32     `protobuf:"varint,5,opt,name=grade,proto3" json:"grade,omitempty"`
	SubjectiveRating       string     `protobuf:"bytes,6,opt,name=subjective_rating,json=subjectiveRating,proto3" json:"subjective_rating,omitempty"`
	Notes                  string     `protobuf:"bytes,7,opt,name=notes,proto3" json:"notes,omitempty"`
	DeterminedEnglishLevel string     `protobuf:"bytes,8,opt,name=determined_english_level,json=determinedEnglishLevel,proto3" json:"determined_english_level,omitempty"`
	MainTask               uint32     `protobuf:"varint,9,opt,name=main_task,json=mainTask,proto3" json:"main_task,omitempty"`
	Question               *anypb.Any `protobuf:"bytes,10,opt,name=question,proto3" json:"question,omitempty"`
}

func (x *Interview) Reset() {
	*x = Interview{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interview_interview_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interview) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interview) ProtoMessage() {}

func (x *Interview) ProtoReflect() protoreflect.Message {
	mi := &file_interview_interview_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interview.ProtoReflect.Descriptor instead.
func (*Interview) Descriptor() ([]byte, []int) {
	return file_interview_interview_proto_rawDescGZIP(), []int{0}
}

func (x *Interview) GetId() uint32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Interview) GetInterviewerId() uint32 {
	if x != nil {
		return x.InterviewerId
	}
	return 0
}

func (x *Interview) GetStudentId() uint32 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *Interview) GetInterviewDate() string {
	if x != nil {
		return x.InterviewDate
	}
	return ""
}

func (x *Interview) GetGrade() uint32 {
	if x != nil {
		return x.Grade
	}
	return 0
}

func (x *Interview) GetSubjectiveRating() string {
	if x != nil {
		return x.SubjectiveRating
	}
	return ""
}

func (x *Interview) GetNotes() string {
	if x != nil {
		return x.Notes
	}
	return ""
}

func (x *Interview) GetDeterminedEnglishLevel() string {
	if x != nil {
		return x.DeterminedEnglishLevel
	}
	return ""
}

func (x *Interview) GetMainTask() uint32 {
	if x != nil {
		return x.MainTask
	}
	return 0
}

func (x *Interview) GetQuestion() *anypb.Any {
	if x != nil {
		return x.Question
	}
	return nil
}

// The basic message containing information about all interviews from db
type Interviews struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interviews []*Interview `protobuf:"bytes,1,rep,name=interviews,proto3" json:"interviews,omitempty"`
}

func (x *Interviews) Reset() {
	*x = Interviews{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interview_interview_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Interviews) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Interviews) ProtoMessage() {}

func (x *Interviews) ProtoReflect() protoreflect.Message {
	mi := &file_interview_interview_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Interviews.ProtoReflect.Descriptor instead.
func (*Interviews) Descriptor() ([]byte, []int) {
	return file_interview_interview_proto_rawDescGZIP(), []int{1}
}

func (x *Interviews) GetInterviews() []*Interview {
	if x != nil {
		return x.Interviews
	}
	return nil
}

// Basic request/response containing id of the interview
type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_interview_interview_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_interview_interview_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_interview_interview_proto_rawDescGZIP(), []int{2}
}

func (x *Id) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_interview_interview_proto protoreflect.FileDescriptor

var file_interview_interview_proto_rawDesc = []byte{
	0x0a, 0x19, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67,
	0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x8b, 0x04, 0x0a, 0x09, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69,
	0x65, 0x77, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x73, 0x74, 0x75, 0x64,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69,
	0x65, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x44, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x72, 0x61, 0x64, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x67, 0x72, 0x61,
	0x64, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65,
	0x5f, 0x72, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73,
	0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x61, 0x74, 0x69, 0x6e, 0x67, 0x12,
	0x14, 0x0a, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x6e, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x38, 0x0a, 0x18, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x5f, 0x6c, 0x65, 0x76, 0x65,
	0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x16, 0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69,
	0x6e, 0x65, 0x64, 0x45, 0x6e, 0x67, 0x6c, 0x69, 0x73, 0x68, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x12,
	0x1b, 0x0a, 0x09, 0x6d, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x08, 0x6d, 0x61, 0x69, 0x6e, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x30, 0x0a, 0x08,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x3a, 0x9e,
	0x01, 0x92, 0x41, 0x9a, 0x01, 0x0a, 0x97, 0x01, 0x2a, 0x10, 0x55, 0x73, 0x65, 0x72, 0x20, 0x6a,
	0x73, 0x6f, 0x6e, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0xd2, 0x01, 0x0e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x65, 0x72, 0x5f, 0x69, 0x64, 0xd2, 0x01, 0x0a, 0x73, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0xd2, 0x01, 0x0e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x5f, 0x64, 0x61, 0x74, 0x65, 0xd2, 0x01, 0x05, 0x67, 0x72, 0x61, 0x64,
	0x65, 0xd2, 0x01, 0x11, 0x73, 0x75, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x72,
	0x61, 0x74, 0x69, 0x6e, 0x67, 0xd2, 0x01, 0x05, 0x6e, 0x6f, 0x74, 0x65, 0x73, 0xd2, 0x01, 0x18,
	0x64, 0x65, 0x74, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x65, 0x64, 0x5f, 0x65, 0x6e, 0x67, 0x6c, 0x69,
	0x73, 0x68, 0x5f, 0x6c, 0x65, 0x76, 0x65, 0x6c, 0xd2, 0x01, 0x09, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x74, 0x61, 0x73, 0x6b, 0xd2, 0x01, 0x08, 0x71, 0x75, 0x65, 0x73, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0x6e, 0x0a, 0x0a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x34, 0x0a,
	0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x14, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x2e, 0x49, 0x6e,
	0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x52, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69,
	0x65, 0x77, 0x73, 0x3a, 0x2a, 0x92, 0x41, 0x27, 0x0a, 0x25, 0x2a, 0x16, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x69, 0x65, 0x77, 0x73, 0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x20, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0xd2, 0x01, 0x0a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x69, 0x65, 0x77, 0x73, 0x22,
	0x30, 0x0a, 0x02, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x3a, 0x1a, 0x92, 0x41, 0x17, 0x0a, 0x15, 0x2a, 0x0e, 0x49, 0x64,
	0x20, 0x6a, 0x73, 0x6f, 0x6e, 0x20, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61, 0xd2, 0x01, 0x02, 0x69,
	0x64, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x4e, 0x45, 0x4b, 0x45, 0x54, 0x53, 0x4b, 0x59, 0x2f, 0x6d, 0x6e, 0x65, 0x6d, 0x6f, 0x73, 0x79,
	0x6e, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x69, 0x65, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_interview_interview_proto_rawDescOnce sync.Once
	file_interview_interview_proto_rawDescData = file_interview_interview_proto_rawDesc
)

func file_interview_interview_proto_rawDescGZIP() []byte {
	file_interview_interview_proto_rawDescOnce.Do(func() {
		file_interview_interview_proto_rawDescData = protoimpl.X.CompressGZIP(file_interview_interview_proto_rawDescData)
	})
	return file_interview_interview_proto_rawDescData
}

var file_interview_interview_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_interview_interview_proto_goTypes = []interface{}{
	(*Interview)(nil),  // 0: interview.Interview
	(*Interviews)(nil), // 1: interview.Interviews
	(*Id)(nil),         // 2: interview.Id
	(*anypb.Any)(nil),  // 3: google.protobuf.Any
}
var file_interview_interview_proto_depIdxs = []int32{
	3, // 0: interview.Interview.question:type_name -> google.protobuf.Any
	0, // 1: interview.Interviews.interviews:type_name -> interview.Interview
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_interview_interview_proto_init() }
func file_interview_interview_proto_init() {
	if File_interview_interview_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_interview_interview_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interview); i {
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
		file_interview_interview_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Interviews); i {
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
		file_interview_interview_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
			RawDescriptor: file_interview_interview_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_interview_interview_proto_goTypes,
		DependencyIndexes: file_interview_interview_proto_depIdxs,
		MessageInfos:      file_interview_interview_proto_msgTypes,
	}.Build()
	File_interview_interview_proto = out.File
	file_interview_interview_proto_rawDesc = nil
	file_interview_interview_proto_goTypes = nil
	file_interview_interview_proto_depIdxs = nil
}
