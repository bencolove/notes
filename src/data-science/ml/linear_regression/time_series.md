# Linear Regression With Time Series
`statsmodels.tsa.deterministic.DeterministicProcess`  
`sklearn.linear_model.LinearRegression`

1. identify a trend
    1. rolling average `df.rolling().mean()`
    2. determine `polynomial order` by looking at it
2. create time-step using the `polynomial order`
    1. `DeterministicProcess.in_sample()` create time-step values from existing datetime index
    2. `.out_of_sample(steps=how-many)` create time-step values for prediction index
3. train model using `LinearRegression.fit(X, y)` using existing time-step values
    1. regress from existing values
    5. predict using prediction index

> Purpose  
Forecasting / prediction

> Involve  
* engineer features to model based on time series (trends, sesons or cycles)
* visualize 
* forecast
* adapt machine learning

> Feature

FEATURE TYPE | DATA SOURCE | IMPLY | HOW-TO
---|---|---|---
raw datetime | raw data | | 
time-step | derived from time index | _time-dependency_ <br> correlated chronologically | `np.arange(len(df.index))`
lag | derived from existing column | _serial-dependency_ <br> correlated with previous value | `df[col].shift(1)`

> Liner regression  
for linear model of two features(variablex denoted as x)
y = w1 * x1 + w2 * x2 + bias
fed with dataset(list of (y, x1, x2)) to deduct
w1, w2 (coefficients)

## Identify a trend

### 1. compute rolling average
```python
rolling_average = series.rolling(
    window=window_width,
    min_periods=min_window_width,
    center=True,
).mean()
```

### 2. DeterministicProcess
```python
from statsmodels.tsa.deterministic import DeterministicProcess

dp = DeterministicProcess(
    index=frequency_index,
    constant=True,
    order=1,
    drop=True,
)
# polynomial order: 
# 0=1               contant 
# 1=x^1(linear)     trend
# 2=x^2(quadratic)  trend_squared
# 3=x^3(cubic)      trend_cubed

# create constant and trends columns based on the existing datetime index
X = dp.in_sample()

# create predict time-step datetime index
X_pred = dp.out_of_sample(steps=how-many)
```



## Draw line directly
`sns.regplot(x=col_name, y=col_name, data=df, scatter_kws=dict(color='RGBA'))`


## ML Time-step feature
> Create time-step(dummy) value manually  
```python
from sklearn.linear_model import LinearRegression

X = df[var_col_name]
# cleansing
X.dropna(inplace=True)

y = df[output_col_name]
# drop corresonding rows against X
y, X = y.align(X, join='inner')

model = LinearRegression()
# type(X)=pd.DataFrame (2D array)
# type(y)=pd.Series
model.fit(X, y)

y_on_x = model.predict(X)

# output coefficients(weights) and intercepts(bias)
print(f'weights={model.coef_}, bias={model.intercept_}')
```

> Create time-step values with `statsmodels`  
```python
from statsmodels.tsa.deterministic import DeterministicProcess

dp = DeterministicProcess(
    index=frequency_index,
    constant=True,
    order=1,
    drop=True,
)
# polynomial order 1=x^1(linear) 2=x^2(quadratic) 3=x^3(cubic)

# create const and trend columns for the given index 
X = dp.in_sample()

X = X[['trend']]
y = df[target_col_name]
model = LinearRegression(fit_intercept=False)
model.fit(X, y)


# use future(out-of-sample) time to predict on instead
# of time-step values
# future 30 sampling points
pred_X = dp.out_of_sample(steps=30)

pred_y = model.predict(X)
```