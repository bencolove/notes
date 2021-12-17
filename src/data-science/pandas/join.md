# join

* concat()
* `join` on _indexes_ (syntax sugar of `merge`)
* `merge` on general columns (implementation)
* align

Method | Usage | Directions | Join
---|---|---|---
`merge` | implementaion of all _join_ methods | columns | 1) _`how`_=_inner_(default) \| _left_ \| _right_ \| _outer_ \| _cross_ <br> 2) _`on`_ \| _`left_on`_ \| _`right_on`_ : using any indexes or columns <br> 3) _`left_index`_ \| _`right_index`_: using indexes 
`join` | _merge_ on keys | columns | `merge` with: <br> _`on`_ and _`how`_
`concat` | union(append) two targets | _`axis`_=0 default appending rows <br> _`axis`_=1 appending columns | _`join`_=_outer_ default append all <br> _`join`_=_inner_ shared indexes 
`align` | shape two targets identically | _`axis`_=0, same rows(indexes) <br> _`axis`_=1, same columns | on: _outer_ \| _inner_ \| _left_ \| _right_  

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