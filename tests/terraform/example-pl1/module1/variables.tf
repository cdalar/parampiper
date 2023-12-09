
variable "resource_group_name" {
  type        = string
  description = "Resource group name"
}

variable "address_space" {
  type        = list(string)
  description = "Resource group name"
}

variable "name" {
  type        = string
  description = "Resource group name"
}
variable "location" {
  type        = string
  description = "location"
  default     = "westeurope"
}
