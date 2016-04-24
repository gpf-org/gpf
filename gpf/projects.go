package gpf

import "github.com/xanzy/go-gitlab"

// ListAllProjects gets a list of all Git projects.
func ListAllProjects(baseURL string, token string) ([]*gitlab.Project, error) {
	git := gitlab.NewClient(nil, token)
	git.SetBaseURL(baseURL)

	result := []*gitlab.Project{}

	options := &gitlab.ListProjectsOptions{
		ListOptions: gitlab.ListOptions{
			PerPage: 100,
		},
	}

	nextPage := true

	for i := 1; nextPage; i++ {
		options.Page = i

		projects, _, err := git.Projects.ListProjects(options)

		if err != nil {
			return nil, err
		}

		count := len(projects)

		switch true {
		case count == 0:
			nextPage = false
		case count > 0:
			result = append(result, projects...)
		}
	}

	return result, nil
}
