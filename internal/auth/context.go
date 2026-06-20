package auth

import (
	"sync"

	"github.com/shubhamranswal/ciphergate/internal/session"
	"github.com/shubhamranswal/ciphergate/internal/user"
)

type Context struct {
	mu      sync.RWMutex
	User    *user.User
	Session *session.Session
}

func (c *Context) Login(user *user.User, session *session.Session) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.User = user
	c.Session = session
}

func (c *Context) Logout() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.User = nil
	c.Session = nil
}

func (c *Context) IsAuthenticated() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.User != nil &&
		c.Session != nil
}
