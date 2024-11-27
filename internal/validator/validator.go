package validator

import (
	"regexp"
	"strings"
	"unicode/utf8"
)

// regex taken from https://html.spec.whatwg.org/multipage/input.html#valid-e-mail-address which is recommended by the W3C
var EmailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

type Validator struct {
	problems map[string]string
}

func (v *Validator) Valid() bool {
	return len(v.problems) == 0
}

func (v *Validator) AddProblem(key, msg string) {
	if v.problems == nil {
		v.problems = make(map[string]string)
	}
	if _, exists := v.problems[key]; !exists {
		v.problems[key] = msg
	}

}

func (v *Validator) Check(ok bool, key, msg string) {
	if !ok {
		v.AddProblem(key, msg)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

// return true if a value contains no more than `count` characters.
func MaxChars(value string, count int) bool {
	return utf8.RuneCountInString(value) <= count
}

func MatchRegex(value string, regex *regexp.Regexp) bool {
	return regex.MatchString(value)
}
