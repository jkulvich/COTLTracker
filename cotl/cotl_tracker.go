package cotl

import (
	"fmt"
	"os"
	"player/android"
	"player/controller"
	"time"
)

const (
	adbTimeout = 2000
)

// Tracker - Воспроизводит последовательность аккордов и нот с таймингами
// Автоматическая калибровка октавы. Смотрим на минимальную октаву и вычитаем её
type Tracker struct { // Am $C4 $B3 Bm C %200 D E %200
	control *controller.Controller
}

// New - Создаёт новый трекер и подключается к устройству
func New() (*Tracker, error) {
	// Подключение к устройству
	dev, err := android.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Минимальная задержка между нажатиями клавиш
	dev.MinTapDelay = 50

	control, err := controller.New(dev)
	if err != nil {
		return nil, err
	}

	return &Tracker{
		control: control,
	}, nil
}

// Play - Воспросизовдит записанные аккорды/ноты/задержки
func (tracker *Tracker) Play(track *Track, speed float32) error {
	for _, block := range track.blocks {
		if block.Note != nil {
			if err := tracker.control.HarpTapNote(block.Note.Octave, block.Note.Tone); err != nil {
				return err
			}
		}
		<-time.After(time.Duration(block.Delay * 1000000))
	}
	return nil
}
