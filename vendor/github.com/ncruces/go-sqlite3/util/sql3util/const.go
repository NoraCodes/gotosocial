package sql3util

const (
	_NONE = iota
	_MEMORY
	_SYNTAX
	_UNSUPPORTEDSQL
)

type ConflictClause uint32

const (
	CONFLICT_NONE ConflictClause = iota
	CONFLICT_ROLLBACK
	CONFLICT_ABORT
	CONFLICT_FAIL
	CONFLICT_IGNORE
	CONFLICT_REPLACE
)

type OrderClause uint32

const (
	ORDER_NONE OrderClause = iota
	ORDER_ASC
	ORDER_DESC
)

type FKAction uint32

const (
	FKACTION_NONE FKAction = iota
	FKACTION_SETNULL
	FKACTION_SETDEFAULT
	FKACTION_CASCADE
	FKACTION_RESTRICT
	FKACTION_NOACTION
)

type FKDefType uint32

const (
	DEFTYPE_NONE FKDefType = iota
	DEFTYPE_DEFERRABLE
	DEFTYPE_DEFERRABLE_INITIALLY_DEFERRED
	DEFTYPE_DEFERRABLE_INITIALLY_IMMEDIATE
	DEFTYPE_NOTDEFERRABLE
	DEFTYPE_NOTDEFERRABLE_INITIALLY_DEFERRED
	DEFTYPE_NOTDEFERRABLE_INITIALLY_IMMEDIATE
)

type StatementType uint32

const (
	CREATE_UNKNOWN StatementType = iota
	CREATE_TABLE
	ALTER_RENAME_TABLE
	ALTER_RENAME_COLUMN
	ALTER_ADD_COLUMN
	ALTER_DROP_COLUMN
)