package auth

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"strings"
	"time"

	"github.com/readytotouch/readytotouch/internal/db/postgres"
	"github.com/readytotouch/readytotouch/internal/domain"
	"github.com/readytotouch/readytotouch/internal/env"
	"github.com/readytotouch/readytotouch/internal/protos/auth"

	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

type Controller struct {
	repository             *postgres.UserRepository
	githubOAuthProvider    domain.OAuthProvider
	gitlabOAuthProvider    domain.OAuthProvider
	bitbucketOAuthProvider domain.OAuthProvider
	jwtService             domain.JwtService
}

func NewController(repository *postgres.UserRepository, githubOAuthProvider domain.OAuthProvider, gitlabOAuthProvider domain.OAuthProvider, bitbucketOAuthProvider domain.OAuthProvider, jwtService domain.JwtService) *Controller {
	return &Controller{repository: repository, githubOAuthProvider: githubOAuthProvider, gitlabOAuthProvider: gitlabOAuthProvider, bitbucketOAuthProvider: bitbucketOAuthProvider, jwtService: jwtService}
}

func (c *Controller) GithubRedirect(ctx *gin.Context) {
	c.redirect(ctx, c.githubOAuthProvider)
}

func (c *Controller) GitlabRedirect(ctx *gin.Context) {
	c.redirect(ctx, c.gitlabOAuthProvider)
}

func (c *Controller) BitbucketRedirect(ctx *gin.Context) {
	c.redirect(ctx, c.bitbucketOAuthProvider)
}

func (c *Controller) GithubCallback(ctx *gin.Context) {
	c.callback(ctx, c.githubOAuthProvider)
}

func (c *Controller) GitlabCallback(ctx *gin.Context) {
	c.callback(ctx, c.gitlabOAuthProvider)
}

func (c *Controller) BitbucketCallback(ctx *gin.Context) {
	c.callback(ctx, c.bitbucketOAuthProvider)
}

func (c *Controller) Logout(ctx *gin.Context) {
	c.jwtService.RemoveToken(ctx)

	ctx.Redirect(http.StatusFound, "/")
}

func (c *Controller) redirect(ctx *gin.Context, provider domain.OAuthProvider) {
	if provider.GetConfig().ClientSecret == "" {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte("Current OAuth2 provider disabled"))

		return
	}

	state, err := c.encodeRedirectState(ctx.Query("redirect"))
	if err != nil {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte("Cannot generate random state for OAuth2 redirect"))

		return
	}

	ctx.SetCookie(c.getStateCookieName(), state, 3600, "/", "", true, true)
	ctx.Redirect(http.StatusFound, provider.GetConfig().AuthCodeURL(state))
}

func (c *Controller) callback(ctx *gin.Context, provider domain.OAuthProvider) {
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

	token, err := provider.GetConfig().Exchange(ctx, queryCode)
	if err != nil {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte("Cannot retrieve access token from OAuth2 provider"))

		return
	}

	user, err := provider.GetUser(token.AccessToken)
	if err != nil {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte("Cannot retrieve user data from OAuth2 provider"))

		return
	}

	userID, err := c.repository.Create(ctx, &domain.SocialProviderUserCreateParams{
		SocialProvider:       provider.SocialProvider(),
		SocialProviderUserID: user.ID,
		Email:                user.Email,
		Username:             user.Username,
		Name:                 user.Name,
	}, time.Now().UTC())
	if err != nil {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte("Cannot create user"))

		return
	}

	err = c.jwtService.AddToken(ctx, domain.JwtUser{
		ID: userID,
	})
	if err != nil {
		ctx.Data(http.StatusInternalServerError, "text/html; charset=utf-8", []byte("Cannot generate JWT token"))

		return
	}

	redirect, err := c.decodeRedirectState(queryState)
	if err != nil {
		// @TODO logging

		ctx.Redirect(http.StatusFound, "/")
	}

	if redirect == "" {
		redirect = "/"
	}

	ctx.Redirect(http.StatusFound, redirect)
}

func (c *Controller) getStateCookieName() string {
	if env.Environment() == env.Development {
		return "dev_oauth_state"
	}

	return "oauth_state"
}

func (c *Controller) encodeRedirectState(redirect string) (string, error) {
	random := make([]byte, 32)
	_, err := rand.Read(random)
	if err != nil {
		return "", nil
	}

	state := &auth.State{
		Payload: &auth.Payload{
			Redirect: "",
			Random:   random,
		},
	}

	if strings.HasPrefix(redirect, "/") {
		state.Payload.Redirect = redirect
	}

	content, err := proto.Marshal(state)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(content), nil
}

func (c *Controller) decodeRedirectState(source string) (string, error) {
	content, err := base64.StdEncoding.DecodeString(source)
	if err != nil {
		return "", err
	}

	state := &auth.State{}
	err = proto.Unmarshal(content, state)
	if err != nil {
		return "", err
	}

	redirect := state.GetPayload().GetRedirect()
	if strings.HasPrefix(redirect, "/") {
		return redirect, nil
	}

	return "", nil
}
