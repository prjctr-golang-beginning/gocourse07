package main

import (
	"concurrency/auditor"
	"context"
	"time"
)

type myRepository struct {
}

type Value string

func (v Value) Value() string {
	return string(v)
}

func (r *myRepository) CreateMany(ctx context.Context, enities []auditor.Valuable) (int, error) {
	return 0, nil
}

func main() {
	a := auditor.New(new(myRepository))
	val1 := Value(`Some value 1`)
	val2 := Value(`Some value 2`)

	a.Update(val1)
	a.Update(val2)

	time.Sleep(time.Second * 5)

	a.Update(val1)
	a.Update(val2)

	time.Sleep(time.Second * 5)
}
