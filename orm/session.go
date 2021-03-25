package orm

import (
	"time"

	"github.com/MarekWojt/GoMan/util/security"
)

type sessionPool map[string]*Session

// Sessions contains the sessions
var Sessions sessionPool = make(map[string]*Session)

// Session holds the session data
type Session struct {
	User    User
	Updated time.Time
}

// ClearSessions deletes sessions that need to be deleted (end of lifetime)
func (sessions *sessionPool) ClearSessions() {
	for key, element := range *sessions {
		if time.Now().Add(24 * time.Hour).After(element.Updated) {
			delete(*sessions, key)
		}
	}
}

// CreateSession creates a session
func (sessions *sessionPool) CreateSession() (key string, session *Session, err error) {
	key, err = security.GenerateRandomString(32)
	if err != nil {
		return
	}

	session = &Session{
		Updated: time.Now(),
	}

	(*sessions)[key] = session
	return
}

// GetSession returns the session by key and replaces this key
func (sessions *sessionPool) GetSession(oldKey string) (key string, session *Session, err error) {
	if oldKey == "" {
		key, session, err = sessions.CreateSession()
		return
	}

	key = oldKey

	session, ok := (*sessions)[key]
	if !ok {
		key, session, err = sessions.CreateSession()
		return
	}

	key, err = security.GenerateRandomString(32)
	if err != nil {
		return
	}

	(*sessions)[key] = session

	delete(*sessions, oldKey)

	return
}
