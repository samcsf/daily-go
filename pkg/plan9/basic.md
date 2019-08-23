# Plan9 基础

## 数据搬运

```s
MOVB $1    , DI // 1 byte
MOVW $0x10 , BX // 2 byte
MOVD $1    , DX // 4 byte
MOVQ $-10  , AX // 8 byte
// ^ 最后一位决定长度
```

## 常见计算

```s
ADDQ  // +
SUBQ  // -
IMULQ // *
```

## 跳转

```s
JMP addr
JMP label
JMP 2(PC)  // 往前跳2行
JMP -2(PC) // 往后跳2行

JNZ target // 如果zero flag被设定就跳转
```

## 变量声明

```s
DATA symbol+offset(SB)/width, value // 定义变量
GLOBL var(SB), flag, value          // 全局变量
//       ^ 如果变量名字后面加"<>"表示谨当前文件有效
```

## 函数声明

```s
TEXT pkg·funcName(SB), NOSPLIT, $n-n
```

## 栈结构

```
-----------------                                           
current func arg0                                           
----------------- <----------- FP(pseudo FP)                
caller ret addr                                            
+---------------+                                           
| caller BP(*)  |                                           
----------------- <----------- SP(pseudo SP，实际上是当前栈帧的 BP 位置)
|   Local Var0  |                                           
-----------------                                           
|   Local Var1  |                                           
-----------------                                           
|   Local Var2  |                                           
-----------------                                           
|   ........    |                                           
-----------------                                           
|   Local VarN  |                                           
-----------------                                           
|               |                                           
|               |                                           
|  temporarily  |                                           
|  unused space |                                           
|               |                                           
|               |                                           
-----------------                                           
|  call retn    |                                           
-----------------                                           
|  call ret(n-1)|                                           
-----------------                                           
|  ..........   |                                           
-----------------                                           
|  call ret1    |                                           
-----------------                                           
|  call argn    |                                           
-----------------                                           
|   .....       |                                           
-----------------                                           
|  call arg3    |                                           
-----------------                                           
|  call arg2    |                                           
|---------------|                                           
|  call arg1    |                                           
-----------------   <------------  hardware SP 位置           
| return addr   |                                           
+---------------+       

```

## 参考链接

- [plan9 汇编入门](http://xargin.com/plan9-assembly/)
- [golang 汇编](https://lrita.github.io/2017/12/12/golang-asm/)
- [汇编 is so easy](https://github.com/cch123/asmshare/blob/master/layout.md)
