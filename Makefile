default: build

test:
	go test -cover -v

bench:
	go test -run=XXX -bench=BenchmarkMatch -benchtime=1m -v

cover:
	go test -cover -coverpkg github.com/dotabuff/manta,github.com/dotabuff/manta/vbkv -coverprofile /tmp/manta.cov -v
	go tool cover -html=/tmp/manta.cov

cpuprofile:
	go test -v -run=TestMatch2159568145 -test.cpuprofile=/tmp/manta.cpuprof
	go tool pprof -svg -output=/tmp/manta.cpuprof.svg manta.test /tmp/manta.cpuprof
	open /tmp/manta.cpuprof.svg

update: update-game-tracking gen-dota-proto generate

update-game-tracking:
	git submodule update --init --remote

gen-dota-proto:
	rm -f dota/*.proto dota/*.pb.go
	cp -f game-tracking/Protobufs/*.proto dota/
	rm -f dota/gametoolevents.proto dota/steammessages_base.proto
	sed -i 's/^\(\s*\)\(optional\|repeated\|required\|extend\)\s*\./\1\2 /' dota/*.proto
	sed -i 's!^\s*rpc\s*\(\S*\)\s*(\.\([^)]*\))\s*returns\s*(\.\([^)]*\))\s*{!rpc \1 (\2) returns (\3) {!' dota/*.proto
	sed -i '1isyntax = "proto2";\n\npackage dota;\n' dota/*.proto
	protoc -I dota --go_out=dota dota/*.proto
	sed -i 's|google/protobuf|github.com/golang/protobuf/protoc-gen-go/descriptor|' dota/*.pb.go

generate:
	go run gen/callbacks.go

sync-replays:
	s3cmd --region=us-west-2 sync ./replays/*.dem s3://manta.dotabuff/
