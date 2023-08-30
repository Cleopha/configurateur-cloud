// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BlockStoragesColumns holds the columns for the "block_storages" table.
	BlockStoragesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "iops", Type: field.TypeInt},
		{Name: "bandwidth", Type: field.TypeFloat64},
		{Name: "price", Type: field.TypeFloat64},
	}
	// BlockStoragesTable holds the schema information for the "block_storages" table.
	BlockStoragesTable = &schema.Table{
		Name:       "block_storages",
		Columns:    BlockStoragesColumns,
		PrimaryKey: []*schema.Column{BlockStoragesColumns[0]},
	}
	// IPAddressesColumns holds the columns for the "ip_addresses" table.
	IPAddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "price", Type: field.TypeFloat64},
	}
	// IPAddressesTable holds the schema information for the "ip_addresses" table.
	IPAddressesTable = &schema.Table{
		Name:       "ip_addresses",
		Columns:    IPAddressesColumns,
		PrimaryKey: []*schema.Column{IPAddressesColumns[0]},
	}
	// InstancesColumns holds the columns for the "instances" table.
	InstancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "cpu", Type: field.TypeInt},
		{Name: "ram_go", Type: field.TypeInt},
		{Name: "stockage_go", Type: field.TypeInt},
		{Name: "gpu", Type: field.TypeString},
		{Name: "type", Type: field.TypeString},
		{Name: "price", Type: field.TypeFloat64},
	}
	// InstancesTable holds the schema information for the "instances" table.
	InstancesTable = &schema.Table{
		Name:       "instances",
		Columns:    InstancesColumns,
		PrimaryKey: []*schema.Column{InstancesColumns[0]},
	}
	// LoadBalancersColumns holds the columns for the "load_balancers" table.
	LoadBalancersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "price", Type: field.TypeFloat64},
	}
	// LoadBalancersTable holds the schema information for the "load_balancers" table.
	LoadBalancersTable = &schema.Table{
		Name:       "load_balancers",
		Columns:    LoadBalancersColumns,
		PrimaryKey: []*schema.Column{LoadBalancersColumns[0]},
	}
	// ObjectStoragesColumns holds the columns for the "object_storages" table.
	ObjectStoragesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "price", Type: field.TypeFloat64},
	}
	// ObjectStoragesTable holds the schema information for the "object_storages" table.
	ObjectStoragesTable = &schema.Table{
		Name:       "object_storages",
		Columns:    ObjectStoragesColumns,
		PrimaryKey: []*schema.Column{ObjectStoragesColumns[0]},
	}
	// PaniersColumns holds the columns for the "paniers" table.
	PaniersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "uid", Type: field.TypeUUID, Unique: true},
	}
	// PaniersTable holds the schema information for the "paniers" table.
	PaniersTable = &schema.Table{
		Name:       "paniers",
		Columns:    PaniersColumns,
		PrimaryKey: []*schema.Column{PaniersColumns[0]},
	}
	// PanierBlockStoragesColumns holds the columns for the "panier_block_storages" table.
	PanierBlockStoragesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "block_storage_panier_block_storage", Type: field.TypeInt},
		{Name: "panier_panier_block_storage", Type: field.TypeInt},
	}
	// PanierBlockStoragesTable holds the schema information for the "panier_block_storages" table.
	PanierBlockStoragesTable = &schema.Table{
		Name:       "panier_block_storages",
		Columns:    PanierBlockStoragesColumns,
		PrimaryKey: []*schema.Column{PanierBlockStoragesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "panier_block_storages_block_storages_panierBlockStorage",
				Columns:    []*schema.Column{PanierBlockStoragesColumns[2]},
				RefColumns: []*schema.Column{BlockStoragesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "panier_block_storages_paniers_panierBlockStorage",
				Columns:    []*schema.Column{PanierBlockStoragesColumns[3]},
				RefColumns: []*schema.Column{PaniersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PanierIPAddressesColumns holds the columns for the "panier_ip_addresses" table.
	PanierIPAddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "IPAddress_id", Type: field.TypeInt},
		{Name: "panier_panier_ip_address", Type: field.TypeInt},
	}
	// PanierIPAddressesTable holds the schema information for the "panier_ip_addresses" table.
	PanierIPAddressesTable = &schema.Table{
		Name:       "panier_ip_addresses",
		Columns:    PanierIPAddressesColumns,
		PrimaryKey: []*schema.Column{PanierIPAddressesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "panier_ip_addresses_ip_addresses_panierIPAddress",
				Columns:    []*schema.Column{PanierIPAddressesColumns[2]},
				RefColumns: []*schema.Column{IPAddressesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "panier_ip_addresses_paniers_panierIPAddress",
				Columns:    []*schema.Column{PanierIPAddressesColumns[3]},
				RefColumns: []*schema.Column{PaniersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PanierInstancesColumns holds the columns for the "panier_instances" table.
	PanierInstancesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "instance_id", Type: field.TypeInt},
		{Name: "panier_panier_instances", Type: field.TypeInt},
	}
	// PanierInstancesTable holds the schema information for the "panier_instances" table.
	PanierInstancesTable = &schema.Table{
		Name:       "panier_instances",
		Columns:    PanierInstancesColumns,
		PrimaryKey: []*schema.Column{PanierInstancesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "panier_instances_instances_panierInstances",
				Columns:    []*schema.Column{PanierInstancesColumns[2]},
				RefColumns: []*schema.Column{InstancesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "panier_instances_paniers_panierInstances",
				Columns:    []*schema.Column{PanierInstancesColumns[3]},
				RefColumns: []*schema.Column{PaniersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PanierLoadBalancersColumns holds the columns for the "panier_load_balancers" table.
	PanierLoadBalancersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "loadBalancer_id", Type: field.TypeInt},
		{Name: "panier_panier_load_balancer", Type: field.TypeInt},
	}
	// PanierLoadBalancersTable holds the schema information for the "panier_load_balancers" table.
	PanierLoadBalancersTable = &schema.Table{
		Name:       "panier_load_balancers",
		Columns:    PanierLoadBalancersColumns,
		PrimaryKey: []*schema.Column{PanierLoadBalancersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "panier_load_balancers_load_balancers_panierLoadBalancer",
				Columns:    []*schema.Column{PanierLoadBalancersColumns[2]},
				RefColumns: []*schema.Column{LoadBalancersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "panier_load_balancers_paniers_panierLoadBalancer",
				Columns:    []*schema.Column{PanierLoadBalancersColumns[3]},
				RefColumns: []*schema.Column{PaniersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// PanierObjectStoragesColumns holds the columns for the "panier_object_storages" table.
	PanierObjectStoragesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "objectStorage_id", Type: field.TypeInt},
		{Name: "panier_panier_object_storage", Type: field.TypeInt},
	}
	// PanierObjectStoragesTable holds the schema information for the "panier_object_storages" table.
	PanierObjectStoragesTable = &schema.Table{
		Name:       "panier_object_storages",
		Columns:    PanierObjectStoragesColumns,
		PrimaryKey: []*schema.Column{PanierObjectStoragesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "panier_object_storages_object_storages_panierObjectStorage",
				Columns:    []*schema.Column{PanierObjectStoragesColumns[2]},
				RefColumns: []*schema.Column{ObjectStoragesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "panier_object_storages_paniers_panierObjectStorage",
				Columns:    []*schema.Column{PanierObjectStoragesColumns[3]},
				RefColumns: []*schema.Column{PaniersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BlockStoragesTable,
		IPAddressesTable,
		InstancesTable,
		LoadBalancersTable,
		ObjectStoragesTable,
		PaniersTable,
		PanierBlockStoragesTable,
		PanierIPAddressesTable,
		PanierInstancesTable,
		PanierLoadBalancersTable,
		PanierObjectStoragesTable,
	}
)

func init() {
	PanierBlockStoragesTable.ForeignKeys[0].RefTable = BlockStoragesTable
	PanierBlockStoragesTable.ForeignKeys[1].RefTable = PaniersTable
	PanierIPAddressesTable.ForeignKeys[0].RefTable = IPAddressesTable
	PanierIPAddressesTable.ForeignKeys[1].RefTable = PaniersTable
	PanierInstancesTable.ForeignKeys[0].RefTable = InstancesTable
	PanierInstancesTable.ForeignKeys[1].RefTable = PaniersTable
	PanierLoadBalancersTable.ForeignKeys[0].RefTable = LoadBalancersTable
	PanierLoadBalancersTable.ForeignKeys[1].RefTable = PaniersTable
	PanierObjectStoragesTable.ForeignKeys[0].RefTable = ObjectStoragesTable
	PanierObjectStoragesTable.ForeignKeys[1].RefTable = PaniersTable
}