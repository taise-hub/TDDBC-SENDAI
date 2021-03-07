package strawbery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StrawveryString(t *testing.T) {
	t.Run("品種:あまおう、サイズ:Lのいちごから文字列表現「あまおう: L」を取得できる", func(t *testing.T) {
		sut := Strawbery{
			kind: "あまおう",
			size: "L",
		}
		actual := sut.String()

		assert.Equal(t, "あまおう: L", actual)
	})

	t.Run("品種:もういっこ、サイズMのいちごから文字列表現「もういっこ: M」を取得できる", func(t *testing.T) {
		sut := Strawbery{
			kind: "もういっこ",
			size: "M",
		}
		actual := sut.String()

		assert.Equal(t, "もういっこ: M", actual)
	})
}
