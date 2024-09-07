package server

import (
	"fmt"

	"github.com/vicuani/go_course/gocourse05/animal"
)

type FullHistoryEpisode struct {
	animals []animal.Animal
}

func CreateFullHistoryEpisode() *FullHistoryEpisode {
	return &FullHistoryEpisode{}
}

func (fh *FullHistoryEpisode) Add(an *animal.Animal) {
	fh.animals = append(fh.animals, *an)
}

func (fh *FullHistoryEpisode) GetData() []animal.Animal {
	return fh.animals
}

type DangerousHistoryEpisode struct {
	animals []animal.Animal
}

func CreateDangerousHistoryEpisode() *DangerousHistoryEpisode {
	return &DangerousHistoryEpisode{}
}

func (dh *DangerousHistoryEpisode) Add(an *animal.Animal) {
	if !an.IsAnimalStateDangerous() {
		return
	}
	dh.animals = append(dh.animals, *an)
}

func (dh *DangerousHistoryEpisode) GetData() []animal.Animal {
	return dh.animals
}

type Server struct {
	fullHistory      []*FullHistoryEpisode
	dangerousHistory []*DangerousHistoryEpisode
}

func NewServer() *Server {
	return &Server{
		fullHistory:      []*FullHistoryEpisode{},
		dangerousHistory: []*DangerousHistoryEpisode{},
	}
}

func (s *Server) AddFullHistoryEpisode(e *FullHistoryEpisode) {
	s.fullHistory = append(s.fullHistory, e)
}

func (s *Server) AddDangerousHistoryEpisode(e *DangerousHistoryEpisode) {
	s.dangerousHistory = append(s.dangerousHistory, e)
}

func (s *Server) SendHistory(otherServer *Server) {
	otherServer.fullHistory = s.fullHistory
	otherServer.dangerousHistory = s.dangerousHistory
}

func (s *Server) PrintCompleteHistoryForID(moveID int) {
	if moveID < 0 || moveID >= len(s.fullHistory) {
		fmt.Printf("current move id (%v) is incorrect\n", moveID)
		return
	}

	fmt.Println("\nFull history:")
	for _, episode := range s.fullHistory[moveID].GetData() {
		fmt.Println(episode)
	}

	fmt.Println("\nDangerous history:")
	for _, episode := range s.dangerousHistory[moveID].GetData() {
		fmt.Println(episode)
	}
}
