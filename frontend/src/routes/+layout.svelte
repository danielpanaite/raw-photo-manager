<script lang="ts">
	import Header from './Header.svelte';
	import favicon from '$lib/assets/favicon.svg';
	import { Styles } from '@sveltestrap/sveltestrap';
	import 'bootstrap/dist/css/bootstrap.min.css';
 	import { onMount } from 'svelte';

	let { children } = $props();

  onMount(async () => {
    await import('bootstrap/dist/js/bootstrap.bundle.min.js');
  });
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
</svelte:head>

<div class="app">
	<Styles />
	<Header />

	<main>
		{@render children()}
	</main>

	<footer>
		<p>
			check out our <a href="https://github.com/danielpanaite/raw-photo-manager">github</a> repo!
		</p>
	</footer>
</div>

<style>
	.app {
		display: flex;
		flex-direction: column;
		min-height: 100vh;
	}

	main {
		flex: 1;
		display: flex;
		flex-direction: column;
		padding: 1rem;
		width: 100%;
		max-width: 64rem;
		margin: 0 auto;
		box-sizing: border-box;
	}

	footer {
		display: flex;
		flex-direction: column;
		justify-content: center;
		align-items: center;
		padding: 12px;
	}

	footer a {
		font-weight: bold;
	}

	@media (min-width: 480px) {
		footer {
			padding: 12px 0;
		}
	}
</style>