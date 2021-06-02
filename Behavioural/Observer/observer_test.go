package observer

import (
	"fmt"
	"testing"
)

func TestObserver(t *testing.T) {
	shirtItem := newItem("Nike Shirt")
	observerFirst := &customer{id: "thomas"}
	observerSecond := &customer{id: "leo"}
	shirtItem.register(observerFirst)
	shirtItem.register(observerSecond)
	shirtItem.updateAvailability(true)
	ExpectedNotification1 := fmt.Sprintf("customer %s got notifed about item %s\n", "thomas", "Nike Shirt")
	if observerFirst.latestNotification != ExpectedNotification1 {
		t.Errorf("Expected notification - `%s` - but got -`%s` ", ExpectedNotification1, observerFirst.latestNotification)
	}
	ExpectedNotification2 := fmt.Sprintf("customer %s got notifed about item %s\n", "leo", "Nike Shirt")
	if observerSecond.latestNotification != ExpectedNotification2 {
		t.Errorf("Expected notification `%s` - but got -`%s` ", ExpectedNotification2, observerSecond.latestNotification)
	}
}
