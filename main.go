package friction

import (
	"database/sql"
	"log"

	// use mysql
	_ "github.com/go-sql-driver/mysql"
)

func InitDB(dsn string) *sql.DB {
	log.Println(dsn)
	db, err := sql.Open("mysql", dsn)

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}

	return db
}

func ShowTables(db *sql.DB) ([]string, error) {

	rows, err := db.Query("SHOW TABLES")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tables := make([]string, 0)
	for rows.Next() {
		var ret string
		err := rows.Scan(&ret)
		if err != nil {
			return nil, err
		}
		tables = append(tables, ret)
	}

	return tables, nil
}

func GetIndexColumns(db *sql.DB, table string) ([]string, error) {
	log.Println(table)

	rows, err := db.Query("show index from " + table)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	columns := make([]string, 0)
	for rows.Next() {

		var (
			table        string
			nonUnique    string
			keyName      string
			seqInIndex   string
			columnName   string
			collation    string
			cardinality  int64
			subPart      sql.NullString
			packed       sql.NullString
			null         string
			indexType    string
			comment      string
			indexComment string
		)

		err := rows.Scan(
			&table,
			&nonUnique,
			&keyName,
			&seqInIndex,
			&columnName,
			&collation,
			&cardinality,
			&subPart,
			&packed,
			&null,
			&indexType,
			&comment,
			&indexComment,
		)

		if err != nil {
			return nil, err
		}
		columns = append(columns, columnName)
	}

	return columns, nil
}

func WarmUp(db *sql.DB, table string, column string, count int) error {

	for i := 0; i < count; i++ {
		log.Println(table, column, i+1)
		rows, err := db.Query("SELECT COUNT(" + column + ") FROM " + table)
		if err != nil {
			return err
		}
		defer rows.Close()
	}

	return nil
}
