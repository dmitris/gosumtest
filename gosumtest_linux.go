package main

import "golang.org/x/sys/unix"

func myMajor(dev uint64) uint32 {
	return unix.Major(dev)
}
