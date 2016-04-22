package gpf

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

// ListAllProjects gets a list of all Git projects.
func ListAllProjects(baseURL string, token string) ([]*gitlab.Project, error) {
	git := gitlab.NewClient(nil, token)
	git.SetBaseURL(baseURL)

	// TODO: handle paging search
	optsList := &gitlab.ListProjectsOptions{}
	projs, _, err := git.Projects.ListAllProjects(optsList)
	if err != nil {
		log.Fatal(err)
	}
	return projs, err
}
