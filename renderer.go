// draftjs exporter for go language
package draftjs

import (
	"bytes"
)

func (contentState *ContentState) render(config *HTMLConfig) string {
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
		buf.WriteString(GetBlockStartTag(block, config))
		buf.WriteString(PerformInlineStylesAndEntities(contentState, block, config))
		buf.WriteString(GetBlockEndTag(block, config))
		prevBlock = block
	}
	if wrapperTag != "" {
		buf.WriteString(GetBlockWrapperEndTag(prevBlock, config))
	}
	return buf.String()
}

func Render(contentState *ContentState, config *HTMLConfig) string {
	if config == nil {
		config = DefaultConfig()
	}
	return contentState.render(config)
}

func (contentState *ContentState) String() string {
	return Render(contentState, nil)
}
