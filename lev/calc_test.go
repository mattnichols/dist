package lev

import "testing"

func TestCalc(t *testing.T) {
	t.Run("Kittens", func(t *testing.T) {
		if Calc("sitting", "kitten") != 3 {
			t.Error("Calc sitting kitten should equal 3")
		}
	})

	t.Run("Blank", func(t *testing.T) {
		if Calc("", "") != 0 {
			t.Error("Calc with blanks should be zero")
		}
		if Calc("kitten", "") != 6 {
			t.Error("Calc with blank should be the length of the non-blank str")
		}
		if Calc("", "sitting") != 7 {
			t.Error("Calc with blank should be the length of the non-blank str")
		}
	})
}
