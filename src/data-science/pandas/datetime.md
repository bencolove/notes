# datetime
1. values of column dtype=datetime64[ns]
1. index
    1. DatetimeIndex
    1. PeriodIndex
    1. RangeIndex

## str and datetime64[ns]

> str => datetime  
`pd.to_datetime(series, format='%Y-%m-%d')`

> datetime => str  
`series.astype(str)`

