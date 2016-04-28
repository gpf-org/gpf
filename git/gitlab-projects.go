package git

import (
	"github.com/xanzy/go-gitlab"
)

// ListAllProjects gets a list of all Git projects.
func (gp GitLabProvider) ListAllProjects() ([]*Project, error) {
	result := []*Project{}

	options := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 100,
		},
	}

	nextPage := true

	for i := 1; nextPage; i++ {
		options.Page = i

		projects, _, err := gp.client.Projects.ListAllProjects(options)

		if err != nil {
			return nil, err
		}

		count := len(projects)

		switch true {
		case count == 0:
			nextPage = false
		case count > 0:
			for _, value := range projects {
				result = append(result, &Project{ID: value.ID, Name: value.Name})
			}
		}
	}

	return result, nil
}
