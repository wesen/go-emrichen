package main

import "testing"

func TestEmrichenFilterTag(t *testing.T) {
	tests := []testCase{
		{
			name: "Filter with conditional !If and !Op",
			inputYAML: `!Filter
  test: !If
    test: !Op
      a: !Lookup chance
      op: gt
      b: .5
    then: !Format "This is {person.name}, hello"
    else: No chance
  over:
  - valid
  - hello
  - 0
  - SSEJ
  - false
  - null`,
			expected: "[valid, hello, 0, SSEJ, false, null]", // Adjust this expected output based on your logic
		},
		{
			name: "Filter with various types of elements",
			inputYAML: `!Filter
  test: !If
    test: !Op
      a: !Lookup chance
      op: gt
      b: .5
    then: !Format "This is {person.name}, hello"
    else: No chance
  over:
  - valid
  - hello
  - 0
  - SSEJ
  - false
  - null`,
			expected: "[valid, hello, 0, SSEJ, false, null]", // Adjust this expected output based on your logic
		},
		{
			name: "Filter using !Not and !Var on a dictionary",
			inputYAML: `!Filter
  as: i
  test: !Not,Var i
  over:
    'yes': true
    no: 0
    nope: false
    oui: 1`,
			expected: "{'yes': true, oui: 1}", // Adjust this expected output based on your logic
		},
		{
			name: "Filter with !Op on a list of integers",
			inputYAML: `!Filter
  test: !Op
    a: !Var item
    op: gt
    b: 4
  over: [1, 7, 2, 5]`,
			expected: "[7, 5]", // This should filter out elements greater than 4
		},
	}

	runTests(t, tests)
}