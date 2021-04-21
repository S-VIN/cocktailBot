package main

const (
	WFINGR = iota
	WFNAME
	WFLIST
	DONE
)

type ClientStatus struct {
	status map[int64]int
	shownCocktails map[int64][]string
}