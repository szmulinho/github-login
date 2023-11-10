package endpoints

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/szmulinho/github-login/internal/model"
	"net/http"
	"time"
)

func (h *handlers) GenerateToken(w http.ResponseWriter, r *http.Request, githubUser model.GitHubLogin, isGithubUser bool) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"githubUserLogin": githubUser.User,
		"isGithubUser":    isGithubUser,
		"exp":             time.Now().Add(time.Hour * time.Duration(1)).Unix(),
	})

	tokenString, err := token.SignedString(model.JwtKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return "", err
	} else {
		fmt.Sprintf("token for user %s generated", model.GitHubLogin{})
	}

	return tokenString, nil
}