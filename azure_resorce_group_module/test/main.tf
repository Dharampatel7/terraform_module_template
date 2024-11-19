terraform {
  required_version = ">= 1.0"
}
module "resource_group" {
  source = "../module"

  resource_group_name     = var.resource_group_name
  resource_group_location = var.resource_group_location
}
