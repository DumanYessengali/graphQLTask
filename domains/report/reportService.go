package report

import (
	"context"
	"errors"
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

func (r *ReportService) ShowTheReportByID(ctx context.Context, id int) (*Report, error) {
	return r.Repository.GetReportByID(id)
}

func (r *ReportService) CreateReportService(ctx context.Context, name, description, comments, seriousness string) (*Report, error) {
	return r.Repository.CreateReport(name, description, comments, seriousness)
}

func (r *ReportService) UpdateReport(name, description, comments, seriousness *string, report *Report, ctx context.Context) error {
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

	err := r.Repository.UpdateReport(name, description, comments, seriousness, report.ID, didUpdateComments)

	if err != nil {
		return err
	}
	return nil
}

func (r *ReportService) DeleteReport(ctx context.Context, id int) error {
	return r.Repository.DeleteReport(id)
}
