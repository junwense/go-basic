package hw1

import "gitee.com/SeanJPK/go-basic/basic/hw/hw1/errs"

// Delete 切片删除方法
func Delete[T any](slice []T, index int) ([]T, T, error) {
	var t T
	len := len(slice)
	if index < 0 || index >= len {
		return nil, t, errs.NewErrIndexOutOfRange(len, index)
	}

	t = slice[index]
	for i := index; i < len-1; i++ {
		slice[i] = slice[i+1]
	}

	slice = slice[:len-1]

	return slice, t, nil
}
