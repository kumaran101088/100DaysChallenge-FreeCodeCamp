package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func extractNames() []string {
	names, err := ioutil.ReadFile("names.txt")
	if err != nil {
		fmt.Println("Error while reading file: ", err)
		os.Exit(1)
	}
	return strings.Split(string(names), "\n")
}

func extractContent() string {
	content, err := ioutil.ReadFile("content.txt")
	if err != nil {
		fmt.Println("Error while reading file: ", err)
		os.Exit(1)
	}
	return string(content)
}

func contentByName(content string, name string) {
	fileName := fmt.Sprintf("%v_content.txt", name)
	err := ioutil.WriteFile(fileName, []byte(content), 0777)
	if err != nil {
		fmt.Println("Error while writing file: ", err)
		os.Exit(1)
	}
}

func main() {
	names := extractNames()
	content := extractContent()
	for _, name := range names {
		newContent := strings.ReplaceAll(content, "{{name}}", string(strings.Title(name)))
		contentByName(newContent, string(name))
	}
}
