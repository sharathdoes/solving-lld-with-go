package database
import (
	"log"
	"path/filepath"
	"runtime"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations(dbURL string) {
	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Join(filepath.Dir(b), "migrations")

	m, err := migrate.New(
		"file://"+basepath,
		dbURL,
	)
	if err != nil {
		log.Fatal("Migration init error:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("Migration run error:", err)
	}

	log.Println("✅ Migrations ran successfully")
}
