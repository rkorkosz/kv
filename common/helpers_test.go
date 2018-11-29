package common

import "testing"

func TestProjectName(t *testing.T) {
	expected := "common"
	actual, err := ProjectName()
	if err != nil {
		t.Fatal(err)
	}
	if actual != expected {
		t.Errorf("Project name incorrect, got: %s, want: %s", actual, expected)
	}
}

func TestFormatKV(t *testing.T) {
	key := "test"
	value := "val123"
	expected := "test=val123"
	actual := FormatKV(key, value)
	if actual != expected {
		t.Errorf("Formatted string incorrect, got: %s, want: %s", actual, expected)
	}
}
