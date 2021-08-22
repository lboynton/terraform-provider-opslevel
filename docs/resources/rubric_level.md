---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opslevel_rubric_level Resource - terraform-provider-opslevel"
subcategory: ""
description: |-
  Manages a rubric level
---

# opslevel_rubric_level (Resource)

Manages a rubric level

## Example Usage

```terraform
resource "opslevel_rubric_level" "example" {
  name = "foo"
  description = "foo level"
}

output "level" {
  value = opslevel_rubric_level.example.id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) The display name of the category.

### Optional

- **description** (String) The description of the category.
- **id** (String) The ID of this resource.
- **last_updated** (String)

## Import

Import is supported using the following syntax:

```shell
terraform import opslevel_rubric_level.example Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS82MDI0
```