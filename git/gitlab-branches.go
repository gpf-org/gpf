package git

func (gp GitLabProvider) ListAllBranches(pid int) ([]*Branch, error) {
	result, _, err := gp.client.Branches.ListBranches(pid)

	if err != nil {
		return nil, err
	}

	branches := make([]*Branch, len(result))

	for i, value := range result {
		branches[i] = &Branch{Name: value.Name, ProjectID: pid}
	}

	return branches, nil
}
