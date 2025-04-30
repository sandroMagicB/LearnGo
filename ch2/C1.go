package main

import "fmt"

func main() {
	var insufficientFundMessage string = "Purchase failed. Insufficient funds."
	var purchaseSuccessMessage string = "Purchase successful."
	var accountBalance float64 = 100.0
	var bulkMessageCost float64 = 75.0
	var isPremiumUser bool = true
	var discountRate float64 = 0.10
	var finalCost float64

	// don't edit above this line

  // Calculate the final cost
  finalCost = bulkMessageCost
  if isPremiumUser {
      finalCost = bulkMessageCost * (1 - discountRate)
  }

  // Check if the user has enough funds
  if accountBalance >= finalCost {
      accountBalance -= finalCost // Deduct the final cost from the account balance
      fmt.Println(purchaseSuccessMessage)
  } else {
      fmt.Println(insufficientFundMessage)
  }
	// don't edit below this line

	fmt.Println("Account balance:", accountBalance)
}

