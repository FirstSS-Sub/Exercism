// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	remark = strings.TrimSpace(remark)
	var lastWord string
	if len(remark) == 0 {
		lastWord = ""
	} else {
		lastWord = remark[len(remark)-1:]
	}

	if lastWord == "?" {
		if strings.ToUpper(remark) == remark && IsLetter(remark) {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Sure."
		}
	} else if strings.ToUpper(remark) == remark && IsLetter(remark) {
		return "Whoa, chill out!"
	} else if remark == "" {
		return "Fine. Be that way!"
	} else {
		return "Whatever."
	}
}

func IsLetter(s string) bool {
	for _, r := range s {
		if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') {
			return true
		}
	}
	return false
}
