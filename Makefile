SANDYPLATFORM=sandy32
SANDYDIR=$(USERPROFILE)\Documents\sandy

GOOS=windows
GOARCH=amd64
GOTAGS=$(SANDYPLATFORM)

.PHONY:buildemu build install runmu run clean
buildemu:
	@set GOOS=$(GOOS)
	@set GOARCH=$(GOARCH)
	@go build -o build/$(SANDYPLATFORM)/sandyemu.exe -tags $(GOTAGS) cmd/sandyemu/main.go

build: buildemu

install: build
	@if exist "$(SANDYDIR)\bin\$(SANDYPLATFORM)" rmdir /S /Q "$(SANDYDIR)\bin\$(SANDYPLATFORM)"
	@xcopy "build\$(SANDYPLATFORM)" "$(SANDYDIR)\bin\$(SANDYPLATFORM)\" /E /C /I >nul
	@if exist"$(SANDYDIR)\resources" rmdir /S /Q "$(SANDYDIR)\resources"
	@xcopy "resources" "$(SANDYDIR)\resources\" /E /C /I >nul

runmu: buildemu
	@if not exist "$(SANDYDIR)\bin\dev" mkdir "$(SANDYDIR)\bin\dev"
	@copy "build\$(SANDYPLATFORM)\sandyemu.exe" "$(SANDYDIR)\bin\dev\sandyemu.exe" >nul
	@cd "build\$(SANDYPLATFORM)" && .\sandyemu.exe -bios "$(SANDYDIR)\resources\bios\SCPH1001.bin"

run: build
	@if exist "$(SANDYDIR)\bin\dev" rmdir /S /Q "$(SANDYDIR)\bin\dev"
	@xcopy "build\$(SANDYPLATFORM)" "$(SANDYDIR)\bin\dev\" /E /C /I >nul
	@cd "build\$(SANDYPLATFORM)" && .\sandyemu.exe -bios "$(SANDYDIR)\resources\bios\SCPH1001.bin"

clean:
	@if exist "build" rmdir /S /Q build