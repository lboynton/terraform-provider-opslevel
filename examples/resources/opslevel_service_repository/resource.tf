data "opslevel_service" "foo" {
  alias = "foo"
}

data "opslevel_repository" "foo" {
  alias = "github.com:example/foo"
}

resource "opslevel_service_repository" "foo" {
  service = data.opslevel_service.foo.id
  repository = data.opslevel_repository.foo.id

  name = "Foo"
  base_directory = "/"
}

resource "opslevel_service_tag" "bar" {
  service_alias = "bar"
  repository_alias = "github.com:example/bar"

  name = "Bar"
  base_directory = "/"
}