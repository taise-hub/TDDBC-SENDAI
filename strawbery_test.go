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

func Test_品種あまおうサイズLのいちごから文字列表現を取得できる(t *testing.T) {
	sut := Strawbery{
		kind: "あまおう",
		size: "L",
	}
	actual := sut.String()

	assert.Equal(t, "あまおう: L", actual)
}

func Test_品種もういっこサイズMのいちごから文字列表現を取得できる(t *testing.T) {
	sut := Strawbery{
		kind: "もういっこ",
		size: "M",
	}
	actual := sut.String()

	assert.Equal(t, "もういっこ: M", actual)
}
