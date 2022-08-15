package main

import (
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"

	"example.com/extended/imports/local_k8s"
	ck8slibrary "github.com/coopnorge/cdk8slibrary"
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
		ServiceEntries: []ck8slibrary.ServiceEntries{
			{
				Name:  jsii.String("google"),
				Hosts: []*string{jsii.String("google.com"), jsii.String("google.no")},
			},
		},
	})

	local_k8s.NewKubeJob(chart, jsii.String("extra-job"), &local_k8s.KubeJobProps{
		Metadata: &local_k8s.ObjectMeta{Name: jsii.String("a-job")},
		Spec: &local_k8s.JobSpec{
			Template: &local_k8s.PodTemplateSpec{
				Metadata: &local_k8s.ObjectMeta{Name: jsii.String("a-job")},
				Spec: &local_k8s.PodSpec{
					Containers: &[]*local_k8s.Container{{
						Image: jsii.String("image/job:tag"),
						Name:  jsii.String("a-job-name"),
					}},
				},
			},
		},
	})

	return chart
}

func main() {
	app := cdk8s.NewApp(nil)
	NewMyChart(app, "extended", nil)
	app.Synth()
}
