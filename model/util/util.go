package util

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type Food int

const (
	Worm Food = iota
	Seed
	Fish
	Rodent
	Berry
	Any
)

func (f Food) String() string {
	return [...]string{"W", "S", "F", "R", "B", "A"}[f]
}

type Nest int // todo
const (
	Platform Nest = iota
	Cup
	Rocks
	Canada
	Star
	None
)

func (n Nest) String() string {
	return [...]string{"Platform", "Cup", "Rocks", "Canada", "Star", "None"}[n]
}

type Region int

const ( // todo right names
	Forest Region = iota
	Grasslands
	Waterlands
)

func (r Region) String() string {
	return [...]string{"Forest", "Grasslands", "Waterlands"}[r]
}

type JoinType int

const (
	And JoinType = iota
	Or
)

func (j JoinType) String() string {
	return [...]string{"And", "Or", ")"}[j]
}

type MoveType int

const (
	PlayBird MoveType = iota
	GainFood
	LayEggs
	DrawCards
	PreGame
)

func (m MoveType) String() string {
	return [...]string{"PlayBird", "GainFood", "LayEggs", "DrawCards", "PreGame"}[m]
}

func Atoi(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

func Combrep(n int, lst []string) [][]string {
	if n == 0 {
		return [][]string{nil}
	}
	if len(lst) == 0 {
		return nil
	}
	r := Combrep(n, lst[1:])
	for _, x := range Combrep(n-1, lst) {
		r = append(r, append(x, lst[0]))
	}
	return r
}

func ParseRegion(s string) []Region {
	r := make([]Region, 0)
	for _, i := range strings.Split(s, "/") {
		r = append(r, Region(Atoi(i)))
	}
	return r
}

type Eggs [][3]int // a list of (region, birdn, #eggs)

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func MapToString(m map[Food]int) string {
	b := new(bytes.Buffer)
	for i := 0; i < 5; i++ {
		f := Food(i)
		fmt.Fprintf(b, "%s=%d,", f.String(), m[f])
	}
	fmt.Fprintf(b, "\n")
	return b.String()
}

func ArrToString(m []Region) []string {
	b := make([]string, len(m))
	for i := 0; i < len(m); i++ {
		b[i] = m[i].String()
	}
	return b
}

//func DisplayBirdArray(arr []model.Bird) string {
//	var br strings.Builder
//
//	for _, i := range [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8} {
//		for _, b := range arr {
//			fmt.Fprintf(&br, "%-65s", b.StringFor(i))
//			fmt.Fprintf(&br, " | ")
//		}
//		fmt.Fprintf(&br, "\n")
//	}
//
//	return br.String()
//}
