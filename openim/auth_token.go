package openim

import (
	"sync"
	"time"
)

type TokenProvider interface {
	getToken() (token string, err error)
	setToken(token string, exp int64) error
}

type MemoryTokenStore struct {
	tokenLock  sync.RWMutex
	token      string
	expireTime time.Time
}

// getToken implements TokenProvider.
func (c *MemoryTokenStore) getToken() (token string, err error) {
	c.tokenLock.RLock()
	defer c.tokenLock.RUnlock()

	if c.expireTime.IsZero() || c.expireTime.After(time.Now()) {
		return
	}

	return c.token, nil
}

// setToken implements TokenProvider.
func (c *MemoryTokenStore) setToken(token string, exp int64) error {
	c.tokenLock.Lock()
	defer c.tokenLock.Unlock()
	c.token = token
	c.expireTime = time.Now().Add(time.Duration(exp) * time.Second)
	return nil
}
