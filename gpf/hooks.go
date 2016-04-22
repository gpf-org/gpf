package gpf

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func CreateProjectHook(baseURL string, token string) (*gitlab.ProjectHook, error) {
	git := gitlab.NewClient(nil, token)
	git.SetBaseURL(baseURL)

	hookURL := baseURL + "/gpf-hooks"

	// TODO: handle paging search
	pid := 1
	optsList := &gitlab.ListProjectHooksOptions{}
	hooks, _, err := git.Projects.ListProjectHooks(pid, optsList)
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

	hook, _, err := git.Projects.AddProjectHook(pid, optsAdd)
	if err != nil {
		log.Fatal(err)
	}

	return hook, err
}
