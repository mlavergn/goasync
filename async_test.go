package goasync

import (
	"testing"
)

func TestPack(t *testing.T) {
	task := Async(func() (interface{}, error) {
		return "ok", nil
	})
	result, err := task.Await()
	if result.(string) != "ok" || err != nil {
		t.Fatal("Error while awaiting result", err)
	}
}
