package pkg

type CtxInfo struct {
	Limit     string `json:"limit,omitempty"`
	Offset    string `json:"offset,omitempty"`
	UserId    string `json:"user_id,omitempty"`
	RequestId string `json:"request_id,omitempty"`
}
