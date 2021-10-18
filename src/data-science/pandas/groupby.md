# Groupby

> DataFrame.groupby(column_or_columns)  
This results in a `groupby` instance

`DataFrame.groupby([str_columns])`  
Slice the df based on the columns.

## Recipts

>select min(price)
GROUP BY points  

>df.groupby('points').price.min()

---

> select count(1) from reviews  
GROUP BY tweets

> reviews.groupby('tweets').size() or  
reviews.groupby('tweents').tweets.count()  

--- 




df.groupby([columns]).apply(lambda_group_df)

1. Select the first wine from each winery

`reviews.groupby('winery').apply(lambda df: df.title.iloc[0])`


2. Pick the best wine(by points) by country and province

`reviews.groupby(['country', 'province']).apply(lambda df: df.loc[df.points.idxmax()])`

3. Summary on slices

`reviews.groupby(['country']).price.agg([len, min, max])`

## MultiIndex
`country_reviews = reviews.groupby(['country', 'province']).description.agg([len])`  
will result in 

`type(country_reviews.index)`  
`pandas.core.indexes.multi.MultiIndex`  

`country_reviews.reset_index()`  
back to single indexing

[advanced_indexing]: https://pandas.pydata.org/pandas-docs/stable/user_guide/advanced.html