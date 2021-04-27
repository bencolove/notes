# Flux and Redux Patterns
[[try to explain][flux_redux]]

It originated from Facebook and aimed at state management issue for SPA application.

---

>Background:

In a typical SPA application, multiple pages share a common set of data( eg. cart contents for product list page and cart overview page ).

>Issue:

Image after an user viewed a product and like to buy it, he clicked the 'add to cart' button. How awkward would it be when he later in the view cart contents page found the item was not there !!   

In order to maintain their consistency, it needs a way to _`synchronize`_ the data among mutiple pages.

>Solutions:
1. flux pattern
1. redux pattern (based on flux and applies three more principles)

---

## Flux Pattern
![flux pattern](https://www.dotnetcurry.com/images/reactjs/redux/flux-pattern.png)

---

## Redux Pattern
![redux pattern](https://www.dotnetcurry.com/images/reactjs/redux/redux.png)


The *`redux`* pattern brings three more principles on *`flux`*:
1. Single source of truth
2. State is read-only
3. Changes are made with pure functions

> Single source of truth

Only one *store* object hold all data as its state and accessible by *getState()*.

> State is read-only

The state (data) is only mutatable by issuing *actions* to the store. like:
```javascript
store.dispatch({
   type: 'ADD_GROCERY_ITEM',
   item: { productName: 'Milk' } 
});
store.dispatch({
   type: 'REMOVE_GROCERY_ITEM',
   index: 3
});
```

> Changes are made with pure functions

When *action*s are just plain objects describing what happened with related data, a *reducer* is responsible for handling the *action*s on receipt. 
```javascript
function groceryItemsReducer(state, action) {
   switch (action.type) {
     case 'ADD_GROCERY_ITEM':
        return object.assign({}, state, {
           groceryItems: [
              action.item,
              …state.groceryItems
           ]
        };
     default:
        return state;
   }
}
```

> Optional *middleware*s

Functions wrapped around the *`dispatching`* phrase in the data flow pipeline, **NOT** the *reducer*.



---

[flux_redux]: https://www.dotnetcurry.com/reactjs/1356/redux-pattern-tutorial