# Data Leakage
Certain information about the target only appear in the stage of training. That leads to accuracy in training but inaccurate performance in production.
1. target leakage  
    causes:  
    * chrnonical  
    * strongly depend on target
1. train-test contamination
    solutions:
    * split first
    * preprocess respectively

> Review feautres before proceeding:  
1. Should features all available at the moment of prediction
1. Should values not changed after target is determined

---

## Target Leakage
Features that will be updated after tagets are determined should not be included.

> Examples:  
1. _`Credit card expenditure`_ only turns positive after a _`card holder`_ is permitted(target)
1. _`Special treament`_ is only took place after a synmdrome is _`confirmed`_(target)

> Look closely to data:  

```python
y = target
suspicous_colmns_y = X.sus_col[y]
suspicous_colmns_n = X.sus_col[~y]

# may => 1.00, all true
(suspicous_columns_y == 0).mean()
# positive
(suspicous_columns_n == 0).mean()
```

## Train-test Contaminiation
Impute for missing values before train_test_split on the full data set. Use pipeline to run them respectively on the train and valid datasets.