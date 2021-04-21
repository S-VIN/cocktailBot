package main

const (
	WFINGR = iota
	WFNAME
	WFLIST
	DONE
)

type keyboardStatus struct {
	chatID int64
	status int
}

type ClientStatus struct {
	status map[int64]int
	shownCocktails map[int64][]string
}