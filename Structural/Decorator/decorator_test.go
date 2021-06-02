package decorator

import "testing"

func TestDecorator(t *testing.T) {
	pizza := &normalPizza{}
	if pizza.getPrice() != 20 {
		t.Errorf("Expected normal pizza price to be 20 but got %d ", pizza.getPrice())
	}
	decoratedPizza := &tomatoTopping{
		pizza: pizza,
	}
	if decoratedPizza.getPrice() != 27 {
		t.Errorf("Expected decorated pizza price to be 27  but got %d ", decoratedPizza.getPrice())
	}
	decoratedPizza2 := &tomatoTopping{
		pizza: decoratedPizza,
	}
	if decoratedPizza2.getPrice() != 34 {
		t.Errorf("Expected decorated pizza price to be 34 but got %d ", decoratedPizza.getPrice())
	}
}
