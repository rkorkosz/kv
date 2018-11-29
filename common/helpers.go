package common

import (
	"fmt"
	"os"
	"strings"
)

// ProjectName gets project name
func ProjectName() (string, error) {
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	name := pwd[strings.LastIndex(pwd, "/")+1:]
	return name, nil
}

// FormatKV formats key and value to env variable as key=value
func FormatKV(key, value string) string {
	return fmt.Sprintf("%s=%s", key, value)
}
