-- Create the database (replace "your_database_name" with your desired name)
CREATEDATABASE your_database_name;
-- Use the newly created database USE your_database_name; Create the Customer table CREATE TABLE customer (
idINTNOTNULLAUTO_INCREMENTPRIMARYKEY,
nameVARCHAR(255) NOTNULL,
passwordVARCHAR(255) NOTNULL-- Login password (consider using hashing for security)
);
-- Create the Account table with a foreign key to Customer
CREATETABLEaccount(
idINTNOTNULLAUTO_INCREMENTPRIMARYKEY,
-- Unique identifier for account
customer_idINTNOTNULL,
amountDECIMAL(10,2)NOTNULLDEFAULT0.00,
FOREIGNKEY(
customer_id
)REFERENCEScustomer(id)-- Establish foreign key relationship

);
