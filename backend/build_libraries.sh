export LD_LIBRARY_PATH=$PWD:$LD_LIBRARY_PATH
gcc -c -fPIC library.c -o build/clib.o
gcc -shared -o libclib.so build/clib.o
ghc --make -dynamic -shared -fPIC library.hs -o libhlib.so










