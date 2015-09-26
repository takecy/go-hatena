package main

import (
	"fmt"

	"github.com/takecy/go-hatena/hatena"
	"golang.org/x/oauth2"
)

func main() {
	// go get golang.org/x/oauth2
	// oauth token

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "... your access token ..."},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	cli := hatena.NewHatena(tc)

	count, err := cli.BookMarks.Count("https://www.google.co.jp")
	if err != nil {
		panic(err)
	}

	fmt.Printf("count:%d", count)
}
