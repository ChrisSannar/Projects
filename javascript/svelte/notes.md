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
```

```

