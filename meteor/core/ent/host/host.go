// Code generated by entc, DO NOT EDIT.

package host

const (
	// Label holds the string label denoting the host type in the database.
	Label = "host"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHostname holds the string denoting the hostname field in the database.
	FieldHostname = "hostname"
	// FieldInterface holds the string denoting the interface field in the database.
	FieldInterface = "interface"
	// FieldLastSeen holds the string denoting the lastseen field in the database.
	FieldLastSeen = "last_seen"

	// EdgeBots holds the string denoting the bots edge name in mutations.
	EdgeBots = "bots"
	// EdgeActions holds the string denoting the actions edge name in mutations.
	EdgeActions = "actions"
	// EdgeMember holds the string denoting the member edge name in mutations.
	EdgeMember = "member"

	// Table holds the table name of the host in the database.
	Table = "hosts"
	// BotsTable is the table the holds the bots relation/edge.
	BotsTable = "bots"
	// BotsInverseTable is the table name for the Bot entity.
	// It exists in this package in order to avoid circular dependency with the "bot" package.
	BotsInverseTable = "bots"
	// BotsColumn is the table column denoting the bots relation/edge.
	BotsColumn = "host_bots"
	// ActionsTable is the table the holds the actions relation/edge.
	ActionsTable = "actions"
	// ActionsInverseTable is the table name for the Action entity.
	// It exists in this package in order to avoid circular dependency with the "action" package.
	ActionsInverseTable = "actions"
	// ActionsColumn is the table column denoting the actions relation/edge.
	ActionsColumn = "host_actions"
	// MemberTable is the table the holds the member relation/edge. The primary key declared below.
	MemberTable = "group_members"
	// MemberInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	MemberInverseTable = "groups"
)

// Columns holds all SQL columns for host fields.
var Columns = []string{
	FieldID,
	FieldHostname,
	FieldInterface,
	FieldLastSeen,
}

var (
	// MemberPrimaryKey and MemberColumn2 are the table columns denoting the
	// primary key for the member relation (M2M).
	MemberPrimaryKey = []string{"group_id", "host_id"}
)

var (
	// DefaultLastSeen holds the default value on creation for the lastSeen field.
	DefaultLastSeen int
	// LastSeenValidator is a validator for the "lastSeen" field. It is called by the builders before save.
	LastSeenValidator func(int) error
)
