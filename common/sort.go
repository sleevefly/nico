package common

type Sort interface {
	int64 | int | int32 | uint32
}

func Order[T Sort](args []T) T {

	return args[0]
}
