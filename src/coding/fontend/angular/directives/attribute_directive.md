# Attribulte Directives
    Attibute directives are used to modify attributes of elements.

---
IMPORTANT:  
Difference between: `[prop]="color"` and `prop="color"`

* adorned with brackets([]), it is first evaluated as a javascript expression and then bind the result to the name.
* **`NO`** brackets, it is bound as-it like string literal to the name.
---

Inorder to support:
```html
<p appHighlight>it is highlighted</p>
```

>1. scaffold  

`ng generate directive hightlight`

>2. import into component

```typescript
import { Directive } from '@angular/core';

@Directive({
    // brackets needed
    selector: '[appHighlight]'
})
export class HighlightDirective {
    constructor() {}
}
```

>3. refer to element context

```typescript
import { Directive, ElementRef } from '@angular/core';

@Directive({
  selector: '[appHighlight]'
})
export class HighlightDirective {
    // modify elemnet via ElementRef
    constructor(el: ElementRef) {
       el.nativeElement.style.backgroundColor = 'yellow';
    }
} 
```

>4. apply the directive

```html
<p appHighlight> oh yeah </p>
```

>5. listen on element's changes

```typescript
import { Directive, ElementRef, HostListener } from '@angular/core';

@HostListener('mouseenter') onMouseEnter() {
    this.highlight('yellow');
  }

  @HostListener('mouseleave') onMouseLeave() {
    this.highlight(null);
  }

  private highlight(color: string) {
    this.el.nativeElement.style.backgroundColor = color;
  }

```

>5. pass values into directive like prop

```typescript
/* template */
<p [appHighlight]="color" extraProp="'orange'" />

@Input('appHighlight') highlightColor: string;
@Input extraProp: string;
```