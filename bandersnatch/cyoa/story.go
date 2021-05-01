package cyoa

import (
	"encoding/json"
	"io"
)

// var defaultHandlerTemplate = `
// <!DOCTYPE html>
// <html lang="en">
// <head>
//     <meta charset="UTF-8">
//     <title>Bandersnatch</title>
// </head>
// <body>
//     <h1>{{.Title}}</h1>
//     {{range .Paragraphs}}
//         <p>{{.}}</p>
//     {{end}}
//     <ul>
//         {{range .Options}}
//         <li> <a href="/{{.Chapter}}"></a>{{.Text}}</li>
//         {{end}}
//     </ul>
// </body>
// </html>`

type Story map[string]Chapter

type Chapter struct {
	Title      string    `json:"title"`
	Paragraphs []string  `json:"story"`
	Options    []Options `json:"options"`
}

type Options struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func JsonStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}
