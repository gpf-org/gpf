package gpf

import "github.com/xanzy/go-gitlab"

// CreateOrUpdateProjectHook creates a hook to a specified project or update it if already exists.
func CreateOrUpdateProjectHook(baseURL string, token string, projectID int) (*gitlab.ProjectHook, error) {
	git := gitlab.NewClient(nil, token)
	git.SetBaseURL(baseURL)

	// TODO: Add parameter to receive the gpf server address
	hookURL := "http://localhost:5544/reload"

	// TODO: handle paging search
	optsList := &gitlab.ListProjectHooksOptions{}
	hooks, _, err := git.Projects.ListProjectHooks(projectID, optsList)

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

			hook, _, err := git.Projects.EditProjectHook(projectID, hooks[i].ID, optsEdit)

			return hook, err
		}
	}

	optsAdd := &gitlab.AddProjectHookOptions{
		URL:                 hookURL,
		PushEvents:          true,
		IssuesEvents:        true,
		MergeRequestsEvents: true,
		TagPushEvents:       true,
	}

	hook, _, err := git.Projects.AddProjectHook(projectID, optsAdd)

	return hook, err
}
