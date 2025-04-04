---
subcategory: "Route 53"
layout: "aws"
page_title: "AWS: aws_route53_zone"
description: |-
  Manages a Route53 Hosted Zone
---


<!-- Please do not edit this file, it is generated. -->
# Resource: aws_route53_zone

Manages a Route53 Hosted Zone. For managing Domain Name System Security Extensions (DNSSEC), see the [`aws_route53_key_signing_key`](route53_key_signing_key.html) and [`aws_route53_hosted_zone_dnssec`](route53_hosted_zone_dnssec.html) resources.

## Example Usage

### Public Zone

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Route53Zone } from "./.gen/providers/aws/route53-zone";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new Route53Zone(this, "primary", {
      name: "example.com",
    });
  }
}

```

### Public Subdomain Zone

For use in subdomains, note that you need to create a
`aws_route53_record` of type `NS` as well as the subdomain
zone.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { Token, TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Route53Record } from "./.gen/providers/aws/route53-record";
import { Route53Zone } from "./.gen/providers/aws/route53-zone";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    const dev = new Route53Zone(this, "dev", {
      name: "dev.example.com",
      tags: {
        Environment: "dev",
      },
    });
    const main = new Route53Zone(this, "main", {
      name: "example.com",
    });
    new Route53Record(this, "dev-ns", {
      name: "dev.example.com",
      records: Token.asList(dev.nameServers),
      ttl: Token.asNumber("30"),
      type: "NS",
      zoneId: main.zoneId,
    });
  }
}

```

### Private Zone

~> **NOTE:** Terraform provides both exclusive VPC associations defined in-line in this resource via `vpc` configuration blocks and a separate [Zone VPC Association](/docs/providers/aws/r/route53_zone_association.html) resource. At this time, you cannot use in-line VPC associations in conjunction with any `aws_route53_zone_association` resources with the same zone ID otherwise it will cause a perpetual difference in plan output. You can optionally use the generic Terraform resource [lifecycle configuration block](https://www.terraform.io/docs/configuration/meta-arguments/lifecycle.html) with `ignore_changes` to manage additional associations via the `aws_route53_zone_association` resource.

~> **NOTE:** Private zones require at least one VPC association at all times.

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Route53Zone } from "./.gen/providers/aws/route53-zone";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    new Route53Zone(this, "private", {
      name: "example.com",
      vpc: [
        {
          vpcId: example.id,
        },
      ],
    });
  }
}

```

## Argument Reference

This resource supports the following arguments:

* `name` - (Required) This is the name of the hosted zone.
* `comment` - (Optional) A comment for the hosted zone. Defaults to 'Managed by Terraform'.
* `delegationSetId` - (Optional) The ID of the reusable delegation set whose NS records you want to assign to the hosted zone. Conflicts with `vpc` as delegation sets can only be used for public zones.
* `forceDestroy` - (Optional) Whether to destroy all records (possibly managed outside of Terraform) in the zone when destroying the zone.
* `tags` - (Optional) A map of tags to assign to the zone. If configured with a provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block) present, tags with matching keys will overwrite those defined at the provider-level.
* `vpc` - (Optional) Configuration block(s) specifying VPC(s) to associate with a private hosted zone. Conflicts with the `delegationSetId` argument in this resource and any [`aws_route53_zone_association` resource](/docs/providers/aws/r/route53_zone_association.html) specifying the same zone ID. Detailed below.

### vpc Argument Reference

* `vpcId` - (Required) ID of the VPC to associate.
* `vpcRegion` - (Optional) Region of the VPC to associate. Defaults to AWS provider region.

## Attribute Reference

This resource exports the following attributes in addition to the arguments above:

* `arn` - The Amazon Resource Name (ARN) of the Hosted Zone.
* `zoneId` - The Hosted Zone ID. This can be referenced by zone records.
* `nameServers` - A list of name servers in associated (or default) delegation set.
  Find more about delegation sets in [AWS docs](https://docs.aws.amazon.com/Route53/latest/APIReference/actions-on-reusable-delegation-sets.html).
* `primaryNameServer` - The Route 53 name server that created the SOA record.
* `tagsAll` - A map of tags assigned to the resource, including those inherited from the provider [`defaultTags` configuration block](https://registry.terraform.io/providers/hashicorp/aws/latest/docs#default_tags-configuration-block).

## Timeouts

[Configuration options](https://developer.hashicorp.com/terraform/language/resources/syntax#operation-timeouts):

* `create` - (Default `30m`)
* `update` - (Default `30m`)
* `delete` - (Default `30m`)

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import Route53 Zones using the zone `id`. For example:

```typescript
// DO NOT EDIT. Code generated by 'cdktf convert' - Please report bugs at https://cdk.tf/bug
import { Construct } from "constructs";
import { TerraformStack } from "cdktf";
/*
 * Provider bindings are generated by running `cdktf get`.
 * See https://cdk.tf/provider-generation for more details.
 */
import { Route53Zone } from "./.gen/providers/aws/route53-zone";
class MyConvertedCode extends TerraformStack {
  constructor(scope: Construct, name: string) {
    super(scope, name);
    Route53Zone.generateConfigForImport(this, "myzone", "Z1D633PJN98FT9");
  }
}

```

Using `terraform import`, import Route53 Zones using the zone `id`. For example:

```console
% terraform import aws_route53_zone.myzone Z1D633PJN98FT9
```

<!-- cache-key: cdktf-0.20.8 input-9c75d13012d076c4e45ee2eed2fab48617563530adc7f1351bafd11b5bd43f70 -->