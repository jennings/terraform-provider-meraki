package provider

import (
	"fmt"
	"math/rand"
	"os"
	"testing"
)

func readTestVars(t *testing.T) *testVars {
	return &testVars{

		OrganizationID: os.Getenv("MERAKI_ORGANIZATION_ID"),
		NetworkID:      os.Getenv("MERAKI_NETWORK_ID"),
	}
}

type testVars struct {
	OrganizationID string
	NetworkID      string
}

func generateName(slug string) string {
	return fmt.Sprintf("tfacc-%s-%d", slug, rand.Uint32())
}
