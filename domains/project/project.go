package project

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
	"twoBinPJ/domains/user"
)

type Project struct {
	ID               int       `json:"id"`
	Name             string    `json:"name"`
	ShortDescription string    `json:"shortDescription"`
	Description      string    `json:"description"`
	Private          bool      `json:"private"`
	Closed           bool      `json:"closed"`
	OrgID            int       `json:"orgId"`
	Created          time.Time `json:"created"`
	Updated          time.Time `json:"updated"`
}

type IProjectRepository interface {
	GetProjectByField(field, value string) (*Project, error)
	GetProjectByID(id int) (*Project, error)
	GetProjectByShortDescription(shortDescription string) (*Project, error)
	GetProjectByName(name string) (*Project, error)
	GetProjectByShortOrgID(OrgID int) (*Project, error)
	CreateProject(name, shortDescription, description string) (*Project, error)
	UpdateProject(name, shortDescription, description *string, id int) error
	DeleteProject(id int) error
}

type IProjectService interface {
	GetProjectByNameService(name string) (*Project, error)
	CreateProjectService(ctx context.Context, name, shortDescription, description string) (*Project, error)
	ShowTheProjectByID(id int) (*Project, error)
	UpdateProject(ctx context.Context, project *Project, name, shortDescription, description *string) error
	DeleteProject(ctx context.Context, id int) error
}

type ProjectModule struct {
	IProjectService
}

func NewProjectModule(Db *sqlx.DB) *ProjectModule {
	userRepository := user.NewUserPostgres(Db)
	userService := user.NewUserService(userRepository)
	projectRepository := NewProjectRepository(Db)
	return &ProjectModule{
		IProjectService: NewProjectService(projectRepository, userService),
	}
}
