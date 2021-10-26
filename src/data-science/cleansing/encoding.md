# Test Encoding
```python
import chardet

with open(file_path, 'rb') as rawdata:
    result = chardet.detect(rawdata.read(10000))

print(result)

>> {'encoding': 'windiows-1252', 'confidence': 0.73, 'language': ''}

pd.read_csv(file_path, encoding='windows-1252')
```