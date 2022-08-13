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
		{3, "Fizz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Fizz"},
		{10, "Buzz"},
		{11, "11"},
		{12, "Fizz"},
		{13, "13"},
		{14, "14"},
		{15, "FizzBuzz"},
		{16, "16"},
		{17, "17"},
		{18, "Fizz"},
		{19, "19"},
		{20, "Buzz"},
		{100, "Buzz"},
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
		{3, "FizzBuzz"},
		{4, "4"},
		{5, "Buzz"},
		{6, "Fizz"},
		{7, "7"},
		{8, "8"},
		{9, "Buzz"},
		{10, "Fizz"},
		{100, "FizzBuzz"},
	}

	for _, test := range testCases {
		output := helper.FizzBuzz(test.index, 100)
		if output != test.output {
			t.Errorf("expected in-(%d) = %s, got %s", test.index, output, test.output)
		}
	}
}
