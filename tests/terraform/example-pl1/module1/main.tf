resource "azurerm_virtual_network" "vnet_module1" {
  name                = var.name
  address_space       = var.address_space
  location            = var.location
  resource_group_name = var.resource_group_name
}