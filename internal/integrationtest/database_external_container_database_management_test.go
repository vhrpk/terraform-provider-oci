// Copyright (c) 2017, 2020, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DatabaseExternalContainerDatabaseManagementRepresentation = map[string]interface{}{
		"external_container_database_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_container_database.test_external_container_database.id}`},
		"external_database_connector_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_database_connector.test_external_database_connector.id}`},
		"license_model":                  acctest.Representation{RepType: acctest.Required, Create: `BRING_YOUR_OWN_LICENSE`},
		"enable_management":              acctest.Representation{RepType: acctest.Required, Create: `true`, Update: `false`},
	}
	DatabaseExternalContainerDatabaseConnectorRepresentation = map[string]interface{}{
		"connection_credentials": acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseExternalDatabaseConnectorConnectionCredentialsRepresentation},
		"connection_string":      acctest.RepresentationGroup{RepType: acctest.Required, Group: DatabaseExternalDatabaseConnectorConnectionStringRepresentation},
		"connector_agent_id":     acctest.Representation{RepType: acctest.Required, Create: `${var.agent_id}`},
		"display_name":           acctest.Representation{RepType: acctest.Required, Create: `myTestConn`},
		"external_database_id":   acctest.Representation{RepType: acctest.Required, Create: `${oci_database_external_container_database.test_external_container_database.id}`},
		"connector_type":         acctest.Representation{RepType: acctest.Optional, Create: `MACS`},
	}

<<<<<<< ours
	DatabaseExternalContainerDatabaseManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_agent", "test_agent", acctest.Required, acctest.Create, agentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_cloud_bridge_environment", "test_environment", acctest.Required, acctest.Create, environmentRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_backup", "test_backup", acctest.Required, acctest.Create, backupRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_database", "test_database", acctest.Required, acctest.Create, databaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_db_home", "test_db_home", acctest.Required, acctest.Create, dbHomeRepresentation) +
		BackupResourceDependencies +
		GenerateResourceFromRepresentationMap("oci_database_backup_destination", "test_backup_destination", Optional, Create, backupDestinationNFSRepresentation) +
		GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", Optional, Update, RepresentationCopyWithNewProperties(exadataInfrastructureActivateRepresentation, map[string]interface{}{"activation_file": Representation{RepType: Optional, Update: activationFilePath}})) +
		GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", Optional, Update, vmClusterNetworkValidateRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_exadata_infrastructure", "test_exadata_infrastructure", acctest.Required, acctest.Create, exadataInfrastructureRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Required, acctest.Create, externalContainerDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, externalDatabaseConnectorRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster_network", "test_vm_cluster_network", acctest.Required, acctest.Create, vmClusterNetworkRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_vm_cluster", "test_vm_cluster", acctest.Required, acctest.Create, vmClusterRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_kms_vault", "test_vault", acctest.Required, acctest.Create, vaultRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_vault_secret", "test_secret", acctest.Required, acctest.Create, secretRepresentation)
=======
	DatabaseExternalContainerDatabaseManagementResourceDependencies = acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database", "test_external_container_database", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_database_connector", "test_external_database_connector", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseConnectorRepresentation)
>>>>>>> theirs
)

// issue-routing-tag: database/default
func TestDatabaseExternalContainerDatabaseManagementResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDatabaseExternalContainerDatabaseManagementResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	agentId := utils.GetEnvSettingWithBlankDefault("connector_agent_id")
	agentIdVariableStr := fmt.Sprintf("variable \"agent_id\" { default = \"%s\" }\n", agentId)

	resourceName := "oci_database_external_container_database_management.test_external_container_database_management"

	resourceCdb := "oci_database_external_container_database.test_external_container_database"

	// Save TF content to Create resource with only required properties. This has to be exactly the same as the config part in the Create step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+agentIdVariableStr+DatabaseExternalContainerDatabaseManagementResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseManagementRepresentation), "database", "externalContainerDatabaseManagement", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify Create (Enable Database Management)
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalContainerDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Enablement
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalContainerDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Required, acctest.Create, DatabaseExternalContainerDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceCdb, "database_management_config.0.database_management_status", "ENABLED"),
			),
		},
		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalContainerDatabaseManagementResourceDependencies,
		},
		// verify Update (Enable Database Management)
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalContainerDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Optional, acctest.Create, DatabaseExternalContainerDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// verify Update (Disable Database Management)
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalContainerDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Optional, acctest.Update, DatabaseExternalContainerDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "external_container_database_id"),
				resource.TestCheckResourceAttrSet(resourceName, "external_database_connector_id"),
			),
		},
		// Verify Disablement
		{
			Config: config + compartmentIdVariableStr + agentIdVariableStr + DatabaseExternalContainerDatabaseManagementResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_database_external_container_database_management", "test_external_container_database_management", acctest.Optional, acctest.Update, DatabaseExternalContainerDatabaseManagementRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceCdb, "database_management_config.0.database_management_status", "NOT_ENABLED"),
			),
		},
	})
}
