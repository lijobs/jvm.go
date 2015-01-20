package instructions

import (
    . "jvmgo/any"
    "jvmgo/rtda"
)

type Instruction interface {
    fetchOperands(bcr *BytecodeReader)
    execute(thread *rtda.Thread)
}

type NoOperandsInstruction struct {
    // empty
}
func (self *NoOperandsInstruction) fetchOperands(bcr *BytecodeReader) {
    // nothing to do
}

type BranchInstruction struct {
    branch int16
}
func (self *BranchInstruction) fetchOperands(bcr *BytecodeReader) {
    self.branch = bcr.readInt16()
}

type Index8Instruction struct {
    index uint
}
func (self *Index8Instruction) fetchOperands(bcr *BytecodeReader) {
    self.index = uint(bcr.readUint8())
}

type Index16Instruction struct {
    index uint
}
func (self *Index16Instruction) fetchOperands(bcr *BytecodeReader) {
    self.index = uint(bcr.readUint16())
}

// todo: move to any.go?
func isLongOrDouble(x Any) (bool) {
    switch x.(type) {
    case int64: return true
    case float64: return true
    default: return false
    }
}
