package server

import (
	"log"

	"github.com/gpf-org/gpf/git"
)

var (
	err      error
	branches []*git.Branch
	mrs      []*git.MergeRequest
)

func (s *Server) Reload() error {
	log.Printf("Reloading the server. It may take awhile.")

	projs, err := s.git.ListAllProjects()
	if err != nil {
		return err
	}

	log.Printf("Projects available: %d", len(projs))

	for _, proj := range projs {
		s.model.UpdateProject(proj)

		log.Printf("Project %s: reloading webhook", *proj.Name)
		s.git.CreateOrUpdateProjectHook(*proj.ID, s.options.PublicURL)

		log.Printf("Project %s: reloading branches", *proj.Name)
		branches, err = s.git.ListAllBranches(*proj.ID)
		if err != nil {
			return err
		}

		s.model.UpdateBranches(branches)

		log.Printf("Project %s: reloading merge requests", *proj.Name)
		mrs, err = s.git.ListMergeRequests(*proj.ID)
		if err != nil {
			return err
		}

		s.model.UpdateMergeRequests(mrs)
	}

	return nil
}
