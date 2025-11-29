package openim

import (
	"sync"
	"time"
)

type TokenProvider interface {
	GetToken() (token string, err error)
	SetToken(token string, exp int64) error
}

type MemoryTokenStore struct {
	tokenLock  sync.RWMutex
	token      string
	expireTime time.Time
}

// getToken implements TokenProvider.
func (c *MemoryTokenStore) GetToken() (token string, err error) {
	c.tokenLock.RLock()
	defer c.tokenLock.RUnlock()

	if c.expireTime.IsZero() || c.expireTime.After(time.Now()) {
		return
	}

	return c.token, nil
}

// setToken implements TokenProvider.
func (c *MemoryTokenStore) SetToken(token string, exp int64) error {
	c.tokenLock.Lock()
	defer c.tokenLock.Unlock()
	c.token = token
	c.expireTime = time.Now().Add(time.Duration(exp) * time.Second)
	return nil
}

type TokenGetter func() (token string, err error)
type TokenSetter func(token string, exp int64) error

type FuncTokenProvider struct {
	getter TokenGetter
	setter TokenSetter
}

func NewFuncTokenProvider(getter TokenGetter, setter TokenSetter) *FuncTokenProvider {
	return &FuncTokenProvider{
		getter: getter,
		setter: setter,
	}
}

func (c *FuncTokenProvider) GetToken() (token string, err error) {
	return c.getter()
}

func (c *FuncTokenProvider) SetToken(token string, exp int64) error {
	return c.setter(token, exp)
}
