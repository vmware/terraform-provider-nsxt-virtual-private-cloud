terraform {
  required_providers {
    nsxt = {
      source  = "terraform.local/local/for-vmware-nsxt-virtual-private-cloud"
      version = "1.0.0"
    }
  }
}

provider "nsxt" {
  host = "10.191.155.100"
  username = "admin"
  password = "r6.Or#F_z_.F"
  org = "default"
  project = "Dev_project"
  vpc = "dev_vpc"
  retry_on_status_codes = [400, 409, 429, 500, 503, 504, 603]
}

# security policy
resource "nsxt_vpc_security_policy" "testSecurityPolicy" {
  nsx_id = "Mysecuritypolicy1"
  sequence_number = 1
  scope = [data.nsxt_vpc_group.test.path]
  tags {
    scope = "os"
    tag = "windows"
  }
  tags {
    scope = "os"
    tag = "linux"
  }
  tags {
    scope = "os"
    tag = "mac"
  }
}

data "nsxt_vpc_group" "test" {
  nsx_id = "test-group-abc-1"
  context_info {
    domain = "test-domain-abc-1"
  }
}

