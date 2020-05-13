package block

import "fmt"

type Blocks []Block

// Block - Описывает представление аудио блока
type Block struct {
	Delay   int
	Note    *Note
	LineNum int
	PosNum  int
}

// NewBlocks - Создаёт новые музыкальные блоки из текстового блока
func NewBlocks(block string, timing int, line, pos int) ([]Block, error) {
	if note, err := NewNote(block); err == nil {
		return note.GetBlocks(), nil
	} else
	if chord, err := NewChord(block); err == nil {
		return chord.GetBlocks(), nil
	} else
	if delay, err := NewDelay(block, timing); err == nil {
		return delay.GetBlocks(), nil
	}
	return nil, fmt.Errorf("unknown block type for %s at %d:%d", block, line, pos)
}
