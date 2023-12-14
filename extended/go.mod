module example.com/extended

go 1.20

require (
	github.com/aws/constructs-go/constructs/v10 v10.1.52
	github.com/aws/jsii-runtime-go v1.63.2
	github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2 v2.3.61
	github.com/coopnorge/cdk8slibrary v0.0.1-alpha4
)

require (
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/creasty/defaults v1.6.0 // indirect
)

//replace github.com/coopnorge/cdk8slibrary v0.0.1-alpha5 => ../../cdk8slibrary
