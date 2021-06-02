package adapter

import (
	"fmt"
	"testing"
)

func TestAdpater(t *testing.T) {
	computer := &computer{}
	legacyPrinter := &legacyPrinter{}
	printerAdapter := &printerAdapter{legacyPrinter: legacyPrinter}
	msg := "hello world"
	if computer.print(printerAdapter, msg) != fmt.Sprintf("printed %s in a legacy printer", msg) {
		t.Errorf("Expected printed message to be %s but got %s ", msg, computer.print(printerAdapter, msg))
	}
}
