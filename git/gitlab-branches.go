package git

func (gp GitLabProvider) ListAllBranches(pid int) ([]*Branch, error) {
	result, _, err := gp.client.Branches.ListBranches(pid)

	branches := make([]*Branch, len(result))

	for _, value := range result {
		branches = append(branches, &Branch{Name: value.Name, ProjectID: pid})
	}

	return branches, err
}
