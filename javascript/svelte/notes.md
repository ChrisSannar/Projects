# Svelte

## Why Svelte?

Because it focuses its functionality on compiling instead of acutally running. That way we only have dev dependencies and the resulting file is just a bundle instead of some program we have to start.

This results in a very small and fast application. 

## Concpets

To disiplay things on the page, use the variable name between curly braces. The same can be applied to attributes
```
<p>{name}</p>
<img {src}>
```

Style tags are scoped and the compiler removes all the unused styles.

Components are similar to React, import them in and have the name of it on the page.
```
<Nested />
```

The use @html tag inside the in-page rendering to render html, but make sure to sanitize
```
{ @html text }
```

Event listeners are similar to Vue:
```
<button on:click={clickFunc}>
    <Child on:click /> // passes the event up
</button>
```

Reactive declaractions: Basically compted values that are watched and updated. They can be single values, blocks, and conditionals
```
$: doubled = count * 2;
$: if (count > 10) { console.log("Too Much!"); count = 9; }
```

Reactivity is triggered by assignments
    Whatever is on the lefthand side gets updated. There are no object bindings

## Binding

we can bind a value when it comes to various types of input:
```
<input bind:value={name}>
<input bind:group={list}>   // You can bind multiple inputs together
{#each things as thing}
    <input bind:value={thing.text}> // You can also bind values in arrays and loops
{/each}
```

## Lifecycle

onMount runs right after the component is rendered to the page. This is the place to put any `fetch` of data
```
import { onMount } from 'svelte';
onMount(async () => fetch('...')}
```

onDestroy runs right when the component is... well destroyed. Used to contain memory leaks.
It also doesn't matter where you call `onDestory`, so it could be used inside a util function.
``` ./util.js
export function asdf() { onDestroy(() => { ... })}
```

beforeEach/afterEach - Good for handling things before the need to happen...?
tick - moves to the next set up?

## Stores

An object that we can subscribe to and update whenever it's imported
```
import { writable } from 'svelte/store';
const count = writable(0);  // Set the store
const unsub = count.subscribe(val => {    // Subscribe to it
    ...
});
count.update(n => n + 1);   // Trigger an update and subscribe will grab it
onDestroy(unsub);   // To prevent memory leaks
```

We can also auto subscribe using `$`, pulling from the example above
```
import { count } from './stores.js'
<p>The count is { $count }</p>
```
**Importing the count needs to be at the top level scope (first imported)**
**Any variable declared with `$` will be treated as a store**

We can also create a "Read Only" store
```
import { readable } from 'svelte/store'
export const time = readable(initialVal, function start(set) {  // Start from the first subscriber
    ... set(...)
    return function stop() {}    // Stop from the last unsubscribe    
});
```

There are "Derived stores" that take their values and update based on the value from another store
``` // Pulling from above example
export const tooLong = derived(
    time,   // Store we're deriving
    $time => Math.round($time - start) > 10000  // function returns the derived value
    // If our time is over 10s, then it's been too long
) 
```

As long as we correctly implement subscribe, we can have anything we want with the store.
This results in custom objects with defined behaviors we can import in (like React Hooks).
```
const { subscribe, update } = writable(0);
return {
    subscribe,
    change: () => update(...)
}
```

If a store has a set method (writable), you can bind to it and change the value
```
const name = writable('name');
{name}
<input bind:value={$name}>
<button on:click="{() => $name += '!'}">!</button>
```



