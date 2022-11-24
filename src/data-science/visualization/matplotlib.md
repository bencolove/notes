# matplotlib.pyplot
* inline mode
* global config
* subplots

## Display Mode
- `%matplotlib notebook`: interactive mode
- `%matplotlib inline`: static mode

## Global config
```python
from warnings import simplefilter
# ignore wanings in cell output
simplefilter("ignore")

# matploglib defaults
plt.style.use('seaborn-whithegrid')
plt.rc('figure', autolayout=True, figsize=(11, 5))
plt.rc('axes',
    labelweight='bold',
    labelsize='large',
    titleweight='bold',
    titlesize=14,
    titlepad=10,
)
# linewidth
# color
# marker
```

## Subplots

### Pick-one-then-plot
```python
# select current subplot
plt.subplot(231) # 2 x 3 subplots and pick no.1 as current
plt.title...
sns.lineplot...

plt.subplot(232) # 2 x 3 subplots and current no.2
plt.title...
sns.barplot...
```

### Axes-oriented
```python
# create 2 x 1 subplots shared x aixs
fig, (ax1, ax2) = plt.subplots(2, 1, figsize=(11, 5.5), sharex=True)

# plot to the axis
ax1.set_title...
ax1.plt...

ax2.set_title...
# specify which axis to plot onto
sns.barplot(..., ax=ax2)
```

## Save Image
```python
fig = plt.figure()

fig.savefig(file_path)

from IPython.display import Image
Image(file_path) # to read it back

```