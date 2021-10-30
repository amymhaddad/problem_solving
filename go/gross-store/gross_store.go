package gross

// Units stores the Gross Store unit measurements.
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen":    6,
		"dozen":              12,
		"small_gross":        120,
		"gross":              144,
		"great_gross":        1728,
	}

}

// NewBill creates a new bill.
func NewBill() map[string]int {
	return make(map[string]int)
}

// AddItem adds an item to customer bill.
func AddItem(bill, units map[string]int, item, unit string) bool {
	_, found := units[unit]
	if !found {
		return false
	}
	bill[item] = units[unit]
	return true
}

// RemoveItem removes an item from customer bill.
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, foundItem := bill[item]
	_, foundUnit := units[unit]

	if !foundItem || !foundUnit {
		return false
	}

	billQuantity := billQuantities(bill, item)
	unitQuantities := Units()
	unitQuantity, _ := unitQuantities[unit]

	if billQuantity < unitQuantity {
		return false
	}

	if billQuantity-unitQuantity == 0 {
		delete(bill, item)
		return true
	}
	bill[item] = billQuantity - unitQuantity
	return true
}

// GetItem returns the quantity of an item that the customer has in his/her bill.
func GetItem(bill map[string]int, item string) (int, bool) {
	total := billQuantities(bill, item)
	if total == 0 {
		return 0, false
	}
	return total, true
}

func billQuantities(bill map[string]int, item string) int {
	return bill[item]
}
