package strawbery

type Strawbery struct {
	kind string
	size string
}

func (berry *Strawbery) String() string {
	return berry.kind + ": " + berry.size
}
