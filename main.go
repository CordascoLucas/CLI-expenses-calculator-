package main

import (
  "log"
  "strconv"
  "flag"
  "simple_cli/commands"
  "fmt"
)

func main() {
  var expenses []float32
  var export string

  flag.StringVar(&export, "export", "", "Exports the details to .txt")

  flag.Parse()

  fmt.Print(`1. Para terminar e imprimir en pantalla introduzca cls
2. Para terminar y exportar sus gastos introduzca export
`)

  for {
	input, err := commands.GetInput()

	if err != nil {
		log.Fatal(err)
	}

	if input == "cls" {
		break
	}

	expense, err := strconv.ParseFloat(input, 32)

	if err != nil {
		continue
	}

	expenses = append(expenses, float32(expense))
  }

  if export == "" {
    commands.ShowInConsole(expenses)
  } else {
    commands.Export(export, expenses)
  }
}
