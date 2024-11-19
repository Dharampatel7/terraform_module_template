package test

import (
    "testing"
    "github.com/gruntwork-io/terratest/modules/terraform"
    "github.com/stretchr/testify/assert"
)

func TestTerraformResourceGroup(t *testing.T) {
    t.Parallel()

    terraformOptions := &terraform.Options{
        TerraformDir: "../module",

        Vars: map[string]interface{}{
            "resource_group_name":     "testResourceGroup",
            "resource_group_location": "West Europe",
        },
    }

    defer terraform.Destroy(t, terraformOptions)

    terraform.InitAndApply(t, terraformOptions)

    resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")

    assert.Equal(t, "testResourceGroup", resourceGroupName)
}
