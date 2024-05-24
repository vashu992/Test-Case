package common

import "fmt"

type WrappedError struct {
	Code    int64
	Message string
	Err     error
}

func (r *WrappedError) Error() string {
	return fmt.Sprintf("Status %d: Message %v: Err %v", r.Code, r.Message, r.Err)
}
