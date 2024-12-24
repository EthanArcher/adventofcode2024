package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var A int
var B int
var C int

var pointerPosition int = 0

var output strings.Builder

func main() {

	program := readFromFile("input.txt")
	fmt.Println(program)
	
	runProgram(program)
	fmt.Println(output.String())

	values := []int{}

	for i:=1; i<8; i++ {
		values = append(values, int(math.Pow(float64(8), float64(15)))*i)
	}

	for j:=14; j>=0; j-- { 
		newValues := []int{}
		power := j

		for _,v := range values {

			for i:=0; i<8; i++ {
				x := v + (i * int(math.Pow(float64(8), float64(power))) )
				resetEverything()
				A = x
				runProgram(program)

				if slicesEqual(program, strings.Split(output.String(), ",")[:len(program)], 16-power) { 
					newValues = append(newValues, x)
				}
			}	
		}
		values = []int{}
		values = append(values, newValues...)
	}
	fmt.Println(values)
	fmt.Println(values[0])
}

func resetEverything() {
	A = 0
	B = 0
	C = 0
	pointerPosition = 0
	output.Reset()
}

func slicesEqual(a []int, b []string, last int) bool {
	if len(a) != len(b){
		return false
	}

	for i:= 0; i<last; i++ {
		p := len(a) - 1-i
		if strconv.Itoa(a[p]) != b[p] {
			return false
		}
	}
	return true
}


func runProgram(program []int) {
	for pointerPosition < len(program) {
		opcode := program[pointerPosition]
		operand := program[pointerPosition+1]
		
		switch opcode {
		case 0:
			adv(operand)
		case 1:
			bxl(operand)
		case 2:
			bst(operand)
		case 3:
			jnz(operand)
		case 4:
			bxc(operand)
		case 5:
			out(operand)
		case 6:
			bdv(operand)
		case 7:
			cdv(operand)
		default:
			panic("Invalid opcode")
		}
	}
}

func combo(operand int) int {
	switch operand {
	case 0: 
		return 0
	case 1: 
		return 1
	case 2:
		return 2
	case 3: 
		return 3
	case 4:
		return A
	case 5:
		return B
	case 6:
		return C
	case 7:
		panic("Invalid operand was 7")
	default:
		panic("Invalid operand")
	}
}

func adv(operand int) {
	operand = combo(operand)
	numerator := A
	denominator := math.Pow(float64(2), float64(operand))
	A = int(math.Floor(float64(numerator) / denominator))
	pointerPosition += 2
}

func bxl(operand int) {
	B = operand ^ B
	pointerPosition += 2
}

func bst(operand int) {
	operand = combo(operand)
	B = operand % 8
	pointerPosition += 2
}

func jnz(operand int) {
	if A!=0 {
		pointerPosition = operand
	} else {
		pointerPosition += 2
	}
}

func bxc(_ int) {
	B = B ^ C
	pointerPosition += 2
}

func out(operand int) {
	operand = combo(operand)
	pointerPosition += 2
	output.WriteString(strconv.Itoa(operand % 8)+",")
}

func bdv(operand int) {
	operand = combo(operand)
	numerator := A
	denominator := math.Pow(float64(2), float64(operand))
	B = int(math.Floor(float64(numerator) / denominator))
	pointerPosition += 2
}

func cdv(operand int) {
	operand = combo(operand)
	numerator := A
	denominator := math.Pow(float64(2), float64(operand))
	C = int(math.Floor(float64(numerator) / denominator))
	pointerPosition += 2
}

func readFromFile(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Errorf("error opening file: %w", err)
		return nil
	}
	defer file.Close()

	var sections []string
	currentBlock := strings.Builder{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" { // Empty line
			sections = append(sections, currentBlock.String())
			currentBlock.Reset() // Start a new block
		} else {
			currentBlock.WriteString(line+"\n")
		}
	}
	sections = append(sections, currentBlock.String())

	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(sections[0], -1)
	A = asANumber(matches[0])
	B = asANumber(matches[1])
	C = asANumber(matches[2])
	
	if err := scanner.Err(); err != nil {
		fmt.Errorf("error reading file: %w", err)
		return nil
	}

	program := []int{}
	for _, v := range strings.Split(sections[1], ",") {
		program = append(program, asANumber(re.FindAllString(v, -1)[0]))
	}

	return program
}

func asANumber(s string) int {
	num, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		fmt.Errorf("error converting string '%s' to int: %w", s, err)	
		return 0
	}
	return num
}
