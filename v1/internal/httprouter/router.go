package httprouter

import (
	"fmt"
	"net/http"

	"github.com/jchen42703/go-api/internal/httprouter/middleware"
	"github.com/jchen42703/go-api/internal/httputil/respond"
	"github.com/jchen42703/go-api/internal/httputil/validate"
	"github.com/labstack/echo/v4"
	echomiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"

	ory "github.com/ory/client-go"
)

func GetSession() echo.HandlerFunc {
	return func(c echo.Context) error {
		session, ok := c.Get(middleware.ORY_SESSION_CONTEXT_KEY).(*ory.Session)
		if !ok || session == nil {
			return c.JSON(http.StatusUnauthorized, respond.NewError(fmt.Errorf("must be signed in")))
		}

		return c.JSON(http.StatusOK, session)
	}
}

// Random protected endpoint
func GetProtected() echo.HandlerFunc {
	return func(c echo.Context) error {
		session, ok := c.Get(middleware.ORY_SESSION_CONTEXT_KEY).(*ory.Session)
		if !ok || session == nil {
			return c.JSON(http.StatusUnauthorized, respond.NewError(fmt.Errorf("must be signed in")))
		}

		return c.JSON(http.StatusOK, "protected")
	}
}

func GetNotProtected() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "not protected")
	}
}

// Test routes to see session data.
func RegisterBaseRoutes(g *echo.Group) {
	g.GET("/session", GetSession())
	g.GET("/protected", GetProtected())
	g.GET("/notprotected", GetNotProtected())
}

func New(oryClient *ory.APIClient, allowedOrigins []string) *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(echomiddleware.RemoveTrailingSlash())
	e.Use(echomiddleware.Logger())
	e.Use(echomiddleware.CORSWithConfig(echomiddleware.CORSConfig{
		AllowOrigins:     allowedOrigins,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true, // This is necessary if you need session cookies because fetch with credentials:true requires CORS to return the credentials header as true
	}))
	e.Use(echomiddleware.CSRF())
	e.Use(middleware.CreateAuthMiddleware(oryClient))
	e.Validator = validate.NewValidator()
	return e
}
