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

func generateName(prefix string) string {
	return fmt.Sprintf("%s%d", prefix, rand.Uint64())
}
