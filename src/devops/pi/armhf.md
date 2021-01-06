# Floating-point Arch
The diff:  
1. `fpu` floating-point processing unit
2. whether push data to `fpu` by using dedicated registers

`gcc -mfloat-abi` parameters | meaning
---|---
soft | donot use `fpu` even equipped
armel (softfp) | use `fpu` with normal registers
armhf (hard) | use `fpu` with dedicated registers     

## Check
```sh
$ readelf -A /proc/self/exe | grep Tag_ABI_VFP_args

Tag_ABI_VFP_args: VFP registers
# VFP Vector Floating-Poiont (hard one armhf)
```