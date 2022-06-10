package umeng

// Result 结果对象
type Result struct {
	Code string            `json:"ret,omitempty"`
	Data map[string]string `json:"data,omitempty"`
}

func (r *Result) Ok() bool {
	return r.Code == "SUCCESS"
}
