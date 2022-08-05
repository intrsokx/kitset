package parseip

import (
	"regexp"
	"strings"
)

func ParseIP(s string) []string {
	arys := strings.Split(s, "\n")

	//matchIP, err := regexp.Compile(`.+\..+\..+\..+`)
	matchIP, err := regexp.Compile("((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}")

	if err != nil {
		panic(err)
	}

	ips := make([]string, 0)
	for _, str := range arys {
		if matchIP.MatchString(str) {
			ips = append(ips, str)
		}
	}

	return ips
}
