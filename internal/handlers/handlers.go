package handlers

import (
	"github.com/ChrisDevOpsOrg/bookings/internal/config"
	"github.com/ChrisDevOpsOrg/bookings/internal/models"
	"github.com/ChrisDevOpsOrg/bookings/internal/render"
	"net/http"
)

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// Repo is repository used by the handlers
var Repo *Repository

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(repo *Repository) {
	Repo = repo
}

// Home is the home page handler
// (m *Repository) is the function receiver, this receiver have access to everything inside repository
// which happens to be application config
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

// About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := map[string]string{
		"test": "Hello Again",
	}
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
