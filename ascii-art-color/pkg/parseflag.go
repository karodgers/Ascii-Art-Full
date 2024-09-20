package pkg

import (
	"regexp"
	"strings"
)

func ParseFlag(sl []string) (string, string, []string, string, error) {
	var color string
	var substring string
	var args []string
	var bannerFile string

	for i, arg := range sl {
		if strings.HasPrefix(arg, "--color=") {
			color = strings.TrimPrefix(arg, "--color=")
			args = append(sl[:i], sl[i+1:]...)
			break
		}
	}

	if len(args) > 0 {
		// Check if the last argument is a valid banner file
		lastArg := args[len(args)-1]
		if lastArg == "shadow" || lastArg == "standard" || lastArg == "thinkertoy" {
			bannerFile = lastArg
			args = args[:len(args)-1]
			// } else if lastArg != "shadow" && lastArg != "standard" && lastArg != "thinkertoy" {
			// 	return "", "", nil, "", fmt.Errorf("invalid use of available banner files")
		}

		// Check for substring
		if len(args) > 1 {
			substring = args[0]
			args = args[1:]
		}

		// If no substring was found, use the entire remaining argument as the text
		if substring == "" && len(args) > 0 {
			args = []string{strings.Join(args, " ")}
		}
	}

	return color, substring, args, bannerFile, nil
}

func Checker(sl []string) bool {
	lastElement := sl[len(sl)-1]

	switch lastElement {
	case "shadow", "standard", "thinkertoy":
		return true
	default:
		return !regexp.MustCompile(`^(sh.{4}|st.{6}|th.{9})$`).MatchString(lastElement)
	}
}
