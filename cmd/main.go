package main

import (
	"context"
	"db-viewer/internal/db"
	manager "db-viewer/internal/engine/connectionManager"
	"db-viewer/internal/engine/drivers/mysql"
	"db-viewer/internal/engine/drivers/postgres"
	"db-viewer/internal/engine/entities"
	"db-viewer/internal/engine/factory"
	"db-viewer/internal/engine/transports"
	"fmt"

	"log"

	tea "charm.land/bubbletea/v2"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Println("This log includes the line number!")

	log.Println("starting ...")
	log.Println("inint factory")
	factory := factory.New()
	log.Println("registring driver to factory mysql")

	factory.Register(mysql.NewDriver())
	log.Println("registring driver to factory pgx")

	factory.Register(postgres.NewDriver())
	log.Println("drivers registered")

	manager := manager.NewConnectionManager()

	config := entities.ConnectionConfig{
		ID:       "local",
		Name:     "Local MySQL",
		Type:     "mysql",
		Host:     "127.0.0.1",
		Port:     3306,
		User:     "app_user",
		Password: "strong_password",
		Database: "app_database",
	}
	transport := transports.NewDirect(config.Host, config.Port)

	conn, err := factory.Create(context.TODO(), config, transport)
	if err != nil {
		log.Fatal("error ", err)
	}
	manager.Add(conn)

	if err := conn.Connect(context.Background()); err != nil {
		log.Fatal("connection failed:", err)
	}
	log.Println("conn status", conn.IsConnected(), conn.Name())


	driver, err := factory.Driver(conn.Type())
	if err != nil {
		log.Fatal("executor selection failed:", err)
	}

	exec := driver.Executor()

	query := `
		SELECT
			e.employeeNumber,
			CONCAT(e.firstName, ' ', e.lastName) AS employee_name,
			o.city,
			COALESCE(o.phone, 'No phone') AS office_phone
		FROM employees e
		LEFT JOIN offices o
			ON e.officeCode = o.officeCode
		ORDER BY employee_name;
	`
	result, err := exec.Execute(
		context.Background(),
		conn,
		query,
	)


	if err != nil {
		log.Fatal("query execution failed:", err)
	}

	log.Println("Query executed successfully")
	log.Println("Duration:", result.Duration)
	log.Println("Rows affected:", result.RowsAffected)
	// fmt.Println("Columns:")
	// for _, col := range result.Columns {
	// 	fmt.Println("Name:", col.Name)
	// 	fmt.Println("Type:", col.DatabaseType)
	// }

	// fmt.Println("Rows:")
	// for _, row := range result.Rows {
	// 	for i, value := range row {
	// 		fmt.Printf("%s = %v\n", result.Columns[i].Name, value)
	// 	}
	// }

	inspect := driver.Inspector()

	tables, err := inspect.ListTables(context.Background(), conn)
	if err != nil {
		fmt.Println("error in insptect tables", err)
	}

	fmt.Println("tables count ", len(tables))

	cols, err := inspect.ListColumns(context.Background(), conn, "customers")
	if err != nil {
		fmt.Println("error in inspect columns", err)
	}
	fmt.Println("columns count", len(cols))
	fmt.Println("columns:", cols)

}




func main_old() {
	_, err := db.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	p := tea.NewProgram(initAppStateModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}