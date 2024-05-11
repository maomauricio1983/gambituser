git add .
git commit -m "Ultimo Commit"
git push

set GOOS=linux
set GOARCH=amd64

go build -tags lambda.enospc -o bootstrap.main.go
del main.zip
tar.exe -a -cf main.zip main
