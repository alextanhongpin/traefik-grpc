
stub:
	find . -name "*.proto" -exec \
	protoc -I/usr/local/include -I. \
	-I${GOPATH}/src \
	-I${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	--gofast_out=plugins=grpc:. \
	--proto_path=. "{}" \;


server-cert:
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./backend.key -out ./backend.cert

client-cert:
	openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout ./frontend.key -out ./frontend.cert