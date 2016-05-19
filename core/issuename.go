package core

import (
	"regexp"
)

type IssueNameExtractor interface {
	IssueNameExtract(value string) (string, bool)
}

type IssueNameRegexp struct {
	regexp *regexp.Regexp
}

func (i *IssueNameRegexp) IssueNameExtract(value string) (string, bool) {
	groups := i.regexp.FindAllStringSubmatch(value, -1)

	if len(groups) == 0 {
		return "", false
	}

	return groups[0][1], true
}

func NewIssueNameRegexp(pattern string) *IssueNameRegexp {
	return &IssueNameRegexp{
		regexp: regexp.MustCompile(pattern),
	}
}
