package consts

type RecordStatus string

const (
	RecordExistStatus    = RecordStatus("EXIST")
	RecordNotExistStatus = RecordStatus("NOT EXIST")
	RecordDeletedStatus  = RecordStatus("DELETED")
)
