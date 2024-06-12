package adb

import (
	"strings"

	"github.com/timoxa0/goadb/internal/errors"
)

type ForwardType uint8

const (
	TypeInvalid ForwardType = iota
	TCP
	LOCAL
)

var forwardTypeStrings = map[string]ForwardType{
	"tcp":   TCP,
	"local": LOCAL,
}

type ForwardTarget string

type ForwardLocal struct {
	ftype  ForwardType
	target ForwardTarget
}

type ForwardRemote struct {
	ftype  ForwardType
	target ForwardTarget
}

type Forward struct {
	local  ForwardLocal
	remote ForwardRemote
}

func parseForward(str string, deviceSerial string) ([]Forward, error) {
	var forwards []Forward
	for _, forwardStr := range strings.Split(str, "\n") {
		if strings.Trim(forwardStr, "\n\t ") == "" {
			continue
		}
		raw_forward := strings.Split(forwardStr, " ")
		serial, local, remote := raw_forward[0], raw_forward[1], raw_forward[2]
		if serial == deviceSerial {
			raw_local := strings.Split(local, ":")
			raw_remote := strings.Split(remote, ":")
			forwards = append(forwards, Forward{
				local: ForwardLocal{
					ftype:  forwardTypeStrings[raw_local[0]],
					target: ForwardTarget(raw_local[1]),
				},
				remote: ForwardRemote{
					ftype:  forwardTypeStrings[raw_remote[0]],
					target: ForwardTarget(raw_remote[1]),
				},
			})
		}
	}
	return forwards, errors.Errorf(errors.ParseError, "invalid device forward")
}
