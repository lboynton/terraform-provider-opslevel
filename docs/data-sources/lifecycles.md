---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opslevel_lifecycles Data Source - terraform-provider-opslevel"
subcategory: ""
description: |-
  
---

# opslevel_lifecycles (Data Source)



## Example Usage

```terraform
data "opslevel_lifecycles" "all" {
}

output "found" {
  value = data.opslevel_lifecycles.all.aliases[0]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **aliases** (List of String)
- **ids** (List of String)
- **indexes** (List of Number)
- **names** (List of String)


