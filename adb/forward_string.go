// Code generated by "stringer -type=ForwardType"; DO NOT EDIT

package adb

import "fmt"

const _ForwardType_name = "InvalidTCPLocal"

var _ForwardType_index = [...]uint8{0, 7, 10, 15}

func (i ForwardType) String() string {
	if i < 0 || i >= ForwardType(len(_ForwardType_index)-1) {
		return fmt.Sprintf("ForwardType(%d)", i)
	}
	return _ForwardType_name[_ForwardType_index[i]:_ForwardType_index[i+1]]
}