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
	// "fmt"
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