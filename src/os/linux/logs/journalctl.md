# `journalctl`

* boot message
* time ranges
* by service(unit)
* follow(tail)
* output format

>Boot message
`journalctl -b -l`

>Time ranges
* after point-of-time  
`--since "2006-01-02 15:03:04"`
* before point-of-time  
`--until "2006-01-02 15:03:04"`
* last period-of-time  
`--since "1 hour ago"`  
`--since "2 days ago"`

>By service(unit)  
`-u service-name`

>Follow
`-f`

>Tail last n entries  
`-n 50`

>Reverse chronologically  
`-r`

>Outpu as JSON  
`-o json-pretty`