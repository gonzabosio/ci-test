package main

import "testing"

func TestPrint(t *testing.T) {
	prnt := PrintText()
	if prnt != "Testing continuous integration in Go" {
		t.Error("The print texts are not the same")
	}
}
