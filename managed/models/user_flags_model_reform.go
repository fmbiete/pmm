// Code generated by gopkg.in/reform.v1. DO NOT EDIT.

package models

import (
	"fmt"
	"strings"

	"gopkg.in/reform.v1"
	"gopkg.in/reform.v1/parse"
)

type userDetailsTableType struct {
	s parse.StructInfo
	z []interface{}
}

// Schema returns a schema name in SQL database ("").
func (v *userDetailsTableType) Schema() string {
	return v.s.SQLSchema
}

// Name returns a view or table name in SQL database ("user_flags").
func (v *userDetailsTableType) Name() string {
	return v.s.SQLName
}

// Columns returns a new slice of column names for that view or table in SQL database.
func (v *userDetailsTableType) Columns() []string {
	return []string{
		"id",
		"tour_done",
		"alerting_tour_done",
		"created_at",
		"updated_at",
	}
}

// NewStruct makes a new struct for that view or table.
func (v *userDetailsTableType) NewStruct() reform.Struct {
	return new(UserDetails)
}

// NewRecord makes a new record for that table.
func (v *userDetailsTableType) NewRecord() reform.Record {
	return new(UserDetails)
}

// PKColumnIndex returns an index of primary key column for that table in SQL database.
func (v *userDetailsTableType) PKColumnIndex() uint {
	return uint(v.s.PKFieldIndex)
}

// UserDetailsTable represents user_flags view or table in SQL database.
var UserDetailsTable = &userDetailsTableType{
	s: parse.StructInfo{
		Type:    "UserDetails",
		SQLName: "user_flags",
		Fields: []parse.FieldInfo{
			{Name: "ID", Type: "int", Column: "id"},
			{Name: "Tour", Type: "bool", Column: "tour_done"},
			{Name: "AlertingTour", Type: "bool", Column: "alerting_tour_done"},
			{Name: "CreatedAt", Type: "time.Time", Column: "created_at"},
			{Name: "UpdatedAt", Type: "time.Time", Column: "updated_at"},
		},
		PKFieldIndex: 0,
	},
	z: new(UserDetails).Values(),
}

// String returns a string representation of this struct or record.
func (s UserDetails) String() string {
	res := make([]string, 5)
	res[0] = "ID: " + reform.Inspect(s.ID, true)
	res[1] = "Tour: " + reform.Inspect(s.Tour, true)
	res[2] = "AlertingTour: " + reform.Inspect(s.AlertingTour, true)
	res[3] = "CreatedAt: " + reform.Inspect(s.CreatedAt, true)
	res[4] = "UpdatedAt: " + reform.Inspect(s.UpdatedAt, true)
	return strings.Join(res, ", ")
}

// Values returns a slice of struct or record field values.
// Returned interface{} values are never untyped nils.
func (s *UserDetails) Values() []interface{} {
	return []interface{}{
		s.ID,
		s.Tour,
		s.AlertingTour,
		s.CreatedAt,
		s.UpdatedAt,
	}
}

// Pointers returns a slice of pointers to struct or record fields.
// Returned interface{} values are never untyped nils.
func (s *UserDetails) Pointers() []interface{} {
	return []interface{}{
		&s.ID,
		&s.Tour,
		&s.AlertingTour,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}

// View returns View object for that struct.
func (s *UserDetails) View() reform.View {
	return UserDetailsTable
}

// Table returns Table object for that record.
func (s *UserDetails) Table() reform.Table {
	return UserDetailsTable
}

// PKValue returns a value of primary key for that record.
// Returned interface{} value is never untyped nil.
func (s *UserDetails) PKValue() interface{} {
	return s.ID
}

// PKPointer returns a pointer to primary key field for that record.
// Returned interface{} value is never untyped nil.
func (s *UserDetails) PKPointer() interface{} {
	return &s.ID
}

// HasPK returns true if record has non-zero primary key set, false otherwise.
func (s *UserDetails) HasPK() bool {
	return s.ID != UserDetailsTable.z[UserDetailsTable.s.PKFieldIndex]
}

// SetPK sets record primary key, if possible.
//
// Deprecated: prefer direct field assignment where possible: s.ID = pk.
func (s *UserDetails) SetPK(pk interface{}) {
	reform.SetPK(s, pk)
}

// check interfaces
var (
	_ reform.View   = UserDetailsTable
	_ reform.Struct = (*UserDetails)(nil)
	_ reform.Table  = UserDetailsTable
	_ reform.Record = (*UserDetails)(nil)
	_ fmt.Stringer  = (*UserDetails)(nil)
)

func init() {
	parse.AssertUpToDate(&UserDetailsTable.s, new(UserDetails))
}
