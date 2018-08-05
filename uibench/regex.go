package uibench

import (
	"regexp"
)

func substringBetween(line string, prefix string, suffix string) string {
	re := regexp.MustCompile(regexp.QuoteMeta(prefix) + "(.*)" + regexp.QuoteMeta(suffix))
	found := re.FindSubmatch([]byte(line))

	if len(found) > 0 {
		return string(found[1])
	} else {
		return ""
	}
}
