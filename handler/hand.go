package handler

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"vitess.io/vitess/go/mysql"
	"vitess.io/vitess/go/sqltypes"
	querypb "vitess.io/vitess/go/vt/proto/query"
)

type MultiHandler struct {
	//db *sql.DB
	db  *mysql.Conn
	db2 *sql.DB
}

func NewMultiHandler() *MultiHandler {
	//vitessdriver.Open()
	//db, err := sql.Open("mysql", "devops:devops_db_2013@tcp(10.0.1.164:3306)/test")
	//if err != nil {
	//	panic(err)
	//}
	//// See "Important settings" section.
	//db.SetConnMaxLifetime(time.Minute * 3)
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(10)
	db, err := mysql.Connect(context.Background(), &mysql.ConnParams{
		Host:  "10.0.1.164",
		Port:  3306,
		Uname: "devops",
		Pass:  "devops_db_2013",
	})
	if err != nil {
		panic(err)
	}
	return &MultiHandler{db: db}
}

// NewConnection is called when a connection is created.
// It is not established yet. The handler can decide to
// set StatusFlags that will be returned by the handshake methods.
// In particular, ServerStatusAutocommit might be set.
func (h *MultiHandler) NewConnection(c *mysql.Conn) {
	fmt.Println("NewConnection")

}

// ConnectionReady is called after the connection handshake, but
// before we begin to process commands.
func (h *MultiHandler) ConnectionReady(c *mysql.Conn) {
	fmt.Println("ConnectionReady")

}

// ConnectionClosed is called when a connection is closed.
func (h *MultiHandler) ConnectionClosed(c *mysql.Conn) {
	fmt.Println("ConnectionClosed")

}

// ComQuery is called when a connection receives a query.
// Note the contents of the query slice may change after
// the first call to callback. So the Handler should not
// hang on to the byte slice.
func (h *MultiHandler) ComQuery(c *mysql.Conn, query string, callback func(*sqltypes.Result) error) error {
	fmt.Println("ComQuery", query)
	result, more, err := h.db.ExecuteFetchMulti(query, 1000, true)
	fmt.Println(more, err)
	if err != nil {
		return err
	}

	return callback(result)
}

// ComPrepare is called when a connection receives a prepared
// statement query.
func (h *MultiHandler) ComPrepare(c *mysql.Conn, query string, bindVars map[string]*querypb.BindVariable) ([]*querypb.Field, error) {
	fmt.Println("ComPrepare", query, bindVars)

	return []*querypb.Field{}, nil
}

// ComStmtExecute is called when a connection receives a statement
// execute query.
func (h *MultiHandler) ComStmtExecute(c *mysql.Conn, prepare *mysql.PrepareData, callback func(*sqltypes.Result) error) error {
	fmt.Println("ComStmtExecute", prepare)
	
	return nil
}

// ComBinlogDumpGTID is called when a connection receives a ComBinlogDumpGTID request
func (h *MultiHandler) ComBinlogDumpGTID(c *mysql.Conn, gtidSet mysql.GTIDSet) error {
	fmt.Println("ComBinlogDumpGTID")
	return nil
}

// WarningCount is called at the end of each query to obtain
// the value to be returned to the client in the EOF packet.
// Note that this will be called either in the context of the
// ComQuery callback if the result does not contain any fields,
// or after the last ComQuery call completes.
func (h *MultiHandler) WarningCount(c *mysql.Conn) uint16 {
	fmt.Println("WarningCount")

	return 0
}

func (h *MultiHandler) ComResetConnection(c *mysql.Conn) {
	fmt.Println("ComResetConnection")
}

type ScanRow sqltypes.Value

func (r *ScanRow) Scan(src any) error {
	//sqltypes.NewDatetime()
	return nil
}
