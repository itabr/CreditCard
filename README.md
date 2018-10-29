# CreditCard

## Requirements
Given a credit card that functions as follows:
Each card has an APR and Credit Limit.
Interest is calculated daily at the close of each day, but not applied.
Interest is applied to the balance at the close of each 30-day period (opening day excluded).
  e.g., asking for the balance on days 28 and 29 will give the same results, but asking for balance on day 30 will give the balance + all interest accrued.
The software should be able to:
Create an account (e.g. opening a new credit card)
Keep track of charges (e.g. card swipes)
Keep track of payments
Provide the outstanding balance for any given day (such as "10 days after account opening")

#### Test Scenario 1
A customer opens a credit card with a $1,000.00 limit at a 35% APR.
The customer charges $500 on opening day (outstanding balance becomes $500).
The customer does not make any more charges or any payments for 30 days.
The total outstanding balance owed 30 days after opening should be $514.38.
500 * (0.35 / 365) * 30 = 14.38

#### Test Scenario 2
A customer opens a credit card with a $1,000.00 limit at a 35% APR.
The customer charges $500 on opening day (outstanding balance becomes $500).
15 days after opening, the customer pays $200 (outstanding balance becomes $300).
25 days after opening, the customer charges another $100 (outstanding balance becomes $400).
The total outstanding balance owed on day 30 should be $411.99.
(500 * 0.35 / 365 * 15) + (300 * 0.35 / 365 * 10) + (400 * 0.35 / 365 * 5) = 11.99

## Solution

#### language and enviroment
This challenge is written in golang(v1.8). all the tests and deployment where done using golang:1.8 docker image. all codes are located in the src folder.

#### Description
The solution includes two structs: bank and account. bank holds list of credit card accounts and the date of the system. using methods define in bank struct, the user can Create an account, Keep track of charges, Keep track of payments and Provide the outstanding balance for any given day. Account struct holds the values associated with credit card and updates them.

#### Usage
Bellow methods are defined under struct Bank.

CreateAccount(apr float32, creditLimit float32) CreateAccount creates new account with specified apr and creditLimit

Charge(id int, c float32) Charge apply a Charge with value c to account with specific id

GetListOfCharges(id int) GetListOfCharges prints and returns the list of Charges([]*transaction.Transaction)

MakePayment(id int, c float32) MakePayment apply a Payment with value c to an account with specific id

GetListOfPayments(id int) GetListOfPayments prints and returns the list of Payments([]*transaction.Transaction)

GetOutstandingBalance(id int) GetOutstandingBalance print and returns the GetoutstandingBalance

Getinterest(id int) Getinterest print and returns the interest

IncrementDateBy(d int) IncrementDateBy increments the Date by specifc number of days

#### Tests
Unit Tests are impelemented for each method of bank and account structs located in the same package as the struct. use go tests to run them.

Test scripts for Test Scenario 1 and Test Scenario 2 are implemented with the same name. use go run [script name] to run them.

## Notes and Assumptions
I assumed that interest is calculated every day and keep increasing until the end of the month and after applying the interest to the outstanding balance it gets reset.
I assumed that bank controls that account and all the charges and payments have to go through the bank and bank controls the date.