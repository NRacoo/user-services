package constants

import "net/textproto"

var (
	XServiceName  = textproto.CanonicalMIMEHeaderKey("x-service-name")
	XApikey       = textproto.CanonicalMIMEHeaderKey("x-apikey")
	XRequestAt     = textproto.CanonicalMIMEHeaderKey("x-request-at")
	Authorization = textproto.CanonicalMIMEHeaderKey("authorization")
)
