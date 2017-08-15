# Draft.js Exporter

[![Go Report Card](https://goreportcard.com/badge/github.com/ejilay/draftjs)](https://goreportcard.com/report/github.com/ejilay/draftjs)

[Draft.js](https://facebook.github.io/draft-js/) is a framework for
building rich text editors. However, it does not support exporting
documents at HTML. This package is designed to take the raw `ContentState`
(output of [`convertToRaw`](https://facebook.github.io/draft-js/docs/api-reference-data-conversion.html#converttoraw))
from Draft.js and convert it to HTML using [Go](https://golang.org). Mostly it
useful for server-side rendering. I should note this package does not provide
any input validation and assumes correct and safe input data.

## Usage

```go
func main() {

	// get your contentState JSON-string
	draftState := exampleDraftStateSource

	// make auxiliary variable
	contentState := draftjs.ContentState{}
	json.Unmarshal([]byte(draftState), &contentState) // don't forget error handling

	// prepare some config (HTML here)
	config := draftjs.DefaultConfig()

	// and just render content state to HTML-string
	s := draftjs.Render(&contentState, config)

	// that's it
	fmt.Println(s)
}
```
For _RawContentState_ like this
```json
{
  "entityMap": {
    "0": {
      "type": "LINK",
      "data": {
        "url": "https://medium.com/@rajaraodv/how-draft-js-represents-rich-text-data-eeabb5f25cf2#.ce9y2wyux"
      }
    }
  },
  "blocks": [
    {
      "text": "Rich text with link",
      "type": "unstyled",
      "depth": 0,
      "inlineStyleRanges": [
        {
          "offset": 0,
          "length": 4,
          "style": "BOLD"
        }, {
          "offset": 2,
          "length": 10,
          "style": "UNDERLINE"
        }, {
          "offset": 5,
          "length": 4,
          "style": "ITALIC"
        }, {
          "offset": 10,
          "length": 4,
          "style": "CODE"
        }
      ],
      "entityRanges": [{
          "offset": 15,
          "length": 4,
          "key": 0
       }]
}]}
```
It will give something like this but without indention:
```html
<p>
	<strong>Ri</strong>
	<strong>
		<ins>ch</ins>
	</strong>
	<ins>
		<em>text</em>
	</ins>
	<ins>
		<code>wi</code>
	</ins>
	<code>th</code>
	<a href="https://medium.com/@rajaraodv/how-draft-js-represents-rich-text-data-eeabb5f25cf2#.ce9y2wyux" target="_blank">link</a>
</p>

```
That look like
<p>
	<strong>Ri</strong>
	<strong>
		<ins>ch</ins>
	</strong>
	<ins></ins>
	<ins>
		<em>text</em>
	</ins>
	<ins></ins>
	<ins>
		<code>wi</code>
	</ins>
	<code>th</code>
	<a href="https://medium.com/@rajaraodv/how-draft-js-represents-rich-text-data-eeabb5f25cf2#.ce9y2wyux" target="_blank">link</a>
</p>




## Setup

You'll need Golang installed first :o)

```bash
go get github.com/ejilay/draftjs
```

## Testing

To test run the following command in project's root directory:

```bash
go test ./...
```
