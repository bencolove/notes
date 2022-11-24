# Load/Save Images
- `scikit-image`

## Loads

### `scikit-image`
```python
from pathlib import Path
def load_data(data_dir):

    subdirs = [d for d in os.listdir(data_dir)
    if os.path.isdir(os.path.join(data_dir, d))
    ]


    subdirs = [d for d in Path(data_dir).itedir() if d.is_dir()]

    files = [for f in subdirs.itedir() if f.is_file() and f.suffix == '.ppm' ]

    for f in files:
        # load the image file
        skimage.data.imread(f)
```

## Plot
```python
import matplotlib.pyplot as plt
# 1row by 4col images and now the first one
plt.subplot(1, 4, 1)

plt.title('title')
plt.axis('off')
# show the image
plt.imshow(image_data)
plt.subplots_adjust(wspace = 0.5)

plot.show()
```