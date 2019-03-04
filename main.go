package main

import (
    "flag"
    "os"

    "github.com/golang/glog"
    "k8s.io/apimachinery/pkg/util/wait"
    "k8s.io/apiserver/pkg/util/logs"

    basecmd "github.com/kubernetes-incubator/custom-metrics-apiserver/pkg/cmd"
    "github.com/kubernetes-incubator/custom-metrics-apiserver/pkg/provider"

    rp10Provider "github.com/keni7385/rp10-custom-metrics/pkg/provider"
)

type Rp10Adapter struct {
	basecmd.AdapterBase

	Message string
}

func (a *Rp10Adapter) makeProviderOrDie() provider.CustomMetricsProvider {
	client, err := a.DynamicClient()
	if err != nil {
		glog.Fatalf("unable to construct dynamic client: %v", err)
	}

	mapper, err := a.RESTMapper()
	if err != nil {
		glog.Fatalf("unable to construct discovery REST mapper: %v", err)
	}

	return rp10Provider.NewProvider(client, mapper)
}

func main() {
	logs.InitLogs()
	defer logs.FlushLogs()

	// Initialize the flags, with one custom flag for the message
	cmd := &Rp10Adapter{}
	cmd.Flags().StringVar(&cmd.Message, "msg", "starting adapter...", "startup message")
	cmd.Flags().AddGoFlagSet(flag.CommandLine)
	cmd.Flags().Parse(os.Args)

	provider := cmd.makeProviderOrDie()
	cmd.WithCustomMetrics(provider)
	// you could also set up external metrics support,
    // if your provider supported it:
    // cmd.WithExternalMetrics(provider)

	glog.Infof(cmd.Message)
	if err := cmd.Run(wait.NeverStop); err != nil {
		glog.Fatalf("unable to run custom metrics adapter: %v", err)
	}
}
