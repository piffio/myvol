package actions

import (
	"github.com/gobuffalo/buffalo"
)

// ProfileShowProfile default implementation.
func ProfileShow(c buffalo.Context) error {
	return c.Render(200, r.HTML("profile/show.html"))
}
