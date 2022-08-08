package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path/filepath"
	"testing"
)

func TestUniqNoArgs(t *testing.T) {

	tmpdir := t.TempDir()
	file_path := "input.txt"
	full_path := filepath.Join(tmpdir, file_path)
	text := []byte("I love music.\nI love music.\nI love music.\n\nI love music of Kartik.\n" +
		"I love music of Kartik.\nThanks.\nI love music of Kartik.\nI love music of Kartik.")

	_, err := os.CreateTemp(tmpdir, file_path)

	if err != nil {
		log.Fatal(err)
	}

	//if _, err := f.Write(text); err != nil {
	//	log.Fatal(err)
	//}

	err = os.WriteFile(full_path, text, 0644)
	if err != nil {
		log.Fatal(err)
	}

	reader, err := os.Open(filepath.Join(tmpdir, file_path))
	if err != nil {
		log.Fatal(err)
	}

	result, err := UniqNoArgs(reader)
	if err != nil {
		log.Fatal(err)
	}

	//result, err := exec.Command("cat input.txt | go run uniq.go").Output()
	//if err != nil {
	//	log.Fatal(err)
	//}

	//expected := "I love music.\n\nI love music of Kartik.\nThanks.\nI love music of Kartik.\n"
	expected := []string{"I love music.\n", "\n", "I love music of Kartik.\n", "Thanks.\n", "I love music of Kartik.\n"}

	assert.Equal(t, expected, result, fmt.Sprintf("Incorrect result. Expected:\n---start---%s---end---\n"+
		"Got:\n---start---%s---end---", expected, result))

}
