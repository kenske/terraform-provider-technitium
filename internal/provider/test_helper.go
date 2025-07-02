package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func GetFileConfig(t *testing.T, file string) string {

	providerBytes, err := os.ReadFile("testdata/provider.tf")
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}

	configBytes, err := os.ReadFile(fmt.Sprintf("testdata/%s", file))
	if err != nil {
		t.Fatalf("failed to read config: %v", err)
	}

	return string(providerBytes) + string(configBytes)
}

var _ plancheck.PlanCheck = expectDrift{}

type expectDrift struct {
	resourceName  string
	attributeName string
	beforeValue   string
}

func (e expectDrift) CheckPlan(ctx context.Context, req plancheck.CheckPlanRequest, resp *plancheck.CheckPlanResponse) {

	for _, rd := range req.Plan.ResourceDrift {
		if rd.Address == e.resourceName {
			beforeMap, ok := rd.Change.After.(map[string]any)
			beforeValue := beforeMap[e.attributeName]

			if !ok {
				resp.Error = fmt.Errorf("before is not a map for resource %s", e.resourceName)
				return
			}

			if beforeValue == e.beforeValue {
				tflog.Info(ctx, fmt.Sprintf("Resource %s has no drift for attribute %s", e.resourceName, e.attributeName))
				return
			}

			resp.Error = fmt.Errorf("resource %s has drift for attribute %s: expected %q, got %q", e.resourceName, e.attributeName, e.beforeValue, beforeValue)
			return
		}
	}

	resp.Error = fmt.Errorf("resource %s not found in drift check", e.resourceName)
	return

}

func ExpectDrift(resourceName string, attributeName string, expectedValue string) plancheck.PlanCheck {

	return expectDrift{
		resourceName:  resourceName,
		attributeName: attributeName,
		beforeValue:   expectedValue,
	}
}
