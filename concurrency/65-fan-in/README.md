# Block Profile

curl "http://localhost:6060/debug/pprof/block" >block.pprof 

go tool pprof -http=localhost:8082 block.pprof


curl "http://localhost:6060/debug/pprof/trace?seconds=5" > trace.out
go tool trace trace.out


