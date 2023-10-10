// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CommentsColumns holds the columns for the "comments" table.
	CommentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "author", Type: field.TypeInt},
		{Name: "postid", Type: field.TypeInt},
		{Name: "username", Type: field.TypeString},
		{Name: "avatar", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "likes", Type: field.TypeInt},
	}
	// CommentsTable holds the schema information for the "comments" table.
	CommentsTable = &schema.Table{
		Name:       "comments",
		Columns:    CommentsColumns,
		PrimaryKey: []*schema.Column{CommentsColumns[0]},
	}
	// PostsColumns holds the columns for the "posts" table.
	PostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "img", Type: field.TypeString},
		{Name: "content", Type: field.TypeString},
		{Name: "author", Type: field.TypeInt},
		{Name: "likes", Type: field.TypeInt, Default: 0},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// PostsTable holds the schema information for the "posts" table.
	PostsTable = &schema.Table{
		Name:       "posts",
		Columns:    PostsColumns,
		PrimaryKey: []*schema.Column{PostsColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "avatar", Type: field.TypeString},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "info", Type: field.TypeString},
		{Name: "bookmarks", Type: field.TypeString},
		{Name: "likes", Type: field.TypeString},
		{Name: "follower", Type: field.TypeString},
		{Name: "following", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeTime},
		{Name: "update_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CommentsTable,
		PostsTable,
		UsersTable,
	}
)

func init() {
}
