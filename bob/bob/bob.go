package bob

import (
	"strings"
)

// Hey is a function to get the remark to Bob and returns Bob's response
func Hey(remark string) string {

	var resp string

	// Check the remark
	question := strings.Contains(remark, "?")
	// Shouting
	remarkUpperCase := strings.ToUpper(remark)

	if remarkUpperCase != remark && question == true {
		resp = "Sure."
	} else if remarkUpperCase == remark && remark != "" && !question {
		resp = "Whoa, chill out!"
	} else if remarkUpperCase == remark && question == true {
		resp = "Calm down, I know what I'm doing!"
	} else if remark == "" {
		resp = "Fine. Be that way!"
	} else {
		resp = "Whatever."
	}

	return resp
}
