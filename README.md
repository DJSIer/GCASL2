# GCASL2

## What is GCASL2?

GCASL2 is CASL2 compiler to COMET II

## 実装状況

### LD
load
- [x] `LD r1,r2`
- [x] `LD r,addr`
- [x] `LD r,LABEL`
- [x] `LD r,addr,x`

### LAD
load address
- [x] `LAD r,addr`
- [x] `LAD r,LABEL`
- [x] `LAD r,addr,x`

 ### ST
 store
- [x] `ST r,addr`
- [x] `ST r,LABEL`
- [x] `ST r,addr,x`

 ### ADDA
 add arthmetic
- [x] `ADDA r1,r2`
- [x] `ADDA r,addr`
- [x] `ADDA r,LABEL`
- [x] `ADDA r,addr,x`

 ### ADDL
 add logical
- [x] `ADDL r1,r2`
- [x] `ADDL r,addr`
- [x] `ADDL r,LABEL`
- [x] `ADDL r,addr,x`

 ### SUBA
 subtract arithmetic
- [x] `SUBA r1,r2`
- [x] `SUBA r,addr`
- [x] `SUBA r,LABEL`
- [x] `SUBA r,addr,x`

 ### SUBL
 substract logical
- [x] `SUBL r1,r2`
- [x] `SUBL r,addr`
- [x] `SUBL r,LABEL`
- [x] `SUBL r,addr,x`

 ### AND
- [x] `AND r1,r2`
- [x] `AND r,addr`
- [x] `AND r,LABEL`
- [x] `AND r,addr,x`

 ### OR
- [x] `OR r1,r2`
- [x] `OR r,addr`
- [x] `OR r,LABEL`
- [x] `OR r,addr,x`

 ### XOR
 exclusive or
- [x] `XOR r1,r2`
- [x] `XOR r,addr`
- [x] `XOR r,LABEL`
- [x] `XOR r,addr,x`

### CPA
compare arithmetic
- [x] `CPA r1,r2`
- [x] `CPA r,addr`
- [x] `CPA r,LABEL`
- [x] `CPA r,addr,x`

### CPL
compare logical
- [x] `CPL r1,r2`
- [x] `CPL r,addr`
- [x] `CPL r,LABEL`
- [x] `CPL r,addr,x`

### SLA
shift left arithmetic
- [x] `SLA r,addr`
- [x] `SLA r,LABEL`
- [x] `SLA r,addr,x`

### SRA
shift right arithmetic
- [x] `SRA r,addr`
- [x] `SRA r,LABEL`
- [x] `SRA r,addr,x`

### SLL
shift left logical
- [x] `SLL r,addr`
- [x] `SLL r,LABEL`
- [x] `SLL r,addr,x`

### SRL
shift right logical
- [x] `SRL r,addr`
- [x] `SRL r,LABEL`
- [x] `SRL r,addr,x`

### JMI
jump on minus
- [x] `JMI addr`
- [x] `JMI addr , x`
- [x] `JMI LABEL`
- [x] `JMI LABEL, x`

### JNZ
jump on nonzero
- [x] `JNZ addr`
- [x] `JNZ addr , x`
- [x] `JNZ LABEL`
- [x] `JNZ LABEL, x`

### JZE
jump on zero
- [x] `JZE addr`
- [x] `JZE addr , x`
- [x] `JZE LABEL`
- [x] `JZE LABEL, x`

### JUMP
unconditional jump
- [x] `JUMP addr`
- [x] `JUMP addr , x`
- [x] `JUMP LABEL`
- [x] `JUMP LABEL, x`

### JPL
jump on plus
- [x] `JPL addr`
- [x] `JPL addr , x`
- [x] `JPL LABEL`
- [x] `JPL LABEL, x`

### JOV
jump on overflow
- [x] `JOV addr`
- [x] `JOV addr , x`
- [x] `JOV LABEL`
- [x] `JOV LABEL, x`

### PUSH
- [x] `PUSH addr`
- [x] `PUSH addr , x`
- [x] `PUSH LABEL`
- [x] `PUSH LABEL, x`

### POP
- [x] `POP r`

### CALL
- [ ] `CALL addr`
- [ ] `CALL addr , x`
- [ ] `CALL LABEL`
- [ ] `CALL LABEL, x`

### RET
return from subroutine
- [x] `RET`
### SVC
- [ ] `SVC addr`
- [ ] `SVC addr , x`
- [ ] `SVC LABEL`
- [ ] `SVC LABEL, x`
### NOP
no operation

Start : 2019/06/06
