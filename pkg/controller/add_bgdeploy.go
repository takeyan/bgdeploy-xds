package controller

import (
	"bgdeploy-xds/pkg/controller/bgdeploy"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, bgdeploy.Add)
}
