default: build

test:
	go test -cover -v

bench:
	go test -run=XXX -bench=. -v

cover:
	go test -cover -coverpkg github.com/dotabuff/manta,github.com/dotabuff/manta/vbkv -coverprofile /tmp/manta.cov -v
	go tool cover -html=/tmp/manta.cov

cpuprofile:
	go test -v -run=TestMatch2159568145 -test.cpuprofile=/tmp/manta.cpuprof
	go tool pprof -svg -output=/tmp/manta.cpuprof.svg manta.test /tmp/manta.cpuprof
	open /tmp/manta.cpuprof.svg

update: update-game-tracking gen-dota-proto generate

game-tracking:
	git init game-tracking
	cd game-tracking && \
	git remote add -f origin https://github.com/SteamDatabase/GameTracking && \
	git config core.sparseCheckout true && \
	echo Protobufs/dota/ >> .git/info/sparse-checkout && \
	echo Protobufs/dota_test/ >> .git/info/sparse-checkout && \
	git pull --depth=1 origin master

update-game-tracking: game-tracking
	git -C game-tracking checkout master
	git -C game-tracking pull origin master

gen-dota-proto: tmp/google/protobuf/descriptor.proto
	rm -f dota/*.proto dota/*.pb.go
	cp -f game-tracking/Protobufs/dota/*.proto dota/
	sed -i 's/^\(\s*\)\(optional\|repeated\|required\|extend\)\s*\./\1\2 /' dota/*.proto
	sed -i 's!^\s*rpc\s*\(\S*\)\s*(\.\([^)]*\))\s*returns\s*(\.\([^)]*\))\s*{!rpc \1 (\2) returns (\3) {!' dota/*.proto
	sed -i '1ipackage dota;\n' dota/*.proto
	protoc -I ./tmp -I dota --go_out=dota dota/*.proto
	sed -i 's|google/protobuf/descriptor.pb|github.com/golang/protobuf/protoc-gen-go/descriptor|' dota/*.pb.go

tmp/google/protobuf/descriptor.proto:
	mkdir -p tmp/google/protobuf
	curl https://raw.githubusercontent.com/google/protobuf/master/src/google/protobuf/descriptor.proto >tmp/google/protobuf/descriptor.proto

generate:
	go generate

sync-replays:
	s3cmd --region=us-west-2 sync ./replays/*.dem s3://manta.dotabuff/
