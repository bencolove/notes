# Categorical Data

Three approches (cardinality num of unique values):
1. drop
1. ordinal encoding (either order matters or low cardinality)
1. one-hot encoding (indicators of all possible values)

## Investigate
By looking at the cardinality of features:
```python 
object_cols = df.columns[df.dtypes == 'object'] 
print(f'object cols: {object_cols}')

cardinalities = {
        col: df[col].nunique()
        for col in object_cols
    }
card_thres = 10

rule_out_cols = [ p[0] for p in cardinalities.items() if p[1] >= card_thres ]
print(f'cols with large(>=10) cardinalities: {rule_out_cols}`)

```

## [How-To](preprocess.md#encode-categorical-feature)

## Tricks

> find cols with missing value  
`df.columns[df.isnull().any()]`

> find categorical features 
`ser = df.columns[df.dtypes == 'object']`  
`df = df.select_dtypes(include=['object'])`

> find low cardinality columns  
```python
low_cardinality_cols = [ col for col in df.columns
    if df[col].nunique() < 10 and df[col].dtype == 'object'
]
```

> check categorical feature cardinality  
`good_cols = [ col for col in df.columns if set(X_valid[col].unique()).issubset(X_train[col].unique())]`  
`bad_cols = list(set(df.columns) - set(good_cols))`    

> set ordinal values  
`df[categorical_cols] = ordinal_enc.fit_transform(df[categorical_cols])`  

> set one-hot values  
```python
# sparse=False makes returned ndarray
onehot_enc = OnehotEncoder(handle_unknown='ignore', sparse=False)
OH_cols_train = pd.DataFrame(onehot_enc.fit_transform(df[oh_cols], index=df.index))
# replace
# remove categorical columns
# add onehot columns
numeric_X_train = X_train.drop(categorical_cols, axis=1)
# concat by columns
preprocessed_X_train = pd.concat([
    numeric_X_train, OH_cols_train
], axis=1)
# one go
onehot_X_train = pd.concat([
    X_train.drop(categorical_cols, axis=1),
    pd.DataFrame(onehot_enc.transfomr(X_train[low_cardinality_cols]), 
        index=X_train.index)
], axis=1)
```