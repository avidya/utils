import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap(t *testing.T) {
	sample := []int{2, 31, 5, 17, 21, 13}
	assert := assert.New(t)
	l := Map[int, float64](func(e int) float64 {
		return float64(e) / 2.0
	}, sample)
	fmt.Println(l)
	assert.True(l[1] > 15 && l[1] < 16)
}

func TestFilter(t *testing.T) {
	sample := []int{2, 31, 5, 17, 16, 21, 13}
	assert := assert.New(t)
	l := Filter[int](func(e int) bool {
		return e%2 == 0
	}, sample)
	fmt.Println(l)
	assert.True(len(l) == 2)
}

func TestReduce(t *testing.T) {
	sample1 := []int{2, 31, 5, 17, 21, 13}
	assert := assert.New(t)
	assert.True(Max[int](sample1) == 31)

	assert.True(Reduce[int](func(a1 int, a2 int) int {
		if a1 > a2 {
			return a1
		} else {
			return a2
		}
	}, sample1, 100) == 100)

	sample2 := []float64{2, 31, 5, 3.14159, 17, 21, 13}
	assert.True(Max[float64](sample2) == 31)

	assert.True(Reduce[int](func(a1 int, a2 int) int {
		return a1 + a2
	}, sample1, nil) == 89)
}

func TestSum(t *testing.T) {
	assert.New(t).True(Sum[int32]([]int32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}) == 45)
	assert.New(t).True(Sum[int32]([]int32{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}) == 55)
	assert.New(t).True(Sum[int32]([]int32{1}) == 1)
	assert.New(t).True(Reduce[int](func(a1 int, a2 int) int {
		return a1 + a2
	}, []int{1}, nil) == 1)
}

func TestInList(t *testing.T) {
	assert.New(t).True(InList[int](5, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}))
}

