package main

func getMonthlyPrice(tier string) int {
	// ?
  switch tier {
    case "basic":
        return 10000    // $100.00 = 10000 pennies
    case "premium":
        return 15000    // $150.00 = 15000 pennies
    case "enterprise":
        return 50000    // $500.00 = 50000 pennies
    default:
        return 0        // Unknown tier returns 0 pennies
  }
}

