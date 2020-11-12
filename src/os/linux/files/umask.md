# How `umask` works
[how-work][example]

> By default

File permission := 666  
Folder permission := 777

> Wanted:

File := 644  
Folder := 755

> Via `umask`

666 - 644 = 022  
777 - 755 = 022


[example]: https://www.howtoforge.com/linux-umask/