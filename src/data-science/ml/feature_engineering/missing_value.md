# 
## Missing Values
Reasons:
1. not existed
1. existed but no records

## Three Approaches
1. Drop column (dropna axis=1)
1. Imputation
    * 0
    * mean
    * next_cell_found
    * regerssion
1. Imputation with indicator

### Drop
> drop columns  
```python
# sum() counts number of True 
not_null_cols = X_train.columns[X_train.isnull().sum() == 0]
not_null_cols = X_train.columns[X_train.notnull().all()]
cols_with_missing = X_train.columns[X_train.isnull().any()]
# select/project
reduced_X_train = X_train[not_null_cols]

# OR
cols_with_missing = [ col for col in X_train.columns if X_train[col].isnull().any() ]
reduced_X_train = X_train.drop(cols_with_missing, axis=1)
```

### Impute with `mean`
`sklearn.impute.SimpleImputer`  
```python
from sklearn.impute import SimpleImputer

mean_imputer = SimpleImputer()
# return np.ndarray
imputed_X_train = pd.DataFrame(mean_imputer.fit_transform(X_train), columns=X_train.columns)
# return np.ndarray
imputed_y_train = pd.DataFrame(mean_imputer.transform(X_valid), columns=X_valid.columns)

```