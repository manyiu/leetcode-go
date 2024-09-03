package designbrowserhistory

import "testing"

func TestConstructor(t *testing.T) {
	homepage := "leetcode.com"
	obj := Constructor(homepage)
	obj.Visit("google.com")
	obj.Visit("facebook.com")
	obj.Visit("youtube.com")
	if got := obj.Back(1); got != "facebook.com" {
		t.Errorf("Back() = %v, want %v", got, "facebook.com")
	}
	if got := obj.Back(1); got != "google.com" {
		t.Errorf("Back() = %v, want %v", got, "google.com")
	}
	if got := obj.Forward(1); got != "facebook.com" {
		t.Errorf("Forward() = %v, want %v", got, "facebook.com")
	}
	obj.Visit("linkedin.com")
	if got := obj.Forward(2); got != "linkedin.com" {
		t.Errorf("Forward() = %v, want %v", got, "linkedin.com")
	}
	if got := obj.Back(2); got != "google.com" {
		t.Errorf("Back() = %v, want %v", got, "google.com")
	}
	if got := obj.Back(7); got != "leetcode.com" {
		t.Errorf("Back() = %v, want %v", got, "leetcode.com")
	}
}
