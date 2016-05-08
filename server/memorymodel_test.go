package server

import (
	"regexp"
	"testing"

	"github.com/gpf-org/gpf/git"
	"github.com/kr/pretty"
)

func TestMemoryModelList(t *testing.T) {
	m := &MemoryModel{
		pattern: regexp.MustCompile("^([^/]+)/.*$"),
		projects: []*git.Project{
			{ID: 1, Name: "Project C"},
			{ID: 2, Name: "Project A"},
			{ID: 3, Name: "Project B"},
		},
		branches: []*git.Branch{
			{Name: "issue-1/remove-pointers", ProjectID: 3},
			{Name: "issue-12/hello-world", ProjectID: 1},
			{Name: "issue-1/abc-123", ProjectID: 1},
			{Name: "issue-30/ola-mundo", ProjectID: 2},
		},
		mergeRequests: []*git.MergeRequest{
			{
				State:           "open",
				SourceBranch:    "issue-12/hello-world",
				TargetBranch:    "develop",
				TargetProjectID: 1,
			},
		},
	}

	expected := []*Feature{
		{
			Name:     "issue-1",
			Commands: []string{"code-review request"},
			Branches: []*FeatureBranch{
				{
					BranchName:  "issue-1/remove-pointers",
					ProjectID:   3,
					ProjectName: "Project B",
				},
				{
					BranchName:  "issue-1/abc-123",
					ProjectID:   1,
					ProjectName: "Project C",
				},
			},
		},
		{
			Name:     "issue-12",
			Commands: []string{},
			Branches: []*FeatureBranch{
				{
					BranchName:  "issue-12/hello-world",
					ProjectID:   1,
					ProjectName: "Project C",
				},
			},
		},
		{
			Name:     "issue-30",
			Commands: []string{"code-review request"},
			Branches: []*FeatureBranch{
				{
					BranchName:  "issue-30/ola-mundo",
					ProjectID:   2,
					ProjectName: "Project A",
				},
			},
		},
	}

	actual := m.List()

	diff := pretty.Diff(expected, actual)
	if len(diff) > 0 {
		t.Error("Expected value differs from actual")
		for i, d := range diff {
			t.Errorf("%d - %s", i, d)
		}
	}
}
