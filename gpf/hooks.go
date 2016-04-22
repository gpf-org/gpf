package gpf

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

// CreateProjectHook creates a hook to a specified project.
func CreateProjectHook(baseURL string, token string, projectID int) (*gitlab.ProjectHook, error) {
	git := gitlab.NewClient(nil, token)
	git.SetBaseURL(baseURL)

	hookURL := baseURL + "/gpf-hooks"

	// TODO: handle paging search
	optsList := &gitlab.ListProjectHooksOptions{}
	hooks, _, err := git.Projects.ListProjectHooks(projectID, optsList)
	if err != nil {
		log.Fatal(err)
	}

	for _, hook := range hooks {
		if hook.URL == hookURL {
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
