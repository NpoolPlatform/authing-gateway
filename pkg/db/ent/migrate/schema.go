// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AppAuthsColumns holds the columns for the "app_auths" table.
	AppAuthsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "resource", Type: field.TypeString},
		{Name: "method", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AppAuthsTable holds the schema information for the "app_auths" table.
	AppAuthsTable = &schema.Table{
		Name:       "app_auths",
		Columns:    AppAuthsColumns,
		PrimaryKey: []*schema.Column{AppAuthsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appauth_app_id_resource_method",
				Unique:  true,
				Columns: []*schema.Column{AppAuthsColumns[1], AppAuthsColumns[2], AppAuthsColumns[3]},
			},
		},
	}
	// AppRoleAuthsColumns holds the columns for the "app_role_auths" table.
	AppRoleAuthsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "role_id", Type: field.TypeUUID},
		{Name: "resource", Type: field.TypeString},
		{Name: "method", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AppRoleAuthsTable holds the schema information for the "app_role_auths" table.
	AppRoleAuthsTable = &schema.Table{
		Name:       "app_role_auths",
		Columns:    AppRoleAuthsColumns,
		PrimaryKey: []*schema.Column{AppRoleAuthsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "approleauth_app_id_role_id_resource_method",
				Unique:  true,
				Columns: []*schema.Column{AppRoleAuthsColumns[1], AppRoleAuthsColumns[2], AppRoleAuthsColumns[3], AppRoleAuthsColumns[4]},
			},
		},
	}
	// AppUserAuthsColumns holds the columns for the "app_user_auths" table.
	AppUserAuthsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID},
		{Name: "resource", Type: field.TypeString},
		{Name: "method", Type: field.TypeString},
		{Name: "create_at", Type: field.TypeUint32},
		{Name: "update_at", Type: field.TypeUint32},
		{Name: "delete_at", Type: field.TypeUint32},
	}
	// AppUserAuthsTable holds the schema information for the "app_user_auths" table.
	AppUserAuthsTable = &schema.Table{
		Name:       "app_user_auths",
		Columns:    AppUserAuthsColumns,
		PrimaryKey: []*schema.Column{AppUserAuthsColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "appuserauth_app_id_user_id_resource_method",
				Unique:  true,
				Columns: []*schema.Column{AppUserAuthsColumns[1], AppUserAuthsColumns[2], AppUserAuthsColumns[3], AppUserAuthsColumns[4]},
			},
		},
	}
	// AuthHistoriesColumns holds the columns for the "auth_histories" table.
	AuthHistoriesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "app_id", Type: field.TypeUUID},
		{Name: "user_id", Type: field.TypeUUID, Nullable: true},
		{Name: "resource", Type: field.TypeString},
		{Name: "method", Type: field.TypeString},
		{Name: "allowed", Type: field.TypeBool},
		{Name: "create_at", Type: field.TypeUint32},
	}
	// AuthHistoriesTable holds the schema information for the "auth_histories" table.
	AuthHistoriesTable = &schema.Table{
		Name:       "auth_histories",
		Columns:    AuthHistoriesColumns,
		PrimaryKey: []*schema.Column{AuthHistoriesColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AppAuthsTable,
		AppRoleAuthsTable,
		AppUserAuthsTable,
		AuthHistoriesTable,
	}
)

func init() {
}
