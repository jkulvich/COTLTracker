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
func New(minDelay int) (*Tracker, error) {
	// Подключение к устройству
	dev, err := android.New()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Минимальная задержка между нажатиями клавиш
	dev.MinTapDelay = minDelay

	control, err := controller.New(dev)
	if err != nil {
		return nil, err
	}

	return &Tracker{
		control: control,
	}, nil
}

// Play - Воспросизовдит записанные аккорды/ноты/задержки
func (tracker *Tracker) Play(track *Track, start int) error {
	for i, block := range track.blocks {
		if i < start {
			continue
		}

		note := ""
		if block.Note != nil {
			note = block.Note.String()
			if err := tracker.control.HarpTapNote(block.Note.Octave, block.Note.Tone); err != nil {
				return err
			}
		}

		fmt.Printf("[%06.2f%%] Block <%2s> [%04d] of [%04d] at %03d:%02d\n",
			float64(i+1)/float64(len(track.blocks))*100,
			note,
			i+1,
			len(track.blocks),
			block.LineNum,
			block.PosNum,
		)

		<-time.After(time.Duration(block.Delay * 1000000))
	}
	return nil
}
