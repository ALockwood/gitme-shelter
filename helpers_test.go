package main

import (
	"os"
	"testing"
)

func TestStringisNilOrEmpty(t *testing.T) {
	nes := "not empty"
	es := ""
	ws := "   "

	if stringIsNilOrEmpty(nes) == true {
		t.Error("Failed to catch non-empty string")
	}

	if stringIsNilOrEmpty(es) == false {
		t.Error("Failed to catch empty string")
	}

	if stringIsNilOrEmpty(ws) == false {
		t.Error("Failed to catch whitespace string")
	}
}

func TestGetEnv(t *testing.T) {
	k := "NOT_A_REAL_ENV_VALUE_I_HOPE__ODOYLE_RULES"
	fb := "Sometimes I feel like an idiot. But I am an idiot, so it kinda works out."
	v := "ODoyle rules"

	if getEnv(k, fb) != fb {
		t.Errorf("Failed to get fallback env value. Or, %s actually exists in your env?!", k)
	}

	os.Setenv(k, v)
	if getEnv(k, fb) != v {
		t.Error("Failed to get preset env var")
	}
}