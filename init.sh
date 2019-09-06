#~/bin/bash
go get -u golang.org/x/lint/golint
go get github.com/pulumi/scripts/gomod-doccopy
GO111MODULE=on go get github.com/pulumi/pulumi-terraform@master
GO111MODULE=on go get github.com/terraform-providers/terraform-provider-panos
GO111MODULE=on go mod vendor
make ensure
