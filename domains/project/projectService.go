package project

import (
	"context"
	"errors"
	"log"
	"twoBinPJ/domains/user"
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
	user, err := user.ForContext(ctx)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return nil, errors.New("INITIALIZING_TOKEN_ERROR")
	}
	if user.Role != 2 {
		log.Printf("You do not have access to create project: %s", err)
		return nil, errors.New("CREATE_PROJECT_ERROR")
	}
	return p.Repository.CreateProject(name, shortDescription, description)
}

func (p *ProjectService) ShowTheProjectByID(id int) (*Project, error) {
	return p.Repository.GetProjectByID(id)
}

func (p *ProjectService) UpdateProject(ctx context.Context, project *Project, name, shortDescription, description *string) error {
	user, err := user.ForContext(ctx)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return errors.New("INITIALIZING_TOKEN_ERROR")
	}
	if user.Role != 2 {
		log.Printf("You do not have access to update project: %s", err)
		return errors.New("UPDATE_PROJECT_ERROR")
	}
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

	err = p.Repository.UpdateProject(name, shortDescription, description, project.ID)

	if err != nil {
		return err
	}
	return nil
}

func (p *ProjectService) DeleteProject(ctx context.Context, id int) error {
	user, err := user.ForContext(ctx)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return errors.New("INITIALIZING_TOKEN_ERROR")
	}
	if user.Role != 2 {
		log.Printf("You do not have access to delete project: %s", err)
		return errors.New("DELETE_PROJECT_ERROR")
	}
	return p.Repository.DeleteProject(id)
}
