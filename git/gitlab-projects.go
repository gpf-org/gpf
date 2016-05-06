package git

import (
	"github.com/xanzy/go-gitlab"
)

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
