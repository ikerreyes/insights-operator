package anonymize

import (
	"regexp"
	"strings"
)

func AnonymizeURLCSV(s string) string {
	strs := strings.Split(s, ",")
	outSlice := AnonymizeURLSlice(strs)
	return strings.Join(outSlice, ",")
}

func AnonymizeURLSlice(in []string) []string {
	var outSlice []string
	for _, str := range in {
		outSlice = append(outSlice, AnonymizeURL(str))
	}
	return outSlice
}

var reURL = regexp.MustCompile(`[^\.\-/\:]`)

func AnonymizeURL(s string) string { return reURL.ReplaceAllString(s, "x") }
