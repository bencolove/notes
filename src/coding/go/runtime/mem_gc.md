# GC
Types:
1. manual: C, C++, Rust
1. GC: Python, Ruby, Java, Go
1. Auto-refcount: Object-C

GC METHOD | HOW-TO
---|---
Mark-Sweep | `Mark` from root <br> Iterate on heap to `Sweep`
Tricolor Marking | Objects are marked with color: <br> * White - weak ref <br> * Black - active with no out-ref <br> * Grey - active with out-ref to white ones  