package commands

import (
  "bufio"
  "os"
  "fmt"
  "strings"
  "simple_cli/expenses"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)

func GetInput() (string, error) {

  fmt.Print("-> ")
	str, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

  str = strings.Replace(str, "\n", "", 1)

	return str, nil
}

func ShowInConsole(expensesList []float32) {
  fmt.Println(ContentString(expensesList))
}

func ContentString(expensesList []float32) string {
  var builder = strings.Builder{}

  fmt.Println("")

  for i, expense := range expensesList {
    builder.WriteString(fmt.Sprintf("Expense: %6.2f\n", expense))

    max, min, avg, total := expensesDetails(expensesList)


    if i == len(expensesList) - 1 {
      builder.WriteString("")
      builder.WriteString("==================\n")
      builder.WriteString(fmt.Sprintf("Total: %6.2f\n", total))
      builder.WriteString(fmt.Sprintf("Max: %6.2f\n", max))
      builder.WriteString(fmt.Sprintf("Min: %6.2f\n", min))
      builder.WriteString(fmt.Sprintf("Average: %6.2f\n", avg))
      builder.WriteString("==================\n")
    }
  }

  return builder.String()
}

func expensesDetails(expensesList []float32) (max, min, avg, total float32) {
  if len(expensesList) == 0 {
    return
  }

  min = expenses.Min(expensesList...)
  max = expenses.Max(expensesList...)
  total = expenses.Sum(expensesList...)
  avg = expenses.Average(expensesList...)

  return
}



func Export(fileName string, list []float32) error {
  f, err := os.Create(fileName)

  if err != nil {
    return err
  }

  defer f.Close()

  w := bufio.NewWriter(f)

  _, err = w.WriteString(ContentString(list))

  if err != nil {
    return err
  }

  return w.Flush()
}
