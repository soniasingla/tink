package migration

import migrate "github.com/rubenv/sql-migrate"

func Get202012041103() *migrate.Migration {
	return &migrate.Migration{
		Id: "202010221010-remove-unique-index",
		Up: []string{`
                ALTER TABLE template DROP CONSTRAINT template_name_key;
                `},
	}
}
