package migration

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

func Run(db *sql.DB) error {

	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS schema_migrations (
			version VARCHAR(255) PRIMARY KEY,
			applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return err
	}

	files, err := filepath.Glob("migrations/*.sql")
	if err != nil {
		return err
	}

	sort.Strings(files)

	for _, file := range files {

		var exists bool
		err = db.QueryRow(
			`SELECT EXISTS(
				SELECT 1
				FROM schema_migrations
				WHERE version=$1
			)`,
			file,
		).Scan(&exists)

		if err != nil {
			return err
		}

		if exists {
			fmt.Printf("⏭️  Migration already applied: %s\n",
				filepath.Base(file))
			continue
		}

		content, err := os.ReadFile(file)
		if err != nil {
			return err
		}

		_, err = db.Exec(string(content))
		if err != nil {
			return fmt.Errorf("%s: %w", file, err)
		}

		_, err = db.Exec(
			`INSERT INTO schema_migrations(version)
			 VALUES($1)`,
			file,
		)

		if err != nil {
			return err
		}

		fmt.Printf("✅ Applied migration %s\n", filepath.Base(file))
	}

	return nil
}
