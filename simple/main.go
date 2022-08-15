package main

import (
	ck8slibrary "github.com/coopnorge/cdk8slibrary"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	cdk8s "github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type MyChartProps struct {
	cdk8s.ChartProps
}

func NewMyChart(scope constructs.Construct, id string, props *MyChartProps) cdk8s.Chart {
	var cprops cdk8s.ChartProps
	if props != nil {
		cprops = props.ChartProps
	}
	chart := cdk8s.NewChart(scope, jsii.String(id), &cprops)

	ck8slibrary.Chart(chart, jsii.String("hello"), &ck8slibrary.ChartProps{
		Name:  jsii.String("ultima"),
		Image: jsii.String("bla/boe:die"),
		Ports: []ck8slibrary.ChartPorts{
			{Port: jsii.Number(3004)},
		},
		//MaxReplicas: jsii.Number(4),
		ServiceEntries: []ck8slibrary.ServiceEntries{
			{
				Name:  jsii.String("google"),
				Hosts: []*string{jsii.String("google.com"), jsii.String("google.no")},
			},
		},
		//PdbMinAvailable: jsii.String("2"),
	})

	return chart
}

func main() {
	app := cdk8s.NewApp(nil)
	NewMyChart(app, "test", nil)
	app.Synth()
}
