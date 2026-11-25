package store

import (
	"database/sql"
	"embed"
	"io/fs"
	"log"
	_ "modernc.org/sqlite"
)

type Store struct {
	db *sql.DB
	Games *GameStore
	Platform *PlatformStore
	Playthrough *PlaythroughStore
}

func NewStore (dbPath string ) (*Store, error) {
	if dbPath == "" {
		dbPath = "journal.db"
	}

	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %s", err.Error())
	}

	log.Println("Database connection successful")

	return &Store{
		db: db, 
		Games: &GameStore{db: db},
		Platform: &PlatformStore{db: db},
		Playthrough: &PlaythroughStore{db: db},
	}, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

//go:embed db_queries/schema.sql
var schemaFS embed.FS

//go:embed db_queries/seed.sql
var seedFS embed.FS
func (s *Store) MakeMigrations() {
	schemaBytes, err := fs.Glob(schemaFS, "db_queries/schema.sql")
	if err != nil {
		log.Fatalf ("Failed to load schema.sql: %v", err)
	}
	
	statement, err := schemaFS.ReadFile(schemaBytes[0])
	_, err = s.db.Exec(string(statement))
	if err != nil {
		log.Fatalf("SOMETHING'S WRONG, I CAN FEEL IT! %s", err)
	}

	seedBytes, err := fs.Glob(seedFS, "db_queries/seed.sql")
	if err != nil {
		log.Fatalf ("Failed to load query.sql: %v", err)
	}

	statement, err = seedFS.ReadFile(seedBytes[0])
	_, err = s.db.Exec(string(statement))
	if err != nil {
		log.Fatalf("SOMETHING'S WRONG, I CAN FEEL IT! %s", err)
	}

	log.Println("Migration successful!")
}
