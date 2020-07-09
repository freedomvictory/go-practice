#include <stdio.h>



int* func_f(void)
{
    int v = 1;
    return &v;
}

int main(void)
{

    int *p = func_f();
    printf("p = %#x, *p = %d\n", p, *p);
    return 0;
}
