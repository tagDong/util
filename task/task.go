package task

import (
	"github.com/yddeng/utils"
)

type Task interface {
	Do() (result []interface{}, err error)
}

type FuncTask struct {
	fn   interface{}
	args []interface{}
}

func NewFuncTask(f interface{}, args ...interface{}) *FuncTask {
	return &FuncTask{
		fn:   f,
		args: args,
	}
}

func (this *FuncTask) Do() (result []interface{}, err error) {
	return utils.CallFunc(this.fn, this.args...)
}
