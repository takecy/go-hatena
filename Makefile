prepare:
    go get -u github.com/tools/godep

setup:
    godep restore

save:
    godep save ./...

update:
    go get -u ./...
    rm -rf Godep
    godep save ./...

test:
    go test ./...