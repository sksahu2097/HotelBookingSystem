package handlers

import (
	"net/http"

	"github.com/sksahu2097/HotelBookingSystem/pkg/config"
	"github.com/sksahu2097/HotelBookingSystem/pkg/models"
	"github.com/sksahu2097/HotelBookingSystem/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// set the AppConfig in the Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func SetRepo(a *Repository) {
	Repo = a
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)
	stringMap := make(map[string]string)
	stringMap["Test"] = "my name is pojo"
	render.RenderTenplate(w, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["Test"] = "my name is pojo"
	remoteIp := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIp
	render.RenderTenplate(w, "aboutUs.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
