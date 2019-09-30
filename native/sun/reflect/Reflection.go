package reflect

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_reflection(getCallerClass, "getCallerClass", "()Ljava/lang/Class;")
	_reflection(getClassAccessFlags, "getClassAccessFlags", "(Ljava/lang/Class;)I")
}

func _reflection(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("sun/reflect/Reflection", name, desc, method)
}

// public static native Class<?> getCallerClass(int i);
// (I)Ljava/lang/Class;
func getCallerClass(frame *rtda.Frame) {
	// top0 is sun/reflect/Reflection
	// top1 is the caller of getCallerClass()
	// top2 is the caller of method
	callerFrame := frame.Thread().TopFrameN(2)
	callerClass := callerFrame.Method().Class.JClass
	frame.PushRef(callerClass)
}

// public static native int getClassAccessFlags(Class<?> type);
// (Ljava/lang/Class;)I
func getClassAccessFlags(frame *rtda.Frame) {
	_type := frame.GetRefVar(0)
	goClass := _type.Extra.(*heap.Class)
	frame.PushInt(int32(goClass.AccessFlags))
}
