package git

import "errors"

// GitProvider is a common interface providing the same public API
// for different Git host providers
type GitProvider interface {
	ListAllProjects() ([]*Project, error)
	ListAllBranches(pid int) ([]*Branch, error)
	CreateOrUpdateProjectHook(pid int) (*ProjectHook, error)
}

type Project struct {
	ID   *int
	Name *string
}

type Branch struct {
	Name string
}

type ProjectHook struct {
	ID                  int
	URL                 string
	ProjectID           int
	PushEvents          bool
	IssuesEvents        bool
	MergeRequestsEvents bool
}

// NewProvider creates a git provider based in the provider type
func NewProvider(baseURL string, token string, providerType string) (GitProvider, error) {
	switch providerType {
	case "gitlab":
		return newProviderGitlab(baseURL, token), nil
	case "github":
		return nil, errors.New("github provider have not been implemented yet")
	}

	return nil, errors.New("invalid git provider")
}
