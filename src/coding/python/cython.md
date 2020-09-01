# Extending Python with C/C++
1. the official way of writing extension/module
2. using Cython to compose extension
3. install extension with setup tools
---
Benchmark before diveing further:  
To mimic itemgetter from itertools:  
```sh
%%cython --annotate

cpdef list cython_getter(list input_list, int pos):
    cdef list ret = []
    for i in range(len(input_list)):
        ret.append(input_list[i][pos])
    return ret
```

```sh
def comprehension(input_list, pos):
    return [ row[0] for row in input_list ]

def basic(input_list, pos):
    ret = []
    for i in range(len(input_list)):
        ret.append(input_list[i][pos])
    return ret

from operator import itemgetter
def map_get(input_list, getter):
    return list(map(getter, input_list))

def compreh_get(input_list, getter):
    return [ getter(row) for row in input_list ]
```
```sh
data = [ [x+y,] for x in range(100) for y in range(100) ]

%timeit basic(data,0)

getter = itemgetter(0)
%timeit comprehension(data,getter)
%timeit map_get(data,getter)

%timeit cython_getter(data, 0)
```
output:  
```sh
711 µs ± 477 ns per loop (mean ± std. dev. of 7 runs, 1000 loops each)
283 µs ± 4.69 µs per loop (mean ± std. dev. of 7 runs, 1000 loops each)
343 µs ± 315 ns per loop (mean ± std. dev. of 7 runs, 1000 loops each)
66.3 µs ± 310 ns per loop (mean ± std. dev. of 7 runs, 10000 loops each)
```

---
## Good old C/C++ for extension
[extension with C/C++][extending_python]  

Write customer extension with pure C/C++, pay close attention to CTypes and Python types as well as how to set and indicate exception state.

---

## Convinent way with Cython
[extension with Cython][cython]

### Install Cython
`$ pip install Cython`

### Build with Setup Tools
Process:  
1. ext.pyx file of extension logics,
1. ext_setup.py file for compiling  
1. compile
1. install

The compile process will do:
1. compile _ext.pyx_ to _ext.c_ (C code)
1. compile _ext.c_ to importable _ext.so_ (linux) or _ext.pyd_ (windows)

#### prepare:  
_ext.pyx_
```python
cpdef list cython_getter(list input_list, int pos):
    cdef list ret = []
    for i in range(len(input_list)):
        ret.append(input_list[i][pos])
    return ret
```
*ext_setup.py*
```python
from setuptools import setup
from Cython.Build import cythonize

setup(
    name='Custom Extension',
    ext_modules=cythonize("ext.pyx"),
    zip_safe=False,
)
```

#### run:  
`$ python ext_setup.py build_ext --inplace`

which ext can be imported as:  
_main.py_
```python
from ext import cython_getter

data = [ [(x*10+y),]  for x in range(10) for y in range(10) ]

print(cython_getter(data, 0))
```
output should look like:
```sh
[0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10 ...
```
#### install:
`$ python ext_setup.py install`  
While the above command is to create an **egg** file for distribution, it will **NOT** work with `cimport` for _pyd_ files. Therefore, argument `zip_safe=False` is included in `setup()` function. 


### Run in Jupyter Notebook
[How to run cython in Jupyter][cython_jupyter]

1. Install Cython in Jupyter
1. Load extension  
`%load_ext cython`
1. Run directly with cell magic `%%Cython`  
```sh
%%Cython

cpdef list cython_getter(list input_list, int pos):
    cdef list ret = []
    for i in range(len(input_list)):
        ret.append(input_list[i][pos])
    return ret
```
1. Call it directly
```sh
cython_getter(data, 0)
```



[extending_python]: https://docs.python.org/3/extending/extending.html
[cython]: https://cython.readthedocs.io/en/latest/index.html
[example]: https://towardsdatascience.com/use-cython-to-get-more-than-30x-speedup-on-your-python-code-f6cb337919b6
[cython_build_extension]: https://cython.readthedocs.io/en/latest/src/quickstart/build.html#building-a-cython-module-using-setuptools
[cython_jupyter]: https://cython.readthedocs.io/en/latest/src/quickstart/build.html#using-the-jupyter-notebook
[cython_windows]: https://github.com/cython/cython/wiki/CythonExtensionsOnWindows