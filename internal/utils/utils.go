package utils

import (
	"golang_htmx_teml/internal/templates/components"
	"golang_htmx_teml/internal/types"

	"github.com/a-h/templ"
)

func SetupPage(body templ.Component) templ.Component {
	var navItems types.NavElements
	navItems.Init()

	navbar := components.Navbar(navItems)
	page := components.PageHeader(navbar, body)
	return page
}
