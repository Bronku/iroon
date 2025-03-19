package main

import (
	"testing"
	"time"
)

func TestCakeOperations(t *testing.T) {
	s, err := NewStore(":memory:")
	if err != nil {
		t.Fatal("Failed to open database", err)
	}
	defer s.db.Close()

	c := cake{
		ID:    -1,
		Name:  "Test Cake",
		Price: 10,
	}

	id, err := s.saveCake(c)
	if id == -1 || err != nil {
		t.Fatal("Failed to save cake", err)
	}

	cakes := s.getCakes()
	if len(cakes) != 1 {
		t.Errorf("Expected 1 cake, got %d", len(cakes))
	}

	if cakes[0].Name != "Test Cake" || cakes[0].Price != 10 {
		t.Errorf("Cake data mismatch. Got %+v", cakes[0])
	}

	updatedCake := cake{
		ID:    id,
		Name:  "Updated Cake",
		Price: 15,
	}

	newID, err := s.saveCake(updatedCake)
	if newID != id {
		t.Errorf("Update returned different ID. Expected %d, got %d", id, newID)
	}

	cakes = s.getCakes()
	if cakes[0].Name != "Updated Cake" || cakes[0].Price != 15 {
		t.Errorf("Cake update failed. Got %+v", cakes[0])
	}
}

func TestOrderOperations(t *testing.T) {
	s, err := NewStore(":memory:")
	if err != nil {
		t.Fatal("Failed to open database", err)
	}
	defer s.db.Close()

	now := time.Now()
	delivery := now.Add(24 * time.Hour)

	o := order{
		ID:       -1,
		Name:     "John",
		Surname:  "Doe",
		Phone:    "1234567890",
		Location: "123 Test St",
		Accepted: now,
		Date:     delivery,
		Status:   "pending",
		Paid:     0,
	}

	id, err := s.saveOrder(o)
	if id == -1 || err != nil {
		t.Fatal("Failed to save order", err)
	}

	retrieved, err := s.getOrder(id)
	if err != nil {
		t.Fatalf("Failed to get order: %v", err)
	}

	if retrieved.Name != "John" || retrieved.Surname != "Doe" {
		t.Errorf("Order names mismatch. Got %+v", retrieved)
	}

	if !timeEqual(retrieved.Date, o.Date) || !timeEqual(retrieved.Accepted, o.Accepted) {
		t.Errorf("Order dates mismatch. \nGot  %+v\nWant %+v", retrieved, o)
	}

	orders := s.getOrders()
	if len(orders) != 1 {
		t.Errorf("Expected 1 order, got %d", len(orders))
	}

	// Test updating an order
	o.ID = id
	o.Status = "completed"

	newID, err := s.saveOrder(o)
	if err != nil {
		t.Fatal("Failed to save order", err)
	}
	if newID != id {
		t.Errorf("Update returned different ID. Expected %d, got %d", id, newID)
	}

	updated, err := s.getOrder(id)
	if err != nil {
		t.Fatalf("Failed to get updated order: %v", err)
	}

	if updated.Status != "completed" {
		t.Errorf("Order update failed. Got %+v", updated)
	}
}

func TestStoreInitialization(t *testing.T) {
	s, err := NewStore(":memory:")
	if err != nil {
		t.Fatal("Failed to open database", err)
	}
	defer s.db.Close()

	cakes := s.getCakes()
	if len(cakes) != 0 {
		t.Errorf("New store should have no cakes, got %d", len(cakes))
	}

	orders := s.getOrders()
	if len(orders) != 0 {
		t.Errorf("New store should have no orders, got %d", len(orders))
	}
}

func timeEqual(t1, t2 time.Time) bool {
	format := "2006-01-02 15:04"
	return t1.Format(format) == t2.Format(format)
}
