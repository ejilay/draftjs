package draftjs

type BlockIterator struct {
	block        *ContentBlock
	index        int
	contentState *ContentState
}

func NewBlockIterator(contentState *ContentState) *BlockIterator {
	bi := new(BlockIterator)
	bi.contentState = contentState
	bi.index = 0
	if len(contentState.Blocks) > 0 {
		bi.block = contentState.Blocks[0]
	}
	return bi
}

func (bi *BlockIterator) HasNext() bool {
	return len(bi.contentState.Blocks) != 0 && bi.index+1 < len(bi.contentState.Blocks)
}

func (bi *BlockIterator) StepNext() *BlockIterator {
	if bi.HasNext() {
		bi.index++
		bi.block = bi.contentState.Blocks[bi.index]
		return bi
	}
	bi.block = nil
	return nil
}

func (bi BlockIterator) NextBlock() *ContentBlock {
	if bi.HasNext() {
		return bi.contentState.Blocks[bi.index+1]
	}
	return nil
}
