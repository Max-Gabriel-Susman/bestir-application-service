package application

import (
	"context"
	"fmt"

	"github.com/Max-Gabriel-Susman/bestir-identity-service/internal/foundation/database"
	"github.com/gocraft/dbr/v2"
)

func NewMySQLStore(conn *dbr.Connection) *MySQLStorage {
	return &MySQLStorage{conn: conn, sess: conn.NewSession(nil)}
}

type MySQLStorage struct {
	conn *dbr.Connection
	sess *dbr.Session
}

var (
	applicationTable = database.NewTable("application", Application{})
)

func (s *MySQLStorage) ListApplications(ctx context.Context) ([]Application, error) {
	query := s.sess.Select(applicationTable.Columns...).
		From(applicationTable.Name)

	applications := []Application{}

	if _, err := query.LoadContext(ctx, &applications); err != nil {
		return applications, database.ClassifyError(err)
	}

	return applications, nil
}

func (s *MySQLStorage) getApplicationByIdempotencyKey(ctx context.Context, idempotencyKey string) (Application, error) {
	var application Application
	err := s.sess.Select(applicationTable.Columns...).
		From(applicationTable.Name).
		Where("idempotency_key = ?", idempotencyKey).
		LoadOneContext(ctx, &application)
	return application, database.ClassifyError(err)
}

func (s *MySQLStorage) CreateApplication(ctx context.Context, application Application) error {
	fmt.Println("Create Account C invoked")
	_, err := s.sess.InsertInto(applicationTable.Name).
		Columns(applicationTable.Columns...).
		Record(application).
		ExecContext(ctx)
	return database.ClassifyError(err)
}
