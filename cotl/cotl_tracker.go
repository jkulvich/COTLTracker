package cotl

import (
	"fmt"
	"github.com/kunaldawn/goandroid/adbutility"
	"github.com/kunaldawn/goandroid/device"
	"github.com/kunaldawn/goandroid/input"
	"player/cotl/chorder"
	"player/cotl/inputter"
	"player/cotl/notter"
	"time"
)

const (
	adbTimeout = 2000
)

// Tracker - Воспроизводит последовательность аккордов и нот с таймингами
// Автоматическая калибровка октавы. Смотрим на минимальную октаву и вычитаем её
type Tracker struct { // Am $C4 $B3 Bm C %200 D E %200
	notter  *notter.Notter
	chorder *chorder.Chorder
}

// New - Создаёт новый трекер и подключается к устройству
func New(adbEndpoint string, serial string) (*Tracker, error) {
	// Подключение к устройству
	dev := device.NewDevice(serial, adbTimeout, adbutility.GetNewAdbEndpoint(adbEndpoint))
	if ok, err := dev.IsAvailable(); !ok {
		return nil, fmt.Errorf("can't connect: %s\n", err)
	}

	// Получение управляющих блоков
	screen := input.NewTouchScreen(dev)
	inputControl := inputter.New(screen)
	notePlayer := notter.New(inputControl)
	chordPlayer := chorder.New(notePlayer)

	return &Tracker{
		notter:  notePlayer,
		chorder: chordPlayer,
	}, nil
}

// Play - Воспросизовдит записанные аккорды/ноты/задержки
func (tracker *Tracker) Play(track *Track, speed float32) error {
	for _, block := range track.blocks {

		switch block.typ {
		case trackBlockDelay:
			delay := float32(int(time.Millisecond)*block.delay) * (1/speed)
			<-time.After(time.Duration(delay))
		case trackBlockNote:
			if err := tracker.notter.PlayNote(block.note, track.octaveShift+track.shift); err != nil {
				return err
			}
		case trackBlockChord:
			if err := tracker.chorder.PlayChord(block.chord); err != nil {
				return err
			}
		}

	}

	return nil
}
