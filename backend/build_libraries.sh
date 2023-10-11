export LD_LIBRARY_PATH=$PWD:$LD_LIBRARY_PATH
gcc -c -fPIC library.c -o build/library.o
gcc -shared -o libclib.so build/library.o









