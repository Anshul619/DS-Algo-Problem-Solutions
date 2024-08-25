package main

import (
	"testing"
)

func TestWordDictionary(t *testing.T) {

	wd := Constructor1()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")

	if wd.Search("pad") {
		t.Errorf("failed1")
	}
	if !wd.Search("bad") {
		t.Errorf("failed2")
	}
	if !wd.Search(".ad") {
		t.Errorf("failed3")
	}
	if !wd.Search("b..") {
		t.Errorf("failed4")
	}
}

func TestWordDictionary1(t *testing.T) {

	wd := Constructor1()
	wd.AddWord("a")
	wd.AddWord("a")

	if !wd.Search(".") {
		t.Errorf("failed1")
	}
	if !wd.Search("a") {
		t.Errorf("failed2")
	}
	if wd.Search("aa") {
		t.Errorf("failed3")
	}

	if wd.Search(".a") {
		t.Errorf("failed4")
	}
	if wd.Search("a.") {
		t.Errorf("failed5")
	}
}

func TestWordDictionary2(t *testing.T) {

	wd := Constructor1()
	wd.AddWord("at")
	wd.AddWord("and")
	wd.AddWord("an")
	wd.AddWord("add")
	wd.AddWord("bat")

	if !wd.Search(".at") {
		t.Errorf("failed1")
	}
}

func TestWordDictionary3(t *testing.T) {

	wd := Constructor1()
	wd.AddWord("at")
	wd.AddWord("and")
	wd.AddWord("an")
	wd.AddWord("add")
	wd.AddWord("bat")

	if wd.Search("b.") {
		t.Errorf("failed1")
	}
	if wd.Search(".") {
		t.Errorf("failed1")
	}
}
