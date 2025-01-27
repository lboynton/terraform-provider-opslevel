---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opslevel_repositories Data Source - terraform-provider-opslevel"
subcategory: ""
description: |-
  
---

# opslevel_repositories (Data Source)



## Example Usage

```terraform
data "opslevel_tier" "tier2" {
    filter {
        field = "alias"
        value = "tier_2"
    }
}

data "opslevel_repositories" "all" {
}

data "opslevel_repositories" "tier2" {
  filter {
    field = "tier"
    value = data.opslevel_tier.tier2.alias
  }
}

output "all" {
  value = data.opslevel_repositories.all.names
}

output "tier2" {
  value = data.opslevel_repositories.tier2.names
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **filter** (Block List, Max: 1) (see [below for nested schema](#nestedblock--filter))
- **id** (String) The ID of this resource.

### Read-Only

- **ids** (List of String)
- **names** (List of String)
- **urls** (List of String)

<a id="nestedblock--filter"></a>
### Nested Schema for `filter`

Required:

- **field** (String) The field of the target resource to filter upon.

Optional:

- **value** (String) The field value of the target resource to match.


