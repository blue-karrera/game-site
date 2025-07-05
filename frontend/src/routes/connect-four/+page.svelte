<script lang="ts">
    import { fly } from 'svelte/transition';
    import { cubicOut } from 'svelte/easing';

    let rows = 6;
    let cols = 7;

    let relative_tile_spacing = 0.5;


    let board: number[][] = $state(Array.from({ length: rows }, () => Array(cols).fill(0)));
    let currentPlayer = $state(true); //player1 true, player2 false
    let win = $state(false);

    async function drop(col: number) {
        if (win) {
            return;
        }

        try {
            const res = await fetch('/connect-four/drop', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify({ player: currentPlayer, column: col })
            });
            const data = await res.json();
            win = data.win;
            currentPlayer = data.currentPlayer;
            board = data.board;
        } catch (e) {
            console.error(e);
        }
    }
</script>



<div>
    <h2>Connect Four</h2>
    <p>
        Current turn:
        <strong style="color: {currentPlayer === true ? 'red' : 'yellow'}">
            {currentPlayer === true ? 'Red' : 'Yellow'}
        </strong>
    </p>
    {#if win}
        <p style="color: green;"><strong>Player {currentPlayer === true ? 2 : 1} wins!</strong></p>
    {/if}
    <div class="board">
    </div>
    
    <div class="back">
        <div class="base">
            <div class="hole">ONE</div>
            <div class="hole">TWO</div>
        </div>
    </div>

    <!--
    <div class="board" style="--rows: {rows}; --cols: {cols}; width: calc({cols} * 64px); height: calc({rows} * 64px)">
        {#each Array(rows * cols) as _, idx (idx)}
            <div class="slot"></div>
        {/each}

        {#each board as row, r (r)}
            {#each row as cell, c (c)}
                {#if cell !== 0}
                    <div class="token player{cell}" style="left: calc({c} * 64px); top: calc({r} * 64px)"></div>
                {/if}
            {/each}
        {/each}

        {#each Array(cols) as _, c (c)}
            <div class="column-overlay" style="left: calc({c} * calc(100% / {cols}))" onclick={() => { const targetRow = getDropRow(c); if (targetRow >= 0) drop(c); }} />
        {/each}
    </div>
    -->

</div>



<style>
    .board {
        position: relative;
        display: grid;
        grid-template-rows: repeat(var(--rows), 60px);
        grid-template-columns: repeat(var(--cols), 60px);
        gap: 4px;
        background-color: blue;
        padding: 4px;
        border-radius: 8px;
    }

    .column-overlay {
        position: absolute;
        top: 0;
        bottom: 0;
        width: calc(100% / var(--cols));
        cursor: pointer;
    }

    .slot {
        width: 60px;
        height: 60px;
        border-radius: 50%;
        background-color: white;
        z-index: 1;
    }

    .token {
        position: absolute;
        width: 56px;
        height: 56px;
        border-radius: 50%;
        margin: 2px;
        z-index: 0;
    }

    .player1 {
        background-color: red;
    }
    .player2 {
        background-color: yellow;
    }




.back {
	background-color: lightblue;
	width: 400px;
	height: 300px;
	/*background-image: repeating-linear-gradient(45deg, white 0px,lightblue 40px);*/
}

.base {
	position: relative;
	left: 10px;
	top: 10px;
	width: 200px;
	height: 200px;
	border: solid 1px black;
	background-color: white;
	mix-blend-mode: hard-light;
}

.hole {
	width: 80px;
	height: 50px;
	margin: 10px;
	border: solid 1px red;	
	border-radius: 10px;
	background-color: gray;
}
</style>
