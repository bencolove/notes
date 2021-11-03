# CommonParameters
CP | Meaning | Usage
---|---|---
-OutVariable -ov | save result object into/append a variable | `ps winword -ov +out` | append result to `out` typed object list
-OutBuffer -ob | buffer results up to (ob+1) before sending downstream in a pipeline| 
-PipelineVariable -pv | save current result into var in a pipeline which can be referred to in downstream | `1..10 | % { $_ + 1 } -PV l | % { 1..2 } -PV r | % { "$l + $r = " + ($l+$r) }`