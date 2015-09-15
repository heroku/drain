package drain

import (
	"fmt"
	"regexp"
	"strconv"
)

type LogplexError struct {
	Code  int // L11, L12, etc.
	Count int // Count of logs referenced in the error
	Msg   string
}

func parseLogplexError(msg string) (*LogplexError, error) {
	// Data:Error L10 (output buffer overflow):
	//    491 messages dropped since 2015-09-15T16:22:24+00:00.
	r := regexp.MustCompile(
		`Error L(?P<num>\d+).*\: (?P<count>\d+) .*`).FindAllStringSubmatch(msg, -1)
	if len(r) < 1 || len(r[0]) < 3 {
		return nil, fmt.Errorf("invalid lerror line")
	}
	num, err := strconv.Atoi(r[0][1])
	if err != nil {
		return nil, err
	}
	count, err := strconv.Atoi(r[0][2])
	if err != nil {
		return nil, err
	}
	return &LogplexError{num, count, msg}, nil
}

func (err LogplexError) Error() string {
	return fmt.Sprintf("L%d: %s", err.Code, err.Msg)
}
