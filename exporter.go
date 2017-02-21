// draftjs exporter for go language
package draftjs

import (
	"bytes"
)

func renderBlocks(contentState *ContentState, config *Config, buf *bytes.Buffer, blockIterator *BlockIterator) {

	var wrapperedBlock *ContentBlock
	wrapperTag := ""
	for blockIterator.block != nil {
		if wrapperTag != GetBlockWrapperTag(blockIterator.block, config) {
			wrapperTag = GetBlockWrapperTag(blockIterator.block, config)
			buf.WriteString(GetBlockWrapperEndTag(wrapperedBlock, config))
			buf.WriteString(GetBlockWrapperStartTag(blockIterator.block, config))
		}

		wrapperedBlock = blockIterator.block
		currentBlock := blockIterator.block

		buf.WriteString(GetBlockStartTag(currentBlock, config))
		buf.WriteString(PerformInlineStylesAndEntities(contentState, currentBlock, config))
		if blockIterator.HasNext() && blockIterator.NextBlock().Depth > blockIterator.block.Depth {
			renderBlocks(contentState, config, buf, blockIterator.StepNext())
		}
		buf.WriteString(GetBlockEndTag(currentBlock, config))

		if blockIterator.HasNext() && blockIterator.NextBlock().Depth < currentBlock.Depth {
			break
		}
		blockIterator.StepNext()
	}
	if wrapperTag != "" && wrapperedBlock != nil {
		buf.WriteString(GetBlockWrapperEndTag(wrapperedBlock, config))
	}
}

// Render renders Draft.js content state to string with config
func Render(contentState *ContentState, config *Config) string {
	var buf bytes.Buffer

	buf.Grow(256 * 1024) // с потолка

	if config == nil {
		config = NewDefaultConfig()
	}

	renderBlocks(contentState, config, &buf, NewBlockIterator(contentState))

	return buf.String()
}

// Interface implementation
func (contentState *ContentState) String() string {
	return Render(contentState, nil)
}
