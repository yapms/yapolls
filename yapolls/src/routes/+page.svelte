<script lang="ts">
	import { resolve } from '$app/paths';
	import type { PageProps } from './$types';
	import { page } from '$app/state';
	import { goto } from '$app/navigation';
	import { SvelteURLSearchParams } from 'svelte/reactivity';
	let { data }: PageProps = $props();

	let pollster_filter: string = $state(page.url.searchParams.get('pollster') ?? '');
	let subject_filter: string = $state(page.url.searchParams.get('subject') ?? '');
	let poll_type_filter: string = $state(page.url.searchParams.get('poll_type') ?? '');

	function filter() {
		const params = new SvelteURLSearchParams(page.url.searchParams);
		params.set('pollster', pollster_filter);
		params.set('subject', subject_filter);
		params.set('poll_type', poll_type_filter);
		goto(resolve(`/?${params.toString()}`));
	}
</script>

<h1 class="navbar">
	<a href={resolve('/')} class="btn-chost btn text-xl"> YAPolls </a>
</h1>

<fieldset class="fieldset" onchange={filter}>
	<legend class="fieldset-legend">Pollsters</legend>
	<select class="select" bind:value={pollster_filter}>
		<option value="">None</option>
		{#each data.pollsters as pollster (pollster)}
			<option>{pollster}</option>
		{/each}
	</select>
</fieldset>

<fieldset class="fieldset" onchange={filter}>
	<legend class="fieldset-legend">Subjects</legend>
	<select class="select" bind:value={subject_filter}>
		<option value="">None</option>
		{#each data.subjects as subjects (subjects.subject)}
			<option>{subjects.subject}</option>
		{/each}
	</select>
</fieldset>

<fieldset class="fieldset" onchange={filter}>
	<legend class="fieldset-legend">Poll Types</legend>
	<select class="select" bind:value={poll_type_filter}>
		<option value="">None</option>
		{#each data.poll_types as poll_type (poll_type)}
			<option>{poll_type}</option>
		{/each}
	</select>
</fieldset>

<ul>
	{#each data.polls as poll (poll.id)}
		<li>{poll.id}</li>
	{/each}
</ul>

<footer class="footer bg-base-200 p-10 text-base-content">
	<aside>
		<p>This is an experimental website to aggregate polling data.</p>
	</aside>
</footer>
