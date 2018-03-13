package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/piffio/myvolumio/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
