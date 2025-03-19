package main

import (
	"testing"
	"time"
)

func TestStore(t *testing.T) {
	s, err := NewStore(":memory:")
	if err != nil {
		t.Fatalf("Failed to create store: %v", err)
	}
	defer s.close()

	// ensure db is empty
	cakes, err := s.getCakes()
	if err != nil {
		t.Errorf("Error getting cakes: %v", err)
	}
	if len(cakes) != 0 {
		t.Errorf("Expected empty cakes list, got: %v", cakes)
	}
	orders, err := s.getOrders()
	if err != nil {
		t.Errorf("Error getting orders: %v", err)
	}
	if len(orders) != 0 {
		t.Errorf("Expected empty orders list, got: %v", orders)
	}

	// create new cake
	newCake := cake{Name: "Chocolate Cake", ID: -1, Price: 2500, Amount: -1}
	newCake.ID, err = s.saveCake(newCake)
	if err != nil {
		t.Fatalf("Failed to save new cake: %v", err)
	}
	if newCake.ID <= 0 {
		t.Errorf("Expected positive cake ID, got: %d", newCake.ID)
	}

	// create another cake
	anotherCake := cake{Name: "Another Cake", ID: -1, Price: 3000, Amount: -1}
	anotherCake.ID, err = s.saveCake(anotherCake)
	if err != nil {
		t.Fatalf("Failed to save new cake: %v", err)
	}
	if anotherCake.ID <= 0 {
		t.Errorf("Expected positive cake ID, got: %d", anotherCake.ID)
	}

	// Update Existing Cake
	newCake.Name = "Updated Cake"
	newCake.Price = 100
	newID, err := s.saveCake(newCake)
	if newID != newCake.ID {
		t.Error("wrong id")
	}
	if err != nil {
		t.Error("error updating cake", err)
	}

	// Update Non-existing Cake
	var updatedCake cake
	updatedCake.ID = 10
	_, err = s.saveCake(updatedCake)
	if err == nil {
		t.Error("did not return an error when attempted to update non existant cake")
	}

	// get cakes
	cakes, err = s.getCakes()
	if err != nil {
		t.Fatalf("Failed to get all cakes after creation: %v", err)
	}
	if len(cakes) != 2 {
		t.Error("Expected two cakes returned, got ", len(cakes))
	}
	for i, c := range cakes {
		if c.ID == newCake.ID {
			newCake.Amount = i
		}
		if c.ID == anotherCake.ID {
			anotherCake.Amount = i
		}
	}
	if newCake.Amount == -1 {
		t.Error("not found newCake")
	}
	if cakes[newCake.Amount].ID != newCake.ID || cakes[newCake.Amount].Name != newCake.Name || cakes[newCake.Amount].Price != newCake.Price {
		t.Errorf("Want %v\nGot  %v", newCake, cakes)
	}
	if anotherCake.Amount == -1 {
		t.Error("not found anotherCake")
	}
	if cakes[anotherCake.Amount].ID != anotherCake.ID || cakes[anotherCake.Amount].Name != anotherCake.Name || cakes[anotherCake.Amount].Price != anotherCake.Price {
		t.Errorf("Want %v\nGot  %v", anotherCake, cakes)
	}

	// create new order
	now := time.Now()
	newOrder := order{
		ID:       -1,
		Name:     "John",
		Surname:  "Doe",
		Phone:    "123-456-7890",
		Location: "Some Location",
		Accepted: now,
		Date:     now.Add(time.Hour * 24),
		Status:   "Pending",
		Paid:     1000,
		Cakes:    []cake{{ID: newCake.ID, Amount: 2}, {ID: anotherCake.ID, Amount: 10}},
	}
	newOrder.ID, err = s.saveOrder(newOrder)
	if err != nil {
		t.Fatalf("Failed to save new order: %v", err)
	}
	if newOrder.ID <= 0 {
		t.Errorf("Expected positive order ID, got: %d", newOrder.ID)
	}

	// create another order
	anotherOrder := order{
		ID:       -1,
		Name:     "Jane",
		Surname:  "Doe",
		Phone:    "123-456-7890",
		Location: "New Location",
		Accepted: now,
		Date:     now.Add(time.Hour * 192),
		Status:   "Accepted",
		Paid:     1500,
		Cakes:    []cake{{ID: newCake.ID, Amount: 100}},
	}
	anotherOrder.ID, err = s.saveOrder(anotherOrder)
	if err != nil {
		t.Fatalf("Failed to save new order: %v", err)
	}
	if anotherOrder.ID <= 0 {
		t.Errorf("Expected positive order ID, got: %d", anotherOrder.ID)
	}

	// update existing order
	newOrder.Name = "James"
	newOrder.Status = "Done"
	newOrder.Cakes = newOrder.Cakes[1:]
	newID, err = s.saveOrder(newOrder)
	if newID != newOrder.ID {
		t.Error("wrong id")
	}
	if err != nil {
		t.Error("error updating order", err)
	}

	// update non existing order
	var updatedOrder order
	updatedOrder.ID = 10
	_, err = s.saveOrder(updatedOrder)
	if err == nil {
		t.Error("did not return an error when attempted to update non existant cake")
	}

	// get orders
	orders, err = s.getOrders()
	if err != nil {
		t.Fatalf("Failed to get all orders after creation: %v", err)
	}
	if len(orders) != 2 {
		t.Error("Expected two orders returned, got ", len(orders))
	}
	newOrderPos := -1
	anotherOrderPos := -1
	for i, o := range orders {
		if o.ID == newOrder.ID {
			newOrderPos = i
		}
		if o.ID == anotherOrder.ID {
			anotherOrderPos = i
		}
	}
	if newOrderPos == -1 {
		t.Error("not found newOrder")
	}
	if newOrder.Name != orders[newOrderPos].Name ||
		newOrder.Surname != orders[newOrderPos].Surname ||
		newOrder.Phone != orders[newOrderPos].Phone ||
		newOrder.Status != orders[newOrderPos].Status ||
		newOrder.Location != orders[newOrderPos].Location ||
		newOrder.Accepted.Format("2006-01-02 15:04") != newOrder.Accepted.Format("2006-01-02 15:04") ||
		newOrder.Date.Format("2006-01-02 15:04") != newOrder.Date.Format("2006-01-02 15:04") {
		t.Errorf("Want %v\nGot  %v", newOrder, orders[newOrderPos])
	}

	if anotherOrderPos == -1 {
		t.Error("not found anotherOrder")
	}
	if anotherOrder.Name != orders[anotherOrderPos].Name ||
		anotherOrder.Surname != orders[anotherOrderPos].Surname ||
		anotherOrder.Phone != orders[anotherOrderPos].Phone ||
		anotherOrder.Status != orders[anotherOrderPos].Status ||
		anotherOrder.Location != orders[anotherOrderPos].Location ||
		anotherOrder.Accepted.Format("2006-01-02 15:04") != anotherOrder.Accepted.Format("2006-01-02 15:04") ||
		anotherOrder.Date.Format("2006-01-02 15:04") != anotherOrder.Date.Format("2006-01-02 15:04") {
		t.Errorf("Want %v\nGot  %v", anotherOrder, orders[anotherOrderPos])
	}
}

func TestSliceComparison(t *testing.T) {
	a := []cake{
		{Name: "ok", Price: 100, ID: 12, Amount: 10},
		{Name: "ok", Price: 11, ID: 13, Amount: 1},
		{Name: "ok", Price: 100, ID: 10, Amount: 11},
	}

	b := []cake{
		{Name: "okas", Price: 110, ID: 12, Amount: 10},
		{Name: "oksa", Price: 11, ID: 13, Amount: 1},
		{Name: "oasdk", Price: 101, ID: 10, Amount: 11},
	}
	if !areCakeSlicesEqual(a, b) {
		t.Error("aren't equal")
	}
}

func areCakesEqual(a, b cake) bool {
	return a.ID == b.ID && a.Amount == b.Amount
}

func areCakeSlicesEqual(a, b []cake) bool {
	if len(a) != len(b) {
		return false
	}

	matched := make([]bool, len(a))

mainLoop:
	for _, e1 := range a {
		for i, e2 := range b {
			if matched[i] {
				continue
			}
			if areCakesEqual(e1, e2) {
				matched[i] = true
				continue mainLoop
			}
		}
	}
	for _, e := range matched {
		if e == false {
			return false
		}
	}
	return true
}
