package adapter

import "fmt"

type legacyPrinter struct {
}

func (l *legacyPrinter) legacyPrint(msg string) string {
	return fmt.Sprintf("printed %s in a legacy printer", msg)
}
