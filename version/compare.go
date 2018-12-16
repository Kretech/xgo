package version

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

var ErrOverflowSection = errors.New("Too long version sections")
var ErrOverSize = errors.New("Too large version number")

const (
	maxSection    = 4
	perSectionBit = 16
	maxPerSection = 1<<perSectionBit - 1
)

const (
	// 跳过用0表示结果，避免默认值带来的影响
	resultEqual = iota + 1
	resultGreater
	resultLess
)

var vCache map[string]uint64

func init() {
	vCache = make(map[string]uint64, 8)
}

// version2Int 把版本号转换成为可比较的数字
// 支持范围：4段4位数，(0.0.0.0, 9999.9999.9999.9999]
func version2Int(sVersion string) (v uint64, err error) {

	sVersion = strings.TrimLeft(sVersion, "Vv")
	sVersion = strings.Replace(sVersion, "-", ".", -1)

	v = vCache[sVersion]
	if v > 0 {
		return
	}

	if strings.Count(sVersion, ".") > maxSection {
		return v, ErrOverflowSection
	}

	svs := strings.Split(sVersion, ".")
	vs := make([]int, 0, len(svs))
	for _, s := range svs {
		i, _ := strconv.Atoi(s)
		if i > maxPerSection {
			return 0, ErrOverSize
		}

		vs = append(vs, i)
	}
	v = uint64(0)

	for i := 0; i < len(vs); i++ {
		v += uint64(vs[i] << (perSectionBit * uint(maxSection-1-i)))
	}

	// fmt.Printf("%-30s%016x\n", sVersion, v)
	vCache[sVersion] = v

	return
}

func Compare(v1, v2 string) (result int, err error) {
	hash1, err := version2Int(v1)
	if err != nil {
		return
	}

	hash2, err := version2Int(v2)
	if err != nil {
		return
	}

	if hash1 > hash2 {
		return resultGreater, nil
	}

	if hash1 < hash2 {
		return resultLess, nil
	}

	if hash1 == hash2 {
		return resultEqual, nil
	}

	return 0, nil
}

func LessThan(v1, v2 string) bool {
	r, err := Compare(v1, v2)
	if err != nil {
		log.Println(`Error In version/compare.go#LessThan`, err)
	}

	return r == resultLess
}

func GreaterThan(v1, v2 string) bool {
	r, err := Compare(v1, v2)
	if err != nil {
		log.Println(`Error In version/compare.go#GreaterThan`, err)
	}

	return r == resultGreater
}

func Equal(v1, v2 string) bool {
	r, err := Compare(v1, v2)
	if err != nil {
		log.Println(`Error In version/compare.go#Equal`, err)
	}

	return r == resultEqual
}
