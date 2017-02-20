// draftjs exporter for go language
package draftjs

import (
	"bytes"
)

func (contentState *ContentState) render(config *Config) string {
	var (
		buf       bytes.Buffer
		prevBlock *ContentBlock
	)
	buf.Grow(256 * 1024) // с потолка
	wrapperTag := ""

	for _, block := range contentState.Blocks {
		if wrapperTag != GetBlockWrapperTag(block, config) {
			wrapperTag = GetBlockWrapperTag(block, config)
			buf.WriteString(GetBlockWrapperEndTag(prevBlock, config))
			buf.WriteString(GetBlockWrapperStartTag(block, config))
		}
		renderBlock(&buf, block, config, contentState)
		prevBlock = block
	}
	if wrapperTag != "" {
		buf.WriteString(GetBlockWrapperEndTag(prevBlock, config))
	}
	return buf.String()
}

func renderBlock(buf *bytes.Buffer, block *ContentBlock, config *Config, contentState *ContentState) {
	buf.WriteString(GetBlockStartTag(block, config))
	buf.WriteString(PerformInlineStylesAndEntities(contentState, block, config))
	buf.WriteString(GetBlockEndTag(block, config))
}

// Render renders Draft.js content state to string with config
func Render(contentState *ContentState, config *Config) string {
	if config == nil {
		config = NewDefaultConfig()
	}
	return contentState.render(config)
}

// Interface implementation
func (contentState *ContentState) String() string {
	return Render(contentState, nil)
}
