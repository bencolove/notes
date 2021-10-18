# Making python project self-contained
1. .py -- cython.exe --> .c (cfile)
1. .py -- cython.exe --> .c (cfile) -- gcc --> .pyd | .so (lib)
1. .py -- cythonize.exe --> .pyd | .so (lib)

```shell
$ cythonize -i main.py
# -i --inplace (place output file next to source)
```

```python
from setuptools import setup
from Cython.Build import cythonize

# fine-granular config
# extension = [
#     Extension...
# ]
setup(
    ext_module = cythonize("source.py"[, include_path=<EXTRA_PATH>])
    name = "optional"
)
```
`$ python setup.py build_ext --inplace`  