# Backup server

In case you are attending a live session and there is no internet connection, you can run a simple go server locally that mocks the hexbot/vexbot api endpoints.
Run `go run main.go` and once the server starts up, you can start performing requests to `http://0.0.0.0:3000/hexbot` and `http://0.0.0.0:3000/vexbot`. Also remember to replace in the code where necessary (e.g the constants `HexbotURL` and `VexbotURL`)