package main

import "testing"

func TestAssignCookies1(t *testing.T) {
	g := []int{1, 2, 3}
	s := []int{1, 1}

	if findContentChildren(g, s) != 1 {
		t.Fatalf("Failed test case")
	}
}

func TestAssignCookies2(t *testing.T) {
	g := []int{1, 1}
	s := []int{1, 2, 3}

	if findContentChildren(g, s) != 2 {
		t.Fatalf("Failed test case")
	}
}

func TestAssignCookies3(t *testing.T) {

	g := []int{1, 2, 3}
	s := []int{}

	if findContentChildren(g, s) != 0 {
		t.Fatalf("Failed test case")
	}
}

func TestAssignCookies4(t *testing.T) {

	g := []int{10, 9, 8, 7}
	s := []int{5, 6, 7, 8}

	if findContentChildren(g, s) != 2 {
		t.Fatalf("Failed test case")
	}
}
