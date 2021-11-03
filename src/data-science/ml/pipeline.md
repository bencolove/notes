# Pipeline
The workflow:
1. identify _numberic_ and _categorical_ features(columns)
1. preprocess
    1. impute _numberic_ missing values
    1. impute _categorical_ missing values and order/one-hot encode 
1. define model
    1. investigate to pick up suitable model
    1. test-run to find best parameters
1. evaluate
    1. cross-validation

## Identify Numeric and Categorical Features
```python
from sklearn.model_selection import train_test_split

# drop predictors(y)
X = data.dropna(axis=0, subset=['Price'])
y = X.Price
X = X.drop(['Price'], axis=1)

# split up dataset
X_train_full, X_valid_full, y_train, y_valid = train_test_split(X, y, train_size=0.8, test_size=0.2, random_state=0)

# identify categorical feature
cat_idfier = [ X_train_full.dtype == 'object' and X_train_full.nunique() < 10 ]
categorical_cols = X_train_full.columns[cat_idfier]

# identify numberical features
num_idfier = X_train_full.dtypes in ['int64', 'float64']
numerical_cols = X_train_full.columns[num_idfier]

# picked columns
model_cols = categorical_cols + numerical_cols

# processed dataset
X_train = X_train_full[model_cols]
X_valid = X_valid_full[model_cols]
```

## Preprocess Features
1. impute missing values
1. order/one-hot encode categorical feature
```python
from sklearn.compose import ColumnTransformer
from sklearn.pipeline import Pipeline
from sklearn.impute import SimpleImputer
from sklearn.preprocess import OneHotEncoder

# preprocess numerical data
numerical_transformer = SimpleImputer(strategy='constant')

# preprocess categorical data
categorical_transformer = Pipeline(steps=[
    ('imputer', SimpleImputer(strategy='most_frequent')),
    ('onehot', OneHotEncoder(handle_unknown='ignore'))
])

# bundle up
preprocessor = ColumnTransformer(transformer=[
    ('num', numerical_transformer, numerical_cols),
    ('cat', categorical_transformer, categorical_cols)
])

```

## Define Model
```python

from sklearn.ensemble import RandomForestRegressor

model = RandomForestRegressor(n_estimators=100, random_state=0)

```

## Evaluate Pipeline
Feed dataset along down the pipeline to fit model and then predict, finally evaluate by errors:
```python
from sklearn.metrics import mean_absolute_error

# bundle preprocess and modeling as pipeline
all_pipeline = Pipeline(steps=[
    ('preprocess', preprocessor),
    ('model', model)
])

# fit(preprocess and fit)
all_pipeline.fit(X_train, y_train)

# predict
preds = all_pipeline.predict(X_valid)

# evaluate
score = mean_absolute_error(y_valid, preds)
```

```python
from sklearn.datasets import make_classification
from sklearn.linear_model import LogisticRegression
from sklearn.model_selection import train_test_split
from sklearn.pipeline import make_pipeline
from sklearn.preprocessing import StandardScaler

X, y = make_classification(random_state=31)
X_train, X_test, y_train, y_test = train_test_split(X, y, random_state=31)
pipe = make_pipeline(StandardScaler(), LogisticRegression())
# scale, train on training dataset
pipe.fit(X_train, y_train)

# scale predict on test dataset
pipe.score(X_test, y_test)

```