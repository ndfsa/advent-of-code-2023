package util

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

func ReadFile(filePath string) (string, error) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return string(file), nil
}

func ReadFileSplit(filePath string) ([]string, error) {
	input, err := ReadFile(filePath)

	if err != nil {
		return nil, err
	}

	lines := strings.SplitN(input, "\n", -1)

	return lines, nil
}

func RunSolution[T comparable](
	t *testing.T,
	solution func(string) (T, error),
	filePath string,
	result T) {

	if _, err := os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		t.Skip(err)
	}

	res, err := solution(filePath)

	if err != nil {
		t.Fatal(err)
	}

	if res != result {
		t.Fatal(fmt.Sprintf("incorrect: %v", res))
	}
	t.Logf("res: %v", res)
}
