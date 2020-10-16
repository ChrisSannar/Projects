<script>
    export let text;
    export let done;
</script>

<div>
    <input type="checkbox" bind:checked={done} />
    <input type="text" bind:value={text} disabled={done}/>
</div>