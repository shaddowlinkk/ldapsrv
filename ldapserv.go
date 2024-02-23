package awesomeProject

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/labstack/gommon/log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	e.GET("/ldap", func(c echo.Context) error {
		l, err := ldap.DialURL("") //need to add ldap url
		if err != nil {
			log.Fatal(err)
		}
		err = l.Bind("cn=Directory Manager", "") //need to add admin password
		if err != nil {
			log.Fatal(err)
		}
		user := "test"
		baseDn := "" // neeed to add base dn
		filter := fmt.Sprintf("(CN=%s)", ldap.EscapeFilter(user))
		searchReq := ldap.NewSearchRequest(baseDn, ldap.ScopeWholeSubtree, 0, 0, 0, false, filter, []string{"uidNumber"}, nil)
		result, err := l.Search(searchReq)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, struct {
				Status string
			}{Status: "error in serach"})
		}
		return c.JSON(http.StatusOK, result.Entries)
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

// Simple implementation of an integer minimum
// Adapted from: https://gobyexample.com/testing-and-benchmarking
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}
