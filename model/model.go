package model

type Model struct {
	table string
}

type Dao interface {
	delete(id []int64)
	pause(id []int64, status int)
}
