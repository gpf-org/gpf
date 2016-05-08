package git

import (
	"github.com/xanzy/go-gitlab"
)

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
	result := []*MergeRequest{}

	options := &gitlab.ListMergeRequestsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 100,
		},
	}

	nextPage := true

	for i := 1; nextPage; i++ {
		options.Page = i

		mrs, _, err := gp.client.MergeRequests.ListMergeRequests(pid, options)
		if err != nil {
			return nil, err
		}

		count := len(mrs)

		switch {
		case count == 0:
			nextPage = false
		case count > 0:
			for _, value := range mrs {
				result = append(result, mapToMergeRequest(value))
			}
		}
	}

	return result, nil
}

func mapToMergeRequest(data *gitlab.MergeRequest) *MergeRequest {
	return &MergeRequest{
		ID:              data.ID,
		ProjectID:       data.ProjectID,
		Title:           data.Title,
		Description:     data.Description,
		WorkInProgress:  data.WorkInProgress,
		State:           data.State,
		TargetBranch:    data.TargetBranch,
		SourceBranch:    data.SourceBranch,
		Upvotes:         data.Upvotes,
		Downvotes:       data.Downvotes,
		TargetProjectID: data.TargetProjectID,
	}
}
