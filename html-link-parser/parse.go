package htmllinkparser

import (
	"fmt"
	"io"

	"golang.org/x/net/html"
)

// Link represents a link (<a href="...">) in an HTML document.
type Link struct {
	Href string
	Text string
}

// Parse will take in an HTML document and will return a slice of links parsed.
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}
	linkNodes(doc, "")
	return nil, nil
}

func linkNodes(n *html.Node, padding string) {
	msg := n.Data
	if n.Type == html.ElementNode {
		msg = "<" + msg + ">"
	}
	fmt.Println(msg)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		linkNodes(n, padding+" ")
	}
}
