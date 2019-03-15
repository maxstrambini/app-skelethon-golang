title buildAPP

rem prompt $p$g
set GOPATH=%~dp0
set GOBIN=%~dp0bin

go build app.go config.go
pause

