package structure

import (
	"fmt"
)

type ApiResponse struct {
	ErrorCode    int
	ErrorMessage string `json:",omitempty"`
	RequestId    string
	FoundRecords int64       `json:",omitempty"`
	Data         interface{} `json:",omitempty"`
}

func (r *ApiResponse) GetError() error {
	if r.ErrorCode != 0 {
		return fmt.Errorf("error code %d, message %q", r.ErrorCode, r.ErrorMessage)
	}
	return nil
}
