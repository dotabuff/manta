SteamKit := $(wildcard ~/github/SteamRE/SteamKit/Resources/Protobufs)

default : proto

proto: dota/google/protobuf/descriptor.pb.go
	rm -f dota/*.proto dota/*.pb.go
	cp -r $(SteamKit)/dota_s2/{client,engine2,server}/*.proto ./dota/
	rm -f dota/matchmaker_common.proto dota/rendermessages.proto dota/*.steamworkssdk.proto
	dos2unix -q dota/*.proto
	sed -i 's/^\(\s*\)\(optional\|repeated\|required\)\s*\./\1\2 /' dota/*.proto
	sed -i '1ipackage dota;\n' dota/*.proto
	protoc -I$(SteamKit) -Idota --go_out=dota dota/*.proto
	sed -i 's|google/protobuf/descriptor.pb|github.com/dotabuff/manta/dota/google/protobuf|' dota/*.pb.go

dota/google/protobuf/descriptor.pb.go : $(SteamKit)/google/protobuf/descriptor.proto
	mkdir -p dota
	protoc -I$(SteamKit)/ --go_out=dota $<
