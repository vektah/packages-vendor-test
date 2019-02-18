package main_test

import (
	"github.com/external/tool"
	"golang.org/x/tools/go/packages"
	"testing"
)

func TestPackageLoading(t *testing.T) {
	// we can reference the vendored package just fine
	if tool.Foo() != true {
		t.Error("vendor imports are broken?")
	}

	// but if we try to load the package via packages
	pkgs, err := packages.Load(nil, "github.com/external/tool")
	if err != nil {
		t.Error(err.Error())
	}

	// we get an error!
	if pkgs[0].Errors != nil {
		t.Error(pkgs[0].Errors)
	}

	// and a blank package name
	if pkgs[0].Name != "tool" {
		t.Error("the package name should be set")
	}
}