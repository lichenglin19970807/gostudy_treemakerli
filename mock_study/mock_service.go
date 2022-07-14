// Package mock_study show my study of go mock.
package mock_study

// mockgen -destination:"mockgen生成的文件存放的位置以及文件的名字" -package:"mocks" modules/package interface
type MockService interface {
	MethodOne(i int) (string, error)
}
