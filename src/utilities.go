package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"sort"

	"github.com/barkimedes/go-deepcopy"

	"golang.org/x/exp/constraints"
	"golang.org/x/exp/maps"
)

type point struct {
	x int
	y int
}

func getInput(fileName string) []string {
	readFile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}
	readFile.Close()

	return fileLines
}

func MinOf[T constraints.Ordered](vars ...T) T {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

func MaxOf[T constraints.Ordered](vars ...T) T {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

func SumOf[T constraints.Ordered](vars ...T) T {
	var sum T

	for _, i := range vars {
		sum += i
	}

	return sum
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func DeepCopy[T any](srcThing T) T {
	return deepcopy.MustAnything(srcThing).(T)
}

func CopyMap[T ~map[K]V, K comparable, V any](srcMap T) T {
	newMap := make(T)
	maps.Copy(newMap, srcMap)
	return newMap
}

func Pop[T any](arr *[]T) T {
	arrLen := len(*arr)
	ret := (*arr)[arrLen-1]
	*arr = (*arr)[:arrLen-1]
	return ret
}

func InsertSorted[T any](val T, arr *[]T, less func(a *T, b *T) bool) {
	index := sort.Search(len(*arr), func(i int) bool { return less(&val, &(*arr)[i]) })
	*arr = append(*arr, val)
	copy((*arr)[index+1:], (*arr)[index:])
	(*arr)[index] = val
}
