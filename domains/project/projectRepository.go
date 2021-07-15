package project

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"time"
)

type ProjectRepository struct {
	DB *sqlx.DB
}

func NewProjectRepository(Db *sqlx.DB) *ProjectRepository {
	return &ProjectRepository{DB: Db}
}

func (p *ProjectRepository) GetProjectByField(field, value string) (*Project, error) {
	var project *Project
	row, err := p.DB.Query(fmt.Sprintf("select*from project where %v=$1", field), value)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var p Project
		err = row.Scan(&p.ID, &p.Name, &p.ShortDescription, &p.Description, &p.Private, &p.Closed, &p.OrgID, &p.Created, &p.Updated)

		if err != nil {
			return nil, err
		}

		project = &p
	}
	return project, nil
}

func (p *ProjectRepository) GetProjectByID(id int) (*Project, error) {
	return p.GetProjectByField("id", strconv.Itoa(id))
}

func (p *ProjectRepository) GetProjectByName(name string) (*Project, error) {
	return p.GetProjectByField("name", name)
}

func (p *ProjectRepository) GetProjectByShortDescription(shortDescription string) (*Project, error) {
	return p.GetProjectByField("short_description", shortDescription)
}

func (p *ProjectRepository) GetProjectByShortOrgID(OrgID int) (*Project, error) {
	return p.GetProjectByField("org_id", strconv.Itoa(OrgID))
}

func (p *ProjectRepository) CreateProject(name, shortDescription, description string) (*Project, error) {
	rows, err := p.DB.Query("insert into project(name, short_description, description,private,closed,org_id,created, updated) values($1,$2,$3,$4,$5,$6,$7,$8) returning *",
		name, shortDescription, description, "true", "true", 1, time.Now(), time.Now())
	if err != nil {
		return nil, err
	}

	var proj *Project
	for rows.Next() {
		var p Project
		err = rows.Scan(&p.ID, &p.Name, &p.ShortDescription, &p.Description, &p.Private, &p.Closed, &p.OrgID, &p.Created, &p.Updated)
		if err != nil {
			return nil, err
		}

		proj = &p
	}
	return proj, nil
}

func (p *ProjectRepository) UpdateProject(name, shortDescription, description *string, id int) error {
	_, err := p.DB.Query("update project set name=$1, short_description=$2, description=$3,updated=$4 where id=$5 ",
		name, shortDescription, description, time.Now(), id)
	return err
}

func (p *ProjectRepository) DeleteProject(id int) error {
	query := "Delete from project where id=$1"
	_, err := p.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
