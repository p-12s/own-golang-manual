package reverse

import "testing"

func TestReverse(t *testing.T) {
	testTable := []struct {
		name, data, result string
	}{
		{
			name:   "hello OK",
			data:   "Hello, world!",
			result: "!dlrow ,olleH",
		},
		{
			name:   "empty OK",
			data:   "",
			result: "",
		},
		{
			name:   "numbers str OK",
			data:   "12345 32",
			result: "23 54321",
		},
		{
			name:   "symbols OK",
			data:   "!@#$%^&*()",
			result: ")(*&^%$#@!",
		},
		{
			name:   "escaped '\\n' OK",
			data:   "\\n",
			result: "n\\",
		},
		{
			name:   "hieroglyphs OK",
			data:   "也施手",
			result: "手施也",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			reverseResult := Reverse(testCase.data)
			if reverseResult != testCase.result {
				t.Fatalf("Error: output: '%s' dosn't equal expected result: '%s'\n",
					reverseResult, testCase.result)
			}
		})
	}
}
