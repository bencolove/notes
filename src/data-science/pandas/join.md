# join

* concat()
* join()
* merge()

# concat() union

pd.concat([canada_youtube, usa_youtube])

# join() on common index
```python
# set index first
left = canada_yt.set_index(['title', 'date'])
right = usa_yt.set_index(['title', 'date'])

# then join on indexes
left.join(right, lsuffix='_CAN', rsuffix='_UK')
```
```SQL
select 
    left.* AS left.*_CAN, 
    right.* AS right.*_UK 
left join right 
on left.title=right.title and left.date=right.date
```