package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("setup")

	res := m.Run()

	fmt.Println("tearDown")

	os.Exit(res)
}

func TestMultiple(t *testing.T) {
	t.Log("simple")
	t.Run("simple", func(t *testing.T) {
		t.Parallel()
		var x, y, excepted = 2, 2, 4

		actual := Multiple(x, y)

		if actual != excepted {
			t.Errorf("%d != %d", excepted, actual)
		}
	})

	t.Log("medium")
	t.Run("medium", func(t *testing.T) {
		t.Parallel()
		var x, y, excepted = 222, 222, 49284

		actual := Multiple(x, y)

		if actual != excepted {
			t.Errorf("%d != %d", excepted, actual)
		}
	})

	t.Log("negative")
	t.Run("negative", func(t *testing.T) {
		t.Parallel()
		var x, y, excepted = 2, -2, -4

		actual := Multiple(x, y)

		if actual != excepted {
			t.Errorf("%d != %d", excepted, actual)
		}
	})
}
