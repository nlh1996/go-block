package ws

import "sync"

// ConnPool .
type ConnPool struct {
	Pool map[int]*Connection
	sync.RWMutex
}

var instance *ConnPool

// GetConnPool .
func GetConnPool() *ConnPool {
	if instance == nil {
		instance = &ConnPool{}
		instance.Pool = make(map[int]*Connection)
	}
	return instance
}

// Set .
func (p *ConnPool) Set(c *Connection) {
	p.Lock()
	defer p.Unlock()
	p.Pool[c.ID] = c
}

// GetConnByID .
func (p *ConnPool) GetConnByID(id int) *Connection {
	p.Lock()
	defer p.Unlock()
	v, ok := p.Pool[id]
	if ok {
		return v
	}
	return nil
}

// DelByID .
func (p *ConnPool) DelByID(id int) {
	p.Lock()
	defer p.Unlock()
	delete(p.Pool, id)
}
