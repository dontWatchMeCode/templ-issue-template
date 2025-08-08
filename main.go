package main

import (
	"bytes"
	"context"

	"github.com/dontWatchMeCode/templ-issue-template/templates"
)

func main() {
	var buf bytes.Buffer
	err := templates.Notifications{}.Index().Render(context.Background(), &buf)
	if err != nil {
		panic(err)
	}

	println(buf.String())
}
