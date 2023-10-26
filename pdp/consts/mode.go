package consts

type UpdateMode string

const (
	UpdateModeInsert = UpdateMode("insert")
	UpdateModeUpdate = UpdateMode("update")
	UpdateModeUpsert = UpdateMode("upsert")
)
