package core

import (
	"reflect"
	"testing"

	"github.com/gpf-org/gpf/git"
)

func TestStoreReset(t *testing.T) {
	t.Errorf("Not implemented")
}

func TestStoreAddBranchInvalidBranchName(t *testing.T) {
	store := NewStore(NewIssueNameRegexp("^issue-([^/]+)/.*$"))

	store.AddProject(&git.Project{ID: 1, Name: "Project A"})

	branch := &git.Branch{Name: "invalid/name"}

	if err := store.AddBranch(branch); err != ErrInvalidBranchName {
		t.Errorf("Expected error [%v], got [%v]", ErrInvalidBranchName, err)
	}
}

func TestStoreAddBranch(t *testing.T) {
	store := NewStore(NewIssueNameRegexp("^issue-([^/]+)/.*$"))

	branch := &git.Branch{
		Name:      "issue-123/fix-label",
		ParentIDs: []string{"1", "2"},
		ProjectID: 1,
	}

	if err := store.AddBranch(branch); err != nil {
		t.Errorf("Unexpected error [%v]", err)
	}
}

func TestStoreGetProjectNotFound(t *testing.T) {
	store := NewStore(nil)

	store.AddProject(&git.Project{
		ID:   1,
		Name: "my-project-site-api",
	})

	if _, err := store.GetProject(2); err != ErrProjectNotFound {
		t.Errorf("Expected error [%v], got [%v]", ErrProjectNotFound, err)
	}
}

func TestStoreGetProject(t *testing.T) {
	store := NewStore(nil)

	store.AddProject(&git.Project{
		ID:   1,
		Name: "my-project-site-api",
	})

	actual, err := store.GetProject(1)
	if err != nil {
		t.Errorf("Unexpected error [%v]", err)
	}

	expected := &git.Project{
		ID:   1,
		Name: "my-project-site-api",
	}

	if *actual != *expected {
		t.Errorf("Expected [%v] to be equal [%v]", actual, expected)
	}
}

func TestStoreListIssues(t *testing.T) {
	store := NewStore(NewIssueNameRegexp("^issue-([^/]+)/.*$"))

	store.AddProject(&git.Project{
		ID:   1,
		Name: "my-project-site-api",
	})
	store.AddBranch(&git.Branch{
		Name:      "issue-branding/fix-route",
		ProjectID: 1,
	})

	store.AddProject(&git.Project{
		ID:   2,
		Name: "my-project-site",
	})
	store.AddBranch(&git.Branch{
		Name:      "issue-abc/change-labels",
		ProjectID: 2,
	})

	actual := store.ListIssues()

	expected := []string{"abc", "branding"}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected [%v], got [%v]", expected, actual)
	}
}

func TestStoreListBranches(t *testing.T) {
	store := NewStore(NewIssueNameRegexp("^issue-([^/]+)/.*$"))

	store.AddBranch(&git.Branch{
		Name: "issue-3/replace-labels",
	})

	if _, err := store.ListBranches("2"); err != ErrIssueNotFound {
		t.Errorf("Expected error [%v], got [%v]", ErrIssueNotFound, err)
	}

	branchA := &git.Branch{
		Name: "issue-2/remove-nested-loop",
	}

	branchB := &git.Branch{
		Name: "issue-2/replace-if-switch",
	}

	store.AddBranch(branchA)
	store.AddBranch(branchB)

	actual, err := store.ListBranches("2")
	if err != nil {
		t.Errorf("Unexpected error [%v]", err)
	}

	expected := []*git.Branch{
		branchA,
		branchB,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected [%v] to be equal [%v]", actual, expected)
	}
}

func TestStoreListMergeRequests(t *testing.T) {
	store := NewStore(NewIssueNameRegexp("^issue-([^/]+)/.*$"))

	mrA := &git.MergeRequest{
		ProjectID:    1,
		SourceBranch: "a",
		TargetBranch: "b",
	}

	mrB := &git.MergeRequest{
		ProjectID:    1,
		SourceBranch: "a",
		TargetBranch: "c",
	}

	mrC := &git.MergeRequest{
		ProjectID:    2,
		SourceBranch: "a",
		TargetBranch: "b",
	}

	mrD := &git.MergeRequest{
		ProjectID:    2,
		SourceBranch: "a",
		TargetBranch: "c",
	}

	store.AddMergeRequest(mrA)
	store.AddMergeRequest(mrB)
	store.AddMergeRequest(mrC)
	store.AddMergeRequest(mrD)

	// filter by ProjectID

	actual := store.ListMergeRequests(&ListMergeRequestsOptions{
		ProjectID: 1,
	})

	expected := []*git.MergeRequest{
		mrA,
		mrB,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected [%v] to be qual [%v]", actual, expected)
	}

	// filter by SourceBranch and TargetBranch

	actual = store.ListMergeRequests(&ListMergeRequestsOptions{
		SourceBranch: "a",
		TargetBranch: "c",
	})

	expected = []*git.MergeRequest{
		mrB,
		mrD,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected [%v] to be equal [%v]", actual, expected)
	}

	// filter by ProjectID and TargetBranch

	actual = store.ListMergeRequests(&ListMergeRequestsOptions{
		ProjectID:    1,
		TargetBranch: "c",
	})

	expected = []*git.MergeRequest{
		mrB,
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected [%v] to be equal [%v]", actual, expected)
	}
}
