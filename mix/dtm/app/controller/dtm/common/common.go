package common

const (
	DefaultHTTPServer = "http://dtm:36789/api/dtmsvr"
	DefaultGrpcServer = "dtm:36790"
)

var SagaTransInAction = "http://localhost:8001/trans/saga/in"
var SagaTransInRevokeAction = "http://localhost:8001/trans/saga/in/revoke"
var SagaTransOutAction = "http://localhost:8001/trans/saga/out"
var SagaTransOutRevokeAction = "http://localhost:8001/trans/saga/out/revoke"

var TccTransInAction = "http://localhost:8001/trans/tcc/in"
var TccTransInConfirmAction = "http://localhost:8001/trans/tcc/in/confirm"
var TccTransInRevokeAction = "http://localhost:8001/trans/tcc/in/revoke"
var TccTransOutAction = "http://localhost:8001/trans/tcc/out"
var TccTransOutConfirmAction = "http://localhost:8001/trans/tcc/out/confirm"
var TccTransOutRevokeAction = "http://localhost:8001/trans/tcc/out/revoke"


var XATransInAction = "http://localhost:8001/trans/xa/in"
var XATransOutAction = "http://localhost:8001/trans/xa/out"

var MsgTransInAction = "http://localhost:8001/trans/msg/in"
var MsgTransOutAction = "http://localhost:8001/trans/msg/out"