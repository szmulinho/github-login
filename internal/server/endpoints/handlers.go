package endpoints

import (
	"github.com/szmulinho/github-login/internal/model"
	"gorm.io/gorm"
	"net/http"
)

type Handlers interface {
	HandleLogin(w http.ResponseWriter, r *http.Request)
	HandleCallback(w http.ResponseWriter, r *http.Request)
	GetUserDataHandler(w http.ResponseWriter, r *http.Request)
	getUserFromToken(tokenString string) (*model.GithubUser, error)
	checkRepoAdminAccess(accessToken string, user model.GithubUser) bool
	getData(accessToken, apiUrl string) (string, error)
	Logged(w http.ResponseWriter, r *http.Request, githubData string)
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
	updateOrCreateGitHubUser(db *gorm.DB, githubUser model.GithubUser) error
	updateOrCreatePublicRepo(db *gorm.DB, publicRepo model.PublicRepo) error
}

type handlers struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) Handlers {
	return &handlers{
		db: db,
	}
}
