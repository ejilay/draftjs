package draftjs

import "fmt"

type Decorator interface {
	RenderBeginning(data map[string]string) string
	RenderEnding(data map[string]string) string
}

type LinkDecorator struct {
}

func (decorator *LinkDecorator) RenderBeginning(data map[string]string) string {
	return fmt.Sprintf("<a href=\"%s\" target=\"_blank\">", data["url"])
}

func (decorator *LinkDecorator) RenderEnding(data map[string]string) string {
	return "</a>"
}
