package report

import "context"

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

func (r *ReportService) CreateReportService(ctx context.Context, name, description, comments, seriousness string, unreadComments bool) (*Report, error) {
	return r.Repository.CreateReport(name, description, comments, seriousness, unreadComments)
}
