.PHONY: clean, build, zip

clean:
	rm -rf dist/lambda

# GOOS=linux go build -o バイナリファイル名 goファイル名
build:
	GOARCH=amd64 GOOS=linux go build -o dist/lambda/regist-user app/lambda/regist-user/main.go

# zip ファイル名 バイナリファイル名
zip:
	zip zip/regist-user-handler.zip dist/lambda/regist-user