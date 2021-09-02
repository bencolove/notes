# Back-tracking Problem

Pseudo code
```python
def is_valid(state) -> bool:
    pass

def get_candidates(state) -> List[Candidate]:
    pass

def search(state, solutions) -> bool:
    if is_valid(state):
        solutions.add(state)
    
    for candidate in get_candidates(state):
        set_candidate(state, candidate)
        if search(state, solutions) != True:
            unset_candidate(state, candidate)
        else:
            return True
    return False

```