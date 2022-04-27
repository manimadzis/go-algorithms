package stack

type Stack interface{
	Push(value interface{}) error
	Pop() (interface{}, error)
	Peek() (interface{}, error)
}