package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func part2() {
	fmt.Println("Start")
	path := "inputs/2.txt"
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	s := string(file)

	ss := strings.Split(s, "\n")
	var calories []int
	calsPerElf := 0
	println(len(ss))

	for n := range ss {
		v := ss[n]
		
		if (v != "" ) {
			cal, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			calsPerElf += cal
		} else {
			calories = append(calories, calsPerElf)
			calsPerElf = 0
		}
	}

	println("size: ", len(calories))

	//sort slice in ascending order
	sort.Slice(calories, func (i, j int) bool {
		return calories[i] < calories[j]
	})

	total := calories[len(calories) -1] + calories[len(calories) -2] + calories[len(calories) -3] 
	println("total: ", total)
}

func part1() {
	fmt.Println("Start")
	path := "inputs/2.txt"
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	s := string(file)

	ss := strings.Split(s, "\n")
	var calories []int
	calsPerElf := 0
	println(len(ss))

	for n := range ss {
		v := ss[n]
		
		if (v != "" ) {
			cal, err := strconv.Atoi(v)
			if err != nil {
				log.Fatal(err)
			}
			calsPerElf += cal
		} else {
			calories = append(calories, calsPerElf)
			calsPerElf = 0
		}
	}

	mostCal := calories[0]
	println("size: ", len(calories))

	for n := range calories {
		//println("elf ", n, " ", calories[n])
		if (calories[n] == 0) {
			break
		}
		if mostCal < calories[n] {
			mostCal = calories[n]
		}
	}
	println("high elf ", mostCal)
}


func main() {
	part2()
}
