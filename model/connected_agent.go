package model

import "time"

type ConnectedAgent struct {
	Uuid    string `json:"uuid"`
	Address string `address:"address"`

	JoinedAt time.Time `json:"joinedAt"`
	ExpireAt time.Time `json:"expireAt"`
}

func NewConnectedAgent(uuid string, address string, expireAt time.Time) *ConnectedAgent {
	return &ConnectedAgent{
		Uuid:     uuid,
		Address:  address,
		JoinedAt: time.Now(),
		ExpireAt: expireAt,
	}
}

func (c *ConnectedAgent) IsActive() bool {
	return !time.Now().After(c.ExpireAt)
}

func (c *ConnectedAgent) ExtendValidity(duration time.Duration) {
	c.ExpireAt = time.Now().Add(duration)
}
