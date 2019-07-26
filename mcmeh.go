package main

import (
	"flag"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	. "gopkg.in/src-d/go-git.v4/_examples"
	"gopkg.in/src-d/go-git.v4/plumbing/object"
	"os"
	"strings"
)

var format = "%s contains %s"
var repo, searchType, word string

func parseCommitHash(co *object.Commit, ch chan string) {
	var hash = co.Hash.String()

	if strings.Contains(hash, word) {
		ch <- fmt.Sprintf(format, hash, word)
	}
}

func parseCommitMessage(c *object.Commit, ch chan string) {
	var m = c.Message

	if strings.Contains(m, word) {
		ch <- fmt.Sprintf(format, m, word)
	}
}

func main() {
	dir, err := os.Getwd()
	CheckIfError(err)

	flag.StringVar(&repo, "repo", dir, "Path to repo")
	flag.StringVar(&searchType, "type", "message", "Type to search. 'message' or 'hash'")
	flag.StringVar(&word, "word", "commit", "Word to find")
	flag.Parse()

	fs := map[string]func(*object.Commit, chan string){"hash": parseCommitHash, "message": parseCommitMessage}
	f := fs[searchType]

	r, err := git.PlainOpen(repo)
	CheckIfError(err)

	ref, err := r.Head()
	CheckIfError(err)

	cIter, err := r.Log(&git.LogOptions{From: ref.Hash()})
	CheckIfError(err)

	ch := make(chan string)

	var cCount int
	err = cIter.ForEach(func(c *object.Commit) error {
		go f(c, ch)
		cCount++
		return nil
	})
	CheckIfError(err)

	for e := range ch {
		fmt.Println(e)
	}
}
