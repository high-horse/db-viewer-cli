package manager

import (
	"context"
	"sync"
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
	"github.com/redis/go-redis/v9"
)

type Manager interface {
	Add(Connection) error
	Get(id string)(Connection, bool)
	Remove(id string) error
	List() []Connection
	CloseAll() error
}

type Connection interface {
	ID() string
	
	Name() string
	Type() string
	DatabaseName() string

	Connect(ctx context.Context) error
	Disconnect() error
	Ping(ctx context.Context) error
	IsConnected() bool
}

type SQLConnection interface {
	Connection

	DB() *sql.DB
}

type NoSQLConnection interface {
	Connection

	Client() *mongo.Client
	DB() *mongo.Database
}

type RedisConnection interface {
    Connection

    Client() *redis.Client
}


type ConnectionManager struct {
	rw sync.RWMutex
	connections map[string]Connection
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{
		connections: make(map[string]Connection),
	}
}


func(m *ConnectionManager) Add(conn Connection) error {
	m.rw.RLock()
	defer m.rw.RUnlock()

	if conn == nil {
		return ErrInvalidConnection
	}

	id := conn.ID()

	if _, exists := m.connections[id]; exists {
		return ErrConnectionExists
	}

	m.connections[id] = conn

	return nil
}


func(m *ConnectionManager)Get(id string)(Connection, bool) {
	m.rw.RLock()
	defer m.rw.RUnlock()

	conn, ok := m.connections[id]
	return conn, ok
}

func (m *ConnectionManager)Remove(id string) error{
	m.rw.RLock()
	defer m.rw.RUnlock()

	conn, ok := m.connections[id]
	if !ok {
		return ErrConnectionNotFound
	}

	if err := conn.Disconnect(); err != nil {
		return  err
	}

	delete(m.connections, id)
	return nil
}

func (m *ConnectionManager)List() []Connection {
	m.rw.RLock()
	defer m.rw.RUnlock()

	list := make([]Connection, 0, len(m.connections))

	for _, c := range m.connections {
		list = append(list, c)
	}

	return list
}

func (m *ConnectionManager) CloseAll() error {
	m.rw.RLock()
	defer m.rw.RUnlock()

	for id, conn := range m.connections {
		_ = conn.Disconnect()
		delete(m.connections, id)
	}

	return  nil
}