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

	return mapToMergeRequest(result), nil
}

// ListMergeRequests fetch all the merge requests for a project
func (gp GitLabProvider) ListMergeRequests(pid int) ([]*MergeRequest, error) {
	// TODO: handle paging search
	opts := &gitlab.ListMergeRequestsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 99999999999,
		},
	}
	result, _, err := gp.client.MergeRequests.ListMergeRequests(pid, opts)

	mrs := []*MergeRequest{}

	for i, value := range result {
		mrs[i] = mapToMergeRequest(value)
	}

	return mrs, err
}

func mapToMergeRequest(data *gitlab.MergeRequest) *MergeRequest {
	return &MergeRequest{
		ID:             data.ID,
		ProjectID:      data.ProjectID,
		Title:          data.Title,
		Description:    data.Description,
		WorkInProgress: data.WorkInProgress,
		State:          data.State,
		TargetBranch:   data.TargetBranch,
		SourceBranch:   data.SourceBranch,
		Upvotes:        data.Upvotes,
		Downvotes:      data.Downvotes,
	}
}
