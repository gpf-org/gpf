package git

import (
	"github.com/xanzy/go-gitlab"
)

// CreateOrUpdateProjectHook creates a hook to a specified project or update it if already exists.
func (gp GitLabProvider) CreateOrUpdateProjectHook(pid int, hookURL string) (*ProjectHook, error) {
	// TODO: handle paging search
	optsList := &gitlab.ListProjectHooksOptions{}
	result, _, err := gp.client.Projects.ListProjectHooks(pid, optsList)
	if err != nil {
		return nil, err
	}

	hooks := mapToProjectHooks(result)

	for i := range hooks {
		if hooks[i].URL == hookURL {
			// ensure it has the right configurations
			optsEdit := &gitlab.EditProjectHookOptions{
				URL:                 hookURL,
				PushEvents:          true,
				IssuesEvents:        true,
				MergeRequestsEvents: true,
				TagPushEvents:       true,
			}

			hook, _, err := gp.client.Projects.EditProjectHook(pid, hooks[i].ID, optsEdit)

			return mapToProjectHook(hook), err
		}
	}

	optsAdd := &gitlab.AddProjectHookOptions{
		URL:                 hookURL,
		PushEvents:          true,
		IssuesEvents:        true,
		MergeRequestsEvents: true,
		TagPushEvents:       true,
	}

	hook, _, err := gp.client.Projects.AddProjectHook(pid, optsAdd)

	return mapToProjectHook(hook), err
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
