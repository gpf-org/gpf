package core

import (
	"reflect"
	"testing"

	"github.com/gpf-org/gpf/git"
)

func TestListIssues(t *testing.T) {
	store := NewStore(NewIssueNameRegexp("^issue-([^/]+)/.*$"))

	store.AddProject(&git.Project{
		ID:   1,
		Name: "my-project-site-api",
	})
	store.AddBranch(&git.Branch{
		Name:      "issue-abc/replace-put-post",
		ProjectID: 1,
	})

	store.AddProject(&git.Project{
		ID:   2,
		Name: "my-project-site",
	})
	store.AddBranch(&git.Branch{
		Name:      "issue-abc/modify-consumer-api",
		ProjectID: 2,
	})

	store.AddMergeRequest(&git.MergeRequest{
		ProjectID:    2,
		SourceBranch: "issue-abc/modify-consumer-api",
		TargetBranch: "develop",
	})

	actual := ListIssues(store)

	expected := []*Issue{
		{
			Name: "abc",
			IssueBranches: []*IssueBranch{
				{
					ProjectName: "my-project-site-api",
					BranchName:  "issue-abc/replace-put-post",
				},
				{
					ProjectName: "my-project-site",
					BranchName:  "issue-abc/modify-consumer-api",
				},
			},
			Commands: []int{CommandCodeReviewRequest},
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected [%v] to be equal [%v]", actual, expected)
	}
}

func TestSupportCodeReviewRequestIssueNotFound(t *testing.T) {
	store := NewStore(nil)

	_, _, err := SupportCodeReviewRequest(store, "issue-1")
	if err != ErrIssueNotFound {
		t.Fatalf("Expected error [%v], got [%v]", ErrIssueNotFound, err)
	}
}

func TestSupportCodeReviewRequest(t *testing.T) {
	store := NewStore(NewIssueNameRegexp("^issue-([^/]+)/.*$"))

	store.AddBranch(&git.Branch{
		Name:      "issue-2/change-label",
		ProjectID: 1,
	})

	store.AddBranch(&git.Branch{
		Name:      "issue-2/fix-something",
		ProjectID: 1,
	})

	store.AddMergeRequest(&git.MergeRequest{
		ProjectID:    1,
		SourceBranch: "issue-2/fix-something",
		TargetBranch: "develop",
	})

	support, actual, err := SupportCodeReviewRequest(store, "2")
	if err != nil {
		t.Errorf("Expected error [%v]", err)
	}

	if !support {
		t.Errorf("Expected to support code review request")
	}

	expected := []*git.Branch{
		{
			Name:      "issue-2/change-label",
			ProjectID: 1,
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected [%v] to be equal [%v]", actual, expected)
	}
}

func TestCodeReviewRequestAlreadyRequested(t *testing.T) {
	store := NewStore(NewIssueNameRegexp("^issue-([^/]+)/.*$"))

	store.AddBranch(&git.Branch{
		Name:      "issue-2/fix-something",
		ProjectID: 1,
	})

	store.AddMergeRequest(&git.MergeRequest{
		ProjectID:    1,
		SourceBranch: "issue-2/fix-something",
		TargetBranch: "develop",
	})

	_, err := CodeReviewRequest(nil, store, "2")
	if err != ErrCodeReviewAlreadyRequested {
		t.Fatalf("Excepted error [%v], got [%v]", ErrCodeReviewAlreadyRequested, err)
	}
}

func TestCodeReviewRequest(t *testing.T) {
	provider := &FakeGitProvider{}

	store := NewStore(NewIssueNameRegexp("^issue-([^/]+)/.*$"))

	store.AddBranch(&git.Branch{
		Name:      "issue-3/fix-something",
		ProjectID: 1,
	})

	store.AddBranch(&git.Branch{
		Name:      "issue-3/do-something-else",
		ProjectID: 2,
	})

	store.AddMergeRequest(&git.MergeRequest{
		ProjectID:    2,
		SourceBranch: "issue-3/do-something-else",
		TargetBranch: "develop",
	})

	affectedBranches, err := CodeReviewRequest(provider, store, "3")
	if err != nil {
		t.Fatalf("Unexpected error [%v]", err)
	}

	if len(affectedBranches) != 1 {
		t.Fatalf("Expected affected branches [%d], got [%d]", 1, len(affectedBranches))
	}

	expected := []*AffectedBranch{
		{
			Branch: &git.Branch{
				Name:      "issue-3/fix-something",
				ProjectID: 1,
			},
			Error: nil,
		},
	}

	if !reflect.DeepEqual(affectedBranches, expected) {
		t.Errorf("Expected [%v] to be equal [%v]", affectedBranches, expected)
	}
}

type FakeGitProvider struct{}

func (f *FakeGitProvider) ListAllProjects() ([]*git.Project, error)       { return nil, nil }
func (f *FakeGitProvider) ListAllBranches(pid int) ([]*git.Branch, error) { return nil, nil }
func (f *FakeGitProvider) CreateOrUpdateProjectHook(pid int, hookURL string) (*git.ProjectHook, error) {
	return nil, nil
}
func (f *FakeGitProvider) CreateMergeRequest(opts *git.CreateMergeRequestOptions) (*git.MergeRequest, error) {
	return nil, nil
}
func (f *FakeGitProvider) ListMergeRequests(pid int) ([]*git.MergeRequest, error) { return nil, nil }
