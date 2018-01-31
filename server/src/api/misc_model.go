package api

// StringIDReq 文字列ID用
type StringIDReq struct {
	ID string `json:"id,string" swagger:",in=path"`
}
