package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	f, err := ioutil.ReadFile("registercommands.txt")
	if err != nil {
		log.Fatalln(err)
	}

	allRegs := Registers{}

	max := 0
	commandLines := strings.Split(string(f), "\n")
	for _, l := range commandLines {
		cmd := ParseCommand(l)
		if cmd.Cond.Check(allRegs, cmd.CondRegister, cmd.CondAmount) {
			changedValue := allRegs.Set(cmd.Register, cmd.Op, cmd.Amount)
			if changedValue > max {
				max = changedValue
			}
		}
	}

	// Pick a random max from the list (in case the lowest is negative)
	maxAtEnd := 0
	for _, val := range allRegs {
		maxAtEnd = val
		break
	}

	// Find the actual max
	for _, val := range allRegs {
		if val > maxAtEnd {
			maxAtEnd = val
		}
	}

	fmt.Printf("Largest register value at end: %d\n", maxAtEnd)
	fmt.Printf("Largest register value at any time: %d\n", max)
}

type Command struct {
	Register     string
	Op           Operation
	Amount       int
	Cond         Condition
	CondRegister string
	CondAmount   int
}

type Operation int

const (
	Inc Operation = iota
	Dec
)

type Condition int

const (
	GT Condition = iota
	LT
	GTE
	LTE
	EQ
	NE
)

type Registers map[string]int

func (rs Registers) Get(reg string) int {
	if _, exists := rs[reg]; !exists {
		rs[reg] = 0
	}
	return rs[reg]
}

func (rs Registers) Set(reg string, op Operation, val int) int {
	switch op {
	case Inc:
		rs.Inc(reg, val)
	case Dec:
		rs.Dec(reg, val)
	}
	return rs.Get(reg)
}

func (rs Registers) Inc(reg string, val int) {
	if _, exists := rs[reg]; !exists {
		rs[reg] = 0
	}
	rs[reg] += val
}

func (rs Registers) Dec(reg string, val int) {
	if _, exists := rs[reg]; !exists {
		rs[reg] = 0
	}
	rs[reg] -= val
}

func (c Condition) Check(rs Registers, reg string, condAmt int) bool {
	switch c {
	case GT:
		return rs.Get(reg) > condAmt
	case LT:
		return rs.Get(reg) < condAmt
	case GTE:
		return rs.Get(reg) >= condAmt
	case LTE:
		return rs.Get(reg) <= condAmt
	case EQ:
		return rs.Get(reg) == condAmt
	case NE:
		fallthrough
	default:
		return rs.Get(reg) != condAmt
	}
}

// The order in the command is
// 0        1         2      3    4            5         6
// register operation amount "if" condRegister condition condAmount
func ParseCommand(s string) Command {
	fields := strings.Fields(s)
	reg := fields[0]

	var op Operation
	switch fields[1] {
	case "inc":
		op = Inc
	case "dec":
		op = Dec
	}

	multBy := 1
	if fields[2][0] == '-' {
		multBy = -1
		fields[2] = fields[2][1:]
	}
	amount, _ := strconv.Atoi(fields[2])
	amount *= multBy

	condReg := fields[4]

	var cond Condition
	switch fields[5] {
	case ">":
		cond = GT
	case "<":
		cond = LT
	case ">=":
		cond = GTE
	case "<=":
		cond = LTE
	case "==":
		cond = EQ
	case "!=":
		cond = NE
	}

	multBy = 1
	if fields[6][0] == '-' {
		multBy = -1
		fields[6] = fields[6][1:]
	}
	condAmount, _ := strconv.Atoi(fields[6])
	condAmount *= multBy

	return Command{
		Register:     reg,
		Op:           op,
		Amount:       amount,
		Cond:         cond,
		CondRegister: condReg,
		CondAmount:   condAmount,
	}
}
