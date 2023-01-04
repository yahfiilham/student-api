package http

type Resp struct {
	Code int         `json:"-"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data"`
}
