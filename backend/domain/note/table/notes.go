package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var Notes = newNotesTable("public", "notes", "")

type notesTable struct {
	postgres.Table

	// Columns
	ID         postgres.ColumnString
	UserID     postgres.ColumnString
	Title      postgres.ColumnString
	Content    postgres.ColumnString
	Icon       postgres.ColumnString
	IsArchived postgres.ColumnBool
	ParentID   postgres.ColumnString
	CoverImage postgres.ColumnString
	CreatedAt  postgres.ColumnTimestamp
	UpdatedAt  postgres.ColumnTimestamp

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type NotesTable struct {
	notesTable

	EXCLUDED notesTable
}

// AS creates new NotesTable with assigned alias
func (a NotesTable) AS(alias string) *NotesTable {
	return newNotesTable(a.SchemaName(), a.TableName(), alias)
}

// FromSchema Schema creates new NotesTable with assigned schema name
func (a NotesTable) FromSchema(schemaName string) *NotesTable {
	return newNotesTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new NotesTable with assigned table prefix
func (a NotesTable) WithPrefix(prefix string) *NotesTable {
	return newNotesTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new NotesTable with assigned table suffix
func (a NotesTable) WithSuffix(suffix string) *NotesTable {
	return newNotesTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newNotesTable(schemaName, tableName, alias string) *NotesTable {
	return &NotesTable{
		notesTable: newNotesTableImpl(schemaName, tableName, alias),
		EXCLUDED:   newNotesTableImpl("", "excluded", ""),
	}
}

func newNotesTableImpl(schemaName, tableName, alias string) notesTable {
	var (
		IDColumn         = postgres.StringColumn("id")
		UserIDColumn     = postgres.StringColumn("user_id")
		TitleColumn      = postgres.StringColumn("title")
		ContentColumn    = postgres.StringColumn("content")
		IconColumn       = postgres.StringColumn("icon")
		IsArchivedColumn = postgres.BoolColumn("is_archived")
		ParentIDColumn   = postgres.StringColumn("parent_id")
		CoverImageColumn = postgres.StringColumn("cover_image")
		CreatedAtColumn  = postgres.TimestampColumn("created_at")
		UpdatedAtColumn  = postgres.TimestampColumn("updated_at")
		allColumns       = postgres.ColumnList{IDColumn, UserIDColumn, TitleColumn, ContentColumn, IconColumn, IsArchivedColumn, ParentIDColumn, CoverImageColumn, CreatedAtColumn, UpdatedAtColumn}
		mutableColumns   = postgres.ColumnList{UserIDColumn, TitleColumn, ContentColumn, IconColumn, IsArchivedColumn, ParentIDColumn, CoverImageColumn, CreatedAtColumn, UpdatedAtColumn}
	)

	return notesTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		ID:         IDColumn,
		UserID:     UserIDColumn,
		Title:      TitleColumn,
		Content:    ContentColumn,
		Icon:       IconColumn,
		IsArchived: IsArchivedColumn,
		ParentID:   ParentIDColumn,
		CoverImage: CoverImageColumn,
		CreatedAt:  CreatedAtColumn,
		UpdatedAt:  UpdatedAtColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
