package server

import (
	"fmt"
	"log"
)

func (s *Server) reload() {
	s.status = StatusMaintenance

	projs, err := s.git.ListAllProjects()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	log.Printf("Projects available: %d", len(projs))

	for _, proj := range projs {
		log.Printf("Project %s: reloading webhook", *proj.Name)
		s.git.CreateOrUpdateProjectHook(*proj.ID)

		log.Printf("Project %s: reloading information", *proj.Name)
		// TODO: save project information into the gpf database
	}

	s.status = StatusOK
}
