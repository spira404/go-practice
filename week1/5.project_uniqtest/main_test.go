package main

import (
	"bytes"
	"testing"
)

var testOk = `1
2
3
4
5`

var testOkResult = `1
2
3
4
5
`

var testFail = `1
2
1`

func TestOk(t *testing.T) {
	in := bytes.NewBufferString(testOk)
	out := bytes.NewBuffer(nil)
	err := uniq(in, out)
	if err != nil {
		t.Errorf("Test OK failed: %s", err)
	}
	
	result := out.String()
	if result != testOkResult {
		t.Errorf("Test OK - does not match")
	}

}


func TestForError(t *testing.T) {
	in := bytes.NewBufferString(testFail)
	out := bytes.NewBuffer(nil)
	err := uniq(in, out)
	if err != nil {
		t.Errorf("test for OK failed - error")
	}
}

