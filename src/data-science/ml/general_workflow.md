# General Workflow
Machine learning is an iterative process with which you regressively go through:
1. load data
    1. fix index
    1. locate predictor column(y)
    1. remove rows where (isnull(y))
    1. X remove y
1. feature engineering 
    1. missing value (impute)
    1. categorical data (impute+ordinal/onehot)
    1. numerical data (impute+standarize)
1. LOOP: parameterized pipeline on
    * model
    * model parameters
    * cross-validation
1. find the configue with least error
