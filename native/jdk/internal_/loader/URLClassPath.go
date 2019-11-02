package loader

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
)

func init() {
	_urlcp(getLookupCacheURLs, "getLookupCacheURLs", "(Ljava/lang/ClassLoader;)[Ljava/net/URL;")
}

func _urlcp(method func(frame *rtda.Frame), name, desc string) {
	heap.RegisterNativeMethod("jdk/internal/loader/URLClassPath", name, desc, method)
}

// private static native URL[] getLookupCacheURLs(ClassLoader var0);
// (Ljava/lang/ClassLoader;)[Ljava/net/URL;
func getLookupCacheURLs(frame *rtda.Frame) {
	frame.PushNull()
}