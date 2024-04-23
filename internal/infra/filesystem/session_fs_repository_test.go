package filesystem_test

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"

	"github.com/TristanSch1/flow/internal/domain/session"
	"github.com/TristanSch1/flow/internal/infra/filesystem"
	"github.com/TristanSch1/flow/pkg/timerange"
)

const (
	TestFolderPath = "./"
)

func setup() {
	os.RemoveAll("./.flow")
}

func TestConstructorCreateFolder_Success(t *testing.T) {
	setup()

	filesystem.NewFileSystemSessionRepository(TestFolderPath)

	path := filepath.Join(TestFolderPath, filesystem.FlowFolderName)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		t.Errorf("File %v not found at location %v", filesystem.FlowFolderName, path)
	}
}

func TestSave_Success(t *testing.T) {
	setup()

	repository := filesystem.NewFileSystemSessionRepository(TestFolderPath)

	repository.Save(session.Session{
		Id:        "1",
		StartTime: time.Now(),
		Project:   "Flow",
	})

	path := filepath.Join(TestFolderPath, filesystem.FlowFolderName)
	if _, err := os.Stat(filepath.Join(path, "1.json")); os.IsNotExist(err) {
		t.Errorf("Session with ID id1 is not correctly saved")
	}
}

func TestFindAllSessions_Success(t *testing.T) {
	setup()

	repository := filesystem.NewFileSystemSessionRepository(TestFolderPath)

	repository.Save(session.Session{
		Id:        "1",
		StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"test-save"},
	})

	repository.Save(session.Session{
		Id:        "2",
		StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
		Project:   "Flow",
	})

	got := repository.FindAllSessions()

	want := []session.Session{
		{
			Id:        "1",
			StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
			Project:   "Flow",
			Tags:      []string{"test-save"},
		},
		{
			Id:        "2",
			StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
			Project:   "Flow",
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FileSystemSessionRepository.FindAll() = %v, want %v", got, want)
	}
}

func TestFindAllSessions_NoSessions_Success(t *testing.T) {
	setup()

	repository := filesystem.NewFileSystemSessionRepository(TestFolderPath)

	got := repository.FindAllSessions()

	want := []session.Session{}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FileSystemSessionRepository.FindAll() = %v, want %v", got, want)
	}
}

func TestFindLastSession_Success(t *testing.T) {
	setup()

	repository := filesystem.NewFileSystemSessionRepository(TestFolderPath)

	repository.Save(session.Session{
		Id:        "1",
		StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"test-save"},
	})

	repository.Save(session.Session{
		Id:        "2",
		StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
		Project:   "Flow",
	})

	got := repository.FindLastSession()

	want := session.Session{
		Id:        "2",
		StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
		Project:   "Flow",
	}

	if !reflect.DeepEqual(*got, want) {
		t.Errorf("FileSystemSessionRepository.FindLastSession() = %v, want %v", *got, want)
	}
}

func TestFindAllProjects(t *testing.T) {
	setup()

	repository := filesystem.NewFileSystemSessionRepository(TestFolderPath)

	repository.Save(session.Session{
		Id:        "1",
		StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"test"},
	})

	repository.Save(session.Session{
		Id:        "2",
		StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 22, 0, 0, 0, time.UTC),
		Project:   "MyTodo",
	})

	repository.Save(session.Session{
		Id:        "3",
		StartTime: time.Date(2024, 4, 17, 23, 0, 0, 0, time.UTC),
		Project:   "MyTodo",
	})

	got, _ := repository.FindAllProjects()

	want := []string{"Flow", "MyTodo"}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("FileSystemSessionRepository.FindAllProjects() = %v, want %v", got, want)
	}
}

func TestFindAllProjectsTags(t *testing.T) {
	setup()

	repository := filesystem.NewFileSystemSessionRepository(TestFolderPath)

	repository.Save(session.Session{
		Id:        "1",
		StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"tests", "integration"},
	})

	repository.Save(session.Session{
		Id:        "2",
		StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 22, 0, 0, 0, time.UTC),
		Project:   "MyTodo",
		Tags:      []string{"add-todo", "update-todo"},
	})

	repository.Save(session.Session{
		Id:        "3",
		StartTime: time.Date(2024, 4, 17, 23, 0, 0, 0, time.UTC),
		Project:   "MyTodo",
		Tags:      []string{"update-todo", "delete-todo"},
	})

	tests := []struct {
		name string
		want []string
	}{
		{
			name: "Flow",
			want: []string{"tests", "integration"},
		},
		{
			name: "MyTodo",
			want: []string{"add-todo", "update-todo", "delete-todo"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, _ := repository.FindAllProjectTags(tt.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileSystemSessionRepository.FindAllProjectTags() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindInTimeRange(t *testing.T) {
	setup()
	repository := filesystem.NewFileSystemSessionRepository(TestFolderPath)
	repository.Save(session.Session{
		Id:        "1",
		StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
		Project:   "Flow",
		Tags:      []string{"tests", "integration"},
	})
	repository.Save(session.Session{
		Id:        "2",
		StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
		EndTime:   time.Date(2024, 4, 17, 23, 0, 0, 0, time.UTC),
		Project:   "MyTodo",
		Tags:      []string{"add-todo", "update-todo"},
	})
	repository.Save(session.Session{
		Id:        "3",
		StartTime: time.Date(2024, 4, 18, 21, 0, 0, 0, time.UTC),
		Project:   "MyTodo",
		Tags:      []string{"delete-todo"},
	})

	tests := []struct {
		name string
		args timerange.TimeRange
		want []session.Session
	}{
		{
			name: "All",
			args: timerange.TimeRange{},
			want: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"tests", "integration"},
				},
				{
					Id:        "2",
					StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 17, 23, 0, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"add-todo", "update-todo"},
				},
				{
					Id:        "3",
					StartTime: time.Date(2024, 4, 18, 21, 0, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"delete-todo"},
				},
			},
		},
		{
			name: "Since",
			args: timerange.TimeRange{
				Since: time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
			},
			want: []session.Session{
				{
					Id:        "2",
					StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 17, 23, 0, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"add-todo", "update-todo"},
				},
				{
					Id:        "3",
					StartTime: time.Date(2024, 4, 18, 21, 0, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"delete-todo"},
				},
			},
		},
		{
			name: "Until",
			args: timerange.TimeRange{
				Until: time.Date(2024, 4, 17, 20, 1, 0, 0, time.UTC),
			},
			want: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"tests", "integration"},
				},
			},
		},
		{
			name: "Since and Until",
			args: timerange.TimeRange{
				Since: time.Date(2024, 4, 17, 17, 0, 0, 0, time.UTC),
				Until: time.Date(2024, 4, 17, 22, 0, 0, 0, time.UTC),
			},
			want: []session.Session{
				{
					Id:        "1",
					StartTime: time.Date(2024, 4, 17, 19, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 17, 20, 0, 0, 0, time.UTC),
					Project:   "Flow",
					Tags:      []string{"tests", "integration"},
				},
				{
					Id:        "2",
					StartTime: time.Date(2024, 4, 17, 21, 0, 0, 0, time.UTC),
					EndTime:   time.Date(2024, 4, 17, 23, 0, 0, 0, time.UTC),
					Project:   "MyTodo",
					Tags:      []string{"add-todo", "update-todo"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repository.FindInTimeRange(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FileSystemSessionRepository.FindInTimeRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
