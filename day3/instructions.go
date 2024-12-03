package main

import (
	"fmt"
	"log"
)

type Instruction struct {
	Name       InstructionName
	StartIndex int
	EndIndex   int
	Param1     int
	Param2     int
}

type InstructionName int
type MachineState int

const (
	MULTIPLY InstructionName = iota
	DO
	DONT
)

var instructName = map[InstructionName]string{
	MULTIPLY: "MULT",
	DO:       "START",
	DONT:     "STOP",
}

func (in InstructionName) String() string {
	return instructName[in]
}

func (i Instruction) String() string {
	if i.Name == MULTIPLY {
		return fmt.Sprintf("%d: %s\t %d * %d", i.StartIndex, i.Name, i.Param1, i.Param2)
	}
	return fmt.Sprintf("%d: %s", i.StartIndex, i.Name)
}

const (
	ON MachineState = iota
	OFF
)

type Machine struct {
	State              MachineState
	InstructionSet     []Instruction
	CurrentInstruction int
	Value              int
}

func NewMachine() *Machine {
	return &Machine{
		State:              ON,
		InstructionSet:     []Instruction{},
		CurrentInstruction: 0,
		Value:              0,
	}
}

func (m *Machine) InitializeProgram(inst []Instruction) {
	m.InstructionSet = inst
}

func (m *Machine) ProcessInstructionAt(instructionIndex int) {
	if instructionIndex > len(m.InstructionSet)-1 || instructionIndex < 0 {
		log.Printf("[MACHINE ERROR] Instruction out of range! %d - %d", instructionIndex, len(m.InstructionSet))
		return
	}
	m.CurrentInstruction = instructionIndex

	switch m.InstructionSet[instructionIndex].Name {
	case MULTIPLY:
		if m.State == ON {
			m.Value += m.InstructionSet[instructionIndex].Param1 * m.InstructionSet[instructionIndex].Param2
		}
		break
	case DO:
		m.State = ON
		break
	case DONT:
		m.State = OFF
		break
	default:
		log.Printf("[ERROR] Machine instruction missing or corrupted!")
	}
	return
}

func (m *Machine) RunProgram() {
	for ind := range m.InstructionSet {
		m.ProcessInstructionAt(ind)
	}
}

func (m *Machine) OutputResult() {
	fmt.Printf("Value: %d\n", m.Value)
}
