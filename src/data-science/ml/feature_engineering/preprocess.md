# Preprocess Data
1. Standarization
    * mean removal
    * variance scaling
1. Non-linear transformation
1. Normalization
1. Encoding categorical features
1. Discretization
1. Imputation of missing values
1. Polynomial features
1. Other transformers

## Standarization
A standarized feature is standard normally distributed: Gaussian with _`ZERO`_ mean and _`UNIT`_ variance.

In practice, it can be done by:
1. X - mean(X)
1. X / stdev(X) (標準差, 方差, _`SD`_, `sqrt((X-mean(X))^2/count(X))`)

> Why?  
Algothrims like `RBF kernel`, `S`upport`V`ector`M`achines, `l1,l2` regularizers of linear models require __`ALL`__ features are zero meaned and with variance in the same order.  
Otherwise features with variance that is orders of magnitude larger than others might dominate.

>How-to
```python
from sklearn.preprocess import StanderScaler
import numpy as np

X_train = np.array([[1., -1., 2.], [2., 0., 0.], [0., 1., -1.]])
# compute mean and SD
scaler = StandardScaler().fit(X_train)
scaler.mean_
scaler.scale_
# transform in-place ( (X- scaler.mean_) / scaler.scale_ )
X_scaled = scaler.transform(X_train)

# results, by-column
X_scaled.mean(axis=0)
# >>> array([0., 0., 0.])
X_scaled.std(axis=0)
# >>> array([1., 1., 1.])
```

## Encode Categorical Feature
* Ordinal
* One-hot

```python
from sklearn.preprocess import OrdinalEncoder, OneHotEncoder
# dataset
X = [['m','us','safari'],['f','eu','chrome']]

# ordinal encoding
enc = OrdinalEncoder()
# fit compute
enc.fit(X)
# transform apply
enc.transform([['f', 'us', 'safari']])
# >>> [0., 1., 1.]

# one-hot encoding
onehot_enc = OneHotEncoder()
# fit compute
onehot_enc.fit(X)
# transform apply -> scipy.sparse.csr.csr_matrix
onehot_enc.transform([['f','us','safari'],['m','eu','safari']]).toarray()
# >>> array([[1.,0.,1.,0.,1.,0],[0.,1.,0.,1.,1.,0]])
# internals
onehot_enc.categories_
# custom categories (some value may not show so far)
genders=['f','m']
locations=['eu','us','asia']
browsers=['safari','chrome','edge']
enc = OneHotEncoder(categories=[
    genders, locations, browsers
])
# ignore missing category
OneHotEncoder(handle_unknow='ignore')
```

## you-should-see