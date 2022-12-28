package middleware

import (
	"fmt"
	"net/http"

	"github.com/jchen42703/go-api/internal/auth"
	"github.com/jchen42703/go-api/internal/httputil/respond"
	echo "github.com/labstack/echo/v4"
	ory "github.com/ory/client-go"
)

// Feel free to change this to something else if you want.
// This is the key you need to use to access the Ory session in the echo.Context
const ORY_SESSION_CONTEXT_KEY = "user_session"

// Checks if a user has an active session (is logged in) by checking the session cookie.
// Then, it adds the session object to the "user_session" key in the echo Context.
func CreateAuthMiddleware(oryClient *ory.APIClient) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			whitelist := []string{
				"/api",
				"/api/notprotected", // dev route for not protected route
				"/api/auth/signup",
				"/api/auth/logout",
			}

			// Skip auth if whitelisted
			for _, whitelistedPath := range whitelist {
				if whitelistedPath == c.Path() {
					return next(c)
				}
			}

			// check if we have a session
			session, resp, err := auth.ValidateSession(oryClient, c)
			if err != nil {
				fmt.Println("validate sess err: ", respond.NewError(err))
				fmt.Println("resp: ", resp)
				// return c.JSON(http.StatusUnauthorized, templates.NewError(err))
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid session")
			}

			isLogin := c.Path() == "/api/auth/login"

			// TODO: refresh with new jwt if valid JWT
			// Give 200 OK and frontend should redirect to the correct page.
			if isLogin {
				return c.JSON(http.StatusOK, respond.NewMessage("already logged in"))
			}

			// so that it can be reused in next handlers
			c.Set(ORY_SESSION_CONTEXT_KEY, session)
			return next(c)
		}
	}
}
