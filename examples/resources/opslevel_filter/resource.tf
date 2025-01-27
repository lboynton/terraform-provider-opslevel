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
