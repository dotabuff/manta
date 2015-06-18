default: build

test:
	go test -cover -v

bench:
	go test -bench=. -v

cover:
	go test -cover -coverprofile /tmp/manta.cov -v
	go tool cover -html=/tmp/manta.cov

update: update-game-tracking gen-dota-proto gen-game-events gen-message-lookup

game-tracking:
	git init game-tracking
	cd game-tracking && \
	git remote add -f origin https://github.com/SteamDatabase/GameTracking && \
	git config core.sparseCheckout true && \
	echo Protobufs/dota/ >> .git/info/sparse-checkout && \
	echo Protobufs/dota_reborn/ >> .git/info/sparse-checkout && \
	echo Protobufs/dota_s2/ >> .git/info/sparse-checkout && \
	echo Protobufs/dota_test/ >> .git/info/sparse-checkout && \
	git pull --depth=1 origin master

update-game-tracking: game-tracking
	git -C game-tracking checkout master
	git -C game-tracking pull origin master

gen-dota-proto: dota/google/protobuf/descriptor.pb.go
	rm -rf dota/*.proto
	cp -f game-tracking/Protobufs/dota_reborn/*/*.proto -t dota/ || true
	sed -i 's/^\(\s*\)\(optional\|repeated\|required\|extend\)\s*\./\1\2 /' dota/*.proto
	sed -i 's!^\s*rpc\s*\(\S*\)\s*(\.\([^)]*\))\s*returns\s*(\.\([^)]*\))\s*{!rpc \1 (\2) returns (\3) {!' dota/*.proto
	sed -i '1ipackage dota;\n' dota/*.proto
	cp -r google dota
	protoc -Idota --go_out=dota dota/*.proto
	sed -i 's|google/protobuf/descriptor.pb|github.com/dotabuff/manta/dota/google/protobuf|' dota/*.pb.go

dota/google/protobuf/descriptor.pb.go: google/protobuf/descriptor.proto
	mkdir -p dota/google/protobuf
	protoc -I. --go_out=dota $<

gen-game-events:
	go run gen/game_event.go fixtures/source_1_legacy_game_events_list.pbmsg game_event_lookup.go

gen-message-lookup:
	go run gen/message_lookup.go dota message_lookup.go

gen-wire-proto:
	protoc --go_out=. wire.proto
	sed -i 's/Wire/wire/g' wire.pb.go

sync-replays:
	s3cmd --region=us-west-2 sync ./replays/*.dem s3://manta.dotabuff/
