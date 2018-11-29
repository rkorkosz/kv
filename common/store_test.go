package common

import (
	"os"
	"testing"
)

const dbPath = "/tmp/kv.db"

func TestBoltStoreSetGet(t *testing.T) {
	store := newBoltStore(dbPath)
	defer func() {
		store.Close()
		os.Remove(dbPath)
	}()
	err := store.Set([2]string{"key", "value"})
	if err != nil {
		t.Errorf("Set failed: %s", err.Error())
	}
	env, err := store.Get("key")
	if err != nil {
		t.Errorf("Get failed: %s", err.Error())
	}
	expected := [2]string{"key", "value"}
	if env != expected {
		t.Errorf("Env incorrect, got: %s, want: %s", env, expected)
	}
}

func TestBoltSetGetMany(t *testing.T) {
	store := newBoltStore(dbPath)
	defer func() {
		store.Close()
		os.Remove(dbPath)
	}()
	envs := [][2]string{
		{"key1", "value1"},
		{"key2", "value2"},
		{"key3", "value3"},
	}
	err := store.SetMany(envs)
	if err != nil {
		t.Errorf("SetMany failed: %s", err.Error())
	}
	actual, err := store.GetMany()
	if err != nil {
		t.Errorf("GetMany failed: %s", err.Error())
	}
	for i, env := range envs {
		if env != actual[i] {
			t.Errorf("Env incorrect: got: %s, want: %s", env, actual[i])
		}
	}
}
