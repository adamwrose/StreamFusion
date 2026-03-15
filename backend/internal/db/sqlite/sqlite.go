package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// DB wraps the SQLite connection for config, themes, and credentials.
type DB struct {
	conn *sql.DB
}

func New(path string) (*DB, error) {
	conn, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, fmt.Errorf("sqlite: open %q: %w", path, err)
	}
	if err := migrate(conn); err != nil {
		return nil, fmt.Errorf("sqlite: migrate: %w", err)
	}
	return &DB{conn: conn}, nil
}

func migrate(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS credentials (
			id       INTEGER PRIMARY KEY AUTOINCREMENT,
			platform TEXT    NOT NULL UNIQUE,
			key      TEXT    NOT NULL,
			secret   TEXT    NOT NULL DEFAULT ''
		);
		CREATE TABLE IF NOT EXISTS themes (
			id      INTEGER PRIMARY KEY AUTOINCREMENT,
			name    TEXT    NOT NULL UNIQUE,
			payload TEXT    NOT NULL  -- JSON blob of CSS variable overrides
		);
		CREATE TABLE IF NOT EXISTS config (
			key   TEXT PRIMARY KEY,
			value TEXT NOT NULL
		);
	`)
	return err
}

func (db *DB) Close() error {
	return db.conn.Close()
}

// GetCredential retrieves the API key/secret for a platform.
func (db *DB) GetCredential(platform string) (key, secret string, err error) {
	row := db.conn.QueryRow(
		`SELECT key, secret FROM credentials WHERE platform = ?`, platform,
	)
	err = row.Scan(&key, &secret)
	if err != nil {
		return "", "", fmt.Errorf("sqlite: get credential for %q: %w", platform, err)
	}
	return key, secret, nil
}

// UpsertCredential stores or updates an API credential.
func (db *DB) UpsertCredential(platform, key, secret string) error {
	_, err := db.conn.Exec(
		`INSERT INTO credentials(platform, key, secret) VALUES(?,?,?)
		 ON CONFLICT(platform) DO UPDATE SET key=excluded.key, secret=excluded.secret`,
		platform, key, secret,
	)
	return err
}

// GetTheme retrieves a theme JSON payload by name.
func (db *DB) GetTheme(name string) (string, error) {
	var payload string
	err := db.conn.QueryRow(`SELECT payload FROM themes WHERE name = ?`, name).Scan(&payload)
	if err != nil {
		return "", fmt.Errorf("sqlite: get theme %q: %w", name, err)
	}
	return payload, nil
}

// UpsertTheme stores or updates a theme JSON payload.
func (db *DB) UpsertTheme(name, payload string) error {
	_, err := db.conn.Exec(
		`INSERT INTO themes(name, payload) VALUES(?,?)
		 ON CONFLICT(name) DO UPDATE SET payload=excluded.payload`,
		name, payload,
	)
	return err
}
