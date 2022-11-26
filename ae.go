package main

import "syscall"

type FileProc func(loop *AeLoop, fd int, extra interface{})
type TimeProc func(loop *AeLoop, fd int, extra interface{})

type FeType int
type TeType int

// File event types
const (
	AE_READABLE FeType = 1
	AE_WRITABLE FeType = 2
)

// Time event types
const (
	AE_NORMAL TeType = 1
	AE_ONCE   TeType = 2
)

// AeFileEvent File event
type AeFileEvent struct {
	fd    int
	mask  FeType
	proc  FileProc
	extra interface{}
}

// AeTimeEvent Time event
type AeTimeEvent struct {
	fd       int
	mask     TeType
	when     int64
	interval int64
	proc     TimeProc
	extra    interface{}
	next     *AeTimeEvent
}

// AeLoop Event loop
type AeLoop struct {
	FileEvents      map[int]*AeFileEvent
	TimeEvents      *AeTimeEvent
	fileEventFd     int
	timeEventNextId int
	stop            bool
}

// File event to epoll
var fe2e = []uint32{0, syscall.EPOLLIN, syscall.EPOLLOUT}

// Get the file event key
func getFeKey(fd int, mask FeType) int {
	if mask == AE_READABLE {
		return fd
	} else {
		return -1 * fd
	}
}
