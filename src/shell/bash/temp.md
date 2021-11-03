# TEMP files

## `/tmp` folder
'''shell
tdir=$(mktmp -d)

tfile=$(mktmp -p $tdir)

# clean up
rm $tfile
rm -R $tdir

'''