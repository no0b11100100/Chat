module server

go 1.16

require (
	command v1.0.0
	github.com/mattn/go-sqlite3 v1.14.8 // indirect
)

replace command => ./../command
