package nqueens

import "testing"

func TestCheck(t *testing.T) {
	vs := []struct {
		queens   []int
		expected bool
	}{
		{
			[]int{0, 0},
			true,
		},
		{
			[]int{1, 1},

			false,
		},
		{
			[]int{1, 2},
			false,
		},
		{
			[]int{2, 4, 1, 3},
			true,
		},
		{
			[]int{2, 5, 0, 0, 0, 0, 0, 0},
			true,
		},
		{
			[]int{5, 2, 7, 0, 0, 0, 0, 0},
			false,
		},
		{
			[]int{5, 2, 4, 7, 3, 8, 6, 1},
			true,
		},
	}
	for _, v := range vs {
		s := new(Solver)
		s.Size = len(v.queens)
		s.Queens = v.queens[:]
		actual := s.Check()
		if actual != v.expected {
			t.Errorf("For Queens %v we expected it be a %v solution, but the check is returning %v solution", v.queens, v.expected, actual)
		}
	}
}

func TestSolve8Queens(t *testing.T) {
	s := new(Solver)
	s.Size = 8
	s.Queens = make([]int, 8)
	result := s.Solve()
	if result == false {
		t.Error("8 Queens has a solution and the solver didn't find it.")
	}
	result = s.Check()
	if result == false {
		t.Errorf("The solution found for 8 Queens is not valid.\nSolution: %v\n", s.Queens)
	}
	for _, q := range s.Queens {
		if q == 0 {
			t.Errorf("The solution found for 8 Queens has missing points.\nSolution: %v\n", s.Queens)
		}
	}
}

func TestCountSolutions(t *testing.T) {
	vs := []struct {
		size     int
		expected int
	}{
		{
			1,
			1,
		},
		{
			2,
			0,
		},
		{
			3,
			0,
		},
		{
			4,
			2,
		},
		{
			5,
			10,
		},
		{
			6,
			4,
		},
		{
			7,
			40,
		},
		{
			8,
			92,
		},
		{
			9,
			352,
		},
		{
			10,
			724,
		},
	}
	for _, v := range vs {
		actual := HowManySolutions(v.size)
		if actual != v.expected {
			t.Errorf("For size %d the function has found %d solutions, and the expected value is %d.\n", v.size, actual, v.expected)
		}
		actualc := HowManySolutionsConcurrent(v.size)
		if actualc != v.expected {
			t.Errorf("For size %d the concurrent function has found %d solutions, and the expected value is %d.\n", v.size, actualc, v.expected)
		}
	}
}
