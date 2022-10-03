package mnemosyne

const (
	Version = "1.0"
)

type Response struct {
	Version string   `json:"version"`
	Texts   []string `json:"texts"`
}

func NewResponse() *Response {
	return &Response{
		Version: Version,
	}
}
