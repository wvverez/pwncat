package utils

import "fmt"

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Magenta = "\033[35m"
)

func RedText(s string) string    { return fmt.Sprintf("%s%s%s", Red, s, Reset) }
func GreenText(s string) string  { return fmt.Sprintf("%s%s%s", Green, s, Reset) }
func YellowText(s string) string { return fmt.Sprintf("%s%s%s", Yellow, s, Reset) }
func MagentaText(s string) string { return fmt.Sprintf("%s%s%s", Magenta, s, Reset) }
