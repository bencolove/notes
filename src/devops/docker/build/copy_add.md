# COPY and ADD

`COPY --chown owner:group  src1 src2 dest`  
`ADD src_or_URL dest`

* unzip into `dest` given `gzip`,`tar` files
* if `dest` ends with `/` copies content into this `folder`
* otherwise copies all content into this `file`
* `COPY` works the `go` way
    * `src` is folder, copies contents excluding the `src` folder itself
    * `src` is pattern with `*` globing, copies matched files and `flattened`