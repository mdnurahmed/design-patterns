package observer

import "fmt"

type customer struct {
	id                 string
	latestNotification string
}

func (c *customer) notify(itemName string) {
	c.latestNotification = fmt.Sprintf("customer %s got notifed about item %s\n", c.id, itemName)
}

func (c *customer) getID() string {
	return c.id
}
