## Requirements
The software should be able to:
Create an account (e.g. opening a new credit card)
Keep track of charges (e.g. card swipes)
Keep track of payments
Provide the outstanding balance for any given day (such as "10 days after account opening")

## Solution

### language and enviroment
This challenge is written in golang(v1.8). all the tests and deployment where done using golang:1.8 docker image. all codes are located in the src folder.

### Description
The solution includes two structs: bank and account. bank holds list of credit card accounts and the date of the system. using methods define in bank struct, the user can Create an account, Keep track of charges, Keep track of payments and Provide the outstanding balance for any given day. Account struct holds the values associated with credit card and updates them.

### Usage
Bellow methods are defined under struct Bank.

CreateAccount(apr float32, creditLimit float32) CreateAccount creates new account with specified apr and creditLimit

Charge(id int, c float32) Charge apply a Charge with value c to account with specific id

GetListOfCharges(id int) GetListOfCharges prints and returns the list of Charges([]*transaction.Transaction)

MakePayment(id int, c float32) MakePayment apply a Payment with value c to an account with specific id

GetListOfPayments(id int) GetListOfPayments prints and returns the list of Payments([]*transaction.Transaction)

GetOutstandingBalance(id int) GetOutstandingBalance print and returns the GetoutstandingBalance

Getinterest(id int) Getinterest print and returns the interest

IncrementDateBy(d int) IncrementDateBy increments the Date by specifc number of days

### Tests
Unit Tests are impelemented for each method of bank and account structs located in the same package as the struct. use go tests to run them.

Test scripts for Test Scenario 1 and Test Scenario 2 are implemented with the same name. use go run [script name] to run them.
