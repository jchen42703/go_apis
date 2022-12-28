package respond

type MessageResponse struct {
	Message string `json:"message"`
}

func NewMessage(msg string) MessageResponse {
	r := MessageResponse{}
	r.Message = msg
	return r
}
