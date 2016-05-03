package git

import "errors"

// GitProvider is a common interface providing the same public API
// for different Git host providers
type GitProvider interface {
	ListAllProjects() ([]*Project, error)
	ListAllBranches(pid int) ([]*Branch, error)
	CreateOrUpdateProjectHook(pid int, hookURL string) (*ProjectHook, error)
	CreateMergeRequest(pid int, opts CreateMergeRequestOptions) (*MergeRequest, error)
	ListMergeRequests(pid int) ([]*MergeRequest, error)
}

type Project struct {
	ID   int
	Name string
}

type Branch struct {
	Name      string
	ProjectID int
}

type ProjectHook struct {
	ID                  int
	URL                 string
	ProjectID           int
	PushEvents          bool
	IssuesEvents        bool
	MergeRequestsEvents bool
}

type MergeRequest struct {
	// It has available more fields.
	// Include only those really used.
	ID             int
	ProjectID      int
	Title          string
	Description    string
	WorkInProgress bool
	State          string
	TargetBranch   string
	SourceBranch   string
	Upvotes        int
	Downvotes      int
}

type CreateMergeRequestOptions struct {
	Title           string
	Description     string
	SourceBranch    string
	TargetBranch    string
	AssigneeID      int
	TargetProjectID int
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
