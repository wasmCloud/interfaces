// sqldb: wasmcloud database capability contract
package sqldb

import (
	"github.com/wasmcloud/actor-tinygo"           //nolint
	cbor "github.com/wasmcloud/tinygo-cbor"       //nolint
	msgpack "github.com/wasmcloud/tinygo-msgpack" //nolint
)

// Metadata about a Column in the result set
type Column struct {
	// column ordinal
	Ordinal uint32
	// Column name in the result
	Name string
	// column data type as reported by the database
	DbType string
}

// MEncode serializes a Column using msgpack
func (o *Column) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("ordinal")
	encoder.WriteUint32(o.Ordinal)
	encoder.WriteString("name")
	encoder.WriteString(o.Name)
	encoder.WriteString("dbType")
	encoder.WriteString(o.DbType)

	return encoder.CheckError()
}

// MDecodeColumn deserializes a Column using msgpack
func MDecodeColumn(d *msgpack.Decoder) (Column, error) {
	var val Column
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "ordinal":
			val.Ordinal, err = d.ReadUint32()
		case "name":
			val.Name, err = d.ReadString()
		case "dbType":
			val.DbType, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a Column using cbor
func (o *Column) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("ordinal")
	encoder.WriteUint32(o.Ordinal)
	encoder.WriteString("name")
	encoder.WriteString(o.Name)
	encoder.WriteString("dbType")
	encoder.WriteString(o.DbType)

	return encoder.CheckError()
}

// CDecodeColumn deserializes a Column using cbor
func CDecodeColumn(d *cbor.Decoder) (Column, error) {
	var val Column
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "ordinal":
			val.Ordinal, err = d.ReadUint32()
		case "name":
			val.Name, err = d.ReadString()
		case "dbType":
			val.DbType, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// List of columns in the result set returned by a Query operation
type Columns []Column

// MEncode serializes a Columns using msgpack
func (o *Columns) MEncode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.MEncode(encoder)
	}

	return encoder.CheckError()
}

// MDecodeColumns deserializes a Columns using msgpack
func MDecodeColumns(d *msgpack.Decoder) (Columns, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]Column, 0), err
	}
	size, err := d.ReadArraySize()
	if err != nil {
		return make([]Column, 0), err
	}
	val := make([]Column, size)
	for i := uint32(0); i < size; i++ {
		item, err := MDecodeColumn(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// CEncode serializes a Columns using cbor
func (o *Columns) CEncode(encoder cbor.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		item_o.CEncode(encoder)
	}

	return encoder.CheckError()
}

// CDecodeColumns deserializes a Columns using cbor
func CDecodeColumns(d *cbor.Decoder) (Columns, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([]Column, 0), err
	}
	size, indef, err := d.ReadArraySize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite arrays not supported")
	}
	if err != nil {
		return make([]Column, 0), err
	}
	val := make([]Column, size)
	for i := uint32(0); i < size; i++ {
		item, err := CDecodeColumn(d)
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// Result of an Execute operation
type ExecuteResult struct {
	// the number of rows affected by the query
	RowsAffected uint64
	// optional error information.
	// If error is included in the QueryResult, other values should be ignored.
	Error *SqlDbError
}

// MEncode serializes a ExecuteResult using msgpack
func (o *ExecuteResult) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("rowsAffected")
	encoder.WriteUint64(o.RowsAffected)
	encoder.WriteString("error")
	if o.Error == nil {
		encoder.WriteNil()
	} else {
		o.Error.MEncode(encoder)
	}

	return encoder.CheckError()
}

// MDecodeExecuteResult deserializes a ExecuteResult using msgpack
func MDecodeExecuteResult(d *msgpack.Decoder) (ExecuteResult, error) {
	var val ExecuteResult
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "rowsAffected":
			val.RowsAffected, err = d.ReadUint64()
		case "error":
			fval, err := MDecodeSqlDbError(d)
			if err != nil {
				return val, err
			}
			val.Error = &fval
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a ExecuteResult using cbor
func (o *ExecuteResult) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("rowsAffected")
	encoder.WriteUint64(o.RowsAffected)
	encoder.WriteString("error")
	if o.Error == nil {
		encoder.WriteNil()
	} else {
		o.Error.CEncode(encoder)
	}

	return encoder.CheckError()
}

// CDecodeExecuteResult deserializes a ExecuteResult using cbor
func CDecodeExecuteResult(d *cbor.Decoder) (ExecuteResult, error) {
	var val ExecuteResult
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "rowsAffected":
			val.RowsAffected, err = d.ReadUint64()
		case "error":
			fval, err := CDecodeSqlDbError(d)
			if err != nil {
				return val, err
			}
			val.Error = &fval
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// An optional list of arguments to be used in the SQL statement.
// When a statement uses question marks '?' for placeholders,
// the capability provider will replace the specified arguments during execution.
// The command must have exactly as many placeholders as arguments, or the request will fail.
// The members are CBOR encoded.
type Parameters [][]byte

// MEncode serializes a Parameters using msgpack
func (o *Parameters) MEncode(encoder msgpack.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		encoder.WriteByteArray(item_o)
	}

	return encoder.CheckError()
}

// MDecodeParameters deserializes a Parameters using msgpack
func MDecodeParameters(d *msgpack.Decoder) (Parameters, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([][]byte, 0), err
	}
	size, err := d.ReadArraySize()
	if err != nil {
		return make([][]byte, 0), err
	}
	val := make([][]byte, size)
	for i := uint32(0); i < size; i++ {
		item, err := d.ReadByteArray()
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

// CEncode serializes a Parameters using cbor
func (o *Parameters) CEncode(encoder cbor.Writer) error {

	encoder.WriteArraySize(uint32(len(*o)))
	for _, item_o := range *o {
		encoder.WriteByteArray(item_o)
	}

	return encoder.CheckError()
}

// CDecodeParameters deserializes a Parameters using cbor
func CDecodeParameters(d *cbor.Decoder) (Parameters, error) {
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return make([][]byte, 0), err
	}
	size, indef, err := d.ReadArraySize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite arrays not supported")
	}
	if err != nil {
		return make([][]byte, 0), err
	}
	val := make([][]byte, size)
	for i := uint32(0); i < size; i++ {
		item, err := d.ReadByteArray()
		if err != nil {
			return val, err
		}
		val = append(val, item)
	}
	return val, nil
}

type PingResult struct {
	// Optional error information.
	Error *SqlDbError
}

// MEncode serializes a PingResult using msgpack
func (o *PingResult) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(1)
	encoder.WriteString("error")
	if o.Error == nil {
		encoder.WriteNil()
	} else {
		o.Error.MEncode(encoder)
	}

	return encoder.CheckError()
}

// MDecodePingResult deserializes a PingResult using msgpack
func MDecodePingResult(d *msgpack.Decoder) (PingResult, error) {
	var val PingResult
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "error":
			fval, err := MDecodeSqlDbError(d)
			if err != nil {
				return val, err
			}
			val.Error = &fval
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a PingResult using cbor
func (o *PingResult) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(1)
	encoder.WriteString("error")
	if o.Error == nil {
		encoder.WriteNil()
	} else {
		o.Error.CEncode(encoder)
	}

	return encoder.CheckError()
}

// CDecodePingResult deserializes a PingResult using cbor
func CDecodePingResult(d *cbor.Decoder) (PingResult, error) {
	var val PingResult
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "error":
			fval, err := CDecodeSqlDbError(d)
			if err != nil {
				return val, err
			}
			val.Error = &fval
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Result of a query
type QueryResult struct {
	// number of rows returned
	NumRows uint64
	// description of columns returned
	Columns Columns
	// result rows, encoded in CBOR as
	// an array (rows) of arrays (fields per row)
	Rows []byte
	// optional error information.
	// If error is included in the QueryResult, other values should be ignored.
	Error *SqlDbError
}

// MEncode serializes a QueryResult using msgpack
func (o *QueryResult) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(4)
	encoder.WriteString("numRows")
	encoder.WriteUint64(o.NumRows)
	encoder.WriteString("columns")
	o.Columns.MEncode(encoder)
	encoder.WriteString("rows")
	encoder.WriteByteArray(o.Rows)
	encoder.WriteString("error")
	if o.Error == nil {
		encoder.WriteNil()
	} else {
		o.Error.MEncode(encoder)
	}

	return encoder.CheckError()
}

// MDecodeQueryResult deserializes a QueryResult using msgpack
func MDecodeQueryResult(d *msgpack.Decoder) (QueryResult, error) {
	var val QueryResult
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "numRows":
			val.NumRows, err = d.ReadUint64()
		case "columns":
			val.Columns, err = MDecodeColumns(d)
		case "rows":
			val.Rows, err = d.ReadByteArray()
		case "error":
			fval, err := MDecodeSqlDbError(d)
			if err != nil {
				return val, err
			}
			val.Error = &fval
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a QueryResult using cbor
func (o *QueryResult) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(4)
	encoder.WriteString("numRows")
	encoder.WriteUint64(o.NumRows)
	encoder.WriteString("columns")
	o.Columns.CEncode(encoder)
	encoder.WriteString("rows")
	encoder.WriteByteArray(o.Rows)
	encoder.WriteString("error")
	if o.Error == nil {
		encoder.WriteNil()
	} else {
		o.Error.CEncode(encoder)
	}

	return encoder.CheckError()
}

// CDecodeQueryResult deserializes a QueryResult using cbor
func CDecodeQueryResult(d *cbor.Decoder) (QueryResult, error) {
	var val QueryResult
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "numRows":
			val.NumRows, err = d.ReadUint64()
		case "columns":
			val.Columns, err = CDecodeColumns(d)
		case "rows":
			val.Rows, err = d.ReadByteArray()
		case "error":
			fval, err := CDecodeSqlDbError(d)
			if err != nil {
				return val, err
			}
			val.Error = &fval
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// Detailed error information from the previous operation
type SqlDbError struct {
	// Type of error.
	// The list of enum variants for this field may be expanded in the future
	// to provide finer-granularity failure information
	Code string
	// error message
	Message string
}

// MEncode serializes a SqlDbError using msgpack
func (o *SqlDbError) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("code")
	encoder.WriteString(o.Code)
	encoder.WriteString("message")
	encoder.WriteString(o.Message)

	return encoder.CheckError()
}

// MDecodeSqlDbError deserializes a SqlDbError using msgpack
func MDecodeSqlDbError(d *msgpack.Decoder) (SqlDbError, error) {
	var val SqlDbError
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "code":
			val.Code, err = d.ReadString()
		case "message":
			val.Message, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a SqlDbError using cbor
func (o *SqlDbError) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(2)
	encoder.WriteString("code")
	encoder.WriteString(o.Code)
	encoder.WriteString("message")
	encoder.WriteString(o.Message)

	return encoder.CheckError()
}

// CDecodeSqlDbError deserializes a SqlDbError using cbor
func CDecodeSqlDbError(d *cbor.Decoder) (SqlDbError, error) {
	var val SqlDbError
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "code":
			val.Code, err = d.ReadString()
		case "message":
			val.Message, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

type Statement struct {
	// Optional database in which the statement must be executed.
	// The value in this field is case-sensitive.
	Database   string
	Parameters *Parameters
	// A sql query or statement that is a non-empty string containing
	// in the syntax of the back-end database.
	Sql string
}

// MEncode serializes a Statement using msgpack
func (o *Statement) MEncode(encoder msgpack.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("database")
	encoder.WriteString(o.Database)
	encoder.WriteString("parameters")
	if o.Parameters == nil {
		encoder.WriteNil()
	} else {
		o.Parameters.MEncode(encoder)
	}
	encoder.WriteString("sql")
	encoder.WriteString(o.Sql)

	return encoder.CheckError()
}

// MDecodeStatement deserializes a Statement using msgpack
func MDecodeStatement(d *msgpack.Decoder) (Statement, error) {
	var val Statement
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, err := d.ReadMapSize()
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "database":
			val.Database, err = d.ReadString()
		case "parameters":
			fval, err := MDecodeParameters(d)
			if err != nil {
				return val, err
			}
			val.Parameters = &fval
		case "sql":
			val.Sql, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// CEncode serializes a Statement using cbor
func (o *Statement) CEncode(encoder cbor.Writer) error {
	encoder.WriteMapSize(3)
	encoder.WriteString("database")
	encoder.WriteString(o.Database)
	encoder.WriteString("parameters")
	if o.Parameters == nil {
		encoder.WriteNil()
	} else {
		o.Parameters.CEncode(encoder)
	}
	encoder.WriteString("sql")
	encoder.WriteString(o.Sql)

	return encoder.CheckError()
}

// CDecodeStatement deserializes a Statement using cbor
func CDecodeStatement(d *cbor.Decoder) (Statement, error) {
	var val Statement
	isNil, err := d.IsNextNil()
	if err != nil || isNil {
		return val, err
	}
	size, indef, err := d.ReadMapSize()
	if err != nil && indef {
		err = cbor.NewReadError("indefinite maps not supported")
	}
	if err != nil {
		return val, err
	}
	for i := uint32(0); i < size; i++ {
		field, err := d.ReadString()
		if err != nil {
			return val, err
		}
		switch field {
		case "database":
			val.Database, err = d.ReadString()
		case "parameters":
			fval, err := CDecodeParameters(d)
			if err != nil {
				return val, err
			}
			val.Parameters = &fval
		case "sql":
			val.Sql, err = d.ReadString()
		default:
			err = d.Skip()
		}
		if err != nil {
			return val, err
		}
	}
	return val, nil
}

// SqlDb - SQL Database connections
// To use this capability, the actor must be linked
// with the capability contract "wasmcloud:sqldb"
type SqlDb interface {
	// Execute an sql statement
	Execute(ctx *actor.Context, arg Statement) (*ExecuteResult, error)
	// Perform select query on database, returning all result rows
	Query(ctx *actor.Context, arg Statement) (*QueryResult, error)
}

// SqlDbHandler is called by an actor during `main` to generate a dispatch handler
// The output of this call should be passed into `actor.RegisterHandlers`
func SqlDbHandler(actor_ SqlDb) actor.Handler {
	return actor.NewHandler("SqlDb", &SqlDbReceiver{}, actor_)
}

// SqlDbContractId returns the capability contract id for this interface
func SqlDbContractId() string { return "wasmcloud:sqldb" }

// SqlDbReceiver receives messages defined in the SqlDb service interface
// SqlDb - SQL Database connections
// To use this capability, the actor must be linked
// with the capability contract "wasmcloud:sqldb"
type SqlDbReceiver struct{}

func (r *SqlDbReceiver) Dispatch(ctx *actor.Context, svc interface{}, message *actor.Message) (*actor.Message, error) {
	svc_, _ := svc.(SqlDb)
	switch message.Method {

	case "Execute":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeStatement(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.Execute(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			resp.MEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			resp.MEncode(enc)
			return &actor.Message{Method: "SqlDb.Execute", Arg: buf}, nil
		}
	case "Query":
		{

			d := msgpack.NewDecoder(message.Arg)
			value, err_ := MDecodeStatement(&d)
			if err_ != nil {
				return nil, err_
			}

			resp, err := svc_.Query(ctx, value)
			if err != nil {
				return nil, err
			}

			var sizer msgpack.Sizer
			size_enc := &sizer
			resp.MEncode(size_enc)
			buf := make([]byte, sizer.Len())
			encoder := msgpack.NewEncoder(buf)
			enc := &encoder
			resp.MEncode(enc)
			return &actor.Message{Method: "SqlDb.Query", Arg: buf}, nil
		}
	default:
		return nil, actor.NewRpcError("MethodNotHandled", "SqlDb."+message.Method)
	}
}

// SqlDbSender sends messages to a SqlDb service
// SqlDb - SQL Database connections
// To use this capability, the actor must be linked
// with the capability contract "wasmcloud:sqldb"
type SqlDbSender struct{ transport actor.Transport }

// NewProvider constructs a client for sending to a SqlDb provider
// implementing the 'wasmcloud:sqldb' capability contract, with the "default" link
func NewProviderSqlDb() *SqlDbSender {
	transport := actor.ToProvider("wasmcloud:sqldb", "default")
	return &SqlDbSender{transport: transport}
}

// NewProviderSqlDbLink constructs a client for sending to a SqlDb provider
// implementing the 'wasmcloud:sqldb' capability contract, with the specified link name
func NewProviderSqlDbLink(linkName string) *SqlDbSender {
	transport := actor.ToProvider("wasmcloud:sqldb", linkName)
	return &SqlDbSender{transport: transport}
}

// Execute an sql statement
func (s *SqlDbSender) Execute(ctx *actor.Context, arg Statement) (*ExecuteResult, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "SqlDb.Execute", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := MDecodeExecuteResult(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// Perform select query on database, returning all result rows
func (s *SqlDbSender) Query(ctx *actor.Context, arg Statement) (*QueryResult, error) {

	var sizer msgpack.Sizer
	size_enc := &sizer
	arg.MEncode(size_enc)
	buf := make([]byte, sizer.Len())

	var encoder = msgpack.NewEncoder(buf)
	enc := &encoder
	arg.MEncode(enc)

	out_buf, _ := s.transport.Send(ctx, actor.Message{Method: "SqlDb.Query", Arg: buf})
	d := msgpack.NewDecoder(out_buf)
	resp, err_ := MDecodeQueryResult(&d)
	if err_ != nil {
		return nil, err_
	}
	return &resp, nil
}

// This file is generated automatically using wasmcloud/weld-codegen 0.4.5
