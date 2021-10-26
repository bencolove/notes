# Decision Tree

## What
Classifier from input variables to outputs as categories.


---

## How
`sklearn.tree.DecisionTreeRegressor`  
1. Define - pick a proper model
1. Fit
1. Predict
1. Evaluate
    * MAE (Mean Absolute Error) abs(a - mean(a))
    `sklearn.metrics.mean_absulote_error`
    * std (Standard Deviation) sqrt(mean(x))

```python
from sklearn.tree import DecisionTreeRegressor

# to ensure same results each run
model = DecisionTreeRegressor(random_state=1)

# fit
y = home_data.Price
X = home_data[['Rooms', 'Bathroom', 'Landsize', 'Lattitude', 'Longtitude']]
model.fit(X, y)

# predict
model.predict(X.head())
```

## Data Set
Split data set into two groups:
1. training
1. validation

With `sklearn.model_selection.train_test_split`  
```python
from sklearn.model_seelction import train_test_split

train_X, val_X, train_y, val_y = train_test_split(X, y, random_state = 0)

model = DecisionRegressor()

# fit with training data
model.fit(train_X, train_y)

# predict with validation data
val_pred = model.predict(val_X)

# validating with MAE
mean_absolute_error(val_y, val_pred)

```
