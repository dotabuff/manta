default: build

test:
	go test -cover -v

update: update-game-tracking gen-dota-proto gen-message-lookup

game-tracking:
	git init game-tracking
	cd game-tracking && \
	git remote add -f origin https://github.com/SteamDatabase/GameTracking && \
	git config core.sparseCheckout true && \
	echo Protobufs/dota/ >> .git/info/sparse-checkout && \
	echo Protobufs/dota_s2/ >> .git/info/sparse-checkout && \
	echo Protobufs/dota_test/ >> .git/info/sparse-checkout && \
	git pull --depth=1 origin master

update-game-tracking: game-tracking
	git -C game-tracking checkout master
	git -C game-tracking pull origin master

gen-dota-proto: dota/google/protobuf/descriptor.pb.go
	rm -rf dota/*.proto
	cp -f game-tracking/Protobufs/dota/*.proto dota/
	cp -f game-tracking/Protobufs/dota_s2/*/*.proto dota/
	sed -i 's/^\(\s*\)\(optional\|repeated\|required\|extend\)\s*\./\1\2 /' dota/*.proto
	sed -i 's!^\s*rpc\s*\(\S*\)\s*(\.\([^)]*\))\s*returns\s*(\.\([^)]*\))\s*{!rpc \1 (\2) returns (\3) {!' dota/*.proto
	sed -i '1ipackage dota;\n' dota/*.proto
	cp -r google dota
	protoc -Idota --go_out=dota dota/*.proto
	sed -i 's|google/protobuf/descriptor.pb|github.com/dotabuff/manta/dota/google/protobuf|' dota/*.pb.go

dota/google/protobuf/descriptor.pb.go: google/protobuf/descriptor.proto
	mkdir -p dota/google/protobuf
	protoc -I. --go_out=dota $<

gen-message-lookup:
	go run gen/packet.go dota message_lookup.go
