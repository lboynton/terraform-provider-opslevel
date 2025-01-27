data "opslevel_lifecycle" "beta" {
    filter {
        field = "alias"
        value = "beta"
    }
}

data "opslevel_tier" "tier3" {
    filter {
        field = "index"
        value = "3"
    }
}

resource "opslevel_team" "foo" {
  name = "foo"
  manager_email = "john.doe@example.com"
  responsibilities = "Responsible for foo frontend and backend"
}

resource "opslevel_service" "foo" {
  name = "foo"

  description = "foo service"
  framework   = "rails"
  language    = "ruby"

  lifecycle_alias = data.opslevel_lifecycle.beta.alias
  tier_alias = data.opslevel_tier.tier3.alias
  owner_alias = opslevel_team.foo.alias
}

output "foo_aliases" {
  value = opslevel_service.example.aliases
}
