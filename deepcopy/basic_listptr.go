package deepcopy

func CloneListPtr_bool(l []*bool) []*bool {
	ret := make([]*bool, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_bool(o))
	}
	return ret
}

func CloneListPtr_byte(l []*byte) []*byte {
	ret := make([]*byte, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_byte(o))
	}
	return ret
}

func CloneListPtr_complex128(l []*complex128) []*complex128 {
	ret := make([]*complex128, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_complex128(o))
	}
	return ret
}

func CloneListPtr_complex64(l []*complex64) []*complex64 {
	ret := make([]*complex64, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_complex64(o))
	}
	return ret
}

func CloneListPtr_error(l []*error) []*error {
	ret := make([]*error, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_error(o))
	}
	return ret
}

func CloneListPtr_float32(l []*float32) []*float32 {
	ret := make([]*float32, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_float32(o))
	}
	return ret
}

func CloneListPtr_float64(l []*float64) []*float64 {
	ret := make([]*float64, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_float64(o))
	}
	return ret
}

func CloneListPtr_int(l []*int) []*int {
	ret := make([]*int, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_int(o))
	}
	return ret
}

func CloneListPtr_int16(l []*int16) []*int16 {
	ret := make([]*int16, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_int16(o))
	}
	return ret
}

func CloneListPtr_int32(l []*int32) []*int32 {
	ret := make([]*int32, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_int32(o))
	}
	return ret
}

func CloneListPtr_int64(l []*int64) []*int64 {
	ret := make([]*int64, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_int64(o))
	}
	return ret
}

func CloneListPtr_int8(l []*int8) []*int8 {
	ret := make([]*int8, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_int8(o))
	}
	return ret
}

func CloneListPtr_rune(l []*rune) []*rune {
	ret := make([]*rune, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_rune(o))
	}
	return ret
}

func CloneListPtr_string(l []*string) []*string {
	ret := make([]*string, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_string(o))
	}
	return ret
}

func CloneListPtr_uint(l []*uint) []*uint {
	ret := make([]*uint, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_uint(o))
	}
	return ret
}

func CloneListPtr_uint16(l []*uint16) []*uint16 {
	ret := make([]*uint16, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_uint16(o))
	}
	return ret
}

func CloneListPtr_uint32(l []*uint32) []*uint32 {
	ret := make([]*uint32, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_uint32(o))
	}
	return ret
}

func CloneListPtr_uint64(l []*uint64) []*uint64 {
	ret := make([]*uint64, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_uint64(o))
	}
	return ret
}

func CloneListPtr_uint8(l []*uint8) []*uint8 {
	ret := make([]*uint8, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_uint8(o))
	}
	return ret
}

func CloneListPtr_uintptr(l []*uintptr) []*uintptr {
	ret := make([]*uintptr, 0, len(l))
	for _, o := range l {
		ret = append(ret, Clone_uintptr(o))
	}
	return ret
}
