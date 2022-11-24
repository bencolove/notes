# Image Preprocessing
- load
- resize
- color

> Read/Load

```python
skimage.data.imread(f)
```

## Resize
```python
from skimage import transform

images28 = [ transform.resize(img, (28, 28)) for img in imgs ]
```

## Grayscale
```python
from skimage.color import rgb2gray

# Convert `images28` to an array
images28 = np.array(images28)

# Convert `images28` to grayscale
images28 = rgb2gray(images28)
```

