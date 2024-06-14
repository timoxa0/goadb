package adb

import (
	"strings"

	"github.com/timoxa0/goadb/internal/errors"
)

type ForwardType int8

const (
	TypeInvalid ForwardType = iota
	ForwardTCP
	ForwardLOCAL
)

var forwardTypeStrings = map[string]ForwardType{
	"tcp":   ForwardTCP,
	"local": ForwardLOCAL,
}

type ForwardTarget string

type ForwardLocal struct {
	FType   ForwardType
	FTarget ForwardTarget
}

type ForwardRemote struct {
	FType   ForwardType
	FTarget ForwardTarget
}

type Forward struct {
	Local  ForwardLocal
	Remote ForwardRemote
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
				Local: ForwardLocal{
					FType:   forwardTypeStrings[raw_local[0]],
					FTarget: ForwardTarget(raw_local[1]),
				},
				Remote: ForwardRemote{
					FType:   forwardTypeStrings[raw_remote[0]],
					FTarget: ForwardTarget(raw_remote[1]),
				},
			})
		}
	}
	return forwards, errors.Errorf(errors.ParseError, "invalid device forward")
}
