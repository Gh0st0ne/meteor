// Code generated by entc, DO NOT EDIT.

package action

const (
	// Label holds the string label denoting the action type in the database.
	Label = "action"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldUUID holds the string denoting the uuid field in the database.
	FieldUUID = "uuid"
	// FieldMode holds the string denoting the mode field in the database.
	FieldMode = "mode"
	// FieldArgs holds the string denoting the args field in the database.
	FieldArgs = "args"
	// FieldQueued holds the string denoting the queued field in the database.
	FieldQueued = "queued"
	// FieldResponded holds the string denoting the responded field in the database.
	FieldResponded = "responded"
	// FieldResult holds the string denoting the result field in the database.
	FieldResult = "result"

	// EdgeTargeting holds the string denoting the targeting edge name in mutations.
	EdgeTargeting = "targeting"

	// Table holds the table name of the action in the database.
	Table = "actions"
	// TargetingTable is the table the holds the targeting relation/edge.
	TargetingTable = "actions"
	// TargetingInverseTable is the table name for the Host entity.
	// It exists in this package in order to avoid circular dependency with the "host" package.
	TargetingInverseTable = "hosts"
	// TargetingColumn is the table column denoting the targeting relation/edge.
	TargetingColumn = "host_actions"
)

// Columns holds all SQL columns for action fields.
var Columns = []string{
	FieldID,
	FieldUUID,
	FieldMode,
	FieldArgs,
	FieldQueued,
	FieldResponded,
	FieldResult,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Action type.
var ForeignKeys = []string{
	"host_actions",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultQueued holds the default value on creation for the "queued" field.
	DefaultQueued bool
	// DefaultResponded holds the default value on creation for the "responded" field.
	DefaultResponded bool
	// DefaultResult holds the default value on creation for the "result" field.
	DefaultResult string
)
