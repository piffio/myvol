package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/piffio/myvol/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
