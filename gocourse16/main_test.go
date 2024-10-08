package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestProcessContacts(t *testing.T) {
	input :=
`John Doe - +380 (99) 123-45-67
Jane Smith - (067) 987 65 43
Joe Biden - +380///44/1///11/335/1
`

	expectedOutput :=
`Name: John Doe, Phone: (099) 12-34-567
Name: Jane Smith, Phone: (067) 98-76-543
Name: Joe Biden, Phone: (044) 11-13-351
`

	reader := strings.NewReader(input)
	var writer bytes.Buffer

	err := processContacts(reader, &writer)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if writer.String() != expectedOutput {
		t.Errorf("unexpected output:\n got: %v\n want: %v", writer.String(), expectedOutput)
	}
}