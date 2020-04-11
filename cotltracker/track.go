package cotltracker

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type trackBlocks []*trackBlock

// Track - Содержит настройки композиции и последовательность музыкальных блоков
type Track struct {
	blocks      trackBlocks
	octaveShift int
}

// NewTrack - Парсит нотное представление и представляет его в виде трека
func NewTrack(stave string) (*Track, error) {
	track := Track{
		blocks:      make(trackBlocks, 0),
		octaveShift: 0,
	}

	stave = removeComments(stave)
	stave = strings.ReplaceAll(stave, "\n", " ")
	runes := strings.Split(stave, " ")

	for _, r := range runes {
		r = strings.TrimSpace(r)
		if len(r) > 0 {
			track.blocks = append(track.blocks, parseBlock(r))
		}
	}

	// Пытаемся автоматически определить сдвиг октавы
	octaveShift, err := detectOctaveShift(track.blocks)
	if err != nil {
		return nil, err
	}
	track.octaveShift = octaveShift

	return &track, nil
}

// removeComments - Удаляет все строки начинающиеся с #
func removeComments(stave string) string {
	lines := strings.Split(stave, "\n")
	nlines := make([]string, 0, len(lines))

	for _, line := range lines {
		if !strings.HasPrefix(line, "#") {
			nlines = append(nlines, line)
		}
	}

	return strings.Join(nlines, "\n")
}

// detectOctaveShift - Автоматически определяет сдвиг октавы для нот
func detectOctaveShift(blocks trackBlocks) (int, error) {
	minOct := math.MaxInt32

	for _, block := range blocks {
		if block.typ == trackBlockNote {

			res := regexp.MustCompile(`([ABCDEFG])(\d*)`).FindStringSubmatch(strings.ToUpper(block.note))
			if len(res) != 3 {
				return 0, fmt.Errorf("unexpected note form: %s", block.note)
			}

			octave, err := strconv.Atoi(res[2])
			if err != nil {
				return 0, fmt.Errorf("unexpected note form: %s", block.note)
			}

			if octave < minOct {
				minOct = octave
			}

		}
	}

	// Если значение не сменилось - в композиции нет нот, не двигаем
	if minOct == math.MaxInt32 {
		minOct = 0
	}

	return -minOct, nil
}
