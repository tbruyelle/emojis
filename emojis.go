package emojis

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"

	"github.com/imdario/go-html-transform/h5"
	"golang.org/x/net/html"
)

type Emojis []string

func (e Emojis) Random() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return e[r.Intn(len(e))]
}

func Load() (Emojis, error) {
	return loadFromApi()
}

func loadFromApi() (Emojis, error) {
	response, err := http.Get("https://api.github.com/emojis")
	if err != nil {
		return nil, err
	}
	var output map[string]string
	err = json.NewDecoder(response.Body).Decode(&output)
	if err != nil {
		return nil, err
	}
	var emojis Emojis
	for k, _ := range output {
		emojis = append(emojis, ":"+k+":")
	}
	return emojis, nil
}

func loadFromWebsite() (Emojis, error) {
	response, err := http.Get("http://www.emoji-cheat-sheet.com")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	tree, err := h5.New(response.Body)
	if err != nil {
		return nil, err
	}
	var emojis Emojis
	tree.Walk(func(n *html.Node) {
		if len(n.Attr) > 0 {
			for i := 0; i < len(n.Attr); i++ {
				if n.Attr[i].Key == "class" && n.Attr[i].Val == "name" {
					// Found an emoji
					emojis = append(emojis, ":"+n.FirstChild.Data+":")
				}
			}
		}
	})
	return emojis, nil
}
