module pinconfig/m

go 1.19

require (
	github.com/AllenDang/cimgui-go v0.0.0-20230130042708-84cf33bf89a2
	github.com/sqweek/dialog v0.0.0-20220809060634-e981b270ebbf
)

replace github.com/sqweek/dialog => ../dialog

require github.com/TheTitanrain/w32 v0.0.0-20180517000239-4f5cfb03fabf // indirect
