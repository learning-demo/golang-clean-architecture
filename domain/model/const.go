package model

import (
	"fmt"
	"strconv"
)

// enumerations

type Season int64

const (
	//Undefined Season = 0
	Summer Season = 1
	Autumn Season = 2
	Winter Season = 3
	Spring Season = 4
)

var (
	SeasonByName = map[string]Season{
		"Season.Summer": Summer,
		"Season.Autumn": Autumn,
		"Season.Winter": Winter,
		"Season.Spring": Spring,
	}
	SeasonByValue = map[Season]string{
		Summer: "Season.Summer",
		Autumn: "Season.Autumn",
		Winter: "Season.Winter",
		Spring: "Season.Spring",
	}
)

func (e Season) String() string {
	name := SeasonByValue[s]
	if name == "" {
		name = fmt.Sprintf("Unknown enum value Season(%d)", e)
	}
	return name
}

func (e *Season) UnmarshalJSON(b []byte) error {
	st := string(b)
	if st[0] == '"' {
		*e = Season(SeasonByName[st[1:len(st)-1]])
		return nil
	}
	i, err := strconv.Atoi(st)
	*e = Season(i)
	return err
}
