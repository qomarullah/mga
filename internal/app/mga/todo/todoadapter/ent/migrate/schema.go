// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebook/ent/dialect/sql/schema"
	"github.com/facebook/ent/schema/field"
)

var (
	// TodoItemsColumns holds the columns for the "todo_items" table.
	TodoItemsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uid", Type: field.TypeString, Unique: true, Size: 26},
		{Name: "title", Type: field.TypeString, Size: 2147483647},
		{Name: "completed", Type: field.TypeBool},
		{Name: "order", Type: field.TypeInt},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// TodoItemsTable holds the schema information for the "todo_items" table.
	TodoItemsTable = &schema.Table{
		Name:        "todo_items",
		Columns:     TodoItemsColumns,
		PrimaryKey:  []*schema.Column{TodoItemsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TodoItemsTable,
	}
)

func init() {
}
