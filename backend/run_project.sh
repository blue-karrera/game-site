export LD_LIBRARY_PATH="`ghc --print-libdir`/include":$PWD:$LD_LIBRARY_PATH

go run main.go


