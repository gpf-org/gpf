package server

import (
	"regexp"
	"sort"

	"github.com/gpf-org/gpf/git"
)

type MemoryModel struct {
	pattern       *regexp.Regexp
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

func (m *MemoryModel) List() []*Feature {
	features := make(map[string]*Feature)

	for _, branch := range m.branches {
		groups := m.pattern.FindAllStringSubmatch(branch.Name, -1)
		if len(groups) == 0 {
			continue
		}

		name := groups[0][1]

		feature, ok := features[name]
		if !ok {
			feature = &Feature{
				Name:     name,
				Commands: []string{},
				Branches: make([]*FeatureBranch, 0),
			}
			features[name] = feature
		}

		featureBranch := &FeatureBranch{
			BranchName:  branch.Name,
			ProjectID:   branch.ProjectID,
			ProjectName: m.getProjectName(branch.ProjectID),
		}

		feature.Branches = append(feature.Branches, featureBranch)
	}

	// command: code-review request
	for _, feature := range features {
		mergeRequestOpenToDevelop := false

		for _, branch := range feature.Branches {
			exists := m.existMergeRequest(&FindMergeRequestOptions{
				SourceBranch:    branch.BranchName,
				TargetBranch:    "develop",
				TargetProjectID: branch.ProjectID,
				State:           "open",
			})

			mergeRequestOpenToDevelop = mergeRequestOpenToDevelop || exists
		}

		if !mergeRequestOpenToDevelop {
			feature.Commands = append(feature.Commands, "code-review request")
		}
	}

	result := make([]*Feature, 0, len(features))

	for _, feature := range features {
		result = append(result, feature)

		sort.Sort(ByFeatureBranch(feature.Branches))
	}

	sort.Sort(ByFeature(result))

	return result
}

type FindMergeRequestOptions struct {
	SourceBranch    string
	TargetBranch    string
	State           string
	TargetProjectID int
}

func (m *MemoryModel) existMergeRequest(options *FindMergeRequestOptions) bool {
	for _, mergeRequest := range m.mergeRequests {
		if !matchIntOption(options.TargetProjectID, mergeRequest.TargetProjectID) {
			continue
		}

		if !matchStringOption(options.SourceBranch, mergeRequest.SourceBranch) {
			continue
		}

		if !matchStringOption(options.TargetBranch, mergeRequest.TargetBranch) {
			continue
		}

		if !matchStringOption(options.State, mergeRequest.State) {
			continue
		}

		return true
	}

	return false
}

func matchStringOption(option string, value string) bool {
	return option == "" || option == value
}

func matchIntOption(option int, value int) bool {
	return option == 0 || option == value
}

func (m *MemoryModel) getProjectName(projectID int) string {
	for _, project := range m.projects {
		if project.ID == projectID {
			return project.Name
		}
	}

	return ""
}
