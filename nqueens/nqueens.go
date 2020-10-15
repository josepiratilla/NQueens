package nqueens

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

	row := 0
	for i := 0; i < s.Size; i++ {
		if s.Queens[i] == 0 {
			break
		}
		row = i
	}

	for {
		if s.Queens[row] == s.Size {
			s.Queens[row] = 0
			row--
			if row < 0 {
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
