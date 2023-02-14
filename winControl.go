package main

import (
	"fmt"
	"os"
    dlg "github.com/sqweek/dialog"
	imgui "github.com/AllenDang/cimgui-go"
)

func save() {
	file, err := dlg.File().Title("Save As").Filter("All Files", "*").Save()
	fmt.Println(file)
    if err != nil {
	    fmt.Println("Error:", err)
    }
}

func load() {
	file, err := dlg.File().Title("Load").Filter("All Files", "*").Load()
	fmt.Println(file)
    if err != nil {
    	fmt.Println("Error:", err)
    }
}

func winControl() {
	imgui.SetNextWindowPosV(imgui.NewVec2(40, 40), imgui.CondOnce, imgui.NewVec2(0, 0))
	imgui.SetNextWindowSizeV(imgui.NewVec2(300, 1000), imgui.CondOnce)
	imgui.BeginV("Control Panel", nil, imgui.WindowFlagsNoResize+imgui.WindowFlagsNoCollapse+imgui.WindowFlagsNoDocking)

	if imgui.Button("exit") {
		os.Exit(0)
	}
    if imgui.Button("save") {
        save()
    }
    if imgui.Button("load") {
        load()
    }
    if imgui.IsKeyPressed(imgui.KeyQ){
        os.Exit(1)
    }

	imgui.Text(fmt.Sprintf("selpin:%v", selected+1))
	imgui.End()
}


