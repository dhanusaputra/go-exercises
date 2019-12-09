package adventure

import (
	"encoding/json"
	"io"
)

//JSONStory ...
func JSONStory(r io.Reader) (Story, error) {
	d := json.NewDecoder(r)
	var story Story
	if err := d.Decode(&story); err != nil {
		return nil, err
	}
	return story, nil
}

//Story ...
type Story map[string]Chapter

//Chapter ...
type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

//Option ...
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}
