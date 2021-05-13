package unpack

import "testing"

func TestUnpackStr(t *testing.T) {
	input := "abc"
	output := "abc"
	if result := UnpackStr(input); result != output {
		t.Fatalf("bad unpack for %s: got %s expected %s", input, result, output)
	}

	input = "ab2c3d"
	output = "abbcccd"
	if result := UnpackStr(input); result != output {
		t.Fatalf("bad unpack for %s: got %s expected %s", input, result, output)
	}

	input = "ab2c3d4"
	output = "abbcccdddd"
	if result := UnpackStr(input); result != output {
		t.Fatalf("bad unpack for %s: got %s expected %s", input, result, output)
	}

	input = "abc3d2"
	output = "abcccdd"
	if result := UnpackStr(input); result != output {
		t.Fatalf("bad unpack for %s: got %s expected %s", input, result, output)
	}

	input = "a"
	output = "a"
	if result := UnpackStr(input); result != output {
		t.Fatalf("bad unpack for %s: got %s expected %s", input, result, output)
	}

	/*input = "5a"
	output = "a"
	result := UnpackStr(input);
	if !t.Failed() {
		t.Fatalf("bad unpack for %s: got %s expected %s", input, result, output)
	}*/
}