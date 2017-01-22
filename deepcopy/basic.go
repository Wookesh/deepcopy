package deepcopy

func Clone_bool(o *bool) *bool {
	if o == nil {
		return nil
	}
	ret := bool(*o)
	return &ret
}

func Clone_byte(o *byte) *byte {
	if o == nil {
		return nil
	}
	ret := byte(*o)
	return &ret
}

func Clone_complex128(o *complex128) *complex128 {
	if o == nil {
		return nil
	}
	ret := complex128(*o)
	return &ret
}

func Clone_complex64(o *complex64) *complex64 {
	if o == nil {
		return nil
	}
	ret := complex64(*o)
	return &ret
}

func Clone_error(o *error) *error {
	if o == nil {
		return nil
	}
	ret := error(*o)
	return &ret
}

func Clone_float32(o *float32) *float32 {
	if o == nil {
		return nil
	}
	ret := float32(*o)
	return &ret
}

func Clone_float64(o *float64) *float64 {
	if o == nil {
		return nil
	}
	ret := float64(*o)
	return &ret
}

func Clone_int(o *int) *int {
	if o == nil {
		return nil
	}
	ret := int(*o)
	return &ret
}

func Clone_int16(o *int16) *int16 {
	if o == nil {
		return nil
	}
	ret := int16(*o)
	return &ret
}

func Clone_int32(o *int32) *int32 {
	if o == nil {
		return nil
	}
	ret := int32(*o)
	return &ret
}

func Clone_int64(o *int64) *int64 {
	if o == nil {
		return nil
	}
	ret := int64(*o)
	return &ret
}

func Clone_int8(o *int8) *int8 {
	if o == nil {
		return nil
	}
	ret := int8(*o)
	return &ret
}

func Clone_rune(o *rune) *rune {
	if o == nil {
		return nil
	}
	ret := rune(*o)
	return &ret
}

func Clone_string(o *string) *string {
	if o == nil {
		return nil
	}
	ret := string(*o)
	return &ret
}

func Clone_uint(o *uint) *uint {
	if o == nil {
		return nil
	}
	ret := uint(*o)
	return &ret
}

func Clone_uint16(o *uint16) *uint16 {
	if o == nil {
		return nil
	}
	ret := uint16(*o)
	return &ret
}

func Clone_uint32(o *uint32) *uint32 {
	if o == nil {
		return nil
	}
	ret := uint32(*o)
	return &ret
}

func Clone_uint64(o *uint64) *uint64 {
	if o == nil {
		return nil
	}
	ret := uint64(*o)
	return &ret
}

func Clone_uint8(o *uint8) *uint8 {
	if o == nil {
		return nil
	}
	ret := uint8(*o)
	return &ret
}

func Clone_uintptr(o *uintptr) *uintptr {
	if o == nil {
		return nil
	}
	ret := uintptr(*o)
	return &ret
}
