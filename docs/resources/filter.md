---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "opslevel_filter Resource - terraform-provider-opslevel"
subcategory: ""
description: |-
  Manages a team
---

# opslevel_filter (Resource)

Manages a team

## Example Usage

```terraform
resource "opslevel_filter" "tier1" {
  name = "foo"
  predicate {
    key = "tier_index"
    type = "equals"
    value = "1"
  }
  connective = "and"
}

resource "opslevel_filter" "tier2_alpha" {
  name = "foo"
  predicate {
    key = "tier_index"
    type = "equals"
    value = "1"
  }
  predicate {
    key = "lifecycle_index"
    type = "equals"
    value = "1"
  }
  connective = "and"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **name** (String) The team's display name.

### Optional

- **connective** (String) The logical operator to be used in conjunction with predicates.
- **id** (String) The ID of this resource.
- **last_updated** (String)
- **predicate** (Block List) The list of predicates used to select which services apply to the filter. (see [below for nested schema](#nestedblock--predicate))

<a id="nestedblock--predicate"></a>
### Nested Schema for `predicate`

Required:

- **key** (String) The condition key used by the predicate. Possible valuas are 'tier_index', 'lifecycle_index', 'language', 'framework', 'product', 'owner_id', 'name' and 'tags'.
- **type** (String) The condition type used by the predicate.

Optional:

- **key_data** (String) Additional data used by the predicate. This field is used by predicates with key = 'tags' to specify the tag key. For example, to create a predicate for services containing the tag 'db:mysql', set keyData = 'db' and value = 'mysql'.
- **value** (String) The condition value used by the predicate.

## Import

Import is supported using the following syntax:

```shell
terraform import opslevel_filter.example Z2lkOi8vb3BzbGV2ZWwvU2VydmljZS82MDI0
```