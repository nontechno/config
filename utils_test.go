package config

import (
	"fmt"
	"os"
	"testing"
)

func Test_Expand(t *testing.T) {

	source := `
one: two
two: ${FAKE}
three: ${MISSING}
`
	expectedRetain := source

	expectedNonRetain := `
one: two
two: 
three: 
`
	expectedRetain2 := `
one: two
two: some.value.here
three: ${MISSING}
`

	retain := expand(source, true)
	if retain != expectedRetain {
		fmt.Println(retain)
		t.Fatalf("failure to retain missing values")
	}

	nonretain := expand(source, false)
	if nonretain != expectedNonRetain {
		fmt.Println(nonretain)
		t.Fatalf("failure to non-retain missing values")
	}

	if err := os.Setenv("FAKE", "some.value.here"); err != nil {
		t.Fatalf("failure to set env.var, (%v)", err)
	}

	retain2 := expand(source, true)
	if retain2 != expectedRetain2 {
		fmt.Println(retain2)
		t.Fatalf("failure to retain missing values")
	}

}
