package main

import (
	"encoding/json"
	"fmt"

	"github.com/ejilay/draftjs"
)

func main() {

	// get your contentState JSON-string
	draftState := exampleDraftStateSource

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

var exampleDraftStateSource = `{
	  "entityMap": {
	    "0": {
	      "type": "LINK",
	      "mutability": "MUTABLE",
	      "data": {
		"url": "https://medium.com/@rajaraodv/how-draft-js-represents-rich-text-data-eeabb5f25cf2#.ce9y2wyux"
	     }}},
	  "blocks": [{
	      "text": "Rich text with link",
	      "type": "unstyled",
	      "depth": 0,
	      "inlineStyleRanges": [
		{
		  "offset": 0,
		  "length": 4,
		  "style": "BOLD"
		},
		{
		  "offset": 2,
		  "length": 10,
		  "style": "UNDERLINE"
		},
		{
		  "offset": 5,
		  "length": 4,
		  "style": "ITALIC"
		},
		{
		  "offset": 10,
		  "length": 4,
		  "style": "CODE"
		}
	      ],
	      "entityRanges": [{
		  "offset": 15,
		  "length": 4,
		  "key": 0
		}],
	      "data": {}
	    }]}`
