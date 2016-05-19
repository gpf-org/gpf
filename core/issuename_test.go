package core

import (
	"testing"
)

func TestIssueNameRegexp(t *testing.T) {
	extractor := NewIssueNameRegexp("^issue-([^/]+)/.*$")

	invalidBranchName := "invalid/abc"

	if _, ok := extractor.IssueNameExtract(invalidBranchName); ok {
		t.Errorf("Expected [%s] to be invalid, got valid", invalidBranchName)
	}

	branchName := "issue-1234/fix-path"

	name, ok := extractor.IssueNameExtract(branchName)
	if !ok {
		t.Errorf("Expected [%s] to be valid, got invalid", branchName)
	}

	expected := "1234"

	if name != "1234" {
		t.Errorf("Expected [%s], got [%s]", expected, name)
	}
}
