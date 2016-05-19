package core

import (
	"errors"
	"fmt"

	"github.com/gpf-org/gpf/git"
)

var (
	ErrCodeReviewAlreadyRequested = errors.New("Code review already requested")
)

type Issue struct {
	Name          string         `json:"name"`
	IssueBranches []*IssueBranch `json:"issue_branches"`
	Commands      []int          `json:"commands"`
}

type IssueBranch struct {
	ProjectName string `json:"project_name"`
	BranchName  string `json:"branch_name"`
}

type AffectedBranch struct {
	Branch *git.Branch
	Error  error
}

func ListIssues(store *Store) []*Issue {
	names := store.ListIssues()

	result := make([]*Issue, 0, len(names))

	for _, name := range names {
		issue := &Issue{Name: name}

		branches, _ := store.ListBranches(name)

		issue.IssueBranches = make([]*IssueBranch, 0, len(branches))

		for _, branch := range branches {
			project, _ := store.GetProject(branch.ProjectID)

			issueBranch := &IssueBranch{
				ProjectName: project.Name,
				BranchName:  branch.Name,
			}

			issue.IssueBranches = append(issue.IssueBranches, issueBranch)
		}

		if ok, _, _ := SupportCodeReviewRequest(store, name); ok {
			issue.Commands = append(issue.Commands, CommandCodeReviewRequest)
		}

		result = append(result, issue)
	}

	return result
}

func SupportCodeReviewRequest(store *Store, issue string) (bool, []*git.Branch, error) {
	branches, err := store.ListBranches(issue)
	if err != nil {
		return false, nil, err
	}

	result := []*git.Branch{}

	for _, branch := range branches {
		options := &ListMergeRequestsOptions{
			ProjectID:    branch.ProjectID,
			SourceBranch: branch.Name,
			TargetBranch: "develop",
		}

		mergeRequests := store.ListMergeRequests(options)

		if len(mergeRequests) > 0 {
			continue
		}

		result = append(result, branch)
	}

	support := len(result) > 0

	return support, result, nil
}

func CodeReviewRequest(provider git.GitProvider, store *Store, issue string) ([]*AffectedBranch, error) {
	support, branches, err := SupportCodeReviewRequest(store, issue)

	if err != nil {
		return nil, err
	}

	if !support {
		return nil, ErrCodeReviewAlreadyRequested
	}

	result := make([]*AffectedBranch, 0, len(branches))

	for _, branch := range branches {
		options := &git.CreateMergeRequestOptions{
			Title:        fmt.Sprintf("Issue: [%s] - %s", issue, branch.Name),
			SourceBranch: branch.Name,
			TargetBranch: "develop",
			ProjectID:    branch.ProjectID,
		}

		mergeRequest, err := provider.CreateMergeRequest(options)

		if mergeRequest != nil {
			store.AddMergeRequest(mergeRequest)
		}

		affectedBranch := &AffectedBranch{
			Branch: branch,
			Error:  err,
		}

		result = append(result, affectedBranch)
	}

	return result, nil
}
