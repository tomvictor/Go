package hello

import (
	"github.com/go-gorf/gorf"
)

func setup() error {
	// Add setup here
	MqRunner()
	return nil
}

var HelloApp = gorf.GorfBaseApp{
	Name:         "Hello",
	RouteHandler: Urls,
	SetUpHandler: setup,
}
