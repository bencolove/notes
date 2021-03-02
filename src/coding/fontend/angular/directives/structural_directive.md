# Directives
* Structural -- insert, remove, modify elements
1. `*ngFor`
1. `ngSwitch`
1. `*ngIf`
* Component
* Attribute -- modify elements' attributes


Structural Directives | Angular | Vue
---|---|---
LOOP | `*ngFor` | `v-for`
SWITCH | `[ngSwitch]` <br> `*ngSwitchCase` <br> `*ngSwitchDefault` | X
CONDITION | `*ngIf` | `v-if` <br> `v-else`

>Asterisk(*) simply syntactic sugar as  

`<div *ngIf="hero" class="name">{{hero.name}}</div>` is translated to:
```html
<ng-template [ngIf]="hero">
    <div class="name">{{ hero.name }}</div>
</ng-template>
```