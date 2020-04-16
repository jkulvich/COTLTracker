package cotl

import (
	"fmt"
	"github.com/kunaldawn/goandroid/adbutility"
	"github.com/kunaldawn/goandroid/device"
	"github.com/kunaldawn/goandroid/input"
	"player/cotl/tapper"
	"time"
)

const (
	adbTimeout = 2000
)

// Tracker - Воспроизводит последовательность аккордов и нот с таймингами
// Автоматическая калибровка октавы. Смотрим на минимальную октаву и вычитаем её
type Tracker struct { // Am $C4 $B3 Bm C %200 D E %200
	tapper *tapper.Tapper
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

	return &Tracker{
		tapper: tapper.New(screen),
	}, nil
}

// Play - Воспросизовдит записанные аккорды/ноты/задержки
func (tracker *Tracker) Play(track *Track, speed float32) error {
	for _, block := range track.blocks {
		if block.Note != nil {
			if err := tracker.tapper.TapNote(*block.Note); err != nil {
				return err
			}
		}
		<-time.After(time.Duration(block.Delay * 1000000))
	}
	return nil
}
