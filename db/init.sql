-- Create a database named 'my_database'
CREATEDATABASE IFNOTEXISTSpersonal;
-- Switch to 'my_database'
USE personal;
-- Create a table for accounts
CREATETABLEIFNOTEXISTSaccounts(
account_idINTAUTO_INCREMENTPRIMARYKEY,
passwordVARCHAR(20)UNIQUE,
balanceDECIMAL(10, 2),
customer_idINT,
FOREIGNKEY(
customer_id
)REFERENCEScustomers(customer_id)
);
-- Create a table for customers
CREATETABLEIFNOTEXISTScustomers(
customer_idINTAUTO_INCREMENTPRIMARYKEY,
first_nameVARCHAR(50),
last_nameVARCHAR(50),
emailVARCHAR(100)UNIQUE,
phone_numberVARCHAR(20)
);
