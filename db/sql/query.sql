psql -U postgres -d teller -a -f "C:\Users\LENOVO\Desktop\log\query_copy_csv.sql"



INSERT INTO apis (name,url,key,value)
VALUES
('skn', 'https://apidev.banksinarmas.com/internal/transactions/transfer/v2.0/skn', 'X-Gateway-APIKey', '97817cac-d589-4d9c-b9bf-a874f0ff943d'),
('inquirytransfer', 'https://apidev.banksinarmas.com/internal/transactions/transfer/v2.0/inquiry', 'X-Gateway-APIKey', '97817cac-d589-4d9c-b9bf-a874f0ff943d'),
('internaltransfer', 'https://apidev.banksinarmas.com/internal/transactions/transfer/v2.0/internal', 'X-Gateway-APIKey', '97817cac-d589-4d9c-b9bf-a874f0ff943d'),
('getdetail', 'https://apidev.banksinarmas.com/internal/accounts/management/v2.0/getDetail', 'X-Gateway-APIKey', '97817cac-d589-4d9c-b9bf-a874f0ff943d'),
('advice', 'https://apidev.banksinarmas.com/internal/transactions/transfer/v2.0/advice', 'X-Gateway-APIKey', '97817cac-d589-4d9c-b9bf-a874f0ff943d');