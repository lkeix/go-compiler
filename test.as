.data
hello: .string "Hello, World"
hello_len = . - hello

.text

.globl _main

_main:
  mov x0, #1
  ldr x1, =hello
  ldr x2, =hello_len
  mov x16, #4
  svc #0

  mov x0, #0
  mov x16, #93
  svc #0
  ret