package server

import (
	"fmt"

	"github.com/vicuani/go_course/gocourse05/animal"
)

// this interface cannot be moved to main to avoid cycle imports
type HistoryEpisode interface {
	Add(an *animal.Animal)
	GetData() animal.Animals
}

type FullHistoryEpisode struct {
	data animal.Animals
}

func CreateFullHistoryEpisode() HistoryEpisode {
	return &FullHistoryEpisode{
		data: animal.Animals{},
	}
}

func (fh *FullHistoryEpisode) Add(an *animal.Animal) {
	fh.data = append(fh.data, *an)
}

func (fh *FullHistoryEpisode) GetData() animal.Animals {
	return fh.data
}

type DangerousHistoryEpisode struct {
	data animal.Animals
}

func CreateDangerousHistoryEpisode() HistoryEpisode {
	return &DangerousHistoryEpisode{
		data: animal.Animals{},
	}
}

func (dh *DangerousHistoryEpisode) Add(an *animal.Animal) {
	if !an.IsAnimalStateDangerous() {
		return
	}
	dh.data = append(dh.data, *an)
}

func (dh *DangerousHistoryEpisode) GetData() animal.Animals {
	return dh.data
}

type Server struct {
	fullHistory      []HistoryEpisode
	dangerousHistory []HistoryEpisode
}

func NewServer() *Server {
	return &Server{
		fullHistory:      []HistoryEpisode{},
		dangerousHistory: []HistoryEpisode{},
	}
}

func (s *Server) AddFullHistoryEpisode(e HistoryEpisode) {
	s.fullHistory = append(s.fullHistory, e)
}

func (s *Server) AddDangerousHistoryEpisode(e HistoryEpisode) {
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
