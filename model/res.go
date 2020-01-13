package model

// Response .
type Response struct {
	Code uint8  `json:"code"`
	Msg  string `json:"msg"`
}

const (
	addblock = 101
	nextblock = 102
)

// Code .
type Code struct {}

// Handle .
func (r *Response) Handle() {

}

func (c *Code) exec() {

}