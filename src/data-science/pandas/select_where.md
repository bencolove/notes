# Projection(`select`) and Filter(`where`)

## select columns
* `df[sera]` => pd.Series
* `df[[sera, serb]]` => pd.DataFrame
* `df[[sera]]` => pd.DataFrame

---

## where
Boolean_Series = df.sera == va
* `df[sera == va]` => pd.DataFrame

> select * from reviews where points=max(points)

> reviews.loc[reviews.points.idxmax()]

---

## `loc` or `iloc`
* ret_df = df.loc[row_selection, col_selection] 

row_selection | col_selection | type(ret_df)
---|---|---
scalar | scalar | scalar
collection | scalar | pd.Series
scalar | collection | pd.Series
collection | collection | pd.DataFrame

