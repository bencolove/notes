# Cross-validation
![](../img/cross_validation.png)

## What
Devide the full dataset into _`fold`_s, and run the num of _`fold`_s rounds. Within each run, pick one _`fold`_ as validation set and the rest as training set. Finally take the average error as the evaluation of the model.

## When
With small datasets

## Error of Cross-validation
```python
from sklearn.model_selection import cross_val_score

scores = -1 * cross_val_score(pipeline, X, y, cv=num_folds, scoring='neg_mean_absolute_error')

model_val_score = scores.mean()
```
