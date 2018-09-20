package nec

type Job struct {
	Path  string
	OnErr string
	*Command
}
