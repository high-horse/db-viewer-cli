```sql
-- For MySQL: index query table in flat strng
SHOW FULL TABLES FROM your_database_name WHERE Table_type = 'BASE TABLE';

-- for postgres
SELECT table_name 
FROM information_schema.tables 
WHERE table_schema = 'public' AND table_type = 'BASE TABLE';

```


preveledge mariadb query
CREATE USER 'app_user'@'localhost' IDENTIFIED BY 'strong_password';
GRANT ALL PRIVILEGES ON app_database.*  TO 'app_user'@'localhost';
FLUSH PRIVILEGES;

CREATE DATABASE app_database;
CREATE USER 'app_user'@'127.0.0.1' IDENTIFIED BY 'strong_password';
GRANT ALL PRIVILEGES ON app_database.* TO 'app_user'@'127.0.0.1';
FLUSH PRIVILEGES;
