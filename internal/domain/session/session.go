package session

import (
	"time"
)

const (
	FlowingStatus = "FLOWING"
	EndedStatus   = "ENDED"
)

type Session struct {
	Id        string
	StartTime time.Time
	EndTime   time.Time
	Project   string
	Tags      []string
}

func (s Session) GetFormattedStartTime() string {
	return s.StartTime.Format(time.DateTime)
}

func (s Session) GetFormattedEndTime() string {
	if s.EndTime.IsZero() {
		return "/"
	}

	return s.EndTime.Format(time.DateTime)
}

func (s Session) Duration() time.Duration {
	if s.EndTime.IsZero() {
		return 0
	}
	return s.EndTime.Sub(s.StartTime).Round(time.Second)
}

func (s Session) Status() string {
	if s.EndTime.IsZero() {
		return FlowingStatus
	}

	return EndedStatus
}

func (s Session) Equals(session Session) bool {
	return s.Id == session.Id
}

func (s Session) HasTag(tag string) bool {
	for _, t := range s.Tags {
		if t == tag {
			return true
		}
	}
	return false
}
