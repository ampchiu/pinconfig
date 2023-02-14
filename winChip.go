package main

import  imgui "github.com/AllenDang/cimgui-go"

var  selected int = 0

func winChip() {

	PINK := imgui.NewVec4(0.95, 0.75, 0.7, 1.0)
	ORIGIN := imgui.NewVec2(0, 0)

	imgui.SetNextWindowPosV(imgui.NewVec2(350, 40), imgui.CondOnce, ORIGIN)
	imgui.SetNextWindowSizeV(imgui.NewVec2(1200, 1000), imgui.CondOnce)
	imgui.BeginV("Chip View", nil, imgui.WindowFlagsNoResize|imgui.WindowFlagsNoCollapse|imgui.WindowFlagsNoDocking)
	for i := 0; i < bound[3]; i++ {
		imgui.SetCursorPos(imgui.NewVec2(pos(i)))
		if imgui.ButtonV(name(i), imgui.NewVec2(size(i))) {
			selected = i
		}
	}
	imgui.SetCursorPos(imgui.NewVec2(600, 500))

	imgui.TextColored(PINK, "HC32L136")
	imgui.SetCursorPos(imgui.NewVec2(600, 550))

	imgui.BeginChildFrame(1, imgui.NewVec2(200, 30))
	imgui.Text("HC32L136")
	imgui.EndChildFrame()

	imgui.End()
}


