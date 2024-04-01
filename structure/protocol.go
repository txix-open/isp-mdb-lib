package structure

type ProtocolVersion string

const (
	SudirV1      ProtocolVersion = "sudir_v1"
	SudirV1Find  ProtocolVersion = "sudir_v1_find"
	SudirV2      ProtocolVersion = "sudir_v2"
	SudirV2Find  ProtocolVersion = "sudir_v2_find"
	JsonProtocol ProtocolVersion = "json"
	ErlProtocol  ProtocolVersion = "erl"
	SmaProtocol  ProtocolVersion = "sma"
)
