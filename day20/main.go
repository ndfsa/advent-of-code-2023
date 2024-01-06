package day20

import (
	"slices"
	"strings"

	"github.com/ndfsa/advent-of-code-2023/util"
)

const (
	SIGNAL_HIGH = iota
	SIGNAL_LOW
)

type Module interface {
	Execute(Signal) []Signal
	GetOutput() []string
}

type Broadcaster struct {
	Output []string
}

type FlipFlop struct {
	Output []string
	State  bool
}

type Nand struct {
	Output []string
	Input  map[string]int
}

type Signal struct {
	Source      string
	Strength    int
	Destination string
}

type Machine struct {
	Queue   []Signal
	Modules map[string]Module
}

func (b *Broadcaster) Execute(signal Signal) []Signal {
	res := []Signal{}
	for _, dest := range b.Output {
		res = append(res, Signal{
			Strength:    signal.Strength,
			Destination: dest,
			Source:      signal.Destination})
	}
	return res
}

func (f *FlipFlop) Execute(signal Signal) []Signal {
	res := []Signal{}
	if signal.Strength == SIGNAL_HIGH {
		return res
	}

	for _, dest := range f.Output {
		if f.State {
			res = append(res, Signal{
				Strength:    SIGNAL_LOW,
				Destination: dest,
				Source:      signal.Destination})
		} else {
			res = append(res, Signal{
				Strength:    SIGNAL_HIGH,
				Destination: dest,
				Source:      signal.Destination})
		}
	}

	f.State = !f.State
	return res
}

func (n *Nand) Execute(signal Signal) []Signal {
	res := []Signal{}
	n.Input[signal.Source] = signal.Strength

	output := SIGNAL_LOW
	for _, st := range n.Input {
		if st == SIGNAL_LOW {
			output = SIGNAL_HIGH
			break
		}
	}

	for _, dest := range n.Output {
		res = append(res, Signal{
			Strength:    output,
			Destination: dest,
			Source:      signal.Destination})
	}
	return res
}

func (b Broadcaster) GetOutput() []string {
	return b.Output
}

func (f FlipFlop) GetOutput() []string {
	return f.Output
}

func (n Nand) GetOutput() []string {
	return n.Output
}

func parseInput(lines []string) Machine {
	modules := map[string]Module{}
	for _, line := range lines {
		name, dest, _ := strings.Cut(line, " -> ")
		output := strings.Split(dest, ", ")
		if name == "broadcaster" {
			modules[name] = &Broadcaster{Output: output}
		} else if name[0] == '%' {
			modules[name[1:]] = &FlipFlop{Output: output, State: false}
		} else if name[0] == '&' {
			modules[name[1:]] = &Nand{Output: output, Input: map[string]int{}}
		}
	}

	for name, module := range modules {
		if nand, ok := module.(*Nand); ok {
			for otherName, other := range modules {
				if name == otherName {
					continue
				}
				if slices.Contains(other.GetOutput(), name) {
					nand.Input[otherName] = SIGNAL_LOW
				}
			}
		}
	}
	return Machine{
		Modules: modules,
		Queue: []Signal{{
			Strength:    SIGNAL_LOW,
			Destination: "broadcaster",
			Source:      "button"}}}
}

func SolvePart1(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	machine := parseInput(lines)
	count := 0
	lowCount, highCount := 0, 0
	for count < 1000 {
		for len(machine.Queue) > 0 {
			signal := machine.Queue[0]
			if signal.Strength == SIGNAL_LOW {
				lowCount++
			} else {
				highCount++
			}
			machine.Queue = machine.Queue[1:]

			if destModule, ok := machine.Modules[signal.Destination]; ok {
				newSignals := destModule.Execute(signal)
				machine.Queue = append(machine.Queue, newSignals...)
			}
		}

		machine.Queue = append(machine.Queue, Signal{
			Strength:    SIGNAL_LOW,
			Destination: "broadcaster",
			Source:      "button"})
		count++
	}

	return lowCount * highCount, nil
}

func SolvePart2(filePath string) (int, error) {
	lines, err := util.ReadFileSplit(filePath)

	if err != nil {
		return 0, err
	}

	// the problem states that a single pulse to rx will turn on a machine the
	// module rx is preceded by a nand gate, to which 4 inputs go in so you
	// have to find the number of button presses for each to turn on all at the
	// same time (LCM)

	machine := parseInput(lines)

	// find the last nand, extract it's inputs
	nandInput := map[string]int{}
	lastNand := ""
	for name, module := range machine.Modules {
		if slices.Contains(module.GetOutput(), "rx") {
			for inName := range module.(*Nand).Input {
				lastNand = name
				nandInput[inName] = 0
			}
			break
		}
	}

	count := 0
	cycle := 0
out:
	for {
		for len(machine.Queue) > 0 {
			signal := machine.Queue[0]
			machine.Queue = machine.Queue[1:]

			if signal.Destination == lastNand && signal.Strength == SIGNAL_HIGH {
				if prev := nandInput[signal.Source]; prev == 0 {
					nandInput[signal.Source] = count
					cycle++
				}

				if cycle >= len(nandInput) {
					break out
				}
			}

			if destModule, ok := machine.Modules[signal.Destination]; ok {
				newSignals := destModule.Execute(signal)
				machine.Queue = append(machine.Queue, newSignals...)
			}
		}

		machine.Queue = append(machine.Queue, Signal{
			Strength:    SIGNAL_LOW,
			Destination: "broadcaster",
			Source:      "button"})
		count++
	}

	// for some reason off by one
	res := []int{}
	for _, val := range nandInput {
		res = append(res, val+1)
	}

	return util.LCM(res), nil
}
