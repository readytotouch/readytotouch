package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"golang.org/x/oauth2"

	"github.com/readytotouch-yaaws/yaaws-go/internal/db/postgres"
	"github.com/readytotouch-yaaws/yaaws-go/internal/env"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	repository      *postgres.UserRepository
	githubConfig    oauth2.Config
	gitlabConfig    oauth2.Config
	bitbucketConfig oauth2.Config
}

func NewController(repository *postgres.UserRepository, githubConfig oauth2.Config, gitlabConfig oauth2.Config, bitbucketConfig oauth2.Config) *Controller {
	return &Controller{
		repository:      repository,
		githubConfig:    githubConfig,
		gitlabConfig:    gitlabConfig,
		bitbucketConfig: bitbucketConfig,
	}
}

func (c *Controller) GithubRedirect(ctx *gin.Context) {
	c.redirect(ctx, c.githubConfig)
}

func (c *Controller) GitlabRedirect(ctx *gin.Context) {
	c.redirect(ctx, c.gitlabConfig)
}

func (c *Controller) BitbucketRedirect(ctx *gin.Context) {
	c.redirect(ctx, c.bitbucketConfig)
}

func (c *Controller) GithubCallback(ctx *gin.Context) {
	c.callback(ctx, c.githubConfig)
}

func (c *Controller) GitlabCallback(ctx *gin.Context) {
	c.callback(ctx, c.gitlabConfig)
}

func (c *Controller) BitbucketCallback(ctx *gin.Context) {
	c.callback(ctx, c.bitbucketConfig)
}

func (c *Controller) redirect(ctx *gin.Context, config oauth2.Config) {
	if config.ClientSecret == "" {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte("Current OAuth2 provider disabled"))

		return
	}

	state, err := c.generateRandomState()
	if err != nil {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte("Cannot generate random state for OAuth2 redirect"))

		return
	}

	ctx.SetCookie(c.getStateCookieName(), state, 3600, "/", "", true, true)
	ctx.Redirect(http.StatusFound, config.AuthCodeURL(state))
}

func (c *Controller) callback(ctx *gin.Context, config oauth2.Config) {
	var (
		queryCode  = ctx.Query("code")
		queryState = ctx.Query("state")
	)

	if queryCode == "" {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte("Query param code required"))

		return
	}
	if queryState == "" {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte("Query param state required"))

		return
	}

	cookieState, err := ctx.Cookie(c.getStateCookieName())
	if err != nil {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte("Cannot get state from cookie"))

		return
	}

	if cookieState != queryState {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte("State from query param and state from cookie must be the same"))

		return
	}

	token, err := config.Exchange(ctx, queryCode)
	if err != nil {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte("Cannot retrieve access token from OAuth2 provider"))

		return
	}

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(token.AccessToken))
}

func (c *Controller) getStateCookieName() string {
	if env.Environment() == env.Development {
		return "dev_oauth_state"
	}

	return "oauth_state"
}

func (c *Controller) generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", nil
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
