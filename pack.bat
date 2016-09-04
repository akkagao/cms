
echo "开始打包$tarfile..."

set GOARCH=amd64
set GOOS=linux

bee pack -exs="pack.sh:pack:bat:nginx.conf" -exr=data

ren cms.tar.gz cms-v1.0.tar.gz
