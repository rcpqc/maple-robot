set GOOS=linux
set GOARCH=arm64
go build
adb push maple-robot /mnt/user
adb shell chmod +x /mnt/user/maple-robot