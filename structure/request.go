package structure

type Identities struct {
	Ids []int64 `json:"ids" valid:"required~Required"`
}
