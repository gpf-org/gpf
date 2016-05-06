package server

import (
	"github.com/gpf-org/gpf/git"
)

type FeatureBranch struct {
	ProjectName string `json:"project_name"`
	BranchName  string `json:"branch_name"`
}

type Feature struct {
	Name     string           `json:"name"`
	Status   string           `json:"status"`
	Branches []*FeatureBranch `json:"branches,omitempty"`
}

type ByFeature []*Feature

func (f ByFeature) Len() int           { return len(f) }
func (f ByFeature) Swap(i, j int)      { f[i], f[j] = f[j], f[i] }
func (f ByFeature) Less(i, j int) bool { return f[i].Name < f[j].Name }

type ByFeatureBranch []*FeatureBranch

func (f ByFeatureBranch) Len() int      { return len(f) }
func (f ByFeatureBranch) Swap(i, j int) { f[i], f[j] = f[j], f[i] }
func (f ByFeatureBranch) Less(i, j int) bool {
	sameProjectName := f[i].ProjectName == f[j].ProjectName

	if sameProjectName {
		return f[i].BranchName < f[j].BranchName
	}

	return f[i].ProjectName < f[j].ProjectName
}

type ServerModel interface {
	UpdateProject(project *git.Project)
	UpdateBranches(branches []*git.Branch)
	UpdateMergeRequests(mergeRequests []*git.MergeRequest)
	List() []*Feature
	// TODO: Add relevant methods to traverse projects, branches and merge requests
}
