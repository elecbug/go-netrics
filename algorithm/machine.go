package algorithm

type ComputeMachine interface {
	ShortestPath()
	Diameter()
}

type UniMachine struct {
	ComputeMachine
}

type ParallelMachine struct {
	ComputeMachine
	maxCore uint
}

func NewUniMachine() *UniMachine {
	return &UniMachine{}
}

func NewParallelMachine(core uint) *ParallelMachine {
	return &ParallelMachine{
		maxCore: core,
	}
}
