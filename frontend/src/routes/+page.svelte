<script lang="ts">
    let name = $state("");
    let message = $state("");

    $effect(async () => {
        if (name === "") {
            message = "";
            return;
        }

        const res = await fetch('http://localhost:8080/greet', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name })
        });

        const data = await res.json();
        message = data.message;
    });
</script>

<main>
    <h1>Greet As You Type</h1>
        <input type="text" bind:value={name} placeholder="Type your name…" autocomplete="off"/>
    {#if message}
        <p>{message}</p>
    {/if}
</main>

<style>
main {
    font-family: sans-serif;
    padding: 2rem;
}

input {
    width: 100%;
    max-width: 300px;
    padding: 0.5rem;
    font-size: 1rem;
}
</style>

