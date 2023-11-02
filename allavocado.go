package allavocado

import (
	"fmt"
	"log/slog"
)

func Main() int {
	slog.Debug("allavocado", "test", true)
	main()

	return 0
}

// CPU represents the CPU component.
type CPU struct{}

func (c *CPU) Start() {
	fmt.Println("CPU is starting")
}

func (c *CPU) Shutdown() {
	fmt.Println("CPU is shutting down")
}

// Memory represents the memory component.
type Memory struct{}

func (m *Memory) Load() {
	fmt.Println("Memory is loading data")
}

func (m *Memory) Clear() {
	fmt.Println("Memory is clearing data")
}

// HardDrive represents the hard drive component.
type HardDrive struct{}

func (h *HardDrive) Read() {
	fmt.Println("Hard drive is reading data")
}

func (h *HardDrive) Write() {
	fmt.Println("Hard drive is writing data")
}

// ComputerFacade is the facade that provides a simplified interface to the subsystem.
type ComputerFacade struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputerFacade() *ComputerFacade {
	return &ComputerFacade{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (cf *ComputerFacade) StartComputer() {
	fmt.Println("Starting computer...")
	cf.cpu.Start()
	cf.memory.Load()
	cf.hardDrive.Read()
	fmt.Println("Computer started successfully")
}

func (cf *ComputerFacade) ShutdownComputer() {
	fmt.Println("Shutting down computer...")
	cf.cpu.Shutdown()
	cf.memory.Clear()
	cf.hardDrive.Write()
	fmt.Println("Computer shut down successfully")
}

func main() {
	computer := NewComputerFacade()

	// Starting the computer using the facade
	computer.StartComputer()

	// Shutting down the computer using the facade
	computer.ShutdownComputer()
}
