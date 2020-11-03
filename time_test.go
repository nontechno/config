package config

import (
	"testing"
	"time"
)

func Test_Time(t *testing.T) {

	min := time.Duration(time.Second * 0)
	max := time.Duration(time.Second * 10)
	fallback := time.Duration(time.Second * 5)

	one := GetDuration("duration1", min, max, fallback)
	if one != time.Duration(time.Millisecond*300) {
		t.Fatalf("failure to get duration1")
	}

	// to small (negative)
	one = GetDuration("duration2", min, max, fallback)
	if one != fallback {
		t.Fatalf("failure to get duration2")
	}

	// too big
	one = GetDuration("duration3", min, max, fallback)
	if one != fallback {
		t.Fatalf("failure to get duration3")
	}

	// empty
	one = GetDuration("duration4", min, max, fallback)
	if one != fallback {
		t.Fatalf("failure to get duration4")
	}

	// completely missing
	one = GetDuration("duration5", min, max, fallback)
	if one != fallback {
		t.Fatalf("failure to get duration5")
	}

}
