package merge

import "regexp"

var descriptionRegex = regexp.MustCompile("# Telegram Description\n(.*)\n# End of Telegram Description")

func sanitizeDescription(description string) string {
	ms := descriptionRegex.FindStringSubmatch(description)
	if len(ms) > 1 {
		return ms[1]
	}
	return ""
}
