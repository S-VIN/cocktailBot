package main

const (
	WFINGR = iota
	WFNAME
	DONE
)


type keyboardStatus struct {
	chatID int64
	status int
}

type ClientStatus struct{
	status map[int64]int
}

var clientStatus ClientStatus