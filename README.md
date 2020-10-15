# NQueens
Solver for the 8 queens problem, but with arbitrary board size.

## Project structure
Everyting interesting is inside `nqueens` folder. The type `Solver` is defined with all logic embeded into it. 
To initialize the solver you must manually initalize its key variable like:
```go
s:= new(Solver)
s.Size = 8
s.Queens = make([]int,8)
```
`Size` defines the size of the board to place the N Queens. In fact it defines N.
I created the next methods:
### Check
```go
func (s *Solver) Check() bool
```
Given a position defined in `Queens[]` this function checks if the position is valid. I use it, for instance, in the test to check the solutions provided.
### Solve
```go
func (s *Solver) Solve() bool
```
Finds the next valid solution, starting from the current position. Solved can be called repeatedly over the same solver. Between iterations, reading `Queens[]` the different solutions can be obtained.
It returns `true` if a solution is found. If there are no solutions left it returns `false`.
### HowManySolutions
```go
func HowManySolutions(size int) int
```
It creates a solver of the given size and goes through all possible soltions for this size. It returns that number of solutions.
### HowManySolutionsConcurrent
```go
func HowManySolutionsConcurrent(size int) int
```
Functionaly equivalent to the previous one, but running in multiple threads. It creates as many threads as the number of CPU's available minus one (to ensure one CPU if left for OS and other tasks, and thus not loosing control of the machine).
Internally it starts a solver for each possible position of the first queen, limiting the number of workers.
## main
Also in adition in the root folder there is a main implementation that endessly computes how many solutions are for each board using the concurrent implementation.
Example:
```
Z:\Code\Go\NQueens>nqueens
Size 1: 1 solutions.
Size 2: 0 solutions.
Size 3: 0 solutions.
Size 4: 2 solutions.
Size 5: 10 solutions.
Size 6: 4 solutions.
Size 7: 40 solutions.
Size 8: 92 solutions.
Size 9: 352 solutions.
Size 10: 724 solutions.
Size 11: 2680 solutions.
Size 12: 14200 solutions.
Size 13: 73712 solutions.
Size 14: 365596 solutions.
Size 15: 2279184 solutions.
```
Until size 14 it's very fast. Starting from size 15 the required time grows exponentially. Results are the same as the ones published in <https://en.wikipedia.org/wiki/Eight_queens_puzzle#Counting_solutions>