# rolling window

A moving average on a series
```python
DataFrame.rolling(
    window=window_width,
    center=True,
    min_periods=min_start
).mean()
```