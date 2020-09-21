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
<button on:click={clickFunc}>click me</button>
```

Reactive declaractions: Basically compted values that are watched and updated. They can be single values, blocks, and conditionals
```
$: doubled = count * 2;
$: if (count > 10) { console.log("Too Much!"); count = 9; }
```

Reactivity is triggered by assignments
    Whatever is on the lefthand side gets updated. There are no object bindings