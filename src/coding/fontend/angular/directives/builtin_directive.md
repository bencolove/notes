# Built-in Directives
* `NgClass` -- modify on CSS classes
* `NgStyle` -- modify HTML styles
* `NgModel` -- two-way data binding

>Without `ngClass` (Single Class)
```html
<div [class.class1]="shouldAddClassOne">
```


>`NgClass` (Multiple Classes)

```typescript
/* template */
<div [ngClass]="currentClasses" />

/* Component */
currentClasses: {};

setCurrentClasses() {
    // className: boolean indicating add/remove the class
    this.currentClasses = {
        class1: shouldAddClassOne,
        class2: shouldRemoveClassTwo
    };
}
```

>Without `ngStyle` (Single Style)

```typescript
<div [style.font-size]="shouldBig ? 'x-large' : 'smaller' " />
```

>`ngStyle` (Multiple Styles)

```typescript
/* template */
<div [ngStyle]="currentStyles" />

/* component */
currentStyles: {};

setCurrentStyles() {
    this.currentStyles = {
        'font-style': sholdBeItalic ? 'italic' : 'normal',
        'font-weight': sholdBold ? 'bold' : 'normal',
        'font-size': big ? '24px': '12px'
    };
}
```