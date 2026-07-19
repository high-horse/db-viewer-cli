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



=====

	// query := `
	// 	-- select count(*) from customers;
	// 	-- desc customers;
	// 	SELECT
	// 		c.customerName,
	// 		c.country,
	// 		COUNT(o.orderNumber) AS total_orders,
	// 		SUM(oD.quantityOrdered * oD.priceEach) AS total_spent
	// 	FROM customers c
	// 	JOIN orders o 
	// 		ON c.customerNumber = o.customerNumber
	// 	JOIN orderdetails oD
	// 		ON o.orderNumber = oD.orderNumber
	// 	GROUP BY
	// 		c.customerNumber,
	// 		c.customerName,
	// 		c.country
	// 	HAVING total_spent > 10000
	// 	ORDER BY total_spent DESC
	// 	LIMIT 10;
	// `
	// query := `
	// 	SELECT
	// 		p.productName,
	// 		p.productLine,
	// 		p.buyPrice,
	// 		(
	// 			SELECT AVG(p2.buyPrice)
	// 			FROM products p2
	// 			WHERE p2.productLine = p.productLine
	// 		) AS line_average_price
	// 	FROM products p
	// 	WHERE p.buyPrice > (
	// 		SELECT AVG(buyPrice)
	// 		FROM products
	// 	)
	// 	ORDER BY p.buyPrice DESC;
	// `

	// query := `
	// 	WITH customer_orders AS (
	// 		SELECT
	// 			customerNumber,
	// 			COUNT(orderNumber) AS order_count
	// 		FROM orders
	// 		GROUP BY customerNumber
	// 	)
	// 	SELECT
	// 		c.customerName,
	// 		c.country,
	// 		co.order_count
	// 	FROM customers c
	// 	JOIN customer_orders co
	// 		ON c.customerNumber = co.customerNumber
	// 	ORDER BY co.order_count DESC;
	// `

	// query := `
	// 	SELECT
	// 		productName,
	// 		productLine,
	// 		buyPrice,
	// 		RANK() OVER (
	// 			PARTITION BY productLine
	// 			ORDER BY buyPrice DESC
	// 		) AS price_rank
	// 	FROM products;
	// `

	// query := `
	// 	SELECT
	// 		TABLE_NAME,
	// 		TABLE_ROWS,
	// 		ENGINE,
	// 		CREATE_TIME
	// 	FROM information_schema.tables
	// 	WHERE TABLE_SCHEMA = DATABASE()
	// 	ORDER BY TABLE_ROWS DESC;
	// `