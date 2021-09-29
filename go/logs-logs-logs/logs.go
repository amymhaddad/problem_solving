package logs

import (
	"regexp"
	"strings"
)

// Message extracts the message from the provided log line.
func Message(line string) string {
	re := regexp.MustCompile(`\[[A-Z]+\]:\s*`)
	result := re.FindStringSubmatch(line)

	line = strings.ReplaceAll(line, result[0], "")
	return strings.TrimSpace(line)
}

// MessageLen counts the amount of characters (runes) in the message of the log line.
func MessageLen(line string) int {
	panic("Please implement the MessageLen() function")
}

// LogLevel extracts the log level string from the provided log line.
func LogLevel(line string) string {
	panic("Please implement the LogLevel() function")
}

// Reformat reformats the log line in the format `message (logLevel)`.
func Reformat(line string) string {
	panic("Please implement the Reformat() function")
}
