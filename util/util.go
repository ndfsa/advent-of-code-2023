package util

import (
	"errors"
	"bufio"
	"fmt"
	"os"
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
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
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

func IsDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

