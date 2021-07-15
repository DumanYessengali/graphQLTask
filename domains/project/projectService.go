package project

import (
	"context"
	"errors"
)

type ProjectService struct {
	Repository IProjectRepository
}

func NewProjectService(repository IProjectRepository) *ProjectService {
	return &ProjectService{Repository: repository}
}

func (p *ProjectService) GetProjectByNameService(name string) (*Project, error) {
	return p.Repository.GetProjectByName(name)
}

func (p *ProjectService) CreateProjectService(ctx context.Context, name, shortDescription, description string) (*Project, error) {
	return p.Repository.CreateProject(name, shortDescription, description)
}

func (p *ProjectService) ShowTheProjectByID(ctx context.Context, id int) (*Project, error) {
	return p.Repository.GetProjectByID(id)
}

func (p *ProjectService) UpdateProject(project *Project, name, shortDescription, description *string) error {
	didUpdateName := false
	didUpdateShortDescription := false
	didUpdateDescription := false
	if len(*name) < 1 {
		*name = project.Name
		didUpdateName = true

	}
	if len(*shortDescription) < 1 {
		*shortDescription = project.ShortDescription
		didUpdateShortDescription = true

	}
	if len(*description) < 1 {
		*description = project.Description
		didUpdateDescription = true
	}
	if didUpdateName && didUpdateShortDescription && didUpdateDescription {
		return errors.New("no update done")
	}
	err := p.Repository.UpdateProject(name, shortDescription, description, project.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProjectService) DeleteProject(ctx context.Context, id int) error {
	return p.Repository.DeleteProject(id)
}
