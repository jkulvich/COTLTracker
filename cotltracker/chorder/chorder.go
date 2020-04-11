package chorder

import "player/cotltracker/notter"

// Chorder - Воспроизводитель аккордов
type Chorder struct {
	notter *notter.Notter
}

// New - Создаёт новый проигрыватель нот
func New(notter *notter.Notter) *Chorder {
	return &Chorder{
		notter: notter,
	}
}
