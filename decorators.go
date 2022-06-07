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

type ImageDecorator struct {
}

func (decorator *ImageDecorator) RenderBeginning(data map[string]string) string {
	if alt, ok := data["alt"]; ok {
		return fmt.Sprintf("<img src=\"%s\" alt=\"%s\">", data["src"], alt)
	}
	return fmt.Sprintf("<img src=\"%s\">", data["src"])
}

func (decorator *ImageDecorator) RenderEnding(data map[string]string) string {
	return "</img>"
}
