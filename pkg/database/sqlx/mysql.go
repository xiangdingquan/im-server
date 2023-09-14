package sqlx

import (
	"context"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"

	"github.com/go-kratos/kratos/pkg/net/netutil/breaker"
	"github.com/go-kratos/kratos/pkg/time"

	"open.chat/pkg/log"
)

// Config mysql config.
type Config struct {
	Addr         string          // for trace
	DSN          string          // write data source name.
	ReadDSN      []string        // read data source name.
	Active       int             // pool
	Idle         int             // pool
	IdleTimeout  time.Duration   // connect max life time.
	QueryTimeout time.Duration   // query sql timeout
	ExecTimeout  time.Duration   // execute sql timeout
	TranTimeout  time.Duration   // transaction sql timeout
	Breaker      *breaker.Config // breaker
}

// NewMySQL new db and retry connection when has error.
func NewMySQL(c *Config) (db *DB) {
	if c.QueryTimeout == 0 || c.ExecTimeout == 0 || c.TranTimeout == 0 {
		panic("mysql must be set query/execute/transction timeout")
	}
	db, err := Open(c)
	if err != nil {
		log.Error("open mysql error(%v)", err)
		panic(err)
	}
	return
}

// TxWrapper
func TxWrapper(ctx context.Context, db *DB, txF func(*Tx, *StoreResult)) *StoreResult {
	result := &StoreResult{}

	tx, err := db.Begin(ctx)
	if err != nil {
		result.Err = err
		return result
	}

	defer func() {
		if result.Err != nil {
			tx.Rollback()
		} else {
			result.Err = tx.Commit()
		}
	}()

	txF(tx, result)
	return result
}

// Check if MySQL error is a Error Code: 1062. Duplicate entry ... for key ...
func IsDuplicate(err error) bool {
	if err == nil {
		return false
	}

	err = errors.Cause(err)
	myerr, ok := err.(*mysql.MySQLError)
	return ok && myerr.Number == 1062
}

func IsMissingDb(err error) bool {
	if err == nil {
		return false
	}
	err = errors.Cause(err)
	myerr, ok := err.(*mysql.MySQLError)
	return ok && myerr.Number == 1049
}
