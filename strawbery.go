package strawbery

import (
	"errors"
	"math"
)

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

func CompareKind(berry1 *Strawbery, berry2 *Strawbery) bool {
	if berry1.kind != berry2.kind {
		return false
	}
	return true
}

func CompareSize(berry1 *Strawbery, berry2 *Strawbery) uint {
	sizeNum1, err := convertSizeToInt(berry1.size)
	if err != nil {
		return 0
	}
	sizeNum2, err := convertSizeToInt(berry2.size)
	if err != nil {
		return 0
	}

	diff := math.Abs(float64(sizeNum1) - float64(sizeNum2))

	return uint(diff)
}

func convertSizeToInt(size string) (int, error) {
	switch {
	case size == "LL":
		return 4, nil
	case size == "L":
		return 3, nil
	case size == "M":
		return 2, nil
	case size == "S":
		return 1, nil
	default:
		return 0, errors.New("おや,サイズがおかしいよ。")
	}
}
