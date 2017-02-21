package tests

import (
	"encoding/json"
	"testing"

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
