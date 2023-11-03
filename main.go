package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	launchdarkly_flag_eval "github.com/angelolapus/terraform-provider-launchdarkly-flag-evaluation/launchdarklytarget"
)

func main() {
	providerserver.Serve(context.Background(), launchdarkly_flag_eval.New, providerserver.ServeOpts{
		Address: "registry.terraform.io/angelolapus/launchdarkly",
	})
}