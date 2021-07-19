package report

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"strconv"
	"time"
)

type ReportRepository struct {
	DB *sqlx.DB
}

func NewReportRepository(Db *sqlx.DB) *ReportRepository {
	return &ReportRepository{DB: Db}
}

func (p *ReportRepository) GetReportByField(field, value string) (*Report, error) {
	var report *Report
	row, err := p.DB.Query(fmt.Sprintf("select*from report where %v=$1", field), value)
	if err != nil {
		return nil, err
	}

	for row.Next() {
		var r Report
		err = row.Scan(&r.ID, &r.Name, &r.Description, &r.Status, &r.Seriousness, &r.Archive, &r.Delete, &r.Reward,
			&r.Point, &r.ProjectID, &r.VulnerabilityID, &r.UserID, &r.Assignee, &r.UnreadComments, &r.Comments,
			&r.SentReportDate, &r.LastCommentTime, &r.Created, &r.Updated)

		if err != nil {
			return nil, err
		}

		report = &r
	}
	return report, nil
}

func (r *ReportRepository) GetReportByID(id int) (*Report, error) {
	return r.GetReportByField("id", strconv.Itoa(id))
}

func (r *ReportRepository) GetReportByName(name string) (*Report, error) {
	return r.GetReportByField("name", name)
}

func (r *ReportRepository) CreateReport(name, description, comments, seriousness string, unreadComments bool) (*Report, error) {
	rows, err := r.DB.Query("insert into report(name, description, status, seriousness, archive, delete, reward, "+
		"point, project_id,vulnerability_id, user_id, assignee, unread_comments, comments, sent_report_date, "+
		"last_comment_time, created, updated) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18) "+
		"returning *", name, description, ReportStatusPending, seriousness, false, false, 1, 10, 1, 1, 1, 1, unreadComments, comments, time.Now(), time.Now(), time.Now(), time.Now())
	if err != nil {
		return nil, err
	}

}
