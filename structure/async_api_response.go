package structure

import "time"

type GetJobStatusResponse struct {
	RequestId       string
	Status          string `json:",omitempty"`
	CreatedAt       time.Time
	FinishedAt      time.Time `json:",omitempty"`
	ExecutedTime    string    `json:",omitempty"`
	TtlUntil        time.Time `json:",omitempty"`
	Description     string
	PackagesCount   int
	ExecutedEntries int64
}
