// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package panos

import (
	"unicode"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pulumi/pulumi-terraform/pkg/tfbridge"
	"github.com/pulumi/pulumi/pkg/resource"
	"github.com/pulumi/pulumi/pkg/tokens"
	panos "github.com/terraform-providers/terraform-provider-panos/panos"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "panos"
	// modules:
	mainMod  = "index"    // the root index
	panorama = "panorama" // Panorama
	firewall = "firewall" // Firewall
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeDataSource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the data source's
// first character.
func makeDataSource(mod string, res string) tokens.ModuleMember {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeMember(mod+"/"+fn, res)
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

// boolRef returns a reference to the bool argument.
/*
func boolRef(b bool) *bool {
	return &b
}
*/

// stringValue gets a string value from a property map if present, else ""
/*
func stringValue(vars resource.PropertyMap, prop resource.PropertyKey) string {
	val, ok := vars[prop]
	if ok && val.IsString() {
		return val.StringValue()
	}
	return ""
}
*/

// preConfigureCallback is called before the providerConfigure function of the underlying provider.
// It should validate that the provider can be configured, and provide actionable errors in the case
// it cannot be. Configuration variables can be read from `vars` using the `stringValue` function -
// for example `stringValue(vars, "accessKey")`.
func preConfigureCallback(vars resource.PropertyMap, c *terraform.ResourceConfig) error {
	return nil
}

// managedByPulumi is a default used for some managed resources, in the absence of something more meaningful.
//var managedByPulumi = &tfbridge.DefaultInfo{Value: "Managed by Pulumi"}

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := panos.Provider().(*schema.Provider)

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:           p,
		Name:        "panos",
		Description: "A Pulumi package for creating and managing panos resources.",
		Keywords:    []string{"pulumi", "panos"},
		License:     "Apache-2.0",
		Homepage:    "https://pulumi.io",
		Repository:  "https://github.com/gpduck/pulumi-panos",
		Config: map[string]*tfbridge.SchemaInfo{
			// Add any required configuration here, or remove the example below if
			// no additional points are required.
			// "region": {
			// 	Type: makeType("region", "Region"),
			// 	Default: &tfbridge.DefaultInfo{
			// 		EnvVars: []string{"AWS_REGION", "AWS_DEFAULT_REGION"},
			// 	},
			// },
			"hostname": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PANOS_HOSTNAME"},
				},
			},
			"username": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PANOS_USERNAME"},
				},
			},
			"password": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PANOS_PASSWORD"},
				},
			},
			"api_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"PANOS_API_KEY"},
				},
			},
		},
		PreConfigureCallback: preConfigureCallback,
		Resources: map[string]*tfbridge.ResourceInfo{
			// Map each resource in the Terraform provider to a Pulumi type. Two examples
			// are below - the single line form is the common case. The multi-line form is
			// needed only if you wish to override types or other default options.
			//
			// "aws_iam_role": {Tok: makeResource(mainMod, "IamRole")}
			//
			// "aws_acm_certificate": {
			// 	Tok: makeResource(mainMod, "Certificate"),
			// 	Fields: map[string]*tfbridge.SchemaInfo{
			// 		"tags": {Type: makeType(mainPkg, "Tags")},
			// 	},
			// },
			"panos_panorama_address_group":               {Tok: makeResource(panorama, "AddressGroup")},
			"panos_panorama_address_object":              {Tok: makeResource(panorama, "Address")},
			"panos_panorama_administrative_tag":          {Tok: makeResource(panorama, "AdministrativeTag")},
			"panos_address_group":                        {Tok: makeResource(firewall, "AddressGroup")},
			"panos_address_object":                       {Tok: makeResource(firewall, "Address")},
			"panos_administrative_tag":                   {Tok: makeResource(firewall, "AdministrativeTag")},
			"panos_aggregate_interface":                  {Tok: makeResource(firewall, "AggregateInterface")},
			"panos_application_group":                    {Tok: makeResource(firewall, "ApplicationGroup")},
			"panos_application_object":                   {Tok: makeResource(firewall, "Application")},
			"panos_application_signature":                {Tok: makeResource(firewall, "ApplicationSignature")},
			"panos_bfd_profile":                          {Tok: makeResource(firewall, "BfdProfile")},
			"panos_bgp":                                  {Tok: makeResource(firewall, "Bgp")},
			"panos_bgp_aggregate":                        {Tok: makeResource(firewall, "BgpAggregate")},
			"panos_bgp_aggregate_advertise_filter":       {Tok: makeResource(firewall, "BgpAggregateAdvertiseFilter")},
			"panos_bgp_aggregate_suppress_filter":        {Tok: makeResource(firewall, "BgpAggregateSuppressFilter")},
			"panos_bgp_auth_profile":                     {Tok: makeResource(firewall, "BgpAuthProfile")},
			"panos_bgp_conditional_adv":                  {Tok: makeResource(firewall, "BgpConditionalAdv")},
			"panos_bgp_conditional_adv_advertise_filter": {Tok: makeResource(firewall, "BgpConditionalAdvAdvertiseFilter")},
			"panos_bgp_conditional_adv_non_exist_filter": {Tok: makeResource(firewall, "BgpConditionalAdvNonExistFilter")},
			"panos_bgp_dampening_profile":                {Tok: makeResource(firewall, "BgpDampeningProfile")},
			"panos_bgp_export_rule_group":                {Tok: makeResource(firewall, "BgpExportRuleGroup")},
			"panos_bgp_import_rule_group":                {Tok: makeResource(firewall, "BgpImportRuleGroup")},
			"panos_bgp_peer":                             {Tok: makeResource(firewall, "BgpPeer")},
			"panos_bgp_peer_group":                       {Tok: makeResource(firewall, "BgpPeerGroup")},
			"panos_bgp_redist_rule":                      {Tok: makeResource(firewall, "BgpRedistRule")},
			"panos_dag_tags":                             {Tok: makeResource(firewall, "DagTags")},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			// Map each resource in the Terraform provider to a Pulumi function. An example
			// is below.
			// "aws_ami": {Tok: makeDataSource(mainMod, "getAmi")},
			"panos_dhcp_interface_info": {Tok: makeDataSource(mainMod, "getDhcpInterfaceInfo")},
			"panos_panorama_plugin":     {Tok: makeDataSource(panorama, "getPlugin")},
			"panos_system_info":         {Tok: makeDataSource(mainMod, "getSystemInfo")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "latest",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=1.0.0,<2.0.0",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}
