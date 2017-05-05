// draftjs exporter for go language
package draftjs

import (
	"bytes"
)

func renderBlocks(contentState *ContentState, config *Config, blockIterator *BlockIterator, buf *bytes.Buffer) {

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
		PerformInlineStylesAndEntities(contentState, currentBlock, config, buf)
		if blockIterator.HasNext() && blockIterator.NextBlock().Depth > blockIterator.block.Depth {
			renderBlocks(contentState, config, blockIterator.StepNext(), buf)
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

	if config == nil {
		config = NewDefaultConfig()
	}

	config.Precache()

	RenderWithBuf(contentState, config, &buf)

	return buf.String()
}

// RenderWithBuf renders Draft.js content state to buffer with config
func RenderWithBuf(contentState *ContentState, config *Config, buf *bytes.Buffer) {
	renderBlocks(contentState, config, NewBlockIterator(contentState), buf)
}

// Interface implementation
func (contentState *ContentState) String() string {
	return Render(contentState, nil)
}

func RenderPlainText(contentState *ContentState) string {
	var buffer bytes.Buffer
	textLen := len(contentState.Blocks)
	for _, block := range contentState.Blocks {
		textLen += len(block.Text)
	}
	buffer.Grow(textLen)
	RenderPlainTextToBuffer(contentState, &buffer)
	return buffer.String()
}

func RenderPlainTextToBuffer(contentState *ContentState, buffer *bytes.Buffer) {
	for _, block := range contentState.Blocks {
		buffer.WriteString(block.Text)
		buffer.WriteString("\n")
	}
}
