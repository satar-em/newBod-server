package jwt

import "time"

type Authentication struct {
	id        string
	userId    uint
	data      map[string]interface{}
	lastUse   time.Time
	expireDur time.Duration
}

func (a *Authentication) Fresh() {
	a.lastUse = time.Now()
}
func (a *Authentication) IsExpire() bool {
	if a.expireDur <= 0 {
		return false
	}
	return time.Now().After(a.lastUse.Add(a.expireDur))
}

func (a *Authentication) SetValue(id string, value interface{}) {
	a.data[id] = value
}
func (a *Authentication) GetValue(id string) interface{} {
	return a.data[id]
}

func (a *Authentication) SetExpireDur(dur time.Duration) {
	a.expireDur = dur
}
func (a *Authentication) GetId() *string {
	return &a.id
}
func (a *Authentication) GetUser() *uint {
	return &a.userId
}
