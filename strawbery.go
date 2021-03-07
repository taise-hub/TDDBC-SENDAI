package strawbery

import "errors"

type Strawbery struct {
	kind   string
	size   string
	weight uint
}

func New(kind string, weight uint) (*Strawbery, error) {

	size, err := calcSize(weight)
	if err != nil {
		return nil, err
	}
	return &Strawbery{
		kind:   kind,
		size:   size,
		weight: weight,
	}, nil
}

func (berry *Strawbery) String() string {
	return berry.kind + ": " + berry.size
}

func (berry *Strawbery) Size() string {
	return berry.size
}

func calcSize(weight uint) (string, error) {
	switch {
	case weight >= 25:
		return "LL", nil
	case (weight >= 20 && weight <= 24):
		return "L", nil
	case (weight >= 10 && weight <= 19):
		return "M", nil
	case (weight >= 1 && weight <= 9):
		return "S", nil
	default:
		return "", errors.New("重さには1以上の整数を入力してください")
	}
}
