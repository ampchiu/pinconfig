package main

import "fmt"

var (
	bound [4]int
	dx    [4]int
	dy    [4]int
	dim   [3]int
)

func pos(pin int) (float32, float32) {
	var x, y float32
	if pin < bound[0] {
		x = float32(600 - dx[0])
		y = float32(500 - dy[0] + pin*dim[0])
	} else if pin < bound[1] {
		x = float32(600 - dx[1] + (pin-bound[0])*dim[0])
		y = float32(500 + dy[1])
	} else if pin < bound[2] {
		x = float32(600 + dx[2])
		y = float32(500 + dy[2] - (pin-bound[1]+1)*dim[0])
	} else {
		x = float32(600 + dx[3] - (pin-bound[2]+1)*dim[0])
		y = float32(500 - dy[3])
	}
	return x, y
}

func size(pin int) (float32, float32) {
	var w, l float32
	if pin < bound[0] {
		w = float32(dim[2])
		l = float32(dim[1])
	} else if pin < bound[1] {
		w = float32(dim[1])
		l = float32(dim[2])
	} else if pin < bound[2] {
		w = float32(dim[2])
		l = float32(dim[1])
	} else {
		w = float32(dim[1])
		l = float32(dim[2])
	}
	return w, l
}

func vert(st string) string {
	var result string
	for i, ch := range st {
		result = fmt.Sprintf("%s%c\n", result, ch)
		_ = i
	}
	return (result)
}

func name(pos int) string {
	if pos < bound[0] {
		return (fmt.Sprintf("pin%v", pos+1))
	} else if pos < bound[1] {
		return (vert(fmt.Sprintf("pin%v", pos+1)))
	} else if pos < bound[2] {
		return (fmt.Sprintf("pin%v", pos+1))
	} else {
		return (vert(fmt.Sprintf("pin%v", pos+1)))
	}
}
