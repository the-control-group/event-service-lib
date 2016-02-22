package lib

import (
    "encoding/json"
    "time"
)

type Event interface {
    ID() string
    Type() string
    Session() Session
    Customer() Customer
    Created() time.Time
}

type EventV1 struct {
	ID string `json:"event_id"`
	Type string `json:"event_name"`
	SessionID string `json:"session_id"`
	Created time.Time `json:"created"`
}

func (e *EventV1) ID() string {return c.ID}
func (e *EventV1) Type() string {return c.Type}
func (e *EventV1) Session() Session {return SessionV1{c.SessionID}}
func (e *EventV1) Created() string {return c.Created}

type EventV2 struct {
	Created time.Time `json:"created"`
	Name `json:"event_name"`
	Customer json.RawMessage `json:"customer"`
	Order json.RawMessage `json:"order"`
	Transaction json.RawMessage `json:"transaction"`
	Device json.RawMessage `json:"device"`
	Location json.RawMessage `json:"location"`
	Lead json.RawMessage `json:"lead"`
	PaymenOption json.RawMessage `json:"payment_option"`
	Page json.RawMessage `json:"page"`
	TrafficSource json.RawMessage `json:"traffic_source"`
}

type Session interface {
    ID() string
    Type() string
    Created() time.Time
}

type SessionV1 struct {
	ID string `json:"session_id"`
	Type string `json:"session_type"`
	Created time.Time `json:"created"`
}

func (e *SessionV1) ID() string {return c.ID}
func (e *SessionV1) Type() string {return c.Type}
func (e *SessionV1) Created() string {return c.Created}

type Customer interface {
    ID() int64
    UUID() string
    FirstName() string
    LastName() string
    Email() string
}

type CustomerV1 struct {
    ID int64 `json:"id"`
    UUID string `json:"uuid"`
    FirstName string `json:"first_name"`
    LastName string `json:"last_name"`
    Email string `json:"email"`
}

func (c *CustomerV1) ID() int64 {return c.ID}
func (c *CustomerV1) UUID() string {return c.UUID}
func (c *CustomerV1) FirstName() string {return c.FirstName}
func (c *CustomerV1) LastName() string {return c.LastName}
func (c *CustomerV1) Email() string {return c.Email}