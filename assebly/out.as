.globl _main
_main:
  movq 0x2000001, %rax
  xorq %rdi, %rdi
  syscall
