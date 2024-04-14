.globl _main

_main:
  movq $0x2000001, %rax
  movq $1, %rdi
  syscall
