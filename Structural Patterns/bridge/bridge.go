package main

import "fmt"

type Printer interface {
	PrintFile()
}

type Computer interface {
	Print()
	SetPrinter(p Printer)
}

type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("PRINT REQUEST FOR MAC")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("PRINT REQUEST FOR WINDOWS")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

type Epson struct{}

func (p *Epson) PrintFile() {
	fmt.Println("PRINTING USING EPSON PRINTER")
}

type Hp struct{}

func (p *Hp) PrintFile() {
	fmt.Println("PRINTING USING HP PRINTER")
}

func main() {

	hpPrinter := &Hp{}
	epsonPrinter := &Epson{}

	macComputer := &Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	windowsComputer := &Windows{}

	windowsComputer.SetPrinter(hpPrinter)
	windowsComputer.Print()
	fmt.Println()

	windowsComputer.SetPrinter(epsonPrinter)
	windowsComputer.Print()
	fmt.Println()

}
