package tests

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/ejilay/draftjs"
)

func TestRender(t *testing.T) {
	var (
		contentState draftjs.ContentState
		err          error
	)
	config := draftjs.NewDefaultConfig()
	for _, test := range GetTestsTable() {
		if err = json.Unmarshal([]byte(test.State), &contentState); err != nil {
			t.Errorf("Failed unmarshal content (test \"%s\"): \n%v\n%s", test.Name, err, test.State)
			return
		}

		s := draftjs.Render(&contentState, config)
		if s != test.Expected {
			t.Errorf("Error (test \"%s\"):\nGot: %s\nExpected: %s", test.Name, s, test.Expected)
		}
	}
}

func TestRenderPlainText(t *testing.T) {
	contentState := draftjs.ContentState{}
	var err error
	for _, test := range GetTestsPlainTable() {

		if err = json.Unmarshal([]byte(test.State), &contentState); err != nil {
			t.Errorf("Failed unmarshal content (test \"%s\"): \n%v\n%s", test.Name, err, test.State)
			return
		}

		s := draftjs.RenderPlainText(&contentState)
		if s != test.Expected {
			t.Errorf("Error (test \"%s\"):\nGot: %s\nExpected: %s", test.Name, s, test.Expected)
		}

	}
}

var S string // preventing compiler optimization

func BenchmarkRender(b *testing.B) {
	var (
		contentStates []draftjs.ContentState
		err           error
	)

	if err = json.Unmarshal([]byte(ExampleDraftStateSource), &contentStates); err != nil {
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
