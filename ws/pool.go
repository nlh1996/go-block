package ws

// ConnPool .
type ConnPool struct {
	Pool map[int]*Connection
}

var instance *ConnPool

// GetInstance .
func GetInstance() *ConnPool {
	if instance == nil {
		instance = &ConnPool{}
		instance.Pool = make(map[int]*Connection)
	}
	return instance
}
