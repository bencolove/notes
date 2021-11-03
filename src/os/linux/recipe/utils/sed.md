# `sed`
[example](https://www.linuxprobe.com/sed-find-replace.html)

>Synopsis:  
>sed -e PATTERN [-e PATTERN] -|FILE
>* -i[SUFFIX] in-place instead of STDOUT, and backup file with SUFFIX
>* -n disable output whole content to STDOUT
>
>PATTERN := RANGE s|SEARCH|REPLACE|MODE  
>RANGE := n | start,end  
>SEARCH := regex  
>    * \b \S \\. \\? \\{ \\} ^ & ...  
>    * [0-9]{2}  
>
>MODE :=   
>* n: replace [nth] occurance only
>* [n]g: replace occurances from [nth] to last 
>* g: replace all occurances
>* p: only print matched lines
>* I: case insenstive

Usecases:
1. only output replacement lines  
`sed -n 's|search|replace|p'`  
mode _p_ : print the affected lines plus by default sed output source. Therefore affected lines will be shown twice   
-n : disable output source 