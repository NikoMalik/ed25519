package ed25519

import (
	"unsafe"

	"github.com/NikoMalik/low-level-functions/constants"
	"github.com/NikoMalik/mutex"
)

type objPool[T any] struct {
	mut      *mutex.MutexExp // lock mutex
	_        [constants.CacheLinePadSize - 8]byte
	obj      []T
	_        [constants.CacheLinePadSize - unsafe.Sizeof([]T{})]byte
	allocate func() T
	_        [constants.CacheLinePadSize - unsafe.Sizeof(func() T { var z T; return z })]byte
	size     int
	_        [constants.CacheLinePadSize - unsafe.Sizeof(int(0))]byte
}

func nObjPool[T any](capt int, f func() T) *objPool[T] {

	pool := &objPool[T]{allocate: f}

	if capt < 0 {
		return nil
	}

	pool.size = capt
	pool.obj = make([]T, 0, capt)

	return pool
}

func (o *objPool[T]) Get() T {
	if len(o.obj) <= 0 {
		return o.allocate()

	}
	if o.mut == nil {

		obj := o.obj[len(o.obj)-1]
		o.obj = o.obj[:len(o.obj)-1]
		return obj

	}
	if o.mut != nil {
		o.mut.Lock()
		obj := o.obj[len(o.obj)-1]
		o.obj = o.obj[:len(o.obj)-1]

		o.mut.Unlock()
		return obj
	}
	var zero T
	return zero

}

func (o *objPool[T]) Put(obj T) {
	if o.mut != nil {
		o.mut.Lock()
		if len(o.obj) < o.size {

			o.obj = append(o.obj, obj)
		}

		o.mut.Unlock()
	}
	if o.mut == nil {

		if len(o.obj) < o.size {

			o.obj = append(o.obj, obj)
		}
	}
}
