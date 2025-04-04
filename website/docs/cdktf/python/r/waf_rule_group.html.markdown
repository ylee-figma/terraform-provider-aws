---
subcategory: "WAF Classic"
layout: "aws"
page_title: "AWS: aws_waf_rule_group"
description: |-
  Provides a AWS WAF rule group resource.
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_waf_rule_group

Provides a WAF Rule Group Resource

## Example Usage

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.waf_rule import WafRule
from imports.aws.waf_rule_group import WafRuleGroup
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        example = WafRule(self, "example",
            metric_name="example",
            name="example"
        )
        aws_waf_rule_group_example = WafRuleGroup(self, "example_1",
            activated_rule=[WafRuleGroupActivatedRule(
                action=WafRuleGroupActivatedRuleAction(
                    type="COUNT"
                ),
                priority=50,
                rule_id=example.id
            )
            ],
            metric_name="example",
            name="example"
        )
        # This allows the Terraform resource name to match the original name. You can remove the call if you don't need them to match.
        aws_waf_rule_group_example.override_logical_id("example")
```

## Argument Reference

This resource supports the following arguments:

* `name` - (Optional) Name of the rule group. If omitted, Terraform will assign a random, unique name. Conflicts with `name_prefix`.
* `metric_name` - (Required) A friendly name for the metrics from the rule group
* `activated_rule` - (Optional) A list of activated rules, see below
* `tags` - (Optional) Key-value map of resource tags. If configured with a provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.

## Nested Blocks

### `activated_rule`

#### Arguments

* `action` - (Required) Specifies the action that CloudFront or AWS WAF takes when a web request matches the conditions in the rule.
    * `type` - (Required) e.g., `BLOCK`, `ALLOW`, or `COUNT`
* `priority` - (Required) Specifies the order in which the rules are evaluated. Rules with a lower value are evaluated before rules with a higher value.
* `rule_id` - (Required) The ID of a [rule](/docs/providers/aws/r/waf_rule.html)
* `type` - (Optional) The rule type, either [`REGULAR`](/docs/providers/aws/r/waf_rule.html), [`RATE_BASED`](/docs/providers/aws/r/waf_rate_based_rule.html), or `GROUP`. Defaults to `REGULAR`.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `id` - The ID of the WAF rule group.
* `arn` - The ARN of the WAF rule group.
* `tags_all` - A map of tags assigned to the resource, including those inherited from the provider [`default_tags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import WAF Rule Group using the id. For example:

```python
# DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
from constructs import Construct
from cdktf import TerraformStack
#
# Provider bindings are generated by running `cdktf get`.
# See https://cdk.tf/provider-generation for more details.
#
from imports.aws.waf_rule_group import WafRuleGroup
class MyConvertedCode(TerraformStack):
    def __init__(self, scope, name):
        super().__init__(scope, name)
        WafRuleGroup.generate_config_for_import(self, "example", "a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc")
```

Using `terraform import`, import WAF Rule Group using the id. For example:

```console
% terraform import aws_waf_rule_group.example a1b2c3d4-d5f6-7777-8888-9999aaaabbbbcccc
```

<!-- cache-key: cdktf-0.20.8 input-7cdba50cf752079762f2c3a705a382a86bcff79ad39e2b4b8ee487d126465107 -->