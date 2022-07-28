package array

import (
	"encoding/json"
	"testing"

	"github.com/kimulaco/cde/pkg/dir"
)

func TestFilter(t *testing.T) {
	dirs := []dir.Dir{
		{Name: "test-1", Path: "/test-1"},
		{Name: "test-2", Path: "/test-2"},
		{Name: "test-3", Path: "/test-3"},
		{Name: "test-4", Path: "/test-4"},
	}
	result := Filter(dirs, func(dir dir.Dir) bool {
		return dir.Name != "test-2"
	})
	if len(result) != 3 ||
		result[0].Name != "test-1" ||
		result[1].Name != "test-3" {
		v, _ := json.Marshal(result)
		t.Error("Recieve:\n" + string(v))
	}
}
