package strawbery

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_StrawberyString(t *testing.T) {
	type args struct {
		kind string
		size string
	}
	tests := map[string]struct {
		args     args
		expected string
	}{
		"品種:あまおう、サイズ:Lのいちごから文字列表現「あまおう: L」を取得できる": {
			args:     args{kind: "あまおう", size: "L"},
			expected: "あまおう: L",
		},
		"品種:あまおう、サイズ:LLのいちごから文字列表現「あまおう: LL」を取得できる": {
			args:     args{kind: "あまおう", size: "LL"},
			expected: "あまおう: LL",
		},
		"品種:とちおとめ、サイズ:Lのいちごから文字列表現「とちおとめ: L」を取得できる": {
			args:     args{kind: "とちおとめ", size: "L"},
			expected: "とちおとめ: L",
		},
		"品種:もういっこ、サイズMのいちごから文字列表現「もういっこ: M」を取得できる": {
			args:     args{kind: "もういっこ", size: "M"},
			expected: "もういっこ: M",
		},
		"品種:もういっこ、サイズMのいちごから文字列表現「もういっこ: S」を取得できる": {
			args:     args{kind: "もういっこ", size: "S"},
			expected: "もういっこ: S",
		},
	}

	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			sut := Strawbery{
				kind: test.args.kind,
				size: test.args.size,
			}
			actual := sut.String()

			assert.Equal(t, test.expected, actual)
		})
	}

}

func Test_StrawberyWeight(t *testing.T) {
	t.Run("重さが0gの時エラー", func(t *testing.T) {
		_, err := New("あまおう", 0)
		assert.NotEmpty(t, err)
	})

	type args struct {
		kind   string
		weight uint
	}
	tests := map[string]struct {
		args     args
		expected string
	}{
		"重さが1gの時サイズはS": {
			args:     args{kind: "あまおう", weight: 1},
			expected: "S",
		},
		"重さが9gの時サイズはS": {
			args:     args{kind: "あまおう", weight: 9},
			expected: "S",
		},
		"重さが10gの時サイズはM": {
			args:     args{kind: "あまおう", weight: 10},
			expected: "M",
		},
		"重さが19gの時サイズはM": {
			args:     args{kind: "あまおう", weight: 19},
			expected: "M",
		},
		"重さが20gの時サイズはL": {
			args:     args{kind: "あまおう", weight: 20},
			expected: "L",
		},
		"重さが24gの時サイズはL": {
			args:     args{kind: "あまおう", weight: 24},
			expected: "L",
		},
		"重さが25gの時サイズはLL": {
			args:     args{kind: "あまおう", weight: 25},
			expected: "LL",
		},
	}
	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			sut, _ := New(test.args.kind, test.args.weight)
			assert.Equal(t, test.expected, sut.Size())
		})
	}
}

func Test_StrawberyCompareKind(t *testing.T) {
	type args struct {
		kind1   string
		kind2   string
		weight1 uint
		weight2 uint
	}
	tests := map[string]struct {
		args     args
		expected bool
	}{
		"あまおうとあまおうを比較した時にtrueを返す": {
			args: args{
				kind1:   "あまおう",
				kind2:   "あまおう",
				weight1: 1,
				weight2: 25,
			},
			expected: true,
		},
		"とちおとめととちおとめを比較した時にtrueを返す": {
			args: args{
				kind1:   "とちおとめ",
				kind2:   "とちおとめ",
				weight1: 1,
				weight2: 25,
			},
			expected: true,
		},
		"もういっこともういっこを比較した時にtrueを返す": {
			args: args{
				kind1:   "もういっこ",
				kind2:   "もういっこ",
				weight1: 1,
				weight2: 25,
			},
			expected: true,
		},
		"あまおうととちおとめを比較した時にfalseを返す": {
			args: args{
				kind1:   "あまおう",
				kind2:   "とちおとめ",
				weight1: 1,
				weight2: 25,
			},
			expected: false,
		},
		"とちおとめとあまおうを比較した時にfalseを返す": {
			args: args{
				kind1:   "とちおとめ",
				kind2:   "あまおう",
				weight1: 1,
				weight2: 25,
			},
			expected: false,
		},
		"とちおとめともういっこを比較した時にfalseを返す": {
			args: args{
				kind1:   "とちおとめ",
				kind2:   "もういっこ",
				weight1: 1,
				weight2: 25,
			},
			expected: false,
		},
	}

	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			berry1, _ := New(test.args.kind1, test.args.weight1)
			berry2, _ := New(test.args.kind2, test.args.weight2)
			actual := CompareKind(berry1, berry2)
			assert.Equal(t, test.expected, actual)
		})
	}
}

func Test_StrawberyCompareSize(t *testing.T) {
	type args struct {
		weight1 uint
		weight2 uint
	}
	tests := map[string]struct {
		args     args
		expected uint
	}{
		"LとLを比較するとuintの0を返す": {
			args: args{
				weight1: 24,
				weight2: 20,
			},
			expected: uint(0),
		},
		"LとLLを比較するとuintの1を返す": {
			args: args{
				weight1: 20,
				weight2: 30,
			},
			expected: uint(1),
		},
		"LLとLを比較するとuintの1を返す": {
			args: args{
				weight1: 30,
				weight2: 20,
			},
			expected: uint(1),
		},
		"LLとSを比較するとuintの3を返す": {
			args: args{
				weight1: 30,
				weight2: 9,
			},
			expected: uint(3),
		},
		"MとSを比較するとuintの1を返す": {
			args: args{
				weight1: 19,
				weight2: 9,
			},
			expected: uint(1),
		},
	}
	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			berry1, _ := New("とちおとめ", test.args.weight1)
			berry2, _ := New("とちおとめ", test.args.weight2)
			actual := CompareSize(berry1, berry2)
			assert.Equal(t, test.expected, actual)
		})
	}

}

func Test_strawberysAlgne(t *testing.T) {
	type args struct {
		kind1 string
		kind2 string
		kind3 string
	}

	tests := map[string]struct {
		args     args
		expected bool
	}{
		"あまおう,あまおう,あまおうからなるパックは純正": {
			args:     args{kind1: "あまおう", kind2: "あまおう", kind3: "あまおう"},
			expected: true,
		},
		"とちおとめ,とちおとめ,とちおとめからなるパックは純正": {
			args:     args{kind1: "とちおとめ", kind2: "とちおとめ", kind3: "とちおとめ"},
			expected: true,
		},
		"あまおう,あまおう,とちおとめからなるパックは純正でない": {
			args:     args{kind1: "あまおう", kind2: "あまおう", kind3: "とちおとめ"},
			expected: false,
		},
		"あまおう,もういっこ,あまおうからなるパックは純正でない": {
			args:     args{kind1: "あまおう", kind2: "もういっこ", kind3: "あまおう"},
			expected: false,
		},
		"あまおう,とちおとめ,もういっこからなるパックは純正でない": {
			args:     args{kind1: "あまおう", kind2: "とちおとめ", kind3: "もういっこ"},
			expected: false,
		},
		"もういっこ,あまおう,とちおとめからなるパックは純正でない": {
			args:     args{kind1: "もういっこ", kind2: "あまおう", kind3: "とちおとめ"},
			expected: false,
		},
	}

	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			berry1, _ := New(test.args.kind1, 1)
			berry2, _ := New(test.args.kind2, 10)
			berry3, _ := New(test.args.kind3, 20)
			strawverys := []*Strawbery{
				berry1, berry2, berry3,
			}

			actual := IsAligned(strawverys)
			assert.Equal(t, test.expected, actual)
		})
	}

}

func Test_StrawberysMinSize(t *testing.T) {
	type args struct {
		size1 uint
		size2 uint
		size3 uint
	}
	tests := map[string]struct {
		args     args
		expected string
	}{
		"あまおう L,とちおとめ L,もういっこ L からなるパック内の最小サイズはL": {
			args: args{
				size1: 20, size2: 20, size3: 20,
			},
			expected: "L",
		},
		"あまおう S,とちおとめ M,もういっこ LL からなるパック内の最小サイズはS": {
			args: args{
				size1: 1, size2: 10, size3: 25,
			},
			expected: "S",
		},
	}

	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			berry1, _ := New("あまおう", test.args.size1)
			berry2, _ := New("とちおとめ", test.args.size2)
			berry3, _ := New("もういっこ", test.args.size3)
			//いちごのパック詰めを表す構造体を作った方がいいのか、配列でいいのか
			strawverys := []*Strawbery{
				berry1, berry2, berry3,
			}
			actual := GetMinSize(strawverys)
			assert.Equal(t, test.expected, actual.size)
		})
	}

}

func Test_StrawberysMaxSize(t *testing.T) {
	type args struct {
		size1 uint
		size2 uint
		size3 uint
	}
	tests := map[string]struct {
		args     args
		expected string
	}{
		"あまおう L,とちおとめ L,もういっこ L からなるパック内の最大サイズはL": {
			args: args{
				size1: 20, size2: 20, size3: 20,
			},
			expected: "L",
		},
		"あまおう S,とちおとめ M,もういっこ LL からなるパック内の最大サイズはLL": {
			args: args{
				size1: 1, size2: 10, size3: 25,
			},
			expected: "LL",
		},
	}

	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			berry1, _ := New("あまおう", test.args.size1)
			berry2, _ := New("とちおとめ", test.args.size2)
			berry3, _ := New("もういっこ", test.args.size3)
			//いちごのパック詰めを表す構造体を作った方がいいのか、配列でいいのか
			strawberys := []*Strawbery{
				berry1, berry2, berry3,
			}
			actual := GetMaxSize(strawberys)
			assert.Equal(t, test.expected, actual.size)
		})
	}

}

func Test_StrawberysDiffSize(t *testing.T) {
	type args struct {
		size1 uint
		size2 uint
		size3 uint
	}
	tests := map[string]struct {
		args     args
		expected uint
	}{
		"あまおう L とちおとめ L もういっこ L からなるパックの最小サイズと最大サイズの差は0": {
			args: args{
				size1: 20, size2: 20, size3: 20,
			},
			expected: uint(0),
		},
		"あまおう S とちおとめ M もういっこ L からなるパックの最小サイズと最大サイズの差は2": {
			args: args{
				size1: 1, size2: 10, size3: 24,
			},
			expected: uint(2),
		},
		"あまおう L とちおとめ M もういっこ L からなるパックの最小サイズと最大サイズの差は1": {
			args: args{
				size1: 24, size2: 10, size3: 24,
			},
			expected: uint(1),
		},
	}

	for tName, test := range tests {
		t.Run(tName, func(t *testing.T) {
			berry1, _ := New("あまおう", test.args.size1)
			berry2, _ := New("とちおとめ", test.args.size2)
			berry3, _ := New("もういっこ", test.args.size3)
			strawberys := []*Strawbery{
				berry1, berry2, berry3,
			}
			actual := GetDiffSize(strawberys)
			assert.Equal(t, test.expected, actual)
		})
	}
}
