

package main

import "testing"

type testpair struct {
  values uint
  fibo   uint
}

var tests = []testpair{
        {1, 1},
	{2, 1},
	{3, 2},
	{4, 3},
	{5, 5},
	{6, 8},
	{7, 13},
}

func Testfib(t *testing.T) {
  for _, pair := range tests {
    v := fib(pair.values)
    if v != pair.fibo {
      t.Error(
        "For", pair.values,
        "expected", pair.fibo,
        "got", v,
      )
    }
  }
}

