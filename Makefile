SED=sed
ifeq ($(shell uname), Darwin)
	SED=gsed
endif

default: build

test:
	go test -cover -v

testnew:
	go test -cover -run=TestMatchNew -v

bench:
	go test -run=XXX -bench=BenchmarkMatch -benchtime=1m -v

cover:
	go test -cover -coverpkg github.com/dotabuff/manta,github.com/dotabuff/manta/vbkv -coverprofile /tmp/manta.cov -v
	go tool cover -html=/tmp/manta.cov

cpuprofile:
	go test -v -run=TestMatch2159568145 -test.cpuprofile=/tmp/manta.cpuprof
	go tool pprof -svg -output=/tmp/manta.cpuprof.svg manta.test /tmp/manta.cpuprof
	open /tmp/manta.cpuprof.svg

memprofile:
	go test -v -run=TestMatch2159568145 -test.memprofile=/tmp/manta.memprof -test.memprofilerate=1
	go tool pprof --alloc_space manta.test /tmp/manta.memprof

update: update-protobufs generate

update-protobufs:
	rm -rf dota
	mkdir -p ./dota/tmp && \
		curl -L -o - https://github.com/SteamDatabase/GameTracking-Dota2/archive/master.tar.gz | tar -xz --strip-components=1 -C ./dota/tmp && \
		cp -a ./dota/tmp/Protobufs/*.proto ./dota/ && \
		rm -rf ./dota/tmp
	rm -rf dota/gametoolevents.proto dota/dota_messages_mlbot.proto dota/dota_gcmessages_common_bot_script.proto dota/steammessages_base.proto dota/steammessages_clientserver_login.proto dota/tensorflow
	$(SED) -i 's/\.CMsgFightingGame_GameData/CMsgFightingGame_GameData/g' dota/dota_fighting_game_p2p_messages.proto
	$(SED) -i 's/^\(\s*\)\(optional\|repeated\|required\|extend\)\s*\./\1\2 /' dota/*.proto
	$(SED) -i 's!^\s*rpc\s*\(\S*\)\s*(\.\([^)]*\))\s*returns\s*(\.\([^)]*\))\s*{!rpc \1 (\2) returns (\3) {!' dota/*.proto
	$(SED) -i '1isyntax = "proto2";\n\npackage dota;\noption go_package = "github.com/dotabuff/manta/dota;dota";\n' dota/*.proto
	$(SED) -i '/^import "google\/protobuf\/valve_extensions\.proto"/d' dota/*.proto
	$(SED) -i '/^option (/d' dota/*.proto
	$(SED) -i 's/\s\[.*\]//g' dota/*.proto
	$(SED) -i 's/\.CMsgSteamLearn/CMsgSteamLearn/g' dota/*.proto
	$(SED) -i 's/\.CMsgShowcaseItem/CMsgShowcaseItem/g' dota/*.proto
	$(SED) -i 's/\.CMsgShowcaseBackground/CMsgShowcaseBackground/g' dota/*.proto
	$(SED) -i 's/\.CMsgSurvivorsUserData/CMsgSurvivorsUserData/g' dota/*.proto
	protoc -I dota --go_out=paths=source_relative:dota  dota/*.proto

generate:
	go run gen/callbacks.go

sync-replays:
	s3cmd --region=us-west-2 sync ./replays/*.dem s3://manta.dotabuff/
