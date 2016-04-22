package gpf

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

// CreateOrUpdateProjectHook creates a hook to a specified project or update it if already exists.
func CreateOrUpdateProjectHook(baseURL string, token string, projectID int) (*gitlab.ProjectHook, error) {
	git := gitlab.NewClient(nil, token)
	git.SetBaseURL(baseURL)

	hookURL := baseURL + "/gpf-hooks"

	// TODO: handle paging search
	optsList := &gitlab.ListProjectHooksOptions{}
	hooks, _, err := git.Projects.ListProjectHooks(projectID, optsList)
	if err != nil {
		log.Fatal(err)
	}

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
			if err != nil {
				log.Fatal(err)
			}

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
	if err != nil {
		log.Fatal(err)
	}

	return hook, err
}
