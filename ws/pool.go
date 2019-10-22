package ws

// ConnPool .
type ConnPool struct {
	Pool []*Connection
}

var instance *ConnPool

// GetInstance .
func GetInstance() *ConnPool {
	if instance == nil {
		instance = &ConnPool{}
	}
	return instance
}
