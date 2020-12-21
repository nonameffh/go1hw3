package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	Browsers []string `json:"browsers"`
}

type Collector struct {
	currentBrowser struct {
		isAndroid bool
		isMSIE    bool
	}
	currentUser struct {
		hasAndroid bool
		hasMSIE    bool
		mail       string
	}
}

// вам надо написать более быструю оптимальную этой функции
func FastSearch(out io.Writer) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	user := &User{}
	line := -1
	browsers := make(map[string]bool, 64)

	collector := Collector{
		currentBrowser: struct {
			isAndroid bool
			isMSIE    bool
		}{},
		currentUser: struct {
			hasAndroid bool
			hasMSIE    bool
			mail       string
		}{},
	}

	fmt.Fprintln(out, "found users:")

	for scanner.Scan() {
		line++

		text := scanner.Bytes()
		err := user.UnmarshalJSON(text)
		if err != nil {
			continue
		}

		collector.currentUser.hasAndroid = false
		collector.currentUser.hasMSIE = false

		for _, browser := range user.Browsers {
			collector.currentBrowser.isAndroid = strings.Contains(browser, "Android")
			collector.currentBrowser.isMSIE = strings.Contains(browser, "MSIE")
			collector.currentUser.hasAndroid = collector.currentUser.hasAndroid || collector.currentBrowser.isAndroid
			collector.currentUser.hasMSIE = collector.currentUser.hasMSIE || collector.currentBrowser.isMSIE

			if !collector.currentBrowser.isAndroid && !collector.currentBrowser.isMSIE {
				continue
			}

			if _, seen := browsers[browser]; !seen {
				browsers[browser] = true
			}
		}

		if !(collector.currentUser.hasAndroid && collector.currentUser.hasMSIE) {
			continue
		}

		collector.currentUser.mail = strings.ReplaceAll(user.Email, "@", " [at] ")
		fmt.Fprintln(out, "["+strconv.Itoa(line)+"] "+user.Name+" <"+collector.currentUser.mail+">")
	}
	fmt.Fprintln(out)
	fmt.Fprintln(out, "Total unique browsers", len(browsers))
}
