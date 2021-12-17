# XGBoost
_`XGBoost`_ for eXtreme Gradint Boosting.
Better work with DataFrame(TSV, Tab Sparated Values). **NOT** exotic types like images and videos.


Ensemble methods:
1. _`RandomForestRegressor`_
1. _`xgboost.XGBRegressor`

Train cycles:
> do  
1. initial native model
1. train and predict then errors
1. find new model (_`gradient descent`_ on the loss function)
1. add new model to ensemble
> until  
* until errors not improved

```python
import pandas as pd
from sklearn.model_selection import train_test_split

# Load
data

# predictor columns
pred_cols = ['Room' ...]
X = data[pred_cols]
# target
y = data.Price

# split dataset
X_train, X_valid, y_train, y_valid = train_test_split(X, y)

from xgboost import XGBRegressor

model = XGBRegressor()
model.fit(X_train, y_train)

from sklearn.metrics import mean_absolute_error

preds = model.predict(X_valid)
mean_absolute_error(preds, y_valid)
```

## Parameter Tuning
Important Parameter | Impact
---|---
`n_estimators` | iterative cycles (100, 1000)
`early_stopping_rounds` | stop after n straight rounds of deteriorating validation scores
`eval_set` | `(Train, Valid)`
`learning_rate` | weights for added model
`n_jobs` | parallel level

> Recipes:  
1. [Big]n_estimators + early_stopping 
1. [Big]n_estimators + [Small]learning_rate