package references

import (
	"github.com/zxh0/jvm.go/instructions/base"
	"github.com/zxh0/jvm.go/rtda"
)

// Throw exception or error
type AThrow struct{ base.NoOperandsInstruction }

func (instr *AThrow) Execute(frame *rtda.Frame) {
	ex := frame.PopRef()
	if ex == nil {
		frame.Thread().ThrowNPE()
		return
	}

	thread := frame.Thread()
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC() - 1

		handler := frame.Method().FindExceptionHandler(ex.Class, pc)
		if handler != nil {
			frame.ClearStack()
			frame.PushRef(ex)
			frame.SetNextPC(handler.HandlerPc)
			return
		}

		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}

	thread.HandleUncaughtException(ex)
}
