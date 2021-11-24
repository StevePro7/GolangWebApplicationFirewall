package checker

import "google.golang.org/genproto/googleapis/rpc/code"

var OK = int32(code.Code_OK)
var PERMISSION_DENIED = int32(code.Code_PERMISSION_DENIED)
var UNAVAILABLE = int32(code.Code_UNAVAILABLE)
var INVALID_ARGUMENT = int32(code.Code_INVALID_ARGUMENT)
var INTERNAL = int32(code.Code_INTERNAL)
