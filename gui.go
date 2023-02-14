package main

import imgui "github.com/AllenDang/cimgui-go"

var window   imgui.GLFWwindow

func afterCreateContext() {
	imgui.PlotCreateContext()
	fonts := imgui.CurrentIO().Fonts()
	fonts.AddFontFromFileTTFV("FiraCode-Medium.ttf", 18, 0, fonts.GlyphRangesChineseFull())
}

func loop() {
	winControl()
	winChip()
}

func beforeDestroyContext() {
	imgui.PlotDestroyContext()
}

func guiMain() {
	imgui.SetAfterCreateContextHook(afterCreateContext)
	imgui.SetBeforeDestroyContextHook(beforeDestroyContext)

	window = imgui.CreateGlfwWindow("Pin Config", 1, 1, imgui.GLFWWindowFlagsTransparent)

	window.Run(loop)
}
