package main

import (
	"encoding/json"
	"fmt"

	"github.com/ejilay/draftjs"
	"github.com/ejilay/draftjs/tests"
)

func main() {

	// get your contentState JSON-string
	draftState := tests.ExampleDraftStateSource

	// make auxiliary variable
	contentState := draftjs.ContentState{}
	json.Unmarshal([]byte(draftState), &contentState) // don't forget error handling

	// prepare some config (HTML here)
	config := draftjs.NewDefaultConfig()

	// and just render content state to HTML-string
	s := draftjs.Render(&contentState, config)

	// that's it
	fmt.Println(s)
}
