package algorithm

type ParallelMachine struct {
	maxCore uint
}

func NewParallelMachine(core uint) *ParallelMachine {
	return &ParallelMachine{
		maxCore: core,
	}
}
