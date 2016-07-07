package parallel_test

import (
	"fmt"
	"github.com/wayt/parallel"
	"testing"
)

func TestGroup(t *testing.T) {

	g := &parallel.Group{}

	g.Go(func() {
		// Do nothing
	})
	g.Go(func() error {
		return nil
	})

	if err := g.Wait(); err != nil {
		t.Fatalf("error: %s", err)
	}
}

func TestGroupError(t *testing.T) {

	g := &parallel.Group{}

	g.Go(func() {
		// Do nothing
	})
	g.Go(func() error {
		return fmt.Errorf("something bad happened")
	})

	if err := g.Wait(); err == nil {
		t.Fatalf("missing error, should have one !")
	}
}
func TestLargeGroup(t *testing.T) {

	g := &parallel.Group{}

	g.Go(func() {
		// Do nothing
	})
	for i := 0; i < 2048; i++ {
		g.Go(func(a, b int) error {
			return nil
		}, i, 42)
	}

	if err := g.Wait(); err != nil {
		t.Fatalf("error: %s", err)
	}
}

func TestBadFunction(t *testing.T) {

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("bad function did not panic")
		}

	}()

	g := &parallel.Group{}

	g.Go(42)
}
