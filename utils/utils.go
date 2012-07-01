package utils

import (
	"regexp"
	"fmt"
)

func NamedRegexpGroup(str string, reg *regexp.Regexp) (ng map[string]string, matched bool) {
	rst := reg.FindStringSubmatch(str)
	fmt.Printf("%s => %s => %s\n\n", reg, str, rst)
	if len(rst) < 1 {
		return
	}
	ng = make(map[string]string)
	lenRst := len(rst)
	sn := reg.SubexpNames()
	for k, v := range sn {
		// SubexpNames contain the none named group,
		// so must filter v == ""
		if k == 0 || v == "" {
			continue
		}
		if k+1 > lenRst {
			break
		}
		ng[v] = rst[k]
	}
	matched = true
	return
}
