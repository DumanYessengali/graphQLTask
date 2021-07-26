package report

import (
	"context"
	"github.com/jmoiron/sqlx"
	"time"
	"twoBinPJ/domains/user"
)

type Report struct {
	ID              int               `json:"Id"`
	Name            string            `json:"Name"`
	Description     string            `json:"Description"`
	Status          ReportStatus      `json:"Status"`
	Seriousness     ReportSeriousness `json:"Seriousness"`
	Archive         bool              `json:"Archive"`
	Delete          bool              `json:"Delete"`
	Reward          int               `json:"Reward"`
	Point           int               `json:"Point"`
	ProjectID       int               `json:"ProjectId"`
	VulnerabilityID int               `json:"VulnerabilityId"`
	UserID          int               `json:"UserId"`
	Assignee        int               `json:"Assignee"`
	UnreadComments  bool              `json:"UnreadComments"`
	Comments        string            `json:"Comments"`
	SentReportDate  time.Time         `json:"SentReportDate"`
	LastCommentTime time.Time         `json:"LastCommentTime"`
	Created         time.Time         `json:"Created"`
	Updated         time.Time         `json:"Updated"`
}

type ReportStatus string

const (
	ReportStatusRejected            ReportStatus = "REJECTED"
	ReportStatusAccepted            ReportStatus = "ACCEPTED"
	ReportStatusPending             ReportStatus = "PENDING"
	ReportStatusNeedMoreInformation ReportStatus = "NEED_MORE_INFORMATION"
	ReportStatusSorted              ReportStatus = "SORTED"
	ReportStatusDuplicate           ReportStatus = "DUPLICATE"
	ReportStatusInformative         ReportStatus = "INFORMATIVE"
	ReportStatusConfirm             ReportStatus = "CONFIRM"
	ReportStatusPaidUp              ReportStatus = "PAID_UP"
)

type ReportSeriousness string

const (
	ReportSeriousnessNone     ReportSeriousness = "NONE"
	ReportSeriousnessLow      ReportSeriousness = "LOW"
	ReportSeriousnessMiddle   ReportSeriousness = "MIDDLE"
	ReportSeriousnessHigh     ReportSeriousness = "HIGH"
	ReportSeriousnessCritical ReportSeriousness = "CRITICAL"
)

type IReportRepository interface {
	GetReportByID(id int) (*Report, error)
	GetReportByField(field, value string) (*Report, error)
	GetReportByName(name string) (*Report, error)
	CreateReport(name, description, comments, seriousness string, userId string) (*Report, error)
	UpdateReport(name, description, comments, seriousness string, id int, didUpdateComments bool) error
	DeleteReport(id int) error
	SelectReportByStatus(status string) ([]*Report, error)
	UpdateReportStatus(user_id, point, report_id int, reportStauts string) (*Report, error)
}

type IReportService interface {
	ShowAllReportByStatus(ctx context.Context, status string) ([]*Report, error)
	GetReportByNameService(name string) (*Report, error)
	ShowTheReportByID(id int) (*Report, error)
	CreateReportService(ctx context.Context, name, description, comments, seriousness string) (*Report, error)
	UpdateReport(name, description, comments, seriousness *string, report *Report, ctx context.Context) error
	DeleteReport(ctx context.Context, id int) error
	VerifyReport(ctx context.Context, id int, report string) (*Report, error)
}

type ReportModule struct {
	IReportService
}

func NewReportModule(Db *sqlx.DB) *ReportModule {
	reportRepository := NewReportRepository(Db)
	userRepository := user.NewUserPostgres(Db)
	userService := user.NewUserService(userRepository)
	return &ReportModule{
		IReportService: NewReportService(reportRepository, userService),
	}
}
