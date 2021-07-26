package report

import (
	"database/sql"
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

func (r *ReportRepository) GetReportByField(field, value string) (*Report, error) {
	var report *Report
	row, err := r.DB.Query(fmt.Sprintf("select*from report where %v=$1", field), value)
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

func (r *ReportRepository) CreateReport(name, description, comments, seriousness string, userId string) (*Report, error) {
	fmt.Println(ReportStatusPending)
	rows, err := r.DB.Query("insert into report(name, description, status, seriousness, archive, delete, reward, "+
		"point, project_id,vulnerability_id, user_id, assignee, unread_comments, comments, sent_report_date, "+
		"last_comment_time, created, updated) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18) "+
		"returning *", name, description, ReportStatusPending, seriousness, false, false, 1, 10, 1, 1, userId, 1, true, comments, time.Now(), time.Now(), time.Now(), time.Now())
	if err != nil {
		return nil, err
	}
	var report *Report
	for rows.Next() {
		var r Report
		err = rows.Scan(&r.ID, &r.Name, &r.Description, &r.Status, &r.Seriousness, &r.Archive, &r.Delete, &r.Reward, &r.Point,
			&r.ProjectID, &r.VulnerabilityID, &r.UserID, &r.Assignee, &r.UnreadComments, &r.Comments, &r.SentReportDate,
			&r.LastCommentTime, &r.Created, &r.Updated)
		if err != nil {
			return nil, err
		}
		report = &r
	}
	return report, nil
}

func (r *ReportRepository) UpdateReport(name, description, comments, seriousness string, id int, didUpdateComments bool) error {
	var err error
	if didUpdateComments {
		_, err = r.DB.Query("update report set name=$1, description=$2,comments=$3,seriousness=$4,updated=$5 where id=$6",
			name, description, comments, seriousness, time.Now(), id)
	} else {
		_, err = r.DB.Query("update report set name=$1, description=$2,comments=$3,seriousness=$4,last_comment_time=$5,updated=$6 where id=$7",
			name, description, comments, seriousness, time.Now(), time.Now(), id)
	}
	return err
}

func (r *ReportRepository) DeleteReport(id int) error {
	query := "Delete from report where id=$1"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReportRepository) SelectReportByStatus(status string) ([]*Report, error) {
	var report []*Report
	row, err := r.DB.Query("select*from report where status=$1", status)
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

		report = append(report, &r)
	}
	return report, nil
}

func (r *ReportRepository) UpdateReportStatus(user_id, point, report_id int, reportStatus string) (*Report, error) {
	var row *sql.Rows
	var err error

	if reportStatus == string(ReportStatusConfirm) {
		_, err = r.DB.Query("update users set point = point+$1 where id=$2", point, user_id)
		if err != nil {
			return nil, err
		}
		_, err = r.DB.Query("update report set status = $1 where id=$2", reportStatus, report_id)
		if err != nil {
			return nil, err
		}
		row, err = r.DB.Query("select*from report where id=$1", report_id)
		if err != nil {
			return nil, err
		}

	} else {
		_, err = r.DB.Query("update report set status = $1 where id=$2", reportStatus, report_id)
		if err != nil {
			return nil, err
		}
		row, err = r.DB.Query("select*from report where id=$1", report_id)
		if err != nil {
			return nil, err
		}
	}
	var report *Report

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
