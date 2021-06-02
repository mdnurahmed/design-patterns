package singleton

import "testing"

func TestGetInstance(t *testing.T) {

	counter1 := GetInstance()
	if counter1 == nil {
		t.Error("Got nil, expected pointer to Singleton")
	}

	expectedCounter := counter1
	currentCount := counter1.AddValue(1)

	if currentCount != 1 {
		t.Errorf("After calling for the first time to count, the count must be 1, but got: %d\n", currentCount)
	}

	counter2 := GetInstance()

	if counter2 != expectedCounter {
		t.Error("Expected same instance in counter2 but it got a different instance")
	}

	currentCount = counter2.AddValue(1)

	if currentCount != 2 {
		t.Errorf("After calling 'AddOne' using the second counter, the current count must be 2 but was: %d\n", currentCount)
	}
}
