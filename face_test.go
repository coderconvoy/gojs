package gojs

import "testing"

func Test_Available(t *testing.T) {
	res, _ := Single.Asset("assets/js/go_menu.js")
	if len(res) == 0 {
		t.Fail()
	}
}
