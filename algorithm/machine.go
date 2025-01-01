package algorithm

type UniMachine struct {
}

type ParallelMachine struct {
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
