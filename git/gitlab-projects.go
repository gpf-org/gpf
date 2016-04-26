package git

import (
	"github.com/xanzy/go-gitlab"
)

// ListAllProjects gets a list of all Git projects.
func (gp GitLabProvider) ListAllProjects() ([]*Project, error) {
	// TODO: handle paging search
	optsList := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			Page:    1,
			PerPage: 99999999999,
		},
	}
	result, _, err := gp.client.Projects.ListAllProjects(optsList)

	projs := make([]*Project, len(result))

	for _, value := range result {
		projs = append(projs, &Project{ID: value.ID, Name: value.Name})
	}

	return projs, err
}
