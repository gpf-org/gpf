package git

import "github.com/xanzy/go-gitlab"

type GitLabProvider struct {
	client *gitlab.Client
}

func newProviderGitlab(baseURL string, token string) GitLabProvider {
	glc := gitlab.NewClient(nil, token)
	glc.SetBaseURL(baseURL)
	return GitLabProvider{glc}
}
