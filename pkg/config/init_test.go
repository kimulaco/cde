package config

import (
	"encoding/json"
	"testing"
)

func TestInit(t *testing.T) {
	c, initErr := Init()
	if initErr != nil {
		t.Error("Error: " + initErr.Error())
	}
	if c.UpdatedAt == "" {
		v, _ := json.Marshal(c)
		t.Error("Recieve:\n" + string(v))
	}
}
