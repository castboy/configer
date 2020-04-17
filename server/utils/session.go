package utils

import (
	"configer/server/constant"
	"github.com/juju/errors"
	"sort"
)

func TrimZero(sess map[int32][]string) map[int32][]string {
	s := map[int32][]string{}

	for weekday := range sess {
		for k := range sess[weekday] {
			if sess[weekday][k] != "00:00-00:00" {
				s[weekday] = append(s[weekday], sess[weekday][k])
			}
		}
	}

	return s
}

const SessionFull int = 3
const DaysWeek = 7

func OrderAndFill(sess map[int32][]string) (map[int32][]string, error) {
	if sess == nil {
		return nil, constant.NewErr(constant.ArgsErr, errors.NotValidf("session, %v", sess))
	}

	s := map[int32][]string{}

	for weekday := range sess {
		for k := range sess[weekday] {
			if sess[weekday][k] != "00:00-00:00" {
				s[weekday] = append(s[weekday], sess[weekday][k])
			}
		}
	}

	for weekday := range s {
		sort.Strings(s[weekday])

		sizeNow := len(s[weekday])

		for i := 0; i < SessionFull-sizeNow; i++ {
			s[weekday] = append(s[weekday], "00:00-00:00")
		}
	}

	for i := 0; i < DaysWeek; i++ {
		if s[int32(i)] == nil {
			s[int32(i)] = []string{"00:00-00:00", "00:00-00:00", "00:00-00:00"}
		}
	}

	return s, nil
}
