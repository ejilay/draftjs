package tests

import (
	"encoding/json"
	"testing"

	"bytes"
	"github.com/ejilay/draftjs"
)

func TestRender(t *testing.T) {
	contentStates := []draftjs.ContentState{}
	var err error
	if err = json.Unmarshal([]byte(TestString), &contentStates); err != nil {
		t.Errorf("Failed unmarshal content: %v", err)
		return
	}

	config := draftjs.NewDefaultConfig()
	i := 0
	for _, block := range contentStates {
		s := draftjs.Render(&block, config)
		if s != NeedString[i] {
			t.Errorf("\n%s\n", s)
			t.Errorf("\n%s\n", NeedString[i])
		}
		i++
	}
}

var S string // preventing compiler optimization

func BenchmarkRender(b *testing.B) {
	contentStates := []draftjs.ContentState{}
	var err error
	if err = json.Unmarshal([]byte(TestString), &contentStates); err != nil {
		b.Errorf("Failed unmarshal content: %v", err)
		return
	}

	config := draftjs.NewDefaultConfig()
	config.Compile()

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
