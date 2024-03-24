package hw1

// Shrink 调整数组
func Shrink[T any](src []T) []T {
	c, l := cap(src), len(src)
	if c <= 64 {
		return src
	}

	if c > 2048 && c/l >= 2 {
		c = int(float64(c) / 1.5)
	} else if c <= 2048 && c/l >= 4 {
		c = c / 2
	} else {

	}

	ts := make([]T, 0, c)

	ts = append(ts, src...)

	return ts
}
