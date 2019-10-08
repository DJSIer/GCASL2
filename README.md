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
- [ ] `ADDL r1,r2`
- [ ] `ADDL r,addr`
- [ ] `ADDL r,LABEL`
- [ ] `ADDL r,addr,x`

 ### SUBA
 subtract arithmetic
- [ ] `SUBA r1,r2`
- [ ] `SUBA r,addr`
- [ ] `SUBA r,LABEL`
- [ ] `SUBA r,addr,x`

 ### SUBL
 substract logical
- [ ] `SUBL r1,r2`
- [ ] `SUBL r,addr`
- [ ] `SUBL r,LABEL`
- [ ] `SUBL r,addr,x`

 ### AND
- [ ] `AND r1,r2`
- [ ] `AND r,addr`
- [ ] `AND r,LABEL`
- [ ] `AND r,addr,x`

 ### OR
- [ ] `OR r1,r2`
- [ ] `OR r,addr`
- [ ] `OR r,LABEL`
- [ ] `OR r,addr,x`

 ### XOR
 exclusive or
- [ ] `XOR r1,r2`
- [ ] `XOR r,addr`
- [ ] `XOR r,LABEL`
- [ ] `XOR r,addr,x`

### CPA
compare arithmetic
- [ ] `CPA r1,r2`
- [ ] `CPA r,addr`
- [ ] `CPA r,LABEL`
- [ ] `CPA r,addr,x`

### CPL
compare logical
- [ ] `CPL r1,r2`
- [ ] `CPL r,addr`
- [ ] `CPL r,LABEL`
- [ ] `CPL r,addr,x`

### SLA
shift left arithmetic
- [ ] `SLA r,addr`
- [ ] `SLA r,LABEL`
- [ ] `SLA r,addr,x`

### SRA
shift right arithmetic
- [ ] `SRA r,addr`
- [ ] `SRA r,LABEL`
- [ ] `SRA r,addr,x`

### SLL
shift left logical
- [ ] `SLL r,addr`
- [ ] `SLL r,LABEL`
- [ ] `SLL r,addr,x`

### SRL
shift right logical
- [ ] `SRL r,addr`
- [ ] `SRL r,LABEL`
- [ ] `SRL r,addr,x`

### JMI
jump on minus
- [ ] `JMI addr`
- [ ] `JMI addr , x`
- [ ] `JMI LABEL`
- [ ] `JMI LABEL, x`

### JNZ
jump on nonzero
- [ ] `JNZ addr`
- [ ] `JNZ addr , x`
- [ ] `JNZ LABEL`
- [ ] `JNZ LABEL, x`

### JZE
jump on zero
- [ ] `JZE addr`
- [ ] `JZE addr , x`
- [ ] `JZE LABEL`
- [ ] `JZE LABEL, x`

### JUMP
unconditional jump
- [ ] `JUMP addr`
- [ ] `JUMP addr , x`
- [ ] `JUMP LABEL`
- [ ] `JUMP LABEL, x`

### JPL
jump on plus
- [ ] `JPL addr`
- [ ] `JPL addr , x`
- [ ] `JPL LABEL`
- [ ] `JPL LABEL, x`

### JOV
jump on overflow
- [ ] `JOV addr`
- [ ] `JOV addr , x`
- [ ] `JOV LABEL`
- [ ] `JOV LABEL, x`

### PUSH
- [ ] `PUSH addr`
- [ ] `PUSH addr , x`
- [ ] `PUSH LABEL`
- [ ] `PUSH LABEL, x`

### POP
- [ ] `POP r`

### CALL
- [ ] `CALL addr`
- [ ] `CALL addr , x`
- [ ] `CALL LABEL`
- [ ] `CALL LABEL, x`

### RET
return from subroutine
### SVC
- [ ] `SVC addr`
- [ ] `SVC addr , x`
- [ ] `SVC LABEL`
- [ ] `SVC LABEL, x`
### NOP
no operation

Start : 2019/06/06
