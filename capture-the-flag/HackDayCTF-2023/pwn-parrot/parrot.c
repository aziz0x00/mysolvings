//----- (0000000000001165) ----------------------------------------------------
int __cdecl main(int argc, const char **argv, const char **envp)
{
  char s[112]; // [rsp+0h] [rbp-70h] BYREF

  puts("PARROT SIMULATOR");
  while ( should_exit != 1 )
  {
    printf("*The parrot stares at you, waiting for something*\n> ");
    fflush(_bss_start);
    fgets(s, 100, stdin);
    printf(s);
  }
  return 0;
}

//----- (00000000000011D9) ----------------------------------------------------
__int64 s()
{
  __int64 savedregs; // [rsp+0h] [rbp+0h]

  __asm { syscall; LINUX - }
  return savedregs;
}
