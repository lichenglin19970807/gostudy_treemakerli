// Package mock_study show my study of go mock.
package mock_study

type MockService interface {
	MethodOne(i int) (string, error)
}
