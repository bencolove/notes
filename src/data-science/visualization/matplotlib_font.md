# 中文 matplotlib

- pandas.dataframe.plot() focuses what to plot
- matplotlib.plt fine controls how to plot

```python

import pandas as pd
from matplotlib import pyplot as plt

df = pd.read_csv(csv_file)

# filter screen and project
rows = (df['col'] == filter_value) & (df['col2'] == filter_value2)
cols = ['col1', 'col2']
result = df.loc[row, cols]
# may reset index
result.set_index('colx', inplace=True)

# plot by dataframe
result.plot(
    kind='bar', # or 'line'
    title = 'title',
    xlabel = 'xlabel',
    ylabel = 'ylabel',
    figsize=(10, 5)
)

plt.show() # or %matplotlib.pylot inline mode

```

## Chinese Characters
1. You have to import a proper font to display CJK like [Google Noto Fonts](https://fonts.google.com/noto). You may search for something like `CJK TC` and will find `Noto Sans CJK TC`.

1. download the `otf` file from Google and put in alongside 
1. instruct `pyplot` to apply the fonts by
```python

from matplotlib.font_manager import FontProperties

font = FontProperties(fname=r'NotoSansCJKtc-Medium.otf')

chart = result.plot(figsize=(10,5))
# setup wherever chinese is needed
for label in chart.get_xticklables():
    label.set_fontproperties(font)

plt.title('中文標題', fontproperties=font)
plt.xlabel('x軸名稱', fontproperties=font)
plt.ylabel('y軸名稱', fontproperties=font)
plt.legend(prop=font)


```


[1]: https://www.learncodewithmike.com/2021/03/pandas-and-matplotlib.html

