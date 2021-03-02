# Dependency Injection
    Dependencies in angular are shareable services or objects needed among multiple classes.

>1. create injectable

`ng generate service dir/MyService`

--_`dir/MyService`_--
```typescript
import { Injectable } from '@angular/core';

@Injectable({
    // means the service is visible app-wide
    providedIn: 'root',
})
export class MyService {
    constructor() {}
}
```

>2. injecting (consuming) service into Component

```typescript
// DI by angular
constructor(myService: MyService)
```

>3. injecting service into other services
Same as used in components:
1. `import`
2. declare in `constructor`


