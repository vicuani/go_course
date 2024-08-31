package server

import (
	"fmt"

	"github.com/vicuani/go_course/gocourse05/animal"
)

type Server struct {
	FullHistory      []animal.HistoryEpisode
	DangerousHistory []animal.HistoryEpisode
}

func NewServer() *Server {
	return &Server{
		FullHistory:      []animal.HistoryEpisode{},
		DangerousHistory: []animal.HistoryEpisode{},
	}
}

func (s *Server) SendHistory(otherServer *Server) {
	otherServer.FullHistory = s.FullHistory
	otherServer.DangerousHistory = s.DangerousHistory
}

func (s *Server) PrintCompleteHistoryForID(moveID int) {
	if moveID < 0 || moveID >= len(s.FullHistory) {
		fmt.Printf("current move id (%v) is incorrect\n", moveID)
		return
	}

	fmt.Println("\nFull history:")
	for _, episode := range s.FullHistory[moveID].GetData() {
		fmt.Println(episode)
	}

	fmt.Println("\nDangerous history:")
	for _, episode := range s.DangerousHistory[moveID].GetData() {
		fmt.Println(episode)
	}
}
