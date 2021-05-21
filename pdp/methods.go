package pdp

const (
	InsertMethod HandleMethod = iota
	UpdateMethod
	UpsertMethod

	UpdateMethodProperty = "CanAddByUpdate"
)

type HandleMethod int
