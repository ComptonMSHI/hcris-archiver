@echo OFF
set BUILDROOT=%cd%

set SOURCE=%BUILDROOT%
set TARGET=%BUILDROOT%\dist
set GOPATH=%SOURCE%

echo Building using %SOURCE%
echo --

echo Checking for target directory.
if not exist %TARGET% mkdir %TARGET%
if exist %TARGET%\win rmdir /S /Q %TARGET%\win
if not exist %TARGET%\win mkdir %TARGET%\win

echo Copying config template.
copy /Y %BUILDROOT%\config.yaml.template %TARGET%\win\config.yaml

echo Setting up packages.
cd %SOURCE%

go get github.com/ComptonMSHI/hcris-tools
go get github.com/mattn/go-sqlite3
go get github.com/yhat/scrape
go get golang.org/x/net/html
go get golang.org/x/net/html/atom
go get gopkg.in/yaml.v2
go get github.com/fatih/color
go get github.com/skratchdot/open-golang/open

echo Creating Windows binary. (Use build.sh for Linux/macOS)
go build -o %TARGET%/win/hcris-archiver.exe

echo Finished.