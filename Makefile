default: build

test:
	go test -cover -v

testnew:
	go test -cover -run=TestMatchNew -v

bench:
	go test -run=XXX -bench=BenchmarkMatch8552595443 -benchmem -benchtime=10x -count=10 -v

cover:
	go test -cover -coverpkg github.com/dotabuff/manta,github.com/dotabuff/manta/vbkv -coverprofile /tmp/manta.cov -v
	go tool cover -html=/tmp/manta.cov

cpuprofile:
	go test -v -run=TestMatchNew8552595443 -test.cpuprofile=/tmp/manta.cpuprof
	go tool pprof -svg -output=/tmp/manta.cpuprof.svg manta.test /tmp/manta.cpuprof
	open /tmp/manta.cpuprof.svg

memprofile:
	go test -v -run=TestMatchNew8552595443 -test.memprofile=/tmp/manta.memprof -test.memprofilerate=1
	go tool pprof --alloc_space manta.test /tmp/manta.memprof

update: update-protobufs generate

update-protobufs:
	go run gen/updateprotos/main.go

generate:
	go run gen/callbacks.go

sync-replays:
	s3cmd --region=us-west-2 sync ./replays/*.dem s3://manta.dotabuff/
