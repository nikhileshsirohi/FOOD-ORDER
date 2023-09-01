package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// make subTotalBill as global variable to make it easily accessible in case customer modifies her order.
// subTotalBill means total bill but excluding taxes.
var subTotalBill float64

// make a map of customerOrder in which we will store the items ordered as "key" and no. of plates as "value".
var customerOrder = make(map[string]uint, 0)

func main() {
	var customerName string
	fmt.Println("What is your first name?")
	fmt.Scan(&customerName)

	greet(customerName)
	orderItems()
	displayGeneratingBill() //just displays that "generating bill" in a fancy manner.
	generateBill()
	modifyOrder()
	printFinalBill()
	visitAgain(customerName)
}

// greet.go
func greet(customerName string) {
	fmt.Printf("%52s %s\n", "Hello ", customerName)
	fmt.Printf("%72s\n", "_/\\_ Welcome to ABC Restaurant ! _/\\_ \n\n")

}

func visitAgain(customerName string) {
	fmt.Println()
	fmt.Printf("%17s", " ")
	fmt.Printf("_/\\_ Thank you %v for visiting ABC Restaurant ! _/\\_\n", customerName)
	fmt.Printf("%55s\n", "We hope to see you again!")
	fmt.Printf("%51s\n", "Have a nice day! \n\n")
}

// menu.go
type Menu struct {
	itemNo    uint
	itemName  string
	itemPrice float64
}

var menu = []Menu{
	//itemNo  itemName   itemPrice
	{1, "Adrakh Chai", 20},
	{2, "Filter Coffee", 35},
	{3, "Chhas", 35},
	{4, "Lassi", 40},
	{5, "Mango Lassi", 60},
	{6, "Paneer Tikka", 100},
	{7, "Tandoori Malai Chaap", 120},
	{8, "Tandoori Afgani Chaap", 130},
	{9, "Honey Chilli Potato", 80},
	{10, "Veg. Sandwich", 50},
	{11, "Veg. Masala Maggi", 60},
	{12, "Crispy Corn", 100},
	{13, "Cream Roll", 25},
}

func printMenu() {
	fmt.Printf("%15s\n", "Menu")
	fmt.Printf("%s\n", strings.Repeat("-", 35))
	fmt.Printf("%-7s %6s    %12s\n", "S.No.", "Price", "Item Name")
	fmt.Printf("%s\n", strings.Repeat("-", 35))

	for _, element := range menu {
		fmt.Printf(" %-7d %.2f    %-4s\n", element.itemNo, element.itemPrice, element.itemName)
	}
	fmt.Printf("%s", strings.Repeat("-", 35))
	fmt.Println()
}

// modify.go
func modifyOrder() {
	for {
		var isOrderOK string
		fmt.Println("Do you want to change your order?[y/n]")
		fmt.Scan(&isOrderOK)
		if isOrderOK != "y" {
			return
		}
		var serialNo uint
		var modifyType uint

		fmt.Println("Please enter the respective no. to proceed further: ")
		fmt.Println("Press '1' to update item quantity.")
		fmt.Println("Press '2' to delete an item from the order list.")
		fmt.Println("Press '3' to add item(s) in the order list.")
		fmt.Scan(&modifyType)

		switch modifyType {
		case 1:
			printMenu()
			fmt.Println("Please enter the S.No. of the item to be updated: ")
			fmt.Scan(&serialNo)
			updateQuantity(serialNo)

		case 2:
			printMenu()
			fmt.Println("Please enter the S.No. of the item to be deleted: ")
			fmt.Scan(&serialNo)
			delFromOrder(serialNo)

		case 3:
			insertIntoOrder()
		default:
			return
		}
		displayGeneratingBill()
		generateBill()
	}
}

func updateQuantity(serialNo uint) {

	var newQuantity uint

	for _, targetItem := range menu {

		if serialNo == targetItem.itemNo {

			itemName := targetItem.itemName        //name of the item whose quantity to be updated
			oldQuantity := customerOrder[itemName] //current quantity of that item

			fmt.Printf("Current quantity of %v is %v.\n", itemName, oldQuantity)
			fmt.Printf("Now, enter the updated quantity of %v to be ordered: \n", itemName)
			fmt.Scan(&newQuantity)

			//in case new quantity is '0' then delete that item from the order
			if newQuantity == 0 {
				delFromOrder(serialNo)
				return
			}
			fmt.Printf("")

			//update quantity of item
			customerOrder[targetItem.itemName] = newQuantity
			fmt.Printf("Updated the quantity of %v from %v to %v.\n", itemName, oldQuantity, newQuantity)

			//update bill
			subTotalBill -= float64(oldQuantity) * float64(targetItem.itemPrice) //delete the item price for old quantity
			subTotalBill += float64(newQuantity) * float64(targetItem.itemPrice) //add the item price for updated quantity

			break
		}
	}

}

func delFromOrder(serialNo uint) {

	for _, targetItem := range menu {
		if serialNo == targetItem.itemNo {
			itemName := targetItem.itemName //name of the item who is to be removed from the list
			//update bill; customerOrder[itemName] -> it results in the quantity of that item
			subTotalBill -= float64(customerOrder[itemName]) * float64(targetItem.itemPrice)
			//delete item
			delete(customerOrder, itemName)
			fmt.Printf("%v removed from the order list.\n", itemName)
			break
		}
	}
}
func insertIntoOrder() {
	//add an item in the order
	orderItems()
}

// order.go
// order using this function
func orderItems() {
	printMenu()
	var itemNumber uint
	var noOfPlates uint

	for {
		fmt.Println()
		fmt.Println("Enter '0' to exit.")
		fmt.Print("Enter the serial no. of the item to order: ")

		fmt.Scan(&itemNumber)
		if itemNumber == 0 {
			break
		}

		var choiceName string
		var itemPrice float64

		for index, item := range menu {
			if index+1 == int(itemNumber) {
				choiceName = item.itemName
				itemPrice = item.itemPrice
				break
			}
		}
		fmt.Printf("How many %v do you want: ", choiceName)
		fmt.Scan(&noOfPlates)
		if noOfPlates == 0 {
			continue
		} else {
			for index /*, item*/ := range menu {
				if index+1 == int(itemNumber) { //convert optionNumber into int becoz it is of type uint
					customerOrder[choiceName] += noOfPlates
					//alternative way
					//customerOrder[item.itemName] += noOfPlates // adding customerOrder[item.itemName] again in case you order that item again
					subTotalBill += itemPrice * float64(noOfPlates)
					//alternative way
					//subTotalBill += item.itemPrice*noOfPlates
					break
				}
			}
			fmt.Printf("\nYou just ordered %v %v which amounts to ₹%v.\n", noOfPlates, choiceName, itemPrice*float64(noOfPlates))
			//print what you ordered till now
			orderTillNow()
		}
		fmt.Println()
	}
}

// print it everytime you add an item
func orderTillNow() {
	//Print what you've ordered till now
	fmt.Println("\nYour order till now: ")
	fmt.Printf("%s\n", strings.Repeat("-", 32))
	fmt.Printf(" %-12s %s\n", "Quantity", "Item")
	fmt.Printf("%s\n", strings.Repeat("-", 32))

	for i := range customerOrder {
		fmt.Printf(" %-12v %v\n", customerOrder[i], i)
	}

	fmt.Printf("%s\n", strings.Repeat("-", 32))
}

// bill.go
// just beautifying my code :P
func displayGeneratingBill() {
	fmt.Println()
	billDisplayText := "************************************* Generating Bill *************************************"
	for _, element := range billDisplayText {
		fmt.Printf("%c", element) // if you use "%v" instead of "%c" then convert element into string, as shown in the comment below
		// fmt.Print("%v", string(element))
		time.Sleep(time.Millisecond * 15)
	}
}

// prints item name, price, quantity and total price and sub total amount.
func generateBill() {

	fmt.Println()
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	fmt.Printf(" %-30s %-20s %-20s %-20s\n", "Item Name", "Price(₹)", "Quantity", "Total Price(₹)")
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))

	//prints the details of the food item that you've ordered.
	printOrderData()

	fmt.Printf("+%s+\n", strings.Repeat("-", 90))

	//print sub total amount in a cool way!
	fmt.Printf("%47s", " ")
	for _, element := range "Sub Total(excluding GST): ₹" {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Printf("%.2f\n", subTotalBill)

}

// prints the data of the items that you ordered.
func printOrderData() {
	for key := range customerOrder {
		//key -> it is the key values
		for _, element := range menu {
			if key == element.itemName {
				//customerOrder[key] -> will provide the no. of plates of that item
				//float64(customerOrder[key])*element.itemPrice -> this will result in the cost of each item
				totalCostOfItem := float64(customerOrder[key]) * element.itemPrice
				fmt.Printf(" %-30s %-20.2f %-20v %-20.2f\n", key, element.itemPrice, customerOrder[key], totalCostOfItem)
			}
		}
	}
	fmt.Println()
}

func printFinalBill() {
	for _, element := range "Here is your final bill:-" {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Println()

	fmt.Printf("\n%52s\n", "ABC Restaurent")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%s\n", strings.Repeat("*", 91))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%86s\n", "Near MMM Engineering College, First Floor, ABC Restaurent, Gorakhpur, Gorakhpur 273010, U.P")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%50s\n", "Tel: 92145623XX")
	fmt.Printf("%60s\n\n", "Email: abc.restaurent@gmail.com")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%s", strings.Repeat("-", 42))
	fmt.Printf("%s", "INVOICE")
	fmt.Printf("%s\n", strings.Repeat("-", 42))
	time.Sleep(time.Millisecond * 200)

	rand.Seed(time.Now().Unix()) //necessary to produce random integers
	fmt.Printf(" Ticket No: %d\n", rand.Intn(550)+1)

	fmt.Printf(" Date: %v\n", time.Now().Local().Format("02-January-2006")) //display date
	fmt.Printf(" Time: %v", time.Now().Local().Format("3:4 PM"))            //display time
	time.Sleep(time.Millisecond * 200)

	generateBill() //prints details of the bill,like, item name, price, quantity and total price and sub total amount.

	tax := 5 * subTotalBill / (100)
	grandTotal := subTotalBill + tax

	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%71s: ₹%.2f\n", "GST", tax) //display tax amount
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%71s: ₹%.2f\n", "Grand Total", grandTotal) //display final bill
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))

}
