# Driverless Printing
1. PDF -> PJL (need printer support PDF print, otherwise PWG)
1. Jpeg -> Ipp

## By PJL
Constraints:
1. printer support direct `PDF` printing

## By IPP
Constraints:
1. printer supports `IPP`
1. printer supports document types `application/pdf` or `image/jpeg`
1. may need to convert pdf to pclm for `image/pwg-raster` or `application/pclm`

>1. check supported document types  
*`get-printer-attributes`*
>2. save file on disk( conversion if needed)  
>4. send file to print  
*`print-job`*


[example]: https://www.qyh.me/articles/android-ipp-print