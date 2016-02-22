package model
import "time"

type RealAgent struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

type DesiredAgent struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Tarball string `json:"tarball"`
	Md5     string `json:"md5"`
	Cmd     string `json:"cmd"`
}

type HeartbeatReques struct {
	Hostname   string       `json:"hostname"`
	RealAgents []*RealAgent `json:"realAgents"`
}

type HeartbeatResponse struct {
	ErrorMessage string          `json:"errorMessage"`
	DesiredAgent []*DesiredAgent `json:"desiredAgents"`
}

//----------------------------sql model-----------------------------------
type Reminder struct {
	Id        int64     `json:"id"`
	Message   string    `sql:"size:1024" json:"message"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"-"`
}