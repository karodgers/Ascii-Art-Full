package handlers

import (
	"strings"
)

func ParseFlag(sl []string) (string, []string, string, error) {
	var align string
	var args []string
	var bannerFile string

	for i, arg := range sl {
		if strings.HasPrefix(arg, "--align=") {
			align = strings.TrimPrefix(arg, "--align=")
			args = append(sl[:i], sl[i+1:]...)
			break
		}
	}

	if len(args) > 0 {
		lastArg := args[len(args)-1]
		if lastArg == "shadow" || lastArg == "standard" || lastArg == "thinkertoy" {
			bannerFile = lastArg
			args = args[:len(args)-1]
		}
	}

	return align, args, bannerFile, nil
}

func Checker(sl []string) bool {
	if len(sl) == 0 {
		return false
	}
	lastElement := sl[len(sl)-1]
	return lastElement == "shadow" || lastElement == "standard" || lastElement == "thinkertoy"
}
