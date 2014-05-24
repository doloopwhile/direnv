package main

import "strings"

type vim int
var VIM vim

// Is this even desirable ?
func (x vim) Hook() string {
	panic("TODO")
}

// FIXME: Make sure this escaping is valid
func (x vim) Escape(str string) string {
	return strings.Replace(str, "'", "''", -1)
}

// FIXME: escape key and value
func (x vim) Export(key, value string) string {
	return "let $" + key + " = '" + x.Escape(value) + "'\n"
}

// FIXME: escape key
func (x vim) Unset(key string) string {
	return "unlet $" + key + "\n"
}

