package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"log"
	"twoBinPJ/apps/api2/graph/generated"
	"twoBinPJ/apps/api2/graph/model"
	"twoBinPJ/domains/project"
	"twoBinPJ/domains/report"
	"twoBinPJ/domains/user"
	"twoBinPJ/domains/vulnerability"
)

func (r *mutationResolver) SignIn(ctx context.Context, input model.SignInUser) (*model.AuthResponse, error) {
	token, user, err := r.AuthModule.SignIn(ctx, input.Username, input.Password)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		AuthTokens: token,
		User:       user,
	}, nil
}

func (r *mutationResolver) SignUp(ctx context.Context, input model.SignUpUser) (*model.Message, error) {
	message, err := r.AuthModule.SignUp(ctx, input.Username, input.Password)
	if err != nil {
		return nil, err
	}
	return &model.Message{Message: message}, nil
}

func (r *mutationResolver) RefreshTokens(ctx context.Context, input model.Refresh) (*model.AuthResponse, error) {
	token, user, err := r.AuthModule.RefreshTokens(ctx, input.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &model.AuthResponse{
		AuthTokens: token,
		User:       user,
	}, nil
}

func (r *mutationResolver) Logout(ctx context.Context, input model.Refresh) (*model.Message, error) {
	message, err := r.AuthModule.Logout(ctx, input.RefreshToken)
	if err != nil {
		return nil, err
	}
	return &model.Message{Message: message}, nil
}

func (r *mutationResolver) ShowTheProjectByID(ctx context.Context, id int) (*project.Project, error) {
	project, err := r.ProjectModule.ShowTheProjectByID(id)
	if err != nil {
		log.Printf("error while select project id: %s", err)
		return nil, errors.New("INITIALIZING_ID_ERROR")
	}
	return project, nil
}

func (r *mutationResolver) CreateProject(ctx context.Context, input model.CreateProject) (*project.Project, error) {
	project, err := r.ProjectModule.CreateProjectService(ctx, input.Name, input.ShortDescription, input.Description)
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (r *mutationResolver) UpdateProject(ctx context.Context, id int, input model.UpdateProject) (*model.Message, error) {
	project, err := r.ProjectModule.ShowTheProjectByID(id)
	if err != nil {
		return nil, err
	}
	err = r.ProjectModule.UpdateProject(ctx, project, input.Name, input.ShortDescription, input.Description)
	if err != nil {
		return nil, err
	}
	return &model.Message{Message: "Project successfully updated"}, nil
}

func (r *mutationResolver) DeleteProject(ctx context.Context, id int) (*model.Message, error) {
	err := r.ProjectModule.DeleteProject(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.Message{Message: "Project successfully deleted"}, nil
}

func (r *mutationResolver) ShowTheVulnerabilityByID(ctx context.Context, id int) (*vulnerability.Vulnerability, error) {
	vulnerability, err := r.VulnerabilityModule.ShowVulnerabilityByID(ctx, id)
	if err != nil {
		log.Printf("error while select vulnerability id: %s", err)
		return nil, errors.New("INITIALIZING_ID_ERROR")
	}
	return vulnerability, nil
}

func (r *mutationResolver) CreateVulnerability(ctx context.Context, input model.CreateVulnerability) (*vulnerability.Vulnerability, error) {
	vulnerability, err := r.VulnerabilityModule.CreateVulnerability(ctx, input.Name, input.Description)
	if err != nil {
		return nil, err
	}
	return vulnerability, nil
}

func (r *mutationResolver) UpdateVulnerability(ctx context.Context, id int, input model.UpdateVulnerability) (*model.Message, error) {
	vulnerability, err := r.VulnerabilityModule.ShowVulnerabilityByID(ctx, id)
	if err != nil {
		return nil, err
	}
	err = r.VulnerabilityModule.UpdateVulnerability(ctx, input.Name, input.Description, vulnerability)
	if err != nil {
		return nil, err
	}
	return &model.Message{Message: "Vulnerability successfully updated"}, nil
}

func (r *mutationResolver) DeleteVulnerability(ctx context.Context, id int) (*model.Message, error) {
	err := r.VulnerabilityModule.DeleteVulnerability(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.Message{Message: "Vulnerability successfully deleted"}, nil
}

func (r *mutationResolver) ShowTheReportByID(ctx context.Context, id int) (*report.Report, error) {
	report, err := r.ReportModule.ShowTheReportByID(id)
	if err != nil {
		log.Printf("error while select report id: %s", err)
		return nil, errors.New("INITIALIZING_ID_ERROR")
	}
	return report, nil
}

func (r *mutationResolver) CreateReport(ctx context.Context, input model.CreateReport) (*report.Report, error) {
	report, err := r.ReportModule.CreateReportService(ctx, input.Name, input.Description, input.Comments, input.Seriousness)
	if err != nil {
		return nil, err
	}
	return report, nil
}

func (r *mutationResolver) UpdateReport(ctx context.Context, id int, input model.UpdateReport) (*model.Message, error) {
	report, err := r.ReportModule.ShowTheReportByID(id)
	if err != nil {
		return nil, err
	}
	err = r.ReportModule.UpdateReport(input.Name, input.Description, input.Comments, input.Seriousness, report, ctx)
	if err != nil {
		return nil, err
	}
	return &model.Message{Message: "Report successfully updated"}, nil
}

func (r *mutationResolver) DeleteReport(ctx context.Context, id int) (*model.Message, error) {
	err := r.ReportModule.DeleteReport(ctx, id)
	if err != nil {
		return nil, err
	}
	return &model.Message{Message: "Report successfully deleted"}, nil
}

func (r *mutationResolver) ShowAllReportByStatus(ctx context.Context, input model.ReportStatus) ([]*report.Report, error) {
	reportsByStatus, err := r.ReportModule.ShowAllReportByStatus(ctx, input.Status)
	if err != nil {
		return nil, err
	}
	return reportsByStatus, nil
}

func (r *mutationResolver) VerifyReport(ctx context.Context, id int, input model.ReportStatus) (*report.Report, error) {
	report, err := r.ReportModule.VerifyReport(ctx, id, string(input.Status))

	if err != nil {
		return nil, err
	}

	return report, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*user.User, error) {
	getUser, err := r.UserModule.GetUserByIDService(id)
	if err != nil {
		return nil, err
	}
	return getUser, nil
}

func (r *reportResolver) Status(ctx context.Context, obj *report.Report) (string, error) {
	return string(obj.Status), nil
}

func (r *reportResolver) Seriousness(ctx context.Context, obj *report.Report) (string, error) {
	return string(obj.Seriousness), nil
}

func (r *userResolver) Role(ctx context.Context, obj *user.User) (int, error) {
	return int(obj.Role), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Report returns generated.ReportResolver implementation.
func (r *Resolver) Report() generated.ReportResolver { return &reportResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type reportResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
