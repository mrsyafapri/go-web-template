package apps

type ResponseSuccess struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type ResponseFail struct {
	Status  int         `json:"status"`
	Message interface{} `json:"message"`
}
