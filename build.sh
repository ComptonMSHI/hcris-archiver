BUILDROOT=`pwd`

SOURCE="${BUILDROOT}"
TARGET="${BUILDROOT}/dist"
GOPATH="${SOURCE}"

echo "Building using ${SOURCE}"
echo "--"

echo "Checking for target directory."
mkdir "${TARGET}"
rm -rf "${TARGET}/mac-lin"
mkdir "${TARGET}/mac-lin"

echo "Copying config template."
cp "${BUILDROOT}/config.yaml.template" "${TARGET}/mac-lin/config.yaml"

echo "Setting up packages."
cd "${SOURCE}"

go get github.com/ComptonMSHI/hcris-tools
go get github.com/mattn/go-sqlite3
go get github.com/yhat/scrape
go get golang.org/x/net/html
go get golang.org/x/net/html/atom
go get gopkg.in/yaml.v2
go get github.com/fatih/color
go get github.com/skratchdot/open-golang/open

echo "Creating linux/macOS binary. (Use build.bat for Windows)"
go build -o "${TARGET}/mac-lin/hcris-archiver"

echo "Finished."