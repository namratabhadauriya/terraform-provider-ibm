// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIVolumeGroupStorageDetailsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIVolumeGroupStorageDetailsDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pi_volume_group_storage_details.testacc_volume_group_storage_details", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIVolumeGroupStorageDetailsDataSourceConfig() string {
	return fmt.Sprintf(`
data "ibm_pi_volume_group_storage_details" "testacc_volume_group_storage_details" {
    pi_volume_group_id   = "%[1]s"
    pi_cloud_instance_id = "%[2]s"
}`, acc.Pi_volume_group_id, acc.Pi_cloud_instance_id)

}