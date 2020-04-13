// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/openbank/openbank/v1/consent/all.proto

package consent

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type Status int32

const (
	Status_UnknownStatus   Status = 0
	Status_INITIATED       Status = 1
	Status_ACCEPTED        Status = 2
	Status_REJECTED        Status = 3
	Status_REVOKED         Status = 4
	Status_RECEIVED        Status = 5
	Status_VALID           Status = 6
	Status_REVOKEDBYPSU    Status = 7
	Status_EXPIRED         Status = 8
	Status_TERMINATEDBYTPP Status = 9
)

var Status_name = map[int32]string{
	0: "UnknownStatus",
	1: "INITIATED",
	2: "ACCEPTED",
	3: "REJECTED",
	4: "REVOKED",
	5: "RECEIVED",
	6: "VALID",
	7: "REVOKEDBYPSU",
	8: "EXPIRED",
	9: "TERMINATEDBYTPP",
}

var Status_value = map[string]int32{
	"UnknownStatus":   0,
	"INITIATED":       1,
	"ACCEPTED":        2,
	"REJECTED":        3,
	"REVOKED":         4,
	"RECEIVED":        5,
	"VALID":           6,
	"REVOKEDBYPSU":    7,
	"EXPIRED":         8,
	"TERMINATEDBYTPP": 9,
}

func (x Status) String() string {
	return proto.EnumName(Status_name, int32(x))
}

func (Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_955aa7c027ca6560, []int{0}
}

// AnswerConsentChallengeRequest is a request mesasge to answer consent challenge request
type AnswerConsentChallengeRequest struct {
	BankID               string   `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"bank_id,omitempty"`
	ConsentID            string   `protobuf:"bytes,2,opt,name=ConsentID,json=consent_id,proto3" json:"consent_id,omitempty"`
	Answer               string   `protobuf:"bytes,3,opt,name=Answer,json=answer,proto3" json:"answer,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AnswerConsentChallengeRequest) Reset()         { *m = AnswerConsentChallengeRequest{} }
func (m *AnswerConsentChallengeRequest) String() string { return proto.CompactTextString(m) }
func (*AnswerConsentChallengeRequest) ProtoMessage()    {}
func (*AnswerConsentChallengeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_955aa7c027ca6560, []int{0}
}

func (m *AnswerConsentChallengeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AnswerConsentChallengeRequest.Unmarshal(m, b)
}
func (m *AnswerConsentChallengeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AnswerConsentChallengeRequest.Marshal(b, m, deterministic)
}
func (m *AnswerConsentChallengeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AnswerConsentChallengeRequest.Merge(m, src)
}
func (m *AnswerConsentChallengeRequest) XXX_Size() int {
	return xxx_messageInfo_AnswerConsentChallengeRequest.Size(m)
}
func (m *AnswerConsentChallengeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AnswerConsentChallengeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AnswerConsentChallengeRequest proto.InternalMessageInfo

func (m *AnswerConsentChallengeRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *AnswerConsentChallengeRequest) GetConsentID() string {
	if m != nil {
		return m.ConsentID
	}
	return ""
}

func (m *AnswerConsentChallengeRequest) GetAnswer() string {
	if m != nil {
		return m.Answer
	}
	return ""
}

// Consent is a response mesasge
type Consent struct {
	ConsentID            string   `protobuf:"bytes,1,opt,name=ConsentID,json=consent_id,proto3" json:"consent_id,omitempty"`
	Jwt                  string   `protobuf:"bytes,2,opt,name=Jwt,json=jwt,proto3" json:"jwt,omitempty"`
	Status               Status   `protobuf:"varint,3,opt,name=Status,json=status,proto3,enum=consent.Status" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Consent) Reset()         { *m = Consent{} }
func (m *Consent) String() string { return proto.CompactTextString(m) }
func (*Consent) ProtoMessage()    {}
func (*Consent) Descriptor() ([]byte, []int) {
	return fileDescriptor_955aa7c027ca6560, []int{1}
}

func (m *Consent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Consent.Unmarshal(m, b)
}
func (m *Consent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Consent.Marshal(b, m, deterministic)
}
func (m *Consent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Consent.Merge(m, src)
}
func (m *Consent) XXX_Size() int {
	return xxx_messageInfo_Consent.Size(m)
}
func (m *Consent) XXX_DiscardUnknown() {
	xxx_messageInfo_Consent.DiscardUnknown(m)
}

var xxx_messageInfo_Consent proto.InternalMessageInfo

func (m *Consent) GetConsentID() string {
	if m != nil {
		return m.ConsentID
	}
	return ""
}

func (m *Consent) GetJwt() string {
	if m != nil {
		return m.Jwt
	}
	return ""
}

func (m *Consent) GetStatus() Status {
	if m != nil {
		return m.Status
	}
	return Status_UnknownStatus
}

// CreateConsentEmailRequest
type CreateConsentEmailRequest struct {
	BankID               string   `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"bank_id,omitempty"`
	For                  string   `protobuf:"bytes,2,opt,name=For,json=for,proto3" json:"for,omitempty"`
	View                 string   `protobuf:"bytes,3,opt,name=View,json=view,proto3" json:"view,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=Email,json=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateConsentEmailRequest) Reset()         { *m = CreateConsentEmailRequest{} }
func (m *CreateConsentEmailRequest) String() string { return proto.CompactTextString(m) }
func (*CreateConsentEmailRequest) ProtoMessage()    {}
func (*CreateConsentEmailRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_955aa7c027ca6560, []int{2}
}

func (m *CreateConsentEmailRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConsentEmailRequest.Unmarshal(m, b)
}
func (m *CreateConsentEmailRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConsentEmailRequest.Marshal(b, m, deterministic)
}
func (m *CreateConsentEmailRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConsentEmailRequest.Merge(m, src)
}
func (m *CreateConsentEmailRequest) XXX_Size() int {
	return xxx_messageInfo_CreateConsentEmailRequest.Size(m)
}
func (m *CreateConsentEmailRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConsentEmailRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConsentEmailRequest proto.InternalMessageInfo

func (m *CreateConsentEmailRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *CreateConsentEmailRequest) GetFor() string {
	if m != nil {
		return m.For
	}
	return ""
}

func (m *CreateConsentEmailRequest) GetView() string {
	if m != nil {
		return m.View
	}
	return ""
}

func (m *CreateConsentEmailRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

// CreateConsentSMSRequest
type CreateConsentSMSRequest struct {
	BankID               string   `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"bank_id,omitempty"`
	For                  string   `protobuf:"bytes,2,opt,name=For,json=for,proto3" json:"for,omitempty"`
	View                 string   `protobuf:"bytes,3,opt,name=View,json=view,proto3" json:"view,omitempty"`
	PhoneNumber          string   `protobuf:"bytes,4,opt,name=PhoneNumber,json=phone_number,proto3" json:"phone_number,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateConsentSMSRequest) Reset()         { *m = CreateConsentSMSRequest{} }
func (m *CreateConsentSMSRequest) String() string { return proto.CompactTextString(m) }
func (*CreateConsentSMSRequest) ProtoMessage()    {}
func (*CreateConsentSMSRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_955aa7c027ca6560, []int{3}
}

func (m *CreateConsentSMSRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateConsentSMSRequest.Unmarshal(m, b)
}
func (m *CreateConsentSMSRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateConsentSMSRequest.Marshal(b, m, deterministic)
}
func (m *CreateConsentSMSRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateConsentSMSRequest.Merge(m, src)
}
func (m *CreateConsentSMSRequest) XXX_Size() int {
	return xxx_messageInfo_CreateConsentSMSRequest.Size(m)
}
func (m *CreateConsentSMSRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateConsentSMSRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateConsentSMSRequest proto.InternalMessageInfo

func (m *CreateConsentSMSRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *CreateConsentSMSRequest) GetFor() string {
	if m != nil {
		return m.For
	}
	return ""
}

func (m *CreateConsentSMSRequest) GetView() string {
	if m != nil {
		return m.View
	}
	return ""
}

func (m *CreateConsentSMSRequest) GetPhoneNumber() string {
	if m != nil {
		return m.PhoneNumber
	}
	return ""
}

// GetConsentsRequest
type GetConsentsRequest struct {
	BankID               string   `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"bank_id,omitempty"`
	NexeStartingIndex    string   `protobuf:"bytes,2,opt,name=NexeStartingIndex,json=next_starting_index,proto3" json:"next_starting_index,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetConsentsRequest) Reset()         { *m = GetConsentsRequest{} }
func (m *GetConsentsRequest) String() string { return proto.CompactTextString(m) }
func (*GetConsentsRequest) ProtoMessage()    {}
func (*GetConsentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_955aa7c027ca6560, []int{4}
}

func (m *GetConsentsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConsentsRequest.Unmarshal(m, b)
}
func (m *GetConsentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConsentsRequest.Marshal(b, m, deterministic)
}
func (m *GetConsentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConsentsRequest.Merge(m, src)
}
func (m *GetConsentsRequest) XXX_Size() int {
	return xxx_messageInfo_GetConsentsRequest.Size(m)
}
func (m *GetConsentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConsentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetConsentsRequest proto.InternalMessageInfo

func (m *GetConsentsRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *GetConsentsRequest) GetNexeStartingIndex() string {
	if m != nil {
		return m.NexeStartingIndex
	}
	return ""
}

// GetConsentsResponse
type GetConsentsResponse struct {
	Consents             []*Consent `protobuf:"bytes,1,rep,name=Consents,json=consents,proto3" json:"consents,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetConsentsResponse) Reset()         { *m = GetConsentsResponse{} }
func (m *GetConsentsResponse) String() string { return proto.CompactTextString(m) }
func (*GetConsentsResponse) ProtoMessage()    {}
func (*GetConsentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_955aa7c027ca6560, []int{5}
}

func (m *GetConsentsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetConsentsResponse.Unmarshal(m, b)
}
func (m *GetConsentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetConsentsResponse.Marshal(b, m, deterministic)
}
func (m *GetConsentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetConsentsResponse.Merge(m, src)
}
func (m *GetConsentsResponse) XXX_Size() int {
	return xxx_messageInfo_GetConsentsResponse.Size(m)
}
func (m *GetConsentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetConsentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetConsentsResponse proto.InternalMessageInfo

func (m *GetConsentsResponse) GetConsents() []*Consent {
	if m != nil {
		return m.Consents
	}
	return nil
}

// RevokeConsentRequest
type RevokeConsentRequest struct {
	BankID               string   `protobuf:"bytes,1,opt,name=BankID,json=bank_id,proto3" json:"bank_id,omitempty"`
	ConsentID            string   `protobuf:"bytes,2,opt,name=ConsentID,json=consent_id,proto3" json:"consent_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RevokeConsentRequest) Reset()         { *m = RevokeConsentRequest{} }
func (m *RevokeConsentRequest) String() string { return proto.CompactTextString(m) }
func (*RevokeConsentRequest) ProtoMessage()    {}
func (*RevokeConsentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_955aa7c027ca6560, []int{6}
}

func (m *RevokeConsentRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RevokeConsentRequest.Unmarshal(m, b)
}
func (m *RevokeConsentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RevokeConsentRequest.Marshal(b, m, deterministic)
}
func (m *RevokeConsentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RevokeConsentRequest.Merge(m, src)
}
func (m *RevokeConsentRequest) XXX_Size() int {
	return xxx_messageInfo_RevokeConsentRequest.Size(m)
}
func (m *RevokeConsentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RevokeConsentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RevokeConsentRequest proto.InternalMessageInfo

func (m *RevokeConsentRequest) GetBankID() string {
	if m != nil {
		return m.BankID
	}
	return ""
}

func (m *RevokeConsentRequest) GetConsentID() string {
	if m != nil {
		return m.ConsentID
	}
	return ""
}

func init() {
	proto.RegisterEnum("consent.Status", Status_name, Status_value)
	proto.RegisterType((*AnswerConsentChallengeRequest)(nil), "consent.AnswerConsentChallengeRequest")
	proto.RegisterType((*Consent)(nil), "consent.Consent")
	proto.RegisterType((*CreateConsentEmailRequest)(nil), "consent.CreateConsentEmailRequest")
	proto.RegisterType((*CreateConsentSMSRequest)(nil), "consent.CreateConsentSMSRequest")
	proto.RegisterType((*GetConsentsRequest)(nil), "consent.GetConsentsRequest")
	proto.RegisterType((*GetConsentsResponse)(nil), "consent.GetConsentsResponse")
	proto.RegisterType((*RevokeConsentRequest)(nil), "consent.RevokeConsentRequest")
}

func init() {
	proto.RegisterFile("github.com/openbank/openbank/v1/consent/all.proto", fileDescriptor_955aa7c027ca6560)
}

var fileDescriptor_955aa7c027ca6560 = []byte{
	// 1561 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x57, 0x4d, 0x70, 0x1b, 0x49,
	0x15, 0x9e, 0x19, 0xd9, 0xb2, 0xdc, 0x5e, 0x6f, 0xe4, 0xf6, 0x2e, 0xab, 0x55, 0xec, 0xd0, 0x91,
	0xb7, 0x82, 0x71, 0xc5, 0x23, 0x59, 0xde, 0xdd, 0xda, 0x72, 0x15, 0x45, 0xc9, 0x92, 0x36, 0x2b,
	0x63, 0x3b, 0x42, 0xfe, 0x21, 0xc9, 0x01, 0xa7, 0x35, 0xd3, 0x92, 0x26, 0x1e, 0x75, 0x4f, 0xa6,
	0x7b, 0x2c, 0x87, 0x54, 0xaa, 0xf8, 0xb9, 0xe4, 0x46, 0xca, 0x1c, 0x38, 0xa4, 0x0a, 0xae, 0x5c,
	0xa8, 0xa2, 0x8a, 0x03, 0x14, 0x5c, 0xb8, 0x71, 0xe1, 0x00, 0x05, 0x07, 0xb8, 0x73, 0xa6, 0x38,
	0x51, 0xb9, 0x41, 0xf5, 0xfc, 0x49, 0x42, 0x13, 0x93, 0x1c, 0x12, 0x4e, 0x96, 0xdf, 0xfb, 0xde,
	0x7b, 0x5f, 0x7f, 0xef, 0xcd, 0xeb, 0x19, 0xb0, 0xd1, 0xb5, 0x44, 0xcf, 0x6b, 0xeb, 0x06, 0xeb,
	0x17, 0x99, 0x43, 0x68, 0x1b, 0xd3, 0xd3, 0xe1, 0x8f, 0xb3, 0x8d, 0xa2, 0xc1, 0x28, 0x27, 0x54,
	0x14, 0xb1, 0x6d, 0xeb, 0x8e, 0xcb, 0x04, 0x83, 0x33, 0xa1, 0x29, 0xbf, 0xd4, 0x65, 0xac, 0x6b,
	0x93, 0x22, 0x76, 0xac, 0x22, 0xa6, 0x94, 0x09, 0x2c, 0x2c, 0x46, 0x79, 0x00, 0xcb, 0xdf, 0xf4,
	0xff, 0x18, 0xeb, 0x5d, 0x42, 0xd7, 0xf9, 0x00, 0x77, 0xbb, 0xc4, 0x2d, 0x32, 0xc7, 0x47, 0x4c,
	0xa2, 0x0b, 0xcf, 0x55, 0xb0, 0x5c, 0xa1, 0x7c, 0x40, 0xdc, 0x6a, 0x90, 0xbd, 0xda, 0xc3, 0xb6,
	0x4d, 0x68, 0x97, 0xb4, 0xc8, 0x43, 0x8f, 0x70, 0x01, 0x57, 0x40, 0x7a, 0x1b, 0xd3, 0xd3, 0x46,
	0x2d, 0xa7, 0x22, 0x75, 0x75, 0x76, 0x1b, 0x64, 0x94, 0x9c, 0xb2, 0xaa, 0x94, 0x94, 0xa6, 0xd2,
	0x9a, 0x91, 0x6c, 0x4f, 0x2c, 0x13, 0x7e, 0x15, 0xcc, 0x86, 0xf1, 0x8d, 0x5a, 0x4e, 0x9b, 0xc0,
	0x81, 0x90, 0xba, 0x84, 0x16, 0x40, 0x3a, 0x28, 0x98, 0x4b, 0x4d, 0xe0, 0xd2, 0xd8, 0xf7, 0x6c,
	0xa5, 0x33, 0x4a, 0x56, 0xc9, 0x29, 0x85, 0x1f, 0xaa, 0x60, 0x26, 0xcc, 0x3b, 0x5e, 0x42, 0xbd,
	0xb4, 0xc4, 0x12, 0x48, 0xed, 0x0c, 0x44, 0x02, 0x8f, 0xd4, 0x83, 0x81, 0x80, 0x9b, 0x20, 0x7d,
	0x20, 0xb0, 0xf0, 0xb8, 0x4f, 0xe0, 0xdd, 0xf2, 0x15, 0x3d, 0x0c, 0xd5, 0x03, 0xf3, 0x38, 0x23,
	0xee, 0xdb, 0x62, 0x46, 0x3f, 0x53, 0xc1, 0x87, 0x55, 0x97, 0x60, 0x41, 0x42, 0x32, 0xf5, 0x3e,
	0xb6, 0xec, 0xd7, 0xd2, 0x6a, 0x09, 0xa4, 0x3e, 0x67, 0x6e, 0x12, 0xbb, 0x0e, 0x73, 0xe1, 0x35,
	0x30, 0x75, 0x6c, 0x91, 0x41, 0x82, 0x38, 0x53, 0x67, 0x16, 0x19, 0x40, 0x04, 0xa6, 0xfd, 0x92,
	0xb9, 0xa9, 0x09, 0xc0, 0x34, 0x91, 0x8e, 0x98, 0xea, 0x2f, 0x55, 0xf0, 0xc1, 0x18, 0xd5, 0x83,
	0xbd, 0x83, 0xb7, 0x48, 0x74, 0x1d, 0xcc, 0x35, 0x7b, 0x8c, 0x92, 0x7d, 0xaf, 0xdf, 0x26, 0x6e,
	0x02, 0xdd, 0x77, 0x1c, 0xe9, 0x3e, 0xa1, 0xbe, 0x3f, 0x66, 0xfd, 0x04, 0xc0, 0x5b, 0x44, 0x84,
	0x8c, 0xf9, 0x6b, 0xf1, 0xdd, 0x02, 0x0b, 0xfb, 0xe4, 0x9c, 0x1c, 0x08, 0xec, 0x0a, 0x8b, 0x76,
	0x1b, 0xd4, 0x24, 0xe7, 0x09, 0xec, 0x17, 0x29, 0x39, 0x17, 0x27, 0x3c, 0x44, 0x9d, 0x58, 0x12,
	0x16, 0x97, 0xff, 0x16, 0x58, 0x1c, 0x2b, 0xcf, 0x1d, 0xf9, 0x13, 0x7e, 0x06, 0x32, 0x91, 0x2d,
	0xa7, 0xa2, 0xd4, 0xea, 0x5c, 0x39, 0x1b, 0x4f, 0x4d, 0xe8, 0x18, 0xab, 0x91, 0x09, 0x9d, 0xc3,
	0xc1, 0xa1, 0xe0, 0xbd, 0x16, 0x39, 0x63, 0xa7, 0x51, 0x33, 0xde, 0xd0, 0xe3, 0x15, 0xd5, 0x5b,
	0xfb, 0x83, 0x1a, 0x8d, 0x39, 0x7c, 0x1f, 0xcc, 0x1f, 0xd1, 0x53, 0xca, 0x06, 0x34, 0x30, 0x64,
	0x95, 0xbc, 0x96, 0x51, 0xe0, 0x02, 0x98, 0x6d, 0xec, 0x37, 0x0e, 0x1b, 0x95, 0xc3, 0x7a, 0x2d,
	0xab, 0xfa, 0xa6, 0x2c, 0xc8, 0x54, 0xaa, 0xd5, 0x7a, 0x53, 0x5a, 0xb4, 0xc8, 0xd2, 0xaa, 0xef,
	0xd4, 0xab, 0xd2, 0x92, 0xf2, 0x2d, 0x57, 0xc0, 0x4c, 0xab, 0x7e, 0x7c, 0xfb, 0x1b, 0xf5, 0x5a,
	0x76, 0x6a, 0x08, 0xa9, 0xd6, 0x1b, 0xc7, 0xf5, 0x5a, 0x76, 0xda, 0xb7, 0xcc, 0x83, 0xe9, 0xe3,
	0xca, 0x6e, 0xa3, 0x96, 0x4d, 0xfb, 0xff, 0xbe, 0x07, 0xde, 0x09, 0x23, 0xb6, 0xef, 0x36, 0x0f,
	0x8e, 0xb2, 0x33, 0x51, 0x9e, 0xfa, 0x9d, 0x66, 0xa3, 0x55, 0xaf, 0x65, 0x33, 0xbe, 0xe1, 0x03,
	0x70, 0xe5, 0xb0, 0xde, 0xda, 0x6b, 0xec, 0x4b, 0x42, 0xdb, 0x77, 0x0f, 0x9b, 0xcd, 0xec, 0xac,
	0x74, 0xe4, 0xb5, 0x9c, 0x52, 0xfe, 0xcb, 0x3c, 0x78, 0x37, 0x1a, 0x63, 0xe2, 0x9e, 0x59, 0x06,
	0x81, 0x2f, 0x34, 0xf0, 0xa5, 0xe4, 0xd5, 0x05, 0x6f, 0xc4, 0xcd, 0xb9, 0x74, 0xb7, 0xe5, 0x27,
	0x9a, 0x58, 0xf8, 0xa9, 0x76, 0x51, 0xf9, 0xdb, 0xc8, 0xd6, 0xf9, 0x28, 0x48, 0x80, 0x44, 0x8f,
	0xa0, 0x10, 0x8b, 0x5c, 0xf2, 0x90, 0x78, 0x5c, 0x20, 0x23, 0x4a, 0x97, 0x7f, 0x09, 0x4a, 0xd6,
	0x18, 0xa2, 0x76, 0xbe, 0x09, 0x52, 0xe5, 0xd2, 0x06, 0xdc, 0x01, 0x37, 0x02, 0x38, 0x31, 0x2f,
	0x0f, 0x80, 0x08, 0x5c, 0xcb, 0x2f, 0xad, 0x14, 0x4d, 0xd2, 0xb1, 0xa8, 0x15, 0x6c, 0xf0, 0x10,
	0x1c, 0xd2, 0x6b, 0x17, 0xc1, 0x3a, 0x48, 0xdf, 0xae, 0x78, 0xa2, 0x57, 0x86, 0x2b, 0xe0, 0x7a,
	0x4f, 0x08, 0x87, 0x6f, 0x15, 0x8b, 0xd8, 0x13, 0x3d, 0xbd, 0x4d, 0x4f, 0x75, 0xc1, 0xa2, 0x08,
	0x7d, 0xe0, 0x5a, 0x82, 0x7c, 0xff, 0xcf, 0x7f, 0xff, 0x91, 0xf6, 0xe9, 0x96, 0xba, 0x56, 0xd8,
	0x90, 0x97, 0x8c, 0x9c, 0x2f, 0x5e, 0x7c, 0x1c, 0x0c, 0xe0, 0x93, 0x08, 0xca, 0x8b, 0x8f, 0xe3,
	0x69, 0x7b, 0x52, 0x8c, 0xe9, 0x3c, 0xd5, 0x94, 0x67, 0x9a, 0x3f, 0x72, 0xf0, 0xf7, 0x1a, 0x80,
	0x93, 0x5b, 0x10, 0x16, 0x86, 0x72, 0xbe, 0x6c, 0x45, 0x26, 0x48, 0xfe, 0x2f, 0xf5, 0xa2, 0xf2,
	0xbb, 0x11, 0xc9, 0xc3, 0xa5, 0x85, 0x30, 0x45, 0xfe, 0x42, 0x8b, 0x04, 0xca, 0x5f, 0x0d, 0x1c,
	0x1c, 0x61, 0x44, 0xc9, 0x60, 0xdc, 0xb9, 0xb3, 0x1f, 0x88, 0x7b, 0x0b, 0xac, 0xd4, 0x47, 0xed,
	0xc8, 0xf0, 0x63, 0x4c, 0xc4, 0x3d, 0xc3, 0x20, 0x9c, 0x77, 0x3c, 0xdb, 0x7e, 0xa4, 0xbf, 0x29,
	0x65, 0x6f, 0x48, 0x65, 0xaf, 0x5f, 0xa6, 0xac, 0x4f, 0x7a, 0x44, 0xc9, 0xdf, 0x6a, 0x20, 0xfb,
	0xdf, 0x4b, 0x1a, 0xa2, 0x64, 0x1d, 0x87, 0xfb, 0x3b, 0x41, 0xc5, 0x7f, 0xa8, 0x17, 0x95, 0xdf,
	0x8c, 0xa8, 0xf8, 0xfe, 0x50, 0x45, 0xde, 0xe7, 0xb1, 0x86, 0x1f, 0x8e, 0x6b, 0x38, 0xe2, 0xda,
	0xd9, 0x0d, 0x14, 0xac, 0x83, 0xeb, 0x07, 0x7b, 0x07, 0xff, 0x1f, 0xfd, 0x3e, 0x92, 0xfa, 0x7d,
	0xf9, 0x32, 0xfd, 0x78, 0x9f, 0x8f, 0xa8, 0xf7, 0x6f, 0x0d, 0xcc, 0x8d, 0xac, 0x6b, 0x78, 0x35,
	0x96, 0x65, 0xf2, 0x0e, 0xc9, 0x2f, 0x25, 0x3b, 0x83, 0x0d, 0x5f, 0xf8, 0x85, 0x76, 0x51, 0xf9,
	0x81, 0x36, 0xd4, 0x6f, 0x61, 0xd7, 0xe2, 0x02, 0x61, 0x3b, 0x9e, 0x24, 0x9e, 0xff, 0x76, 0x8b,
	0x08, 0xcf, 0xa5, 0x52, 0x3b, 0x5b, 0x3a, 0x0d, 0x46, 0x05, 0xb6, 0xa8, 0x45, 0xbb, 0xc8, 0x73,
	0x90, 0x60, 0xa8, 0x5c, 0x8a, 0xc1, 0x3a, 0xba, 0x9f, 0x70, 0xcd, 0xdc, 0x47, 0x06, 0xa6, 0xa8,
	0x4d, 0x90, 0xc7, 0x89, 0x89, 0x3a, 0xcc, 0x45, 0x0e, 0xee, 0x5a, 0xd4, 0x7f, 0x23, 0xd3, 0x77,
	0x8e, 0x64, 0x03, 0x4a, 0x70, 0x1f, 0x5c, 0x0b, 0x29, 0x23, 0x72, 0x4e, 0x0c, 0x6f, 0x52, 0xfd,
	0x9b, 0x60, 0x2d, 0xbf, 0x9a, 0xa8, 0x7e, 0xc2, 0xc9, 0xda, 0x3a, 0xb8, 0x19, 0x77, 0xa2, 0x00,
	0xd0, 0x65, 0x9d, 0x70, 0x09, 0x36, 0xfd, 0x46, 0x2c, 0xc3, 0xab, 0x97, 0x74, 0x61, 0xa4, 0x03,
	0xbf, 0xd2, 0xc0, 0xfc, 0xd8, 0xbd, 0x06, 0x97, 0x63, 0x99, 0x93, 0xee, 0xbb, 0x84, 0xc9, 0xfd,
	0xa7, 0x7a, 0x51, 0xf9, 0xf9, 0xc8, 0xe4, 0x2e, 0x06, 0x71, 0xfe, 0x6e, 0x8c, 0xe7, 0x16, 0x8e,
	0x18, 0xa3, 0x81, 0xfd, 0x22, 0x18, 0xd8, 0x0a, 0x58, 0xae, 0xc6, 0x2b, 0x54, 0x82, 0xde, 0xd2,
	0xb0, 0x6e, 0xca, 0x61, 0xd5, 0x5f, 0x75, 0x8d, 0x06, 0xd4, 0x86, 0xca, 0xe5, 0x53, 0x4f, 0x35,
	0x65, 0xfb, 0x4f, 0xe9, 0x8b, 0xca, 0x8f, 0xd3, 0x20, 0x55, 0xd6, 0x4b, 0x70, 0x17, 0xcc, 0x45,
	0xe7, 0xa8, 0x34, 0x1b, 0xf0, 0xd3, 0xa6, 0xcb, 0xce, 0x2c, 0x93, 0x70, 0x54, 0x6d, 0x1d, 0xd5,
	0x10, 0x73, 0x88, 0x1b, 0xbc, 0xb9, 0x23, 0x46, 0xc7, 0xee, 0x0d, 0x07, 0xbb, 0xf2, 0xe4, 0x9c,
	0x79, 0xae, 0x41, 0xf4, 0xf2, 0xf4, 0x86, 0x5e, 0xd2, 0x4b, 0x6b, 0xaa, 0x56, 0xce, 0x62, 0xc7,
	0xb1, 0x2d, 0xc3, 0x8f, 0x2a, 0x3e, 0xe0, 0x8c, 0x6e, 0x4d, 0x58, 0x5a, 0x27, 0x20, 0xf5, 0x71,
	0xa9, 0x04, 0xef, 0x80, 0xe3, 0x60, 0xb8, 0x89, 0x89, 0x06, 0x3d, 0x12, 0x14, 0x88, 0x2e, 0xa4,
	0x36, 0x33, 0x1f, 0x21, 0x8b, 0xa3, 0x3e, 0xb6, 0x3b, 0xcc, 0xed, 0x63, 0x21, 0x27, 0x92, 0xb9,
	0xc8, 0x64, 0x84, 0x23, 0xca, 0x04, 0xea, 0x63, 0x61, 0xf4, 0xfc, 0x10, 0x72, 0xee, 0x10, 0x43,
	0xba, 0xc3, 0x58, 0xbd, 0xb5, 0x2b, 0x0b, 0xc8, 0xad, 0x52, 0x7d, 0x79, 0x81, 0x38, 0x51, 0xf8,
	0x44, 0x71, 0xdf, 0xeb, 0x71, 0xe2, 0x7e, 0x85, 0xcb, 0x05, 0x64, 0x12, 0x2a, 0x2c, 0x6c, 0x73,
	0xbd, 0xd5, 0x94, 0xd9, 0x36, 0x61, 0x03, 0xdc, 0x9a, 0xcc, 0x26, 0xf1, 0xc3, 0x54, 0x3d, 0x7c,
	0x46, 0x90, 0x43, 0xdc, 0xbe, 0xc5, 0xb9, 0x25, 0x15, 0x63, 0x08, 0xfb, 0x43, 0x11, 0x56, 0x0e,
	0xe5, 0x6a, 0x6d, 0xc8, 0x8c, 0x1f, 0xc3, 0x35, 0xb0, 0x9a, 0xc4, 0x2f, 0x40, 0xc9, 0xc3, 0xcb,
	0x9c, 0x1d, 0xe6, 0x51, 0x53, 0x6f, 0x7d, 0x0e, 0x52, 0x9f, 0x94, 0x4a, 0xf0, 0xeb, 0xe0, 0x6b,
	0xe3, 0x21, 0x98, 0x22, 0x8f, 0xc6, 0x0a, 0x10, 0xd7, 0x65, 0x2e, 0x62, 0x86, 0xe1, 0xc9, 0x8b,
	0x3e, 0xec, 0x19, 0x27, 0xee, 0x19, 0x71, 0x11, 0xb7, 0x4c, 0xa2, 0xdf, 0xfb, 0x89, 0x06, 0x9e,
	0x6b, 0xf1, 0xdc, 0x3d, 0xd3, 0x32, 0x29, 0xf8, 0x3d, 0xb5, 0x12, 0x92, 0x64, 0x23, 0xef, 0x05,
	0x01, 0x0b, 0x2e, 0x69, 0xb8, 0x84, 0x0b, 0xd7, 0xf2, 0x0b, 0xc8, 0x13, 0x79, 0xa2, 0x27, 0xb5,
	0x31, 0xfc, 0x15, 0x2d, 0x05, 0xe0, 0x3a, 0x3a, 0xec, 0x91, 0x51, 0x87, 0x3c, 0xbc, 0xe3, 0x32,
	0x3f, 0x6b, 0x87, 0xd9, 0x36, 0x1b, 0x04, 0x12, 0xc8, 0xaa, 0xcc, 0xb5, 0xbe, 0x13, 0x20, 0xaa,
	0xcc, 0x24, 0xa8, 0x63, 0xb3, 0x81, 0xbe, 0x3a, 0x55, 0xce, 0xc8, 0xa1, 0x96, 0x29, 0xb6, 0x66,
	0xe5, 0x2f, 0xc1, 0x4e, 0x09, 0xdd, 0x7e, 0x08, 0x3e, 0xf9, 0xdf, 0xcb, 0x03, 0x2e, 0xc8, 0x6f,
	0x82, 0x98, 0xba, 0x89, 0x05, 0x06, 0x9f, 0xbd, 0xc2, 0x03, 0x05, 0x17, 0xf7, 0x30, 0xc5, 0x5d,
	0x32, 0x1e, 0x09, 0x05, 0xf3, 0xd1, 0xd8, 0xb1, 0xa2, 0xad, 0xf1, 0x85, 0xda, 0x54, 0xef, 0x45,
	0x5f, 0xc2, 0xdf, 0x55, 0x95, 0xa7, 0xaa, 0xf2, 0x4c, 0x55, 0x7e, 0xad, 0x2a, 0x7f, 0x55, 0x95,
	0x17, 0xaa, 0xf2, 0x47, 0x4d, 0x69, 0xa7, 0xfd, 0x2f, 0xdb, 0xcd, 0xff, 0x04, 0x00, 0x00, 0xff,
	0xff, 0xb2, 0x44, 0xd1, 0x84, 0x63, 0x0f, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ConsentServiceClient is the client API for ConsentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConsentServiceClient interface {
	// AnswerConsentChallenge sends the answer to the challenge for consnet challenge request
	AnswerConsentChallenge(ctx context.Context, in *AnswerConsentChallengeRequest, opts ...grpc.CallOption) (*Consent, error)
	// CreateConsentEmail creates a new email consent
	CreateConsentEmail(ctx context.Context, in *CreateConsentEmailRequest, opts ...grpc.CallOption) (*Consent, error)
	// CreateConsentSMS creates a new sms consent
	CreateConsentSMS(ctx context.Context, in *CreateConsentSMSRequest, opts ...grpc.CallOption) (*Consent, error)
	// GetConsents returns a list containing up to 20 consents.
	GetConsents(ctx context.Context, in *GetConsentsRequest, opts ...grpc.CallOption) (*GetConsentsResponse, error)
	// RevokeConsent revokes the consent
	RevokeConsent(ctx context.Context, in *RevokeConsentRequest, opts ...grpc.CallOption) (*Consent, error)
}

type consentServiceClient struct {
	cc *grpc.ClientConn
}

func NewConsentServiceClient(cc *grpc.ClientConn) ConsentServiceClient {
	return &consentServiceClient{cc}
}

func (c *consentServiceClient) AnswerConsentChallenge(ctx context.Context, in *AnswerConsentChallengeRequest, opts ...grpc.CallOption) (*Consent, error) {
	out := new(Consent)
	err := c.cc.Invoke(ctx, "/consent.ConsentService/AnswerConsentChallenge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consentServiceClient) CreateConsentEmail(ctx context.Context, in *CreateConsentEmailRequest, opts ...grpc.CallOption) (*Consent, error) {
	out := new(Consent)
	err := c.cc.Invoke(ctx, "/consent.ConsentService/CreateConsentEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consentServiceClient) CreateConsentSMS(ctx context.Context, in *CreateConsentSMSRequest, opts ...grpc.CallOption) (*Consent, error) {
	out := new(Consent)
	err := c.cc.Invoke(ctx, "/consent.ConsentService/CreateConsentSMS", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consentServiceClient) GetConsents(ctx context.Context, in *GetConsentsRequest, opts ...grpc.CallOption) (*GetConsentsResponse, error) {
	out := new(GetConsentsResponse)
	err := c.cc.Invoke(ctx, "/consent.ConsentService/GetConsents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *consentServiceClient) RevokeConsent(ctx context.Context, in *RevokeConsentRequest, opts ...grpc.CallOption) (*Consent, error) {
	out := new(Consent)
	err := c.cc.Invoke(ctx, "/consent.ConsentService/RevokeConsent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConsentServiceServer is the server API for ConsentService service.
type ConsentServiceServer interface {
	// AnswerConsentChallenge sends the answer to the challenge for consnet challenge request
	AnswerConsentChallenge(context.Context, *AnswerConsentChallengeRequest) (*Consent, error)
	// CreateConsentEmail creates a new email consent
	CreateConsentEmail(context.Context, *CreateConsentEmailRequest) (*Consent, error)
	// CreateConsentSMS creates a new sms consent
	CreateConsentSMS(context.Context, *CreateConsentSMSRequest) (*Consent, error)
	// GetConsents returns a list containing up to 20 consents.
	GetConsents(context.Context, *GetConsentsRequest) (*GetConsentsResponse, error)
	// RevokeConsent revokes the consent
	RevokeConsent(context.Context, *RevokeConsentRequest) (*Consent, error)
}

// UnimplementedConsentServiceServer can be embedded to have forward compatible implementations.
type UnimplementedConsentServiceServer struct {
}

func (*UnimplementedConsentServiceServer) AnswerConsentChallenge(ctx context.Context, req *AnswerConsentChallengeRequest) (*Consent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnswerConsentChallenge not implemented")
}
func (*UnimplementedConsentServiceServer) CreateConsentEmail(ctx context.Context, req *CreateConsentEmailRequest) (*Consent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConsentEmail not implemented")
}
func (*UnimplementedConsentServiceServer) CreateConsentSMS(ctx context.Context, req *CreateConsentSMSRequest) (*Consent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateConsentSMS not implemented")
}
func (*UnimplementedConsentServiceServer) GetConsents(ctx context.Context, req *GetConsentsRequest) (*GetConsentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConsents not implemented")
}
func (*UnimplementedConsentServiceServer) RevokeConsent(ctx context.Context, req *RevokeConsentRequest) (*Consent, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeConsent not implemented")
}

func RegisterConsentServiceServer(s *grpc.Server, srv ConsentServiceServer) {
	s.RegisterService(&_ConsentService_serviceDesc, srv)
}

func _ConsentService_AnswerConsentChallenge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnswerConsentChallengeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsentServiceServer).AnswerConsentChallenge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consent.ConsentService/AnswerConsentChallenge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsentServiceServer).AnswerConsentChallenge(ctx, req.(*AnswerConsentChallengeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsentService_CreateConsentEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConsentEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsentServiceServer).CreateConsentEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consent.ConsentService/CreateConsentEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsentServiceServer).CreateConsentEmail(ctx, req.(*CreateConsentEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsentService_CreateConsentSMS_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateConsentSMSRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsentServiceServer).CreateConsentSMS(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consent.ConsentService/CreateConsentSMS",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsentServiceServer).CreateConsentSMS(ctx, req.(*CreateConsentSMSRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsentService_GetConsents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConsentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsentServiceServer).GetConsents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consent.ConsentService/GetConsents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsentServiceServer).GetConsents(ctx, req.(*GetConsentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ConsentService_RevokeConsent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RevokeConsentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConsentServiceServer).RevokeConsent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/consent.ConsentService/RevokeConsent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConsentServiceServer).RevokeConsent(ctx, req.(*RevokeConsentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConsentService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "consent.ConsentService",
	HandlerType: (*ConsentServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AnswerConsentChallenge",
			Handler:    _ConsentService_AnswerConsentChallenge_Handler,
		},
		{
			MethodName: "CreateConsentEmail",
			Handler:    _ConsentService_CreateConsentEmail_Handler,
		},
		{
			MethodName: "CreateConsentSMS",
			Handler:    _ConsentService_CreateConsentSMS_Handler,
		},
		{
			MethodName: "GetConsents",
			Handler:    _ConsentService_GetConsents_Handler,
		},
		{
			MethodName: "RevokeConsent",
			Handler:    _ConsentService_RevokeConsent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "github.com/openbank/openbank/v1/consent/all.proto",
}
