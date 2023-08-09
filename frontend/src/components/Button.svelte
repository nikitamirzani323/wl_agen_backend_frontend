<script>
    import { createEventDispatcher } from "svelte";
    export let button_title = ""
    export let button_css = ""
    export let button_function = ""

    let dispatch = createEventDispatcher()
</script>

<button
    on:click={() => dispatch('click',button_function)}
    type="button"
    class="btn {button_css} btn-sm">
    {button_title}
</button>