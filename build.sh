go build -o a.out main.go
./a.out "$1" > out.as

as -o out.o out.as
ld -o go-compiler -L /Library/Developer/CommandLineTools/SDKs/MacOSX14.4.sdk/usr/lib -lSystem out.o

rm -rf out.s out.o