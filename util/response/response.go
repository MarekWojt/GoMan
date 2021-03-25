package response

// Response is the response of an api call
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

func New(code int, data interface{}) Response {
	return Response{code, data}
}

func Ok(data interface{}) Response {
	return Response{200, data}
}
