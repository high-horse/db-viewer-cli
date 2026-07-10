CREATE TABLE IF NOT EXISTS connections (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    driver TEXT,
    host TEXT,
    port INTEGER,
    user TEXT,
    password TEXT,
    dbname TEXT
);

CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT,
    password TEXT
);

-- INSERT INTO connections (name, driver, host, port, user, password, dbname)
-- VALUES
--     ('Local SQLite', 'sqlite', '', NULL, '', '', 'app.db'),
--     ('Local PostgreSQL', 'postgres', 'localhost', 5432, 'postgres', 'password123', 'myapp'),
--     ('Development MySQL', 'mysql', '127.0.0.1', 3306, 'root', 'root123', 'dev_db'),
--     ('Production MariaDB', 'mariadb', '192.168.1.100', 3306, 'admin', 'securepass', 'production'),
--     ('SQL Server', 'sqlserver', 'db-server.local', 1433, 'sa', 'StrongPass123!', 'SalesDB'),
--     ('Oracle XE', 'oracle', 'localhost', 1521, 'system', 'oracle123', 'XE'),
--     ('Local MariaDB', 'mariadb', 'localhost', 3306, 'app_user', 'strong_password', 'app_database');