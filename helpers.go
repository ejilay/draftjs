package draftjs

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"text/template"
	"unicode/utf8"
)

func GetDescriptorFromMap(key string, sourceMap map[string]*HTMLDescriptor) *HTMLDescriptor {
	if sourceMap == nil {
		return nil
	}
	if v, ok := sourceMap[key]; ok {
		return v
	}
	return nil
}

func GetBlockWrapperTag(block *ContentBlock, config *HTMLConfig) string {
	if block == nil || config == nil {
		return ""
	}
	options := config.GetBlockMapElement(block.Type)
	if options == nil {
		return ""
	}
	return options.Wrapper
}

func GetBlockWrapperStartTag(block *ContentBlock, config *HTMLConfig) string {
	tagName := GetBlockWrapperTag(block, config)
	if tagName == "" {
		return ""
	}
	return fmt.Sprintf("<%s>", tagName)
}

func GetBlockWrapperEndTag(block *ContentBlock, config *HTMLConfig) string {
	tagName := GetBlockWrapperTag(block, config)
	if tagName == "" {
		return ""
	}
	return fmt.Sprintf("</%s>", tagName)
}

func GetBlockTag(block *ContentBlock, config *HTMLConfig) string {
	if block == nil || config == nil {
		return ""
	}
	options := config.GetBlockMapElement(block.Type)
	if options == nil {
		return ""
	}
	return options.Element
}

func GetBlockStartTag(block *ContentBlock, config *HTMLConfig) string {
	tagName := GetBlockTag(block, config)
	if tagName == "" {
		return ""
	}
	return fmt.Sprintf("<%s>", tagName)
}

func GetBlockEndTag(block *ContentBlock, config *HTMLConfig) string {
	tagName := GetBlockTag(block, config)
	if tagName == "" {
		return ""
	}
	return fmt.Sprintf("</%s>", tagName)
}

func GetStylemapElement(style *InlineStyleRange, config *HTMLConfig) string {
	if style == nil || config == nil {
		return ""
	}
	options := config.GetStyleMapElement(style.Style)
	if options == nil {
		return ""
	}
	return options.Element
}

func GetStyleStartTag(style *InlineStyleRange, config *HTMLConfig) string {
	tagName := GetStylemapElement(style, config)
	if tagName == "" {
		return ""
	}
	return fmt.Sprintf("<%s>", tagName)
}

func GetStyleEndTag(style *InlineStyleRange, config *HTMLConfig) string {
	tagName := GetStylemapElement(style, config)
	if tagName == "" {
		return ""
	}
	return fmt.Sprintf("</%s>", tagName)
}

func GetEntityDecorator(content *ContentState, entityRange *EntityRange, config *HTMLConfig) (Decorator, *Entity) {
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

func GetEntityStartTag(content *ContentState, entityRange *EntityRange, config *HTMLConfig) string {
	decorator, entity := GetEntityDecorator(content, entityRange, config)
	if decorator == nil {
		return ""
	}
	return decorator.RenderBeginning(entity.Data)
}

func GetEntityEndTag(content *ContentState, entityRange *EntityRange, config *HTMLConfig) string {
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

func PerformInlineStylesAndEntities(content *ContentState, block *ContentBlock, config *HTMLConfig) string {
	ranges, noStyles := GetRanges(block)
	if noStyles {
		return template.HTMLEscapeString(block.Text)
	}

	var buf bytes.Buffer
	buf.Grow(256 * 1024) // с потолка
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

	return buf.String()
}

func GetEntityForRange(r *Range, block *ContentBlock) []*EntityRange {
	res := make([]*EntityRange, 0, len(block.EntityRanges))
	if block.EntityRanges == nil || len(block.EntityRanges) == 0 {
		return res
	}
	for _, entityRange := range block.EntityRanges {
		if r.Offset >= entityRange.Offset && r.Offset+r.Length <= entityRange.Offset+entityRange.Length {
			res = append(res, entityRange)
		}
	}
	return res
}

func GetStyleForRange(r *Range, block *ContentBlock) []*InlineStyleRange {
	res := make([]*InlineStyleRange, 0, len(block.InlineStyleRanges))
	if block.InlineStyleRanges == nil || len(block.InlineStyleRanges) == 0 {
		return res
	}
	for _, styleRange := range block.InlineStyleRanges {
		if r.Offset >= styleRange.Offset && r.Offset+r.Length <= styleRange.Offset+styleRange.Length {
			res = append(res, styleRange)
		}
	}
	return res
}

// bool == fullstring (no styles)
func GetRanges(block *ContentBlock) ([]*Range, bool) {
	res := make([]*Range, 1)
	res[0] = new(Range)
	res[0].Offset = 0
	res[0].Length = utf8.RuneCountInString(block.Text)

	if len(block.InlineStyleRanges)+len(block.EntityRanges) == 0 {
		return res, true
	}
	breakPoints := GetBreakPoints(block)
	prev := 0
	res = make([]*Range, 0, len(breakPoints))
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
		if lastRange.Length+lastRange.Offset < utf8.RuneCountInString(block.Text) {
			t := new(Range)
			t.Offset = lastRange.Offset + lastRange.Length
			t.Length = utf8.RuneCountInString(block.Text) - t.Offset
			res = append(res, t)
		}
	}
	return res, false
}

func GetBreakPoints(block *ContentBlock) []int {
	breakPoints := make([]int, 0, utf8.RuneCountInString(block.Text))

	inArray := func(v int, arr []int) bool {
		for i := 0; i < len(arr); i++ {
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

	for _, styleRange := range ranges {
		if !inArray(styleRange.Offset, breakPoints) {
			breakPoints = append(breakPoints, styleRange.Offset)
		}
		if !inArray(styleRange.Offset+styleRange.Length, breakPoints) {
			breakPoints = append(breakPoints, styleRange.Offset+styleRange.Length)
		}
	}

	sort.Ints(breakPoints)

	return breakPoints
}
