package auth

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	ory "github.com/ory/client-go"
	"github.com/samber/lo"
)

func NewOryAPIClient(baseKratosUrl string) *ory.APIClient {
	if baseKratosUrl == "" {
		baseKratosUrl = "http://localhost:4433"
	}

	// register a new Ory client with the URL set to the Ory CLI Proxy
	// we can also read the URL from the env or a config file
	c := ory.NewConfiguration()
	c.Servers = ory.ServerConfigurations{{URL: baseKratosUrl}}
	return ory.NewAPIClient(c)
}

func ValidateSession(client *ory.APIClient, ctx echo.Context) (*ory.Session, *http.Response, error) {
	_, err := ctx.Cookie("ory_kratos_session")
	if err != nil {
		return nil, nil, fmt.Errorf("ValidateSession: could not find session cookie 'ory_kratos_session'")
	}

	token := ""
	if h := ctx.Request().Header.Get("Authorization"); strings.HasPrefix(strings.ToLower(h), "bearer ") {
		token = h[7:]
	}

	cookies := strings.Join(lo.Map(ctx.Cookies(), func(cookie *http.Cookie, index int) string {
		return cookie.String()
	}), "; ")

	session, resp, err := client.FrontendApi.ToSession(ctx.Request().Context()).
		XSessionToken(token).
		Cookie(cookies).
		Execute()

	if (err != nil && session == nil) || (err == nil && !*session.Active) {
		// Not logged in
		return nil, nil, fmt.Errorf("ValidateSession: no valid session: %s", err)
	}

	return session, resp, nil
}
