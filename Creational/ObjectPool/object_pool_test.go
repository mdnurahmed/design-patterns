package object_pool

import (
	"strconv"
	"testing"
)

func TestObjectPool(t *testing.T) {
	connections := make([]iPoolObject, 0)
	for i := 0; i < 3; i++ {
		c := &connection{id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	pool, err1 := initPool(connections) //creating a pool with 3 iPoolobjects
	if err1 != nil {
		t.Errorf("Got error %s, expected error to be nil\n", err1.Error())
	}
	conn1, err2 := pool.loan()
	if err2 != nil {
		t.Errorf("Got error %s, expected error to be nil\n", err2.Error())
	}
	conn2, err3 := pool.loan()
	if conn1.getID() == conn2.getID() {
		t.Errorf("Got same object id %s, expected to be different objects\n", conn1.getID())
	}
	if err3 != nil {
		t.Errorf("Got error %s, expected error to be nil\n", err3.Error())
	}
	err4 := pool.receive(conn1)
	if err4 != nil {
		t.Errorf("Got error %s, expected error to be nil\n", err4.Error())
	}
	err5 := pool.receive(conn2)
	if err5 != nil {
		t.Errorf("Got error %s, expected error to be nil\n", err5.Error())
	}
}
