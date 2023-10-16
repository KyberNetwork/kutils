package kutils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	sep              = ","
	splitElem        = strings.Repeat("0", 40)
	splittableElem   = splitElem + sep + splitElem + sep + splitElem
	strSplitInMiddle = []string{splitElem, splitElem, splittableElem, splitElem, splittableElem}
	strNoSplit       = []string{splitElem, splitElem, splitElem, splitElem, splitElem}
	strAllSplittable = []string{splittableElem, splittableElem, splittableElem, splittableElem}
)

// BenchmarkSplitListElem/split_in_middle
// BenchmarkSplitListElem/split_in_middle-16                3464385               3
// 39.5 ns/op vs 51.0 ns/op using strings.Contains
// BenchmarkSplitListElem/no_split
// BenchmarkSplitListElem/no_split-16                       5432838               2
// 22.2 ns/op vs 67.71 ns/op using strings.Contains
// BenchmarkSplitListElem/all_splittable
// BenchmarkSplitListElem/all_splittable-16                 2987763               3
// 95.3 ns/op vs 16.8 ns/op using strings.Contains
func BenchmarkSplitListElem(b *testing.B) {
	b.ResetTimer()
	b.Run("split in middle", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SplitListElem(strSplitInMiddle, sep)
		}
	})
	b.Run("no split", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SplitListElem(strNoSplit, sep)
		}
	})
	b.Run("all splittable", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SplitListElem(strAllSplittable, sep)
		}
	})
}

func TestSplitListElem(t *testing.T) {
	type args struct {
		lst []string
		sep string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"split in middle",
			args{
				lst: strSplitInMiddle,
				sep: sep,
			},
			[]string{splitElem, splitElem, splitElem, splitElem, splitElem, splitElem, splitElem, splitElem, splitElem},
		},
		{
			"no split",
			args{
				lst: strNoSplit,
				sep: sep,
			},
			[]string{splitElem, splitElem, splitElem, splitElem, splitElem},
		},
		{
			"split in middle",
			args{
				lst: strAllSplittable,
				sep: sep,
			},
			[]string{splitElem, splitElem, splitElem, splitElem, splitElem, splitElem, splitElem, splitElem, splitElem,
				splitElem, splitElem, splitElem},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, SplitListElem(tt.args.lst, tt.args.sep), "SplitListElem(%v, %v)", tt.args.lst,
				tt.args.sep)
		})
	}
}
