package strawbery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_品種とサイズを与えていちごを作成する(t *testing.T) {

	actual := Strawbery{
		kind: "あまおう",
		size: "L",
	}
	assert.NotEmpty(t, actual)
}
