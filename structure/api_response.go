package structure

type ApiResponse struct {
	ErrorCode    int
	ErrorMessage string `json:",omitempty"`
	RequestId    string
	FoundRecords int64       `json:",omitempty"`
	Data         interface{} `json:",omitempty"`
}
