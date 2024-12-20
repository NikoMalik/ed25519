#include "textflag.h"
#include "go_asm.h"
#include "funcdata.h"



TEXT ·memcopy_avx2_32(SB), $0-16

	MOVQ addr+0(FP), DI
	MOVQ addr1+8(FP), SI


    XORQ AX,AX 
    
    PCALIGN $32

LOOP:

    VMOVDQU 0(SI)(AX*1), Y0
    VMOVDQU Y0, 0(DI)(AX*1)
    
    ADDQ $32, AX
    CMPQ AX, $32
    JL   LOOP
    VZEROUPPER
    //PCALIGN $32
    RET

TEXT ·memcopy_avx2_64(SB), $0-16

    MOVQ addr+0(FP), DI
    MOVQ addr1+8(FP), SI

    XORQ AX,AX 
    PCALIGN $32

LOOP:

    VMOVDQU 0(SI)(AX*1), Y0
    VMOVDQU 32(SI)(AX*1), Y1
    VMOVDQU Y0, 0(DI)(AX*1)
    VMOVDQU Y1, 32(DI)(AX*1)

    
    ADDQ $64, AX
    CMPQ AX, $64
    JL   LOOP
    VZEROUPPER
    //PCALIGN $32
    RET



