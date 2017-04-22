package draftjs

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"text/template"
	"unicode/utf8"
)

func GetDescriptorFromMap(key string, sourceMap map[string]*Descriptor) *Descriptor {
	if sourceMap == nil {
		return nil
	}
	if v, ok := sourceMap[key]; ok {
		return v
	}
	return nil
}

func GetBlockWrapperTag(block *ContentBlock, config *Config) string {
	if block == nil || config == nil {
		return ""
	}
	options := config.GetBlockMapElement(block.Type)
	if options == nil {
		return ""
	}
	return options.Wrapper
}

func GetBlockWrapperStartTag(block *ContentBlock, config *Config) string {
	if block == nil {
		return ""
	}
	const cacheKey = "GetBlockWrapperStartTag"
	if tag, exist := config.GetFromCache(cacheKey, block.Type); exist {
		return tag
	}
	tagName := GetBlockWrapperTag(block, config)
	var tag string
	if tagName == "" {
		tag = ""
	} else {
		tag = fmt.Sprintf("<%s>", tagName)
	}
	config.SetToCache(cacheKey, block.Type, tag)
	return tag
}

func GetBlockWrapperEndTag(block *ContentBlock, config *Config) string {
	if block == nil {
		return ""
	}
	const cacheKey = "GetBlockWrapperEndTag"
	if tag, exist := config.GetFromCache(cacheKey, block.Type); exist {
		return tag
	}
	tagName := GetBlockWrapperTag(block, config)
	var tag string
	if tagName == "" {
		tag = ""
	} else {
		tag = fmt.Sprintf("</%s>", tagName)
	}
	config.SetToCache(cacheKey, block.Type, tag)
	return tag
}

func GetBlockTag(block *ContentBlock, config *Config) string {
	if block == nil || config == nil {
		return ""
	}
	options := config.GetBlockMapElement(block.Type)
	if options == nil {
		return ""
	}
	return options.Element
}

func GetBlockStartTag(block *ContentBlock, config *Config) string {
	if block == nil {
		return ""
	}
	const cacheKey = "GetBlockStartTag"
	if tag, exist := config.GetFromCache(cacheKey, block.Type); exist {
		return tag
	}
	tagName := GetBlockTag(block, config)
	var tag string
	if tagName == "" {
		tag = ""
	} else {
		tag = fmt.Sprintf("<%s>", tagName)
	}
	config.SetToCache(cacheKey, block.Type, tag)
	return tag
}

func GetBlockEndTag(block *ContentBlock, config *Config) string {
	if block == nil {
		return ""
	}
	const cacheKey = "GetBlockEndTag"
	if tag, exist := config.GetFromCache(cacheKey, block.Type); exist {
		return tag
	}
	tagName := GetBlockTag(block, config)
	var tag string
	if tagName == "" {
		tag = ""
	} else {
		tag = fmt.Sprintf("</%s>", tagName)
	}
	config.SetToCache(cacheKey, block.Type, tag)
	return tag
}

func GetStylemapElement(style *InlineStyleRange, config *Config) string {
	if style == nil || config == nil {
		return ""
	}
	options := config.GetStyleMapElement(style.Style)
	if options == nil {
		return ""
	}
	return options.Element
}

func GetStyleStartTag(style *InlineStyleRange, config *Config) string {
	if style == nil {
		return ""
	}
	const cacheKey = "GetStyleStartTag"
	if tag, exist := config.GetFromCache(cacheKey, style.Style); exist {
		return tag
	}
	tagName := GetStylemapElement(style, config)
	var tag string
	if tagName == "" {
		tag = ""
	} else {
		tag = fmt.Sprintf("<%s>", tagName)
	}
	config.SetToCache(cacheKey, style.Style, tag)
	return tag
}

func GetStyleEndTag(style *InlineStyleRange, config *Config) string {
	if style == nil {
		return ""
	}
	const cacheKey = "GetStyleEndTag"
	if tag, exist := config.GetFromCache(cacheKey, style.Style); exist {
		return tag
	}
	tagName := GetStylemapElement(style, config)
	var tag string
	if tagName == "" {
		tag = ""
	} else {
		tag = fmt.Sprintf("</%s>", tagName)
	}
	config.SetToCache(cacheKey, style.Style, tag)
	return tag
}

func GetEntityDecorator(content *ContentState, entityRange *EntityRange, config *Config) (Decorator, *Entity) {
	var (
		entity *Entity
		ok     bool
	)
	if entity, ok = content.EntityMap[strconv.Itoa(entityRange.Key)]; !ok || entity == nil {
		return nil, nil
	}
	descriptor := config.GetEntityDecorator(entity.Type)
	if descriptor == nil || descriptor.Decorator == nil {
		return nil, nil
	}
	return descriptor.Decorator, entity
}

func GetEntityStartTag(content *ContentState, entityRange *EntityRange, config *Config) string {
	decorator, entity := GetEntityDecorator(content, entityRange, config)
	if decorator == nil {
		return ""
	}
	return decorator.RenderBeginning(entity.Data)
}

func GetEntityEndTag(content *ContentState, entityRange *EntityRange, config *Config) string {
	decorator, entity := GetEntityDecorator(content, entityRange, config)
	if decorator == nil {
		return ""
	}
	return decorator.RenderEnding(entity.Data)
}

// без конвертации к рунам и доп памяти
func substring(s string, start int, end int) string {
	start_str_idx := 0
	i := 0
	for j := range s {
		if i == start {
			start_str_idx = j
		}
		if i == end {
			return s[start_str_idx:j]
		}
		i++
	}
	return s[start_str_idx:]
}

func PerformInlineStylesAndEntities(content *ContentState, block *ContentBlock, config *Config, buf *bytes.Buffer) {
	ranges, noStyles := GetRanges(block)
	if noStyles {
		buf.WriteString(template.HTMLEscapeString(block.Text))
		return
	}

	for _, rng := range ranges {
		styles := GetStyleForRange(rng, block)
		entities := GetEntityForRange(rng, block)
		for i := 0; i < len(entities); i++ {
			buf.WriteString(GetEntityStartTag(content, entities[i], config))
		}
		for i := 0; i < len(styles); i++ {
			buf.WriteString(GetStyleStartTag(styles[i], config))
		}
		buf.WriteString(template.HTMLEscapeString(substring(block.Text, rng.Offset, rng.Offset+rng.Length)))
		for i := len(styles) - 1; i >= 0; i-- {
			buf.WriteString(GetStyleEndTag(styles[i], config))
		}
		for i := len(entities) - 1; i >= 0; i-- {
			buf.WriteString(GetEntityEndTag(content, entities[i], config))
		}
	}

}

func GetEntityForRange(r *Range, block *ContentBlock) []*EntityRange {
	if block.EntityRanges == nil || len(block.EntityRanges) == 0 {
		return nil
	}
	res := make([]*EntityRange, 0, 0)
	for _, entityRange := range block.EntityRanges {
		if r.Offset >= entityRange.Offset && r.Offset+r.Length <= entityRange.Offset+entityRange.Length {
			res = append(res, entityRange)
		}
	}
	return res
}

func GetStyleForRange(r *Range, block *ContentBlock) []*InlineStyleRange {

	if block.InlineStyleRanges == nil || len(block.InlineStyleRanges) == 0 {
		return nil
	}
	res := make([]*InlineStyleRange, 0, 0)
	for _, styleRange := range block.InlineStyleRanges {
		if r.Offset >= styleRange.Offset && r.Offset+r.Length <= styleRange.Offset+styleRange.Length {
			res = append(res, styleRange)
		}
	}
	return res
}

// bool == fullstring (no styles)
func GetRanges(block *ContentBlock) ([]*Range, bool) {
	if len(block.InlineStyleRanges)+len(block.EntityRanges) == 0 {
		return nil, true
	}

	breakPoints, runeCount := GetBreakPoints(block)
	prev := 0
	res := make([]*Range, 0, 0)
	var lastRange *Range
	for _, v := range breakPoints {
		if v == prev {
			continue
		}
		t := new(Range)
		t.Offset = prev
		t.Length = v - prev
		prev = v
		res = append(res, t)
		lastRange = t
	}
	if lastRange != nil {
		if lastRange.Length+lastRange.Offset < runeCount {
			t := new(Range)
			t.Offset = lastRange.Offset + lastRange.Length
			t.Length = utf8.RuneCountInString(block.Text) - t.Offset
			res = append(res, t)
		}
	}
	return res, false
}

func GetBreakPoints(block *ContentBlock) ([]int, int) {
	runeCount := utf8.RuneCountInString(block.Text)
	breakPoints := make([]int, runeCount, runeCount)

	inArray := func(v int, arr []int) bool {
		for i := len(arr) - 1; i >= 0; i-- {
			if v == arr[i] {
				return true
			}
		}
		return false
	}

	ranges := make([]*Range, 0, len(block.InlineStyleRanges)+len(block.EntityRanges))
	for _, styleRange := range block.InlineStyleRanges {
		ranges = append(ranges, &styleRange.Range)
	}
	for _, entityRange := range block.EntityRanges {
		ranges = append(ranges, &entityRange.Range)
	}

	breakPointsCount := 0
	for _, styleRange := range ranges {
		if !inArray(styleRange.Offset, breakPoints[:breakPointsCount]) {
			breakPoints[breakPointsCount] = styleRange.Offset
			breakPointsCount++
		}
		if !inArray(styleRange.Offset+styleRange.Length, breakPoints[:breakPointsCount]) {
			breakPoints[breakPointsCount] = styleRange.Offset + styleRange.Length
			breakPointsCount++
		}
	}

	breakPoints = breakPoints[:breakPointsCount]
	sort.Ints(breakPoints)

	return breakPoints, runeCount
}
