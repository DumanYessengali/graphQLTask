package graph

import (
	"twoBinPJ/domains/auth"
	"twoBinPJ/domains/project"
	"twoBinPJ/domains/report"
	"twoBinPJ/domains/user"
	"twoBinPJ/domains/vulnerability"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	AuthModule          auth.IAuthService
	UserModule          user.IUserService
	VulnerabilityModule vulnerability.IVulnerabilityService
	ProjectModule       project.IProjectService
	ReportModule        report.IReportService
}
