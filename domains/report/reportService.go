package report

import (
	"context"
	"errors"
	"log"
	"strconv"
	"twoBinPJ/domains/user"
)

type ReportService struct {
	Repository  IReportRepository
	UserService user.IUserService
}

func NewReportService(repo IReportRepository, userService *user.UserService) *ReportService {
	return &ReportService{
		Repository:  repo,
		UserService: userService,
	}
}

func (r *ReportService) GetReportByNameService(name string) (*Report, error) {
	return r.Repository.GetReportByName(name)
}

func (r *ReportService) ShowTheReportByID(id int) (*Report, error) {
	return r.Repository.GetReportByID(id)
}

func (r *ReportService) CreateReportService(ctx context.Context, name, description, comments, seriousness string) (*Report, error) {
	users, err := user.ForContext(ctx)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return nil, errors.New("INITIALIZING_TOKEN_ERROR")
	}
	currentUser, err := r.UserService.GetUserByIDService(users.Id)
	if err != nil {
		log.Printf("error while initializing user error")
		return nil, err
	}
	if currentUser.Role != 1 {
		log.Printf("You do not have access to create report: %s", err)
		return nil, errors.New("CREATE_REPORT_ERROR")
	}
	return r.Repository.CreateReport(name, description, comments, seriousness, currentUser.Id)
}

func (r *ReportService) UpdateReport(name, description, comments, seriousness *string, report *Report, ctx context.Context) error {
	users, err := user.ForContext(ctx)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return errors.New("INITIALIZING_TOKEN_ERROR")
	}
	currentUser, err := r.UserService.GetUserByIDService(users.Id)
	if err != nil {
		log.Printf("error while initializing user error")
		return err
	}
	if currentUser.Role != 1 {
		log.Printf("You do not have access to update report: %s", err)
		return errors.New("UPDATE_REPORT_ERROR")
	}
	didUpdate := false
	didUpdateComments := false
	if name == nil {
		report.Name = report.Name
	} else {
		report.Name = *name
		didUpdate = true
	}

	if description == nil {
		report.Description = report.Description
	} else {
		report.Description = *description
		didUpdate = true
	}

	if comments == nil {
		report.Comments = report.Comments
	} else {
		report.Comments = *comments
		didUpdate = true
		didUpdateComments = true
	}

	if seriousness == nil {
		report.Seriousness = report.Seriousness
	} else {
		report.Seriousness = ReportSeriousness(*seriousness)
		didUpdate = true
	}

	if !didUpdate {
		return errors.New("no update done")
	}

	err = r.Repository.UpdateReport(report.Name, report.Description, report.Comments, string(report.Seriousness), report.ID, didUpdateComments)

	if err != nil {
		return err
	}
	return nil
}

func (r *ReportService) DeleteReport(ctx context.Context, id int) error {
	users, err := user.ForContext(ctx)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return errors.New("INITIALIZING_TOKEN_ERROR")
	}
	currentUser, err := r.UserService.GetUserByIDService(users.Id)
	if err != nil {
		log.Printf("error while initializing user error")
		return err
	}
	if currentUser.Role != 1 {
		log.Printf("You do not have access to delete report: %s", err)
		return errors.New("DELETE_REPORT_ERROR")
	}
	return r.Repository.DeleteReport(id)
}

func (r *ReportService) ShowAllReportByStatus(ctx context.Context, status string) ([]*Report, error) {
	users, err := user.ForContext(ctx)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return nil, errors.New("INITIALIZING_TOKEN_ERROR")
	}
	currentUser, err := r.UserService.GetUserByIDService(users.Id)
	if err != nil {
		log.Printf("error while initializing user error")
		return nil, err
	}
	if currentUser.Role != 2 {
		log.Printf("user do not have access to show report by status: %s", err)
		return nil, errors.New("ACCESS_INITIALIZING_ERROR")
	}
	return r.Repository.SelectReportByStatus(status)
}

func (r *ReportService) VerifyReport(ctx context.Context, id int, reportStatus string) (*Report, error) {
	users, err := user.ForContext(ctx)
	if err != nil {
		log.Printf("token is incorrect or wrong: %s", err)
		return nil, errors.New("INITIALIZING_TOKEN_ERROR")
	}
	currentUser, err := r.UserService.GetUserByIDService(users.Id)
	if err != nil {
		log.Printf("error while initializing user error")
		return nil, err
	}
	if currentUser.Role != 2 {
		log.Printf("user do not have access to show report by status: %s", err)
		return nil, errors.New("ACCESS_INITIALIZING_ERROR")
	}
	reportWhichWantToChange, err := r.Repository.GetReportByID(id)
	if err != nil {
		log.Printf("error while initializing report: %s", err)
		return nil, errors.New("INITIALIZING_REPORT_ERROR")
	}
	userID := strconv.Itoa(reportWhichWantToChange.UserID)
	currentUser1, err := r.UserService.GetUserByIDService(userID)
	if err != nil {
		log.Printf("error while initializing user error")
		return nil, err
	}
	if reportWhichWantToChange.Status != ReportStatusConfirm {
		verifiedReport, err := r.Repository.UpdateReportStatus(reportWhichWantToChange.UserID, reportWhichWantToChange.Point, id, reportStatus, currentUser1.Point)
		if err != nil {
			log.Printf("error while updating report: %s", err)
			return nil, errors.New("UPDATING_REPORT_ERROR")
		}
		return verifiedReport, nil
	}

	return nil, errors.New("UPDATING_REPORT_ERROR")
}
