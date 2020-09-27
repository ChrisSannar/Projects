<script>
    import App from '../App.svelte';
import Todo from './Todo.svelte';

    $: todos = [
        { id: 1, text: "I'm one thing", done: false },
        { id: 2, text: "I'm another thing", done: false },
        { id: 3, text: "I'm the last thing", done: false },
    ];
</script>

<div>
<!-- <input type="text" placeholder="Search" on:change={e => todos.filter(todo => todo.text.includes(e.value))}> -->
{#each todos.filter(todo => !todo.done) as todo} 
    <Todo bind:text={todo.text} bind:done={todo.done} />
{/each}
<button on:click={e => todos = [ ...todos, { id: todos[todos.length - 1].id + 1, text: "", done: false } ]}>Add Todo</button>
{#each todos.filter(todo => todo.done) as todo} 
    <Todo bind:text={todo.text} bind:done={todo.done} />
{/each}
</div>