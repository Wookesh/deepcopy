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

func CloneList_bool(l []bool) []bool {
	ret := make([]bool, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_byte(l []byte) []byte {
	ret := make([]byte, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_complex128(l []complex128) []complex128 {
	ret := make([]complex128, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_complex64(l []complex64) []complex64 {
	ret := make([]complex64, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_error(l []error) []error {
	ret := make([]error, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_float32(l []float32) []float32 {
	ret := make([]float32, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_float64(l []float64) []float64 {
	ret := make([]float64, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_int(l []int) []int {
	ret := make([]int, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_int16(l []int16) []int16 {
	ret := make([]int16, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_int32(l []int32) []int32 {
	ret := make([]int32, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_int64(l []int64) []int64 {
	ret := make([]int64, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_int8(l []int8) []int8 {
	ret := make([]int8, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_rune(l []rune) []rune {
	ret := make([]rune, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_string(l []string) []string {
	ret := make([]string, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_uint(l []uint) []uint {
	ret := make([]uint, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_uint16(l []uint16) []uint16 {
	ret := make([]uint16, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_uint32(l []uint32) []uint32 {
	ret := make([]uint32, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_uint64(l []uint64) []uint64 {
	ret := make([]uint64, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_uint8(l []uint8) []uint8 {
	ret := make([]uint8, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}

func CloneList_uintptr(l []uintptr) []uintptr {
	ret := make([]uintptr, 0, len(l))
	for _, o := range l {
		ret = append(ret, o)
	}
	return ret
}
