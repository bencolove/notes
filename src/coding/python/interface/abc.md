# `abc`

```python
class abc.ABC(metaclass=abc.ABCMeta):
    pass
```
So, `abc.ABC` and `abc.ABCMeta` are effectively same as
```python
class WithMeta(metaclass=abc.ABCMeta):
    pass


class WithABC(abc.ABC):
    pass
```