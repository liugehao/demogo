package main

import (
	"database/sql"
	"errors"
	_ "github.com/lib/pq"
)

//var db *sql.DB = nil
//
//func OpenSql() error {
//  var err error = nil
//  db, err = sql.Open("postgres", "port=5432 user=postgres password=123456 dbname=postgres sslmode=disable")
//  return err
//}
//
//func GetDB() (*sql.DB, error) {
//  if db == nil {
//      return nil, errors.New("db hadn't open")
//  }
//  return db, nil
//}

var dbQueue chan *sql.DB

func Init(queue chan *sql.DB) {
	//dbQueue = queue
}
func open() *sql.DB {
	return <-dbQueue
}
func close(db *sql.DB) {
	dbQueue <- db
}

func dealResult(result sql.Result) error {
	affect, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affect <= 0 {
		return errors.New("DBExec no affect")
	}
	return nil
}
func ExecSql(Sql string, args ...interface{}) error {
	db := open()
	defer close(db)
	stmt, err := db.Prepare(Sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(args...)
	if err != nil {
		return err
	}
	return dealResult(result)
}

func QuerySql(Sql string, args ...interface{}) (*sql.Rows, error) {
	db := open()
	defer close(db)
	stmt, err := db.Prepare(Sql)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}
	return rows, err
}

func TxQuerySql(tx *sql.Tx, Sql string, args ...interface{}) (*sql.Stmt, *sql.Rows, error) {
	stmt, err := tx.Prepare(Sql)
	if err != nil {
		return nil, nil, err
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, nil, err
	}
	return stmt, rows, err
}

func TxExecSql(tx *sql.Tx, Sql string, args ...interface{}) error {
	stmt, err := tx.Prepare(Sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(args...)
	if err != nil {
		return err
	}
	return dealResult(result)
}

func ExecMultiSql(Sql string, member []string, args ...interface{}) error {
	db := open()
	defer close(db)
	stmt, err := db.Prepare(Sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, val := range member {
		allArgs := make([]interface{}, 0)
		allArgs = append(allArgs, val)
		allArgs = append(allArgs, args...)
		result, err := stmt.Exec(allArgs...)
		if err != nil {
			return err
		}
		err = dealResult(result)
		if err != nil {
			return err
		}
	}
	return nil
}

func TxExecMultiSql(tx *sql.Tx, Sql string, member []string, args ...interface{}) error {
	stmt, err := tx.Prepare(Sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, val := range member {
		allArgs := make([]interface{}, 0)
		allArgs = append(allArgs, val)
		allArgs = append(allArgs, args...)
		result, err := stmt.Exec(allArgs...)
		if err != nil {
			return err
		}
		err = dealResult(result)
		if err != nil {
			return err
		}
	}
	return nil
}