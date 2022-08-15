module example.com/simple

go 1.18

require (
	github.com/aws/constructs-go/constructs/v10 v10.1.65
	github.com/aws/jsii-runtime-go v1.63.2
	github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2 v2.3.74
)

require (
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/creasty/defaults v1.6.0 // indirect
)

require github.com/coopnorge/cdk8slibrary v0.0.1-alpha4

//replace example.com/cdk8slibrary v0.0.0 => /Users/atze/coopx/git/github/coopnorge/cdk8slibrary
