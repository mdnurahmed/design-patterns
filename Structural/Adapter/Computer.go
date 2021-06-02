package adapter

type computer struct {
}

func (c *computer) print(printer IModernPrinter, msg string) string {
	return printer.modernPrint(msg)
}
