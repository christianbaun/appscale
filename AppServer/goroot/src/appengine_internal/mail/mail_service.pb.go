// Code generated by protoc-gen-go from "mail_service.proto"
// DO NOT EDIT!

package appengine

import proto "goprotobuf.googlecode.com/hg/proto"
import "math"
import "os"


// Reference proto, math & os imports to suppress error if they are not otherwise used.
var _ = proto.GetString
var _ = math.Inf
var _ os.Error


type MailServiceError_ErrorCode int32

const (
	MailServiceError_OK                      = 0
	MailServiceError_INTERNAL_ERROR          = 1
	MailServiceError_BAD_REQUEST             = 2
	MailServiceError_UNAUTHORIZED_SENDER     = 3
	MailServiceError_INVALID_ATTACHMENT_TYPE = 4
)

var MailServiceError_ErrorCode_name = map[int32]string{
	0: "OK",
	1: "INTERNAL_ERROR",
	2: "BAD_REQUEST",
	3: "UNAUTHORIZED_SENDER",
	4: "INVALID_ATTACHMENT_TYPE",
}
var MailServiceError_ErrorCode_value = map[string]int32{
	"OK":                      0,
	"INTERNAL_ERROR":          1,
	"BAD_REQUEST":             2,
	"UNAUTHORIZED_SENDER":     3,
	"INVALID_ATTACHMENT_TYPE": 4,
}

func NewMailServiceError_ErrorCode(x int32) *MailServiceError_ErrorCode {
	e := MailServiceError_ErrorCode(x)
	return &e
}
func (x MailServiceError_ErrorCode) String() string {
	return proto.EnumName(MailServiceError_ErrorCode_name, int32(x))
}

type MailServiceError struct {
	XXX_unrecognized []byte
}

func (this *MailServiceError) Reset() {
	*this = MailServiceError{}
}

type MailAttachment struct {
	FileName         *string "PB(bytes,1,req)"
	Data             []byte  "PB(bytes,2,req)"
	XXX_unrecognized []byte
}

func (this *MailAttachment) Reset() {
	*this = MailAttachment{}
}

type MailMessage struct {
	Sender           *string           "PB(bytes,1,req)"
	ReplyTo          *string           "PB(bytes,2,opt)"
	To               []string          "PB(bytes,3,rep)"
	Cc               []string          "PB(bytes,4,rep)"
	Bcc              []string          "PB(bytes,5,rep)"
	Subject          *string           "PB(bytes,6,req)"
	TextBody         *string           "PB(bytes,7,opt)"
	HtmlBody         *string           "PB(bytes,8,opt)"
	Attachment       []*MailAttachment "PB(bytes,9,rep)"
	XXX_unrecognized []byte
}

func (this *MailMessage) Reset() {
	*this = MailMessage{}
}

func init() {
	proto.RegisterEnum("appengine.MailServiceError_ErrorCode", MailServiceError_ErrorCode_name, MailServiceError_ErrorCode_value)
}