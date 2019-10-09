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
