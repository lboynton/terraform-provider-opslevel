---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opslevel_teams Data Source - terraform-provider-opslevel"
subcategory: ""
description: |-
  
---

# opslevel_teams (Data Source)



## Example Usage

```terraform
data "opslevel_teams" "all" {
}

data "opslevel_teams" "leet" {
  filter {
    field = "manager-email"
    value = "0p5l3v3l@example.com"
  }
}

output "found" {
  value = data.opslevel_teams.all.ids[3]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **filter** (Block List, Max: 1) (see [below for nested schema](#nestedblock--filter))
- **id** (String) The ID of this resource.

### Read-Only

- **aliases** (List of String)
- **ids** (List of String)
- **names** (List of String)

<a id="nestedblock--filter"></a>
### Nested Schema for `filter`

Required:

- **field** (String) The field of the target resource to filter upon.

Optional:

- **value** (String) The field value of the target resource to match.


