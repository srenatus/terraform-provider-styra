---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "styra_policy Resource - terraform-provider-styra"
subcategory: ""
description: |-
  Policy Resource
---

# styra_policy (Resource)

Policy Resource

## Example Usage

```terraform
resource "styra_policy" "my_policy" {
  if_none_match = "...my_if_none_match..."
  modules = {
    "East"          = "..."
    "revolutionize" = "..."
  }
  policy = "...my_policy..."
  signature = {
    excluded   = "{ \"see\": \"documentation\" }"
    signatures = "{ \"see\": \"documentation\" }"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `modules` (Map of String) module file name to rego (and also data.json/data.yaml if enabled for the tenant) contents dictionary
- `policy` (String) policy name

### Optional

- `if_none_match` (String) etag
- `signature` (Attributes) (see [below for nested schema](#nestedatt--signature))

### Read-Only

- `request_id` (String)
- `result` (String) Parsed as JSON.

<a id="nestedatt--signature"></a>
### Nested Schema for `signature`

Optional:

- `excluded` (String) Parsed as JSON.
- `signatures` (String) Parsed as JSON.

## Import

Import is supported using the following syntax:

```shell
terraform import styra_policy.my_styra_policy ""
```