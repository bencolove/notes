# JSONPath
[jsonpath-expr][jsonpath-expr]  
[jsonpath-eval][jsonpath-eval]  

---
`JSONPath` | Description
---|---
`$` | the root object/element
`@` | the current object/element
`. or []` | child operator
`..` | recursive descent. JSONPath borrows this syntax from E4X.
`*` | wildcard. All objects/elements regardless their names.
`[]` | subscript operator. XPath uses it to iterate over element collections and for predicates. In Javascript and JSON it is the native array operator.
`[,]` | Union operator in XPath results in a combination of node sets. JSONPath allows alternate names or array indices as a set.
`[start:end:step]` | array slice operator borrowed from ES4.
`?()` | applies a filter (script) expression.
`()` | script expression, using the underlying script engine.


---
JsonPath | Result
---|---
`$.store.book[*].author` | The authors of all books
`$..author` | All authors
`$.store.*` | All things, both books and bicycles
`$.store..price` | The price of everything
`$..book[2]` | The third book
`$..book[-2]` | The second to last book
`$..book[0,1]` | The first two books
`$..book[:2]` | All books from index 0 (inclusive) until index 2 (exclusive)
`$..book[1:2]` | All books from index 1 (inclusive) until index 2 (exclusive)
`$..book[-2:]` | Last two books
`$..book[2:]` | Book number two from tail
`$..book[?(@.isbn)]` | All books with an ISBN number
`$.store.book[?(@.price < 10)]` | All books in store cheaper than 10
`$..book[?(@.price <= $['expensive'])]` | All books in store that are not "expensive"
`$..book[?(@.author =~ /.*REES/i)]` | All books matching regex (ignore case)
`$..*` | Give me every thing
`$..book.length()` | The number of books

[jsonpath-expr]: https://goessner.net/articles/JsonPath/index.html#e2
[jsonpath-eval]: https://jsonpath.com/