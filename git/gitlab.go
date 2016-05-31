package git

import "github.com/xanzy/go-gitlab"

type GitLabProvider struct {
	client *gitlab.Client
}

func newProviderGitlab(baseURL string, token string) GitLabProvider {
	glc := gitlab.NewClient(nil, token)
	glc.SetBaseURL(baseURL)
	return GitLabProvider{glc}
}

func (gp GitLabProvider) ListAllBranches(pid int) ([]*Branch, error) {
	result, _, err := gp.client.Branches.ListBranches(pid)

	if err != nil {
		return nil, err
	}

	branches := make([]*Branch, len(result))

	for i, value := range result {
		branches[i] = &Branch{Name: value.Name, ProjectID: pid}
	}

	return branches, nil
}

// CreateOrUpdateProjectHook creates a hook to a specified project or update it if already exists.
func (gp GitLabProvider) CreateOrUpdateProjectHook(pid int, hookURL string) (*ProjectHook, error) {
	var result *gitlab.ProjectHook

	listOptions := &gitlab.ListProjectHooksOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 100,
		},
	}

	for i := 1; ; i++ {
		listOptions.Page = i

		hooks, _, err := gp.client.Projects.ListProjectHooks(pid, listOptions)
		if err != nil {
			return nil, err
		}

		if len(hooks) == 0 {
			break
		}

		for _, hook := range hooks {
			if hook.URL != hookURL {
				continue
			}

			result = hook
		}
	}

	if result == nil {
		addOptions := &gitlab.AddProjectHookOptions{
			URL:                 hookURL,
			PushEvents:          true,
			MergeRequestsEvents: true,
		}

		var err error
		result, _, err = gp.client.Projects.AddProjectHook(pid, addOptions)
		if err != nil {
			return nil, err
		}
	}

	return mapToProjectHook(result), nil
}

func mapToProjectHook(data *gitlab.ProjectHook) *ProjectHook {
	return &ProjectHook{
		ID:                  data.ID,
		URL:                 data.URL,
		ProjectID:           data.ProjectID,
		PushEvents:          data.PushEvents,
		IssuesEvents:        data.IssuesEvents,
		MergeRequestsEvents: data.MergeRequestsEvents,
	}
}

func mapToProjectHooks(data []*gitlab.ProjectHook) []*ProjectHook {
	hooks := make([]*ProjectHook, len(data))

	for i, value := range data {
		hooks[i] = mapToProjectHook(value)
	}

	return hooks
}

// CreateMergeRequest creates a new merge request for a branch.
func (gp GitLabProvider) CreateMergeRequest(opts *CreateMergeRequestOptions) (*MergeRequest, error) {
	mropts := &gitlab.CreateMergeRequestOptions{
		Title:           opts.Title,
		SourceBranch:    opts.SourceBranch,
		TargetBranch:    opts.TargetBranch,
		TargetProjectID: opts.ProjectID,
	}

	result, _, err := gp.client.MergeRequests.CreateMergeRequest(opts.ProjectID, mropts)
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
		ID:           data.ID,
		ProjectID:    data.ProjectID,
		TargetBranch: data.TargetBranch,
		SourceBranch: data.SourceBranch,
	}
}

// ListAllProjects gets a list of all Git projects.
func (gp GitLabProvider) ListAllProjects() ([]*Project, error) {
	user, _, _ := gp.client.Users.CurrentUser()

	result := []*Project{}

	options := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 100,
		},
	}

	nextPage := true

	for i := 1; nextPage; i++ {
		options.Page = i

		projects, _, err := gp.listProjects(user.IsAdmin, options)
		if err != nil {
			return nil, err
		}

		count := len(projects)

		switch true {
		case count == 0:
			nextPage = false
		case count > 0:
			for _, value := range projects {
				if value.DefaultBranch == nil {
					continue
				}
				result = append(result, &Project{ID: *value.ID, Name: *value.Name})
			}
		}
	}

	return result, nil
}

func (gp GitLabProvider) listProjects(isAdmin bool, options *gitlab.ListProjectsOptions) ([]*gitlab.Project, *gitlab.Response, error) {
	if isAdmin {
		return gp.client.Projects.ListAllProjects(options)
	} else {
		return gp.client.Projects.ListProjects(options)
	}
}
