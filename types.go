package draftjs

// ContentState
// https://github.com/facebook/draft-js/blob/master/src/model/encoding/RawDraftContentState.js
type ContentState struct {
	Blocks    []*ContentBlock    `json:"blocks"`
	EntityMap map[string]*Entity `json:"entityMap"`
}

// ContentBlock
// https://github.com/facebook/draft-js/blob/master/src/model/encoding/RawDraftContentBlock.js
type ContentBlock struct {
	Key               string              `json:"key"`
	Type              string              `json:"type"`
	Text              string              `json:"text"`
	Depth             int                 `json:"depth"`
	InlineStyleRanges []*InlineStyleRange `json:"inlineStyleRanges"`
	EntityRanges      []*EntityRange      `json:"entityRanges"`
	Data              interface{}         `json:"data"`
}

// Entity
// https://github.com/facebook/draft-js/blob/master/src/model/encoding/RawDraftEntity.js
type Entity struct {
	Type       string            `json:"type"`
	Mutability string            `json:"mutability"`
	Data       map[string]string `json:"data"`
}

type Range struct {
	Offset int `json:"offset"`
	Length int `json:"length"`
}

// InlineStyleRange
// https://github.com/facebook/draft-js/blob/master/src/model/encoding/InlineStyleRange.js
type InlineStyleRange struct {
	Style string `json:"style"`
	Range
}

// EntityRange
// https://github.com/facebook/draft-js/blob/master/src/model/encoding/EntityRange.js
type EntityRange struct {
	Key int `json:"key"`
	Range
}
