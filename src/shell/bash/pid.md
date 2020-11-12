# Display Process ID
1. `$$` => current ProcessID
2. `$PPID` => parent ProcessID
3. `pstree -p` => tree of ps
4. `$!` => last backgrounded process

>Trick

`echo $PPID; echo $$; pstree -p`
To check whether it is right or not.