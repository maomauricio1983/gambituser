git add .
git commit -m "Ultimo Commit"
git push

set GOOS=linux
set GOARCH=amd64

go build main.go
del main.zip
zip main.zip main
