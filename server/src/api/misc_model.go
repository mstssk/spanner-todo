package api

// Int64IDReq int64のID用
type Int64IDReq struct {
	ID int64 `json:"id,string" swagger:",in=path"`
}
