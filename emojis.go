package emojis

import (
	"math/rand"
	"net/http"
	"time"

	"code.google.com/p/go-html-transform/h5"
	"code.google.com/p/go.net/html"
)

type Emojis []string

func (e Emojis) Random() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return e[r.Intn(len(e))]
}

func Load() (Emojis, error) {
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
