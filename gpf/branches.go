package gpf

import (
	"log"

	"github.com/xanzy/go-gitlab"
)

func ListBranches(token string) ([]*gitlab.Branch, error) {
	git := gitlab.NewClient(nil, token)

	pid := 1
	branches, _, err := git.Branches.ListBranches(pid)
	if err != nil {
		log.Fatal(err)
	}

	return branches, err
}
