buildc:
	mkdir -p ./build
	env GOOS=windows GOARCH=amd64 go build -a -o palcd.exe .
	zip ./build/windows-amd64-palcd.zip palcd.exe 
	rm palcd.exe

	env GOOS=windows GOARCH=386 go build -a -o palcd.exe .
	zip ./build/windows-386-palcd.zip palcd.exe 
	rm palcd.exe

	env GOOS=darwin GOARCH=amd64 go build -a -o palcd .
	chmod +x palcd
	zip ./build/macos-amd64-palcd.zip palcd 
	rm palcd

	env GOOS=darwin GOARCH=arm64 go build -a -o palcd .
	chmod +x palcd
	zip ./build/macos-arm64-palcd.zip palcd 
	rm palcd