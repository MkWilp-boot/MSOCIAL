package main

import (
	_ "github.com/go-sql-driver/mysql"
)

// AlreadyIn tells if users is alread connected
type AlreadyIn struct {
	logged bool
}

type sessionLogged string

var in = AlreadyIn{logged: false}

// ChangeLogState changes Session state
func (a *AlreadyIn) ChangeLogState(op bool) {
	if op {
		a.logged = true
	} else {
		a.logged = false
	}
}

// SelectLogState returns the value of session
func (a *AlreadyIn) SelectLogState() bool {
	return a.logged
}
