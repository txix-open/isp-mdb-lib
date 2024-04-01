package delta

type Operation string

const (
	ChangeField     = Operation("CHANGE_FIELD")
	DeleteField     = Operation("DELETE_FIELD")
	DeleteArrayItem = Operation("DELETE_ARRAY_ITEM")
	AddArrayItem    = Operation("ADD_ARRAY_ITEM")
	DeleteBlock     = Operation("DELETE_BLOCK")
	ExtraOperation  = Operation("EXTRA")
)

type Changelog struct {
	Path      Path
	Field     string
	Operation Operation
	NewValue  interface{}
}

type Path struct {
	BlockName string
	IsArray   bool
	ItemId    string
}
