package actions

import (
	"github.com/gobuffalo/buffalo"
)

// ProfileShow show the user profile page
func ProfileShow(c buffalo.Context) error {
	return c.Render(200, r.HTML("profile/show.html"))
}
