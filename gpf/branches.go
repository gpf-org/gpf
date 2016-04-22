package gpf

import "github.com/xanzy/go-gitlab"

func ListBranches(baseURL string, token string) ([]*gitlab.Branch, error) {
	git := gitlab.NewClient(nil, token)
	git.SetBaseURL(baseURL)

	pid := 1
	branches, _, err := git.Branches.ListBranches(pid)

	return branches, err
}
