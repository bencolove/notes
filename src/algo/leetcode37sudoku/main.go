package main

import "fmt"

const (
	SHAPE = 9
	MAX   = SHAPE * SHAPE
	BOX   = 3
	EMPTY = '.'
)

var (
	DEBUG = false
)

type board struct {
	data [MAX]byte
}

func NewBoard(input [][]byte) *board {
	b := &board{}
	for i := 0; i < SHAPE; i++ {
		for j := 0; j < SHAPE; j++ {
			v := byte(EMPTY)
			if input[i][j] != EMPTY {
				v = input[i][j] - '0'
			}
			b.data[i*SHAPE+j] = v
		}
	}
	return b
}

func (b *board) get(i int) byte {
	return b.data[i]
}

func (b *board) set(i int, v byte) {
	b.data[i] = v
}

func (b *board) unset(i int) {
	b.data[i] = EMPTY
}

func (b *board) isValid() bool {
	// check rows
	for i := 0; i < SHAPE; i++ {
		rowSum := 0
		for j := 0; j < SHAPE; j++ {
			if b.data[i*SHAPE+j] == EMPTY {
				return false
			}
			rowSum += int(b.data[i*SHAPE+j])
		}
		if rowSum != 45 {
			return false
		}
	}
	// check cols
	for i := 0; i < SHAPE; i++ {
		rowSum := 0
		for j := 0; j < SHAPE; j++ {
			rowSum += int(b.data[j*SHAPE+i])
		}
		if rowSum != 45 {
			return false
		}
	}
	// check sub-box
	for i := 0; i < SHAPE; i += BOX {
		for j := 0; j < SHAPE; j += BOX {
			rowSum := 0
			for x := 0; x < BOX; x++ {
				for y := 0; y < BOX; y++ {
					rowSum += int(b.data[(i*BOX+x)*SHAPE+j*BOX+y])
				}
			}
			if rowSum != 45 {
				return false
			}
		}
	}
	return true
}

func (b *board) getCandidates(i int) []byte {
	if i >= MAX {
		return []byte{}
	}

	candidates := make([]uint8, 9, 9)

	row, col := i/SHAPE, i%SHAPE

	// check rows
	debug("  find cans at row %d: ", i)
	for j := 0; j < SHAPE; j++ {
		idx := row*SHAPE + j
		if b.data[idx] != EMPTY {
			debug("%d ", b.data[idx])
			candidates[int(b.data[idx])-1] = 1
		}
	}
	debugln()

	// check cols
	debug("  find cans at col %d: ", i)
	for j := 0; j < SHAPE; j++ {
		idx := j*SHAPE + col
		if b.data[idx] != EMPTY {
			debug("%d ", b.data[idx])
			candidates[int(b.data[idx])-1] = 1
		}
	}
	debugln()

	// check sub-box
	subi, subj := i/SHAPE/3, i%SHAPE/3

	debug("  find cans at box %d,%d: ", subi, subj)
	for x := 0; x < BOX; x++ {
		for y := 0; y < BOX; y++ {
			idx := (subi*BOX+x)*SHAPE + subj*BOX + y
			if b.data[idx] != EMPTY {
				debug("%d ", b.data[idx])
				candidates[int(b.data[idx])-1] = 1
			}
		}
	}
	debugln()

	debug("used digits at %d: %v\n", i, candidates)
	left := make([]byte, 0, 9)
	for i, v := range candidates {
		if v == 0 {
			left = append(left, byte(i+1))
		}
	}
	debug("candidates at %d: %v\n", i, left)
	return left
}

func (b *board) searchN(i int) (bool, *board) {
	if b.isValid() {
		return true, b
	}

	if i >= MAX {
		return false, nil
	}

	for ; i < MAX && b.data[i] != EMPTY; i++ {
	}

	candidates := b.getCandidates(i)

	pi, pj := i/SHAPE+1, i%SHAPE+1
	debug("search at [%d,%d]: candidates %v\n", pi, pj, candidates)

	for _, v := range candidates {
		// set
		debug("set %d at [%d,%d]\n", v, pi, pj)
		b.set(i, v)
		if found, solution := b.searchN(i + 1); !found {
			// unset
			b.set(i, EMPTY)
			continue
		} else {
			// found solution
			return true, solution
		}
	}
	// exhausted
	return false, nil
}

func (b *board) search() bool {
	if b.isValid() {
		return true
	}

	for i := 0; i < MAX; i++ {
		if b.get(i) == EMPTY {
			pi, pj := i/SHAPE+1, i%SHAPE+1
			for _, v := range b.getCandidates(i) {
				// set
				debug("set %d at [%d,%d]\n", v, pi, pj)
				b.set(i, v)
				if b.search() {
					return true
				} else {
					// unset
					b.unset(i)
				}
			}
			return false
		}
	}
	return true
}

func (b *board) searchFrom(pos int) bool {
	if b.isValid() {
		return true
	}

	for pos := 0; pos < MAX && b.get(pos) != EMPTY; pos++ {
	}

	if pos >= MAX {
		return true
	} else {
		for _, v := range b.getCandidates(pos) {
			b.set(pos, v)
			if b.searchFrom(pos + 1) {
				return true
			} else {
				b.unset(pos)
			}
		}
		return false
	}

	// for i := 0; i < MAX; i++ {
	// 	if b.get(i) == EMPTY {
	// 		pi, pj := i/SHAPE+1, i%SHAPE+1
	// 		for _, v := range b.getCandidates(i) {
	// 			// set
	// 			debug("set %d at [%d,%d]\n", v, pi, pj)
	// 			b.set(i, v)
	// 			if b.search() {
	// 				return true
	// 			} else {
	// 				// unset
	// 				b.unset(i)
	// 			}
	// 		}
	// 		return false
	// 	}
	// }
	return true
}

func (b *board) solution() [][]byte {
	ret := make([][]byte, SHAPE)
	for i := 0; i < SHAPE; i++ {

		ret[i] = make([]byte, SHAPE)

		for j := 0; j < SHAPE; j++ {
			ret[i][j] = b.data[i*SHAPE+j] + '0'
		}
	}
	return ret
}

func (b *board) populateSolution(output [][]byte) {
	for i := 0; i < SHAPE; i++ {
		for j := 0; j < SHAPE; j++ {
			output[i][j] = b.data[i*SHAPE+j] + '0'
		}
	}
}

func debug(format string, args ...interface{}) {
	if DEBUG {
		fmt.Printf(format, args...)
	}
}
func debugln() {
	if DEBUG {
		fmt.Println()
	}
}

func main() {
	input := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	b := NewBoard(input)
	if b.search() {
		fmt.Printf("found: %v\n", b.solution())
	} else {
		fmt.Printf("no result\n")
	}
}
