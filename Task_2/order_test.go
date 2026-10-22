package order

import (
	"errors"
	"slices"
	"strings"
	"testing"
)

func TestPriceIsZero(t *testing.T) {
	order, err := MakeOrder(1)
	if err != nil {
		t.Fatalf("MakeOrder: %v", err)
	}
	totalPrice, err := order.CountTotalPrice()
	if err != nil {
		t.Fatalf("CountTotalPrice: %v", err)
	}
	if totalPrice != 0 {
		t.Fatalf("total price = %d; want 0", totalPrice)
	}
}

func TestMakeOrder(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		wantErr error
	}{
		{name: "positive ID", id: 42},
		{name: "zero ID", id: 0, wantErr: ErrInvalidId},
		{name: "negative ID", id: -1, wantErr: ErrInvalidId},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order, err := MakeOrder(tt.id)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("MakeOrder(%d) error = %v; want %v", tt.id, err, tt.wantErr)
			}
			if tt.wantErr != nil {
				return
			}
			if order.id != tt.id || order.status != StatusNew || len(order.items) != 0 {
				t.Errorf("unexpected new order: %+v", order)
			}
		})
	}
}

func TestAddItemsAndCountTotalPrice(t *testing.T) {
	order, err := MakeOrder(1)
	if err != nil {
		t.Fatalf("MakeOrder: %v", err)
	}

	items := []Item{
		{name: "bread", price: 100},
		{name: "milk", price: 75},
	}
	for _, item := range items {
		if err := order.AddItem(item); err != nil {
			t.Fatalf("AddItem(%+v): %v", item, err)
		}
	}
	if !slices.Equal(order.items, items) {
		t.Errorf("items = %+v; want %+v", order.items, items)
	}
	total, err := order.CountTotalPrice()
	if err != nil {
		t.Fatalf("CountTotalPrice: %v", err)
	}
	if total != 175 {
		t.Errorf("total price = %d; want 175", total)
	}
	if !slices.Equal(order.items, items) {
		t.Error("CountTotalPrice changed the items")
	}
}

func TestAddItemInvalidData(t *testing.T) {
	tests := []struct {
		name    string
		item    Item
		wantErr error
	}{
		{name: "empty name", item: Item{price: 100}, wantErr: ErrEmtyName},
		{name: "whitespace name", item: Item{name: " \t\n", price: 100}, wantErr: ErrEmtyName},
		{name: "negative price", item: Item{name: "bread", price: -1}, wantErr: ErrNegPrice},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			order := Order{id: 1, status: StatusNew, items: []Item{{name: "milk", price: 75}}}
			before := slices.Clone(order.items)
			if err := order.AddItem(tt.item); !errors.Is(err, tt.wantErr) {
				t.Fatalf("AddItem error = %v; want %v", err, tt.wantErr)
			}
			if order.id != 1 || order.status != StatusNew || !slices.Equal(order.items, before) {
				t.Errorf("invalid item changed the order: %+v", order)
			}
		})
	}
}

func TestAddItemNilOrder(t *testing.T) {
	var order *Order
	if err := order.AddItem(Item{name: "bread", price: 100}); !errors.Is(err, ErrNilOrder) {
		t.Fatalf("AddItem error = %v; want %v", err, ErrNilOrder)
	}
}

func TestAddItemZeroPrice(t *testing.T) {
	order := Order{id: 1, status: StatusNew}
	item := Item{name: "gift", price: 0}
	if err := order.AddItem(item); err != nil {
		t.Fatalf("AddItem with zero price: %v", err)
	}
	if len(order.items) != 1 || order.items[0] != item {
		t.Fatalf("free item was not added: %+v", order.items)
	}
	total, err := order.CountTotalPrice()
	if err != nil || total != 0 {
		t.Errorf("CountTotalPrice = (%d, %v); want (0, nil)", total, err)
	}
}

func TestCountTotalPriceNilOrder(t *testing.T) {
	var order *Order
	total, err := order.CountTotalPrice()
	if !errors.Is(err, ErrNilOrder) || total != 0 {
		t.Fatalf("CountTotalPrice = (%d, %v); want (0, %v)", total, err, ErrNilOrder)
	}
}

func TestOrderCopy(t *testing.T) {
	original := Order{id: 1, status: StatusNew, items: []Item{{name: "bread", price: 100}}}
	copied := original
	copied.id = 2
	copied.status = StatusCompleted

	if original.id != 1 || original.status != StatusNew {
		t.Errorf("changing the copy changed the original: %+v", original)
	}

	copied.items = slices.Clone(copied.items)
	copied.items[0].price = 200
	if original.items[0].price != 100 {
		t.Error("changing a cloned item changed the original")
	}
}

func TestStatusDescriptions(t *testing.T) {
	for _, status := range []string{StatusNew, StatusProcessing, StatusCompleted, StatusCancelled} {
		t.Run(status, func(t *testing.T) {
			description, ok := StatusDescription[status]
			if !ok || strings.TrimSpace(description) == "" {
				t.Errorf("missing description for status %q", status)
			}
		})
	}
}
