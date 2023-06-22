build:
	@./src/scripts/build.sh
win-build:
	@powershell.exe -noexit "& '.\src\scripts\win-build.ps1'"

.PHONY: proto
proto:
#	@powershell.exe -noexit "& '.\scripts\gen-proto.ps1'"
	@./src/scripts/gen-proto.sh