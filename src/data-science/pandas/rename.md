# rename

> rename index or column names  
DataFrame.rename(
    index=mapper,
    columns=mapper
)


> rename axis names  
DataFrame.rename_axis(row_name, axis='rows')
DataFrame.rename_axis(column_name, axis='columns')

Or more often, `set_index()`