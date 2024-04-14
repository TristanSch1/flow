package infra

import (
	"slices"
	"time"

	"github.com/TristanSch1/flow/internal/domain/session"
)

type InMemorySessionRepository struct {
	Sessions []session.Session
}

func (r *InMemorySessionRepository) Save(s session.Session) error {
	startedSessionIndex := slices.IndexFunc(r.Sessions, func(s session.Session) bool {
		return s.StartTime.Equal(s.StartTime)
	})

	if startedSessionIndex == -1 {
		r.Sessions = append(r.Sessions, s)
	} else {
		r.Sessions[startedSessionIndex] = s
	}

	return nil
}

func (r *InMemorySessionRepository) FindByStartTime(startTime time.Time) *session.Session {
	index := slices.IndexFunc(r.Sessions, func(s session.Session) bool {
		return s.StartTime.Equal(startTime)
	})

	return &r.Sessions[index]
}

func (r *InMemorySessionRepository) FindAllByProject(project string) ([]session.Session, error) {
	sessions := []session.Session{}

	for _, session := range r.Sessions {
		if session.Project == project {
			sessions = append(sessions, session)
		}
	}

	return sessions, nil
}

func (r *InMemorySessionRepository) FindLastSession() (*session.Session, error) {
	if len(r.Sessions) == 0 {
		return nil, nil
	}

	return &r.Sessions[len(r.Sessions)-1], nil
}

func (r *InMemorySessionRepository) FindAllSessions() ([]session.Session, error) {
	return r.Sessions, nil
}

func (r *InMemorySessionRepository) FindAllProjects() ([]string, error) {
	sessions, _ := r.FindAllSessions()

	projects := []string{}

	for _, session := range sessions {
		if slices.Contains(projects, session.Project) {
			continue
		}

		projects = append(projects, session.Project)
	}

	return projects, nil
}

func (r *InMemorySessionRepository) FindAllProjectTags(project string) ([]string, error) {
	sessionsForProject, _ := r.FindAllByProject(project)

	tags := []string{}

	for _, session := range sessionsForProject {
		for _, tag := range session.Tags {
			if slices.Contains(tags, tag) {
				continue
			}

			tags = append(tags, tag)
		}
	}

	return tags, nil
}
