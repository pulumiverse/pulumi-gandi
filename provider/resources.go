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

package gandi

import (
	"fmt"
	"path/filepath"

	"github.com/go-gandi/terraform-provider-gandi/v2/gandi"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"

	"github.com/pulumiverse/pulumi-gandi/provider/v2/pkg/version"
)

const (
	mainPkg = "gandi"
)

// Provider returns additional overlaid schema and metadata associated with the provider..
func Provider() tfbridge.ProviderInfo {
	// Instantiate the Terraform provider
	p := shimv2.NewProvider(gandi.Provider())

	// Create a Pulumi provider mapping
	prov := tfbridge.ProviderInfo{
		P:    p,
		Name: "gandi",
		// DisplayName is a way to be able to change the casing of the provider
		// name when being displayed on the Pulumi registry
		DisplayName: "Gandi",
		// The default publisher for all packages is Pulumi.
		// Change this to your personal name (or a company name) that you
		// would like to be shown in the Pulumi Registry if this package is published
		// there.
		Publisher: "Pulumiverse",
		// LogoURL is optional but useful to help identify your package in the Pulumi Registry
		// if this package is published there.
		//
		// You may host a logo on a domain you control or add an SVG logo for your package
		// in your repository and use the raw content URL for that file as your logo URL.
		LogoURL: "https://raw.githubusercontent.com/pulumiverse/pulumi-gandi/main/logo.svg",
		// PluginDownloadURL is an optional URL used to download the Provider
		// for use in Pulumi programs
		// e.g https://github.com/org/pulumi-provider-name/releases/
		PluginDownloadURL: "github://api.github.com/pulumiverse",
		Description:       "A Pulumi package for creating and managing gandi cloud resources.",
		// category/cloud tag helps with categorizing the package in the Pulumi Registry.
		// For all available categories, see `Keywords` in
		// https://www.pulumi.com/docs/guides/pulumi-packages/schema/#package.
		Keywords:   []string{"pulumi", "gandi", "category/cloud"},
		License:    "Apache-2.0",
		Homepage:   "https://www.pulumi.com",
		Repository: "https://github.com/pulumiverse/pulumi-gandi",
		// The GitHub Org for the provider - defaults to `terraform-providers`. Note that this
		// should match the TF provider module's require directive, not any replace directives.
		GitHubOrg:               "go-gandi",
		TFProviderModuleVersion: "v2",
		Config: map[string]*tfbridge.SchemaInfo{
			"key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"GANDI_KEY"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"gandi_domain": {
				Tok: tfbridge.MakeResource(mainPkg, "domains", "Domain"),
			},
			"gandi_nameservers": {
				Tok: tfbridge.MakeResource(mainPkg, "domains", "Nameservers"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"nameservers": {
						CSharpName: "Servers",
					},
				},
			},
			"gandi_dnssec_key": {
				Tok: tfbridge.MakeResource(mainPkg, "domains", "DNSSecKey"),
			},
			"gandi_glue_record": {
				Tok: tfbridge.MakeResource(mainPkg, "domains", "GlueRecord"),
			},

			"gandi_mailbox": {
				Tok: tfbridge.MakeResource(mainPkg, "email", "Mailbox"),
			},
			"gandi_email_forwarding": {
				Tok: tfbridge.MakeResource(mainPkg, "email", "Forwarding"),
			},

			"gandi_livedns_domain": {
				Tok: tfbridge.MakeResource(mainPkg, "livedns", "Domain"),
			},
			"gandi_livedns_record": {
				Tok: tfbridge.MakeResource(mainPkg, "livedns", "Record"),
			},

			"gandi_simplehosting_instance": {
				Tok: tfbridge.MakeResource(mainPkg, "simplehosting", "Instance"),
			},
			"gandi_simplehosting_vhost": {
				Tok: tfbridge.MakeResource(mainPkg, "simplehosting", "VHost"),
			},
		},
		DataSources: map[string]*tfbridge.DataSourceInfo{
			"gandi_domain": {
				Tok: tfbridge.MakeDataSource(mainPkg, "domains", "getDomain"),
			},
			"gandi_glue_record": {
				Tok: tfbridge.MakeDataSource(mainPkg, "domains", "getGlueRecord"),
			},

			"gandi_livedns_domain": {
				Tok: tfbridge.MakeDataSource(mainPkg, "livedns", "getDomain"),
			},
			"gandi_livedns_domain_ns": {
				Tok: tfbridge.MakeDataSource(mainPkg, "livedns", "getDomainNameserver"),
			},

			"gandi_mailbox": {
				Tok: tfbridge.MakeDataSource(mainPkg, "email", "getMailbox"),
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName: "@pulumiverse/gandi",
			// List any npm dependencies and their versions
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // so we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
			// See the documentation for tfbridge.OverlayInfo for how to lay out this
			// section, or refer to the AWS provider. Delete this section if there are
			// no overlay files.
			//Overlay: &tfbridge.OverlayInfo{},
		},
		Python: &tfbridge.PythonInfo{
			PackageName: "pulumiverse_gandi",
			// List any Python dependencies and their version ranges
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumiverse/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
