package report

import (
	"context"
	"errors"
	"log"
	"twoBinPJ/domains/user"
)

type ReportService struct {
	Repository IReportRepository
}

func NewReportService(repo IReportRepository) *ReportService {
	return &ReportService{Repository: repo}
}

func (r *ReportService) GetReportByNameService(name string) (*Report, error) {
	return r.Repository.GetReportByName(name)
}

func (r *ReportService) ShowTheReportByID(id int) (*Report, error) {
	return r.Repository.GetReportByID(id)
}

func (r *ReportService) CreateReportService(ctx context.Context, name, description, comments, seriousness string) (*Report, error) {
	user, err := user.ForContext(ctx)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return nil, errors.New("INITIALIZING_TOKEN_ERROR")
	}
	if user.Role != 1 {
		log.Printf("You do not have access to create report: %s", err)
		return nil, errors.New("CREATE_REPORT_ERROR")
	}
	return r.Repository.CreateReport(name, description, comments, seriousness, user.Id)
}

func (r *ReportService) UpdateReport(name, description, comments, seriousness *string, report *Report, ctx context.Context) error {
	user, err := user.ForContext(ctx)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return errors.New("INITIALIZING_TOKEN_ERROR")
	}
	if user.Role != 1 {
		log.Printf("You do not have access to update report: %s", err)
		return errors.New("UPDATE_REPORT_ERROR")
	}
	didUpdateName := false
	didUpdateDescription := false
	didUpdateComments := false
	didUpdateSeriousness := false

	if len(*name) < 1 {
		*name = report.Name
		didUpdateName = true
	}

	if len(*description) < 1 {
		*description = report.Description
		didUpdateDescription = true
	}

	if len(*comments) < 1 {
		*comments = report.Comments
		didUpdateComments = true
	}

	if len(*seriousness) < 1 {
		*seriousness = string(report.Seriousness)
		didUpdateSeriousness = true
	}

	if didUpdateName && didUpdateDescription && didUpdateComments && didUpdateSeriousness {
		return errors.New("no update done")
	}

	err = r.Repository.UpdateReport(name, description, comments, seriousness, report.ID, didUpdateComments)

	if err != nil {
		return err
	}
	return nil
}

func (r *ReportService) DeleteReport(ctx context.Context, id int) error {
	user, err := user.ForContext(ctx)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return errors.New("INITIALIZING_TOKEN_ERROR")
	}
	if user.Role != 1 {
		log.Printf("You do not have access to delete report: %s", err)
		return errors.New("DELETE_REPORT_ERROR")
	}
	return r.Repository.DeleteReport(id)
}
