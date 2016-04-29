package git

import "github.com/xanzy/go-gitlab"

// CreateMergeRequest creates a new merge request for a branch.
func (gp GitLabProvider) CreateMergeRequest(pid int, opts CreateMergeRequestOptions) (*MergeRequest, error) {
	mropts := &gitlab.CreateMergeRequestOptions{
		Title:           opts.Title,
		Description:     opts.Description,
		SourceBranch:    opts.SourceBranch,
		TargetBranch:    opts.TargetBranch,
		AssigneeID:      opts.AssigneeID,
		TargetProjectID: opts.TargetProjectID,
	}

	result, _, err := gp.client.MergeRequests.CreateMergeRequest(pid, mropts)
	if err != nil {
		return nil, err
	}

	mr := &MergeRequest{
		ID:             result.ID,
		ProjectID:      result.ProjectID,
		Title:          result.Title,
		Description:    result.Description,
		WorkInProgress: result.WorkInProgress,
		State:          result.State,
		TargetBranch:   result.TargetBranch,
		SourceBranch:   result.SourceBranch,
		Upvotes:        result.Upvotes,
		Downvotes:      result.Downvotes,
	}

	return mr, nil
}
