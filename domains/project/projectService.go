package project

import (
	"context"
	"errors"
	"log"
	"twoBinPJ/domains/user"
)

type ProjectService struct {
	Repository  IProjectRepository
	UserService user.IUserService
}

func NewProjectService(repository IProjectRepository, userService *user.UserService) *ProjectService {
	return &ProjectService{
		Repository:  repository,
		UserService: userService,
	}
}

func (p *ProjectService) GetProjectByNameService(name string) (*Project, error) {
	return p.Repository.GetProjectByName(name)
}

func (p *ProjectService) CreateProjectService(ctx context.Context, name, shortDescription, description string) (*Project, error) {
	users, err := user.ForContext(ctx)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return nil, errors.New("INITIALIZING_TOKEN_ERROR")
	}
	currentUser, err := p.UserService.GetUserByIDService(users.Id)
	if err != nil {
		log.Printf("error while initializing user error")
		return nil, err
	}
	if currentUser.Role != 2 {
		log.Printf("You do not have access to create project: %s", err)
		return nil, errors.New("CREATE_PROJECT_ERROR")
	}
	return p.Repository.CreateProject(name, shortDescription, description)
}

func (p *ProjectService) ShowTheProjectByID(id int) (*Project, error) {
	return p.Repository.GetProjectByID(id)
}

func (p *ProjectService) UpdateProject(ctx context.Context, project *Project, name, shortDescription, description *string) error {
	users, err := user.ForContext(ctx)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return errors.New("INITIALIZING_TOKEN_ERROR")
	}
	currentUser, err := p.UserService.GetUserByIDService(users.Id)
	if err != nil {
		log.Printf("error while initializing user error")
		return err
	}
	if currentUser.Role != 2 {
		log.Printf("You do not have access to update project: %s", err)
		return errors.New("UPDATE_PROJECT_ERROR")
	}
	didUpdate := false
	if name == nil {
		project.Name = project.Name
	} else {
		project.Name = *name
		didUpdate = true
	}

	if shortDescription == nil {
		project.ShortDescription = project.ShortDescription
	} else {
		project.ShortDescription = *shortDescription
		didUpdate = true
	}

	if description == nil {
		project.Description = project.Description
	} else {
		project.Description = *description
		didUpdate = true
	}

	if !didUpdate {
		return errors.New("no update done")
	}

	err = p.Repository.UpdateProject(name, shortDescription, description, project.ID)

	if err != nil {
		return err
	}
	return nil
}

func (p *ProjectService) DeleteProject(ctx context.Context, id int) error {
	users, err := user.ForContext(ctx)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return errors.New("INITIALIZING_TOKEN_ERROR")
	}
	currentUser, err := p.UserService.GetUserByIDService(users.Id)
	if err != nil {
		log.Printf("error while initializing user error")
		return err
	}
	if currentUser.Role != 2 {
		log.Printf("You do not have access to delete project: %s", err)
		return errors.New("DELETE_PROJECT_ERROR")
	}
	return p.Repository.DeleteProject(id)
}
