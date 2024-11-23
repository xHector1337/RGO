#include <sys/ptrace.h>
#include <stdlib.h>

void check(){
	if (ptrace(PTRACE_TRACEME,0,1,0) == -1){
		exit(1);
	}
}
