package server

import (
	"fmt"
	"log"

	"github.com/gpf-org/gpf/git"
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
		projects: projs,
		branches: []*git.Branch{},
	}

	for _, proj := range projs {
		log.Printf("Project %s: reloading webhook", *proj.Name)
		s.git.CreateOrUpdateProjectHook(*proj.ID)

		log.Printf("Project %s: reloading information", *proj.Name)
		branches, err := s.git.ListAllBranches(*proj.ID)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		s.data.branches = append(s.data.branches, branches...)
	}

	s.status = StatusOK
}
