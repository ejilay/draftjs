package tests

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/ejilay/draftjs"
)

func TestRender(t *testing.T) {
	var (
		contentStates []draftjs.ContentState
		err           error
	)

	if err = json.Unmarshal([]byte(testString), &contentStates); err != nil {
		t.Errorf("Failed unmarshal content: %v", err)
		return
	}

	config := draftjs.NewDefaultConfig()
	i := 0
	for _, block := range contentStates {
		s := draftjs.Render(&block, config)
		if s != needStrings[i] {
			t.Errorf("\n%s\n", s)
			t.Errorf("\n%s\n", needStrings[i])
		}
		i++
	}
}

func TestRenderOneSymbol(t *testing.T) {
	var (
		contentState = new(draftjs.ContentState)
		err          error
	)

	if err = json.Unmarshal([]byte(testStringOneSymbol), contentState); err != nil {
		t.Errorf("Failed unmarshal content: %v", err)
		return
	}

	config := draftjs.NewDefaultConfig()
	s := draftjs.Render(contentState, config)
	if s != testStringOneSymbolExpected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", testStringOneSymbolExpected, s)
	}
}

func TestRenderWrongRanges(t *testing.T) {
	var (
		contentState = new(draftjs.ContentState)
		err          error
	)

	if err = json.Unmarshal([]byte(testStringWrongRanges), contentState); err != nil {
		t.Errorf("Failed unmarshal content: %v", err)
		return
	}

	config := draftjs.NewDefaultConfig()
	s := draftjs.Render(contentState, config)
	if s != testStringWrongRangesExpected {
		t.Errorf("\nExpected:\n%s\nGot:\n%s", testStringWrongRangesExpected, s)
	}
}

var S string // preventing compiler optimization

func BenchmarkRender(b *testing.B) {
	var (
		contentStates []draftjs.ContentState
		err           error
	)

	if err = json.Unmarshal([]byte(testString), &contentStates); err != nil {
		b.Errorf("Failed unmarshal content: %v", err)
		return
	}

	config := draftjs.NewDefaultConfig()
	config.Precache()

	var buf bytes.Buffer
	buf.Grow(10 * 1024 * 1024)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		for _, block := range contentStates {
			draftjs.RenderWithBuf(&block, config, &buf)
		}
		buf.Reset()
	}
	S = buf.String()
}
