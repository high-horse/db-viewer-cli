```sql
-- For MySQL: index query table in flat strng
SHOW FULL TABLES FROM your_database_name WHERE Table_type = 'BASE TABLE';

-- for postgres
SELECT table_name 
FROM information_schema.tables 
WHERE table_schema = 'public' AND table_type = 'BASE TABLE';

```