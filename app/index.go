package app

import (
	"strings"
	"upload/context"
	"upload/context/global"
)

func Index(c *context.Context) {
	if global.Conf.Release == 1 {
		c.Text("It works!")
	} else {
		c.Html("index", strings.Join(global.Conf.AllowFileType, ", "))
	}
}
