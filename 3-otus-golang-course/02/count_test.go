package foobar

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCount(t *testing.T) {

	testTable := []struct {
		str                         string
		findibleRune                rune
		expectedDuplicatedRuneCount int
	}{
		{
			str:                         "qwerasdfe",
			findibleRune:                'e',
			expectedDuplicatedRuneCount: 2,
		},
		{
			str:                         "",
			findibleRune:                'e',
			expectedDuplicatedRuneCount: 0,
		},
		{
			str:                         "aaa",
			findibleRune:                's',
			expectedDuplicatedRuneCount: 0,
		},
		{
			str:                         "aaa",
			findibleRune:                'a',
			expectedDuplicatedRuneCount: 3,
		},
	}

	for _, testCase := range testTable {
		result := Count(testCase.str, testCase.findibleRune)

		t.Logf("Calling Count(%s), finded rune '%c', result %d\n", testCase.str, testCase.findibleRune, result)

		assert.Equal(t, testCase.expectedDuplicatedRuneCount, result,
			fmt.Sprintf("bad count for %s: got %d expected %d", testCase.str, result, testCase.expectedDuplicatedRuneCount))
		/*if c := Count(testCase.str, testCase.findibleRune); c != testCase.expectedDuplicatedRuneCount {
			t.Fatalf("bad count for %s: got %d expected %d", testCase.str, c , testCase.findibleRune)
		}*/
	}
}
