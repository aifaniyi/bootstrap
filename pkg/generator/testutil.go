package generator

import "testing"

func fail(t *testing.T, expected, actual interface{}) {
	t.Fatalf("\nexpected:\t%v\nfound:\t\t%v", expected, actual)
}
