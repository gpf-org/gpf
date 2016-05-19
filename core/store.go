package core

import (
	"errors"
	"sort"

	"github.com/gpf-org/gpf/git"
)

var (
	ErrIssueNotFound     = errors.New("Issue not found")
	ErrProjectNotFound   = errors.New("Project not found")
	ErrInvalidBranchName = errors.New("Invalid branch name")
)

type ListMergeRequestsOptions struct {
	ProjectID    int
	SourceBranch string
	TargetBranch string
}

type Store struct {
	extractor     IssueNameExtractor
	projects      map[int]*git.Project
	issues        map[string][]*git.Branch
	mergeRequests []*git.MergeRequest
}

func NewStore(extractor IssueNameExtractor) *Store {
	return &Store{
		extractor: extractor,
		projects:  make(map[int]*git.Project),
		issues:    make(map[string][]*git.Branch),
	}
}

func (s *Store) Reset() {
	s.projects = make(map[int]*git.Project)
	s.issues = make(map[string][]*git.Branch)
	s.mergeRequests = make([]*git.MergeRequest, 0)
}

func (s *Store) AddProject(project *git.Project) *git.Project {
	s.projects[project.ID] = project

	return project
}

func (s *Store) AddBranch(branch *git.Branch) error {
	issueName, ok := s.extractor.IssueNameExtract(branch.Name)
	if !ok {
		return ErrInvalidBranchName
	}

	s.issues[issueName] = append(s.issues[issueName], branch)

	return nil
}

func (s *Store) AddMergeRequest(mergeRequest *git.MergeRequest) *git.MergeRequest {
	s.mergeRequests = append(s.mergeRequests, mergeRequest)

	return mergeRequest
}

func (s *Store) GetProject(id int) (*git.Project, error) {
	project, ok := s.projects[id]
	if !ok {
		return nil, ErrProjectNotFound
	}

	return project, nil
}

func (s *Store) ListIssues() []string {
	result := make([]string, 0, len(s.issues))

	for name, _ := range s.issues {
		result = append(result, name)
	}

	sort.Strings(result)

	return result
}

func (s *Store) ListBranches(issue string) ([]*git.Branch, error) {
	branches, ok := s.issues[issue]
	if !ok {
		return nil, ErrIssueNotFound
	}

	return branches, nil
}

func (s *Store) ListMergeRequests(options *ListMergeRequestsOptions) []*git.MergeRequest {
	result := []*git.MergeRequest{}

	for _, mergeRequest := range s.mergeRequests {
		if options.ProjectID != 0 && options.ProjectID != mergeRequest.ProjectID {
			continue
		}

		if options.SourceBranch != "" && options.SourceBranch != mergeRequest.SourceBranch {
			continue
		}

		if options.TargetBranch != "" && options.TargetBranch != mergeRequest.TargetBranch {
			continue
		}

		result = append(result, mergeRequest)
	}

	return result
}
