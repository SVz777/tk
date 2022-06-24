#include "textflag.h"

TEXT ·GetGoID(SB),NOSPLIT,$0-8
    MOVQ (TLS),AX
    MOVQ ·offset(SB),BX
    MOVQ (AX)(BX*1),CX
    MOVQ CX,g+0(FP)
    RET
