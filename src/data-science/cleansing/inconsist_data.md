# Inconsist Data Entry

## Categorical Data
1. unique and sort a Series
1. lower
1. strip
1. similar content

```python
countries = df['Country'].unique()
print(countries.sort())

df['Country'] = df['Country'].str.lower()
df['Country'] = df['Country'].str.strip()

import fuzzywuzzy
from fuzzywuzzy import process
# to group 'southkorea' and 'south korea'

matches = fuzzywuzzy.process.extract('south korea', countries, limit=10, scorer=fuzzywuzzy.fuzz.token_sort_ration)

close_matches = [m[0] for m in matches if m[1] >= min_ratio]

row_index = df[column].isin(close_matches)

df.loc[row_idnex, column] = string_to_match

```
