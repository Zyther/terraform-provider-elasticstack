package integration_ds

import (
	"github.com/elastic/terraform-provider-elasticstack/generated/kbapi"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type integrationDataSourceModel struct {
	ID         types.String `tfsdk:"id"`
	Name       types.String `tfsdk:"name"`
	Prerelease types.Bool   `tfsdk:"prerelease"`
	Version    types.String `tfsdk:"version"`
}

func (m *integrationDataSourceModel) populateFromAPI(pkgName string, packages []kbapi.PackageListItem) {
	m.Version = types.StringNull()
	for _, pkg := range packages {
		if pkg.Name == pkgName {
			m.Version = types.StringValue(pkg.Version)
			return
		}
	}
}
