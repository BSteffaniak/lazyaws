package controllers

import "github.com/BSteffaniak/lazyaws/pkg/gui/types"

func AttachControllers(context types.Context, controllers ...types.IController) {
	for _, controller := range controllers {
		context.AddKeybindingsFn(controller.GetKeybindings)
		context.AddMouseKeybindingsFn(controller.GetMouseKeybindings)
		context.AddOnClickFn(controller.GetOnClick())
	}
}
