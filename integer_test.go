package config

import (
	"testing"
)

func Test_Integer(t *testing.T) {

	min := 20
	max := 200
	fallback := 100

	one := GetInt("int1", min, max, fallback)
	if one != fallback {
		t.Fatalf("failure to get int1")
	}

	// to small (negative)
	one = GetInt("int2", min, max, fallback)
	if one != 123 {
		t.Fatalf("failure to get int2")
	}

	// too big
	one = GetInt("int3", min, max, fallback)
	if one != fallback {
		t.Fatalf("failure to get int3")
	}

	// empty
	one = GetInt("int4", min, max, fallback)
	if one != fallback {
		t.Fatalf("failure to get int4")
	}

	// completely missing
	one = GetInt("int5", min, max, fallback)
	if one != fallback {
		t.Fatalf("failure to get int5")
	}

	one = GetInt("int6", min, max, fallback)
	if one != fallback {
		t.Fatalf("failure to get int6")
	}
}
