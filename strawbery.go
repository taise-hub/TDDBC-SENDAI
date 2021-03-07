package strawbery

import "errors"

type Strawbery struct {
	kind string
	size string
}

func New(kind string, weight uint) (*Strawbery, error) {
	if weight == 0 {
		return nil, errors.New("重さには1以上の整数を入力してください")
	}
	return &Strawbery{
		kind: kind,
		size: "L",
	}, nil
}

func (berry *Strawbery) String() string {
	return berry.kind + ": " + berry.size
}

func (berry *Strawbery) Size() string {
	return "S"
}
