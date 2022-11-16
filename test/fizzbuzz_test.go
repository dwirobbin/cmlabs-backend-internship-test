package test

import (
	"testing"

	"github.com/dwirobbin/cmlabs-backend-internship-test/helper"
)

func TestFizzBuzzSuccess(t *testing.T) {
	testCases := []struct {
		index  int
		output string
	}{
		{1, "1"},
		{2, "2"},
		{3, "fizz"},
		{4, "4"},
		{5, "buzz"},
		{6, "fizz"},
		{7, "7"},
		{8, "8"},
		{9, "fizz"},
		{10, "buzz"},
		{11, "11"},
		{12, "fizz"},
		{13, "13"},
		{14, "14"},
		{15, "fizzbuzz"},
		{16, "16"},
		{17, "17"},
		{18, "fizz"},
		{19, "19"},
		{20, "buzz"},
		{100, "buzz"},
	}

	for _, test := range testCases {
		output := helper.FizzBuzz(test.index, 100)
		if output != test.output {
			t.Errorf("expected in-(%d) = %s, got %s", test.index, output, test.output)
		}
	}
}

func TestFizzBuzzFail(t *testing.T) {
	testCases := []struct {
		index  int
		output string
	}{
		{1, "1"},
		{2, "2"},
		{3, "fizzbuzz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Buzz"},
		{10, "Fizz"},
		{100, "fizzbuzz"},
	}

	for _, test := range testCases {
		output := helper.FizzBuzz(test.index, 100)
		if output != test.output {
			t.Errorf("expected in-(%d) = %s, got %s", test.index, output, test.output)
		}
	}
}
