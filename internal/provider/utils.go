package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

func convertStringList(items []string) []types.String {

	if len(items) == 0 {
		return nil
	}

	var attr []types.String
	for _, item := range items {
		attr = append(attr, types.StringValue(item))
	}

	return attr

}
