---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_db_system_connectors"
sidebar_current: "docs-oci-datasource-database_management-external_db_system_connectors"
description: |-
  Provides the list of External Db System Connectors in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_db_system_connectors
This data source provides the list of External Db System Connectors in Oracle Cloud Infrastructure Database Management service.

Lists the external connectors in the specified external DB system.

## Example Usage

```hcl
data "oci_database_management_external_db_system_connectors" "test_external_db_system_connectors" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.external_db_system_connector_display_name
	external_db_system_id = oci_database_management_external_db_system.test_external_db_system.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to only return the resources that match the entire display name.
* `external_db_system_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system.


## Attributes Reference

The following attributes are exported:

* `external_db_system_connector_collection` - The list of external_db_system_connector_collection.

### ExternalDbSystemConnector Reference

The following attributes are exported:

* `agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent used for the external DB system connector. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_failure_message` - The error message indicating the reason for connection failure or `null` if the connection was successful. 
* `connection_info` - The connection details required to connect to an external DB system component.
	* `component_type` - The component type.
	* `connection_credentials` - The credentials used to connect to the ASM instance. Currently only the `DETAILS` type is supported for creating MACS connector credentials. 
		* `credential_name` - The name of the credential information that used to connect to the DB system resource. The name should be in "x.y" format, where the length of "x" has a maximum of 64 characters, and length of "y" has a maximum of 199 characters. The name strings can contain letters, numbers and the underscore character only. Other characters are not valid, except for the "." character that separates the "x" and "y" portions of the name. *IMPORTANT* - The name must be unique within the Oracle Cloud Infrastructure region the credential is being created in. If you specify a name that duplicates the name of another credential within the same Oracle Cloud Infrastructure region, you may overwrite or corrupt the credential that is already using the name.

			For example: inventorydb.abc112233445566778899 
		* `credential_type` - The type of credential used to connect to the ASM instance.
		* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
		* `role` - The role of the user connecting to the ASM instance.
		* `ssl_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the SSL keystore and truststore details.
		* `user_name` - The user name used to connect to the ASM instance.
	* `connection_string` - The Oracle Database connection string. 
		* `host_name` - The host name of the database or the SCAN name in case of a RAC database.
		* `hosts` - The list of host names of the ASM instances.
		* `port` - The port used to connect to the ASM instance.
		* `protocol` - The protocol used to connect to the ASM instance.
		* `service` - The service name of the ASM instance.
	* `database_credential` - The credential to connect to the database to perform tablespace administration tasks.
		* `credential_type` - The type of the credential for tablespace administration tasks.
		* `named_credential_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the named credential where the database password metadata is stored. 
		* `password` - The database user's password encoded using BASE64 scheme.
		* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the database password is stored. 
		* `role` - The role of the database user.
		* `username` - The user to connect to the database.
* `connection_status` - The status of connectivity to the external DB system component.
* `connector_type` - The type of connector.
* `display_name` - The user-friendly name for the external connector. The name does not have to be unique.
* `external_db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system that the connector is a part of.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system connector.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current lifecycle state of the external DB system connector.
* `time_connection_status_last_updated` - The date and time the connectionStatus of the external DB system connector was last updated.
* `time_created` - The date and time the external DB system connector was created.
* `time_updated` - The date and time the external DB system connector was last updated.

