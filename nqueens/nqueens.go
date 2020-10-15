package nqueens

import (
	"runtime"
	"sync"
)

//Solver is the main data type to solve the NQueens problem
type Solver struct {
	Size   int
	Queens []int
}

//Check reviews the queens position and check if it is a valid solution
func (s *Solver) Check() bool {
	for i, v := range s.Queens {
		if v == 0 {
			return true
		}
		for j := 0; j < i; j++ {
			if s.Queens[j] == v || abs(s.Queens[j]-v) == abs(i-j) {
				return false
			}
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

//Solve solves the next N Queens problem, starting form the position given by the solver.
func (s *Solver) Solve() bool {

	return s.solveFromRow(0)

}

func (s *Solver) solveFromRow(fixedRow int) bool {
	row := fixedRow
	for i := 0; i < s.Size; i++ {
		row = i
		if s.Queens[i] == 0 {
			break
		}

	}

	for {
		if s.Queens[row] == s.Size {
			s.Queens[row] = 0
			row--
			if row < fixedRow {
				return false
			}
			continue
		}
		s.Queens[row]++
		if s.Check() {
			row++
			if row >= s.Size {
				return true
			}
		}
	}
}

//HowManySolutions returns the total number of solutions available for the input size.
func HowManySolutions(size int) int {

	s := new(Solver)

	s.Size = size
	s.Queens = make([]int, size)

	var i int
	for i = 0; s.Solve(); i++ {
	}

	return i
}

//HowManySolutionsConcurrent returns the total number of solutions available for the input size starting as many threads as the size
func HowManySolutionsConcurrent(size int) int {

	sols := make([]int, size)
	var wg sync.WaitGroup

	parallelThreads := runtime.NumCPU() - 1
	if parallelThreads <= 0 {
		parallelThreads = 1
	}

	execs := make(chan int)

	go func() {
		for i := 0; i < size; i++ {
			execs <- i
		}
		close(execs)
	}()

	for i := 0; i < parallelThreads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for exec := range execs {
				s := new(Solver)
				s.Size = size
				s.Queens = make([]int, size)
				s.Queens[0] = exec
				for ; s.solveFromRow(1); sols[exec]++ {
				}
			}
		}()
	}
	wg.Wait()
	i := 0
	for _, sol := range sols {
		i += sol
	}

	return i
}
