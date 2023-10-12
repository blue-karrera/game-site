export LD_LIBRARY_PATH="`ghc --print-libdir`/include":$PWD:$LD_LIBRARY_PATH
echo $LD_LIBRARY_PATH

ghc --make -no-hs-main -v -optc-O test.c Safe -o test

gcc -c -fPIC library.c -o build/clib.o
gcc -shared -o libclib.so build/clib.o

#ghc --make -dynamic -shared -fPIC library.hs -o libhlib.so
ghc -c -O Safe.hs -odir build
cp Safe.hi build/
rm Safe.hi
cp Safe_stub.h build/
rm Safe_stub.h










