<script>
	let value = $state('123+1');
	let opertatorUsed = $state(false);
	function add(char) {
		value = value + char;
	}
	async function requestResult() {
		let bodyObj = { foo: value };
		const req = new Request('http://localhost:8080/calculate', {
			method: 'PUT',
			body: JSON.stringify(bodyObj)
		});

		await fetch(req)
			.then((res) => res.json())
			.then((data) => (value = data));
	}
</script>

<h1>This app is a Calculator</h1>

<textarea readonly name="" id="textarea">{value}</textarea>
<div class="buttonContainer">
	{#each [1, 2, 3, 4, 5, 6, 7, 8, 9, 0] as char}
		<button
			onclick={() => {
				add(char);
			}}
			class="number">{char}</button
		>
	{/each}

	{#each ['/', '*', '+', '-'] as operator}
		<button
			onclick={() => {
				value = value + operator;
				opertatorUsed = true;
			}}
			class="operation">{operator}</button
		>
	{/each}
	<button
		onclick={() => {
			value = '';
			opertatorUsed = false;
		}}
		class="clear">C</button
	>
	<button onclick={requestResult} class="result">=</button>
</div>

<style>
	textarea {
		width: 325px;
		resize: none;
		font-size: 2rem;
		height: 100px;
		background: gray;
		border: solid 2px black;
		text-align: end;
	}
	.buttonContainer {
		width: fit-content;
		height: fit-content;
		display: grid;
		grid-template-columns: repeat(5, 1fr);
		grid-template-rows: repeat(4, 1fr);
		background-color: black;
		padding: 0.5rem;
		border-radius: 0 0 10px 10px;
	}
	.buttonContainer > * {
		width: 60px;
		height: 60px;
		text-align: center;
		font-size: 1.5rem;
		border: none;
		border-radius: 10px;
		margin: 0.1rem;
	}
	.buttonContainer > *:hover {
		cursor: pointer;
	}
	.number {
		background-color: white;
	}
	.operation {
		background: gray;
	}
	.clear {
		background-color: orange;
	}
	.result {
		width: 100%;
		grid-column: 1/5;
		background-color: orangered;
	}
</style>
