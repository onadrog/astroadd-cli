package helpers

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Get the os platform separator
var sep = os.PathSeparator
var basePath = fmt.Sprintf("src%cpages%c", sep, sep)

// Return the title without dashes and the directory path for the given path.
// Or an error
func Sanitize(path string) (title string, err error) {
	dir, title := filepath.Split(path)
	ext := filepath.Ext(title)
	title = strings.ReplaceAll(title, ext, "")

	err = os.MkdirAll(fmt.Sprintf("%s%s", basePath,dir), 0750)
	if err != nil && !os.IsExist(err) {
		return "", err
	}

	re := regexp.MustCompile(`_|-|/|\\|\W`)
	title = re.ReplaceAllString(title, " ")
    title = strings.ToUpper(title[:1]) + title[1:]
	return title, nil
}

func WriteContent(path string) {

	title, err := Sanitize(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	date := time.Now().Format(time.RFC3339)
	content := []byte(
		fmt.Sprintf(`---
title: "%s"
pubDate: %s
description: ""
tags: []
draft: true
---`,
			title, date))

	err = os.WriteFile(fmt.Sprintf("%s%s", basePath, path), content, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	out := fmt.Sprintf("File successfully created in %s%s",basePath, path)
	fmt.Println(out)
}
