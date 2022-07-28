package config

import (
	"encoding/json"
	"testing"
)

func TestIsNotDir(t *testing.T) {
	noDir := ConfigDir{}
	if !IsNotDir(noDir) {
		v, _ := json.Marshal(noDir)
		t.Error("Recieve:\n" + string(v))
	}

	testDir := ConfigDir{Name: "test", Path: "/test/dir"}
	if IsNotDir(testDir) {
		v, _ := json.Marshal(testDir)
		t.Error("Recieve:\n" + string(v))
	}
}
