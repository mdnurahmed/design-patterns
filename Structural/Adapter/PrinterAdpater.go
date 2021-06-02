package adapter

type printerAdapter struct {
	legacyPrinter ILegacyPrinter
}

func (p *printerAdapter) modernPrint(msg string) string {
	return p.legacyPrinter.legacyPrint(msg)
}
