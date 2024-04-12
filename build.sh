go run main.go > out.as
as -o out.o out.as
ld -o a.out -L /Library/Developer/CommandLineTools/SDKs/MacOSX12.1.sdk/usr/lib -lSystem out.o

rm -rf out.s out.o