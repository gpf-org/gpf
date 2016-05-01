package server

import (
	"github.com/gpf-org/gpf/git"
)

type MemoryModel struct {
	projects      []*git.Project
	branches      []*git.Branch
	mergeRequests []*git.MergeRequest
}

func (m *MemoryModel) UpdateProject(project *git.Project) {
	// TODO: Add logic to really update projects, not just add them
	m.projects = append(m.projects, project)
}

func (m *MemoryModel) UpdateBranches(branches []*git.Branch) {
	// TODO: Add logic to really update branches, not just add them
	m.branches = append(m.branches, branches...)
}

func (m *MemoryModel) UpdateMergeRequests(mergeRequests []*git.MergeRequest) {
	// TODO: Add logic to really update merge requests, not just add them
	m.mergeRequests = append(m.mergeRequests, mergeRequests...)
}
