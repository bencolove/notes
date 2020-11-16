# Contract Interface (Virtual Subclass)
>Pros:  
Subclass does not need to declare subclass of the interface(class) but still qualified as an instace of a subclass implementing the interface.

>Cons:
The interface(inherited parent class) will not be shown from `__mro__` as opposited to proper inheritance.  

```python
# class Contract(metaclass=abc.ABCMeta): or
class Contract(abc.ABC):
    @classmethod
    def __subclasscheck__(cls, subclass):
        # called by isinstance(obj, interface)
        # default implementation super().__subclasshook__(subclass)
        return cls.__subclasshook__(subclass)

    @classmethod
    def __subclasshook__(cls, subclass):
        """check whether cls has contracted methods
           this method will only be called once per class, since 
           results are cached
        """
        # called by issubclass(type(impl), interface)
        required_attrs = ['attr_1',]
        required_fns = ['fn_1', 'fn_2']

        for P in cls.__mro__:
            for attr in required_attrs:
                if attr not in P.__dict__:
                    return NotImplemented
            for fn in required_fns:
                if fn not in P.__dict__ or P.__dict__[fn] is None:
                        return NotImplemented
        return True
```


Examples:
```python
import abc

class Lengthy(abc.ABC):
    
    @classmethod
    def __subclasshook__(cls, subclass):
        print('__subclasshook__')
        # look for 'length()' in  __dict__ from each class in subclass.__mro__
        if any( 'length' in P.__dict__ for P in subclass.__mro__ ):
            return True
        return NotImplemented

    # NOT necessary but clear with what method should be implemented if inheriting it
    @abc.abstractmethod
    def length(self):
        raise NotImplemented
        
        
class NoLengthy:
    pass

class WithLengthy:
    def length(self):
        return 0
        

n = NoLengthy()
w = WithLengthy()

# False
print(f'isinstance(n, Lengthy)={isinstance(n, Lengthy)}')
# False
print(f'issubclass(NoLengthy, Lengthy)={issubclass(NoLengthy, Lengthy)}')

# True
print(f'isinstance(w, Lengthy)={isinstance(w, Lengthy)}')
# True
print(f'issubclass(WithLengthy, Lengthy)={issubclass(WithLengthy, Lengthy)}')
```

It works pretty much the same way as `GO` does with `interface` and `duck typing`. It only checks whether the subclass has the required set of attributes/methods.

One more advantages of this manner is that it will be possible to make existing legacy class to **`VIRTUALLY`** subclass an class introduced later behind.