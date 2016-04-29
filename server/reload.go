package server

import (
	"fmt"
	"log"

	"github.com/gpf-org/gpf/git"
)

var (
	err      error
	branches []*git.Branch
	mrs      []*git.MergeRequest
)

func (s *Server) reload() {
	s.status = StatusMaintenance
	log.Printf("Reloading the server. It may take awhile.")

	projs, err := s.git.ListAllProjects()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	log.Printf("Projects available: %d", len(projs))

	s.data = GitData{
		projects:      projs,
		branches:      []*git.Branch{},
		mergeRequests: []*git.MergeRequest{},
	}

	for _, proj := range projs {
		log.Printf("Project %s: reloading webhook", *proj.Name)
		s.git.CreateOrUpdateProjectHook(*proj.ID)

		log.Printf("Project %s: reloading branches", *proj.Name)
		branches, err = s.git.ListAllBranches(*proj.ID)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		s.data.branches = append(s.data.branches, branches...)

		log.Printf("Project %s: reloading merge requests", *proj.Name)
		mrs, err = s.git.ListMergeRequests(*proj.ID)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		s.data.mergeRequests = append(s.data.mergeRequests, mrs...)
	}

	s.status = StatusOK
}
