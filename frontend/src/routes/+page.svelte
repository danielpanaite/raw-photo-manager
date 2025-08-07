<script lang="ts">
  import { DeleteFiles, GetOS, ListDir, PickFolder } from '$lib/wailsjs/go/main/App';
  import { onMount } from 'svelte';

  let files: string[] = [];
  let error = '';
  let selected: string[] = [];
  let path = 'C:\\Users\\Public\\Pictures';
  // let systemos = '';

  onMount(async () => {
    try {
      const result = await ListDir(path);
      files = result;
    } catch (e) {
      if (e instanceof Error) {
        error = e.message;
      }
    }
  });

  // GetOS().then(result => {
  //   systemos = result;
  // }).catch(error => {
  //   systemos = 'Error: ' + error;
  // });

  function toggleCheckbox(value: string) {
    if (selected.includes(value)) {
      selected = selected.filter((v) => v !== value);
    } else {
      selected = [...selected, value];
    }
  }

  async function submitForm() {
    try {
      await DeleteFiles(selected)
      selected = [];
    } catch (e) {
      if (e instanceof Error) {
        error = e.message;
      }
    }
    try {
      const result = await ListDir(path);
      files = result;
    } catch (e) {
      if (e instanceof Error) {
        error = e.message;
      }
    }
  }

  async function selectFolder() {
    path = await PickFolder();
    try {
      const result = await ListDir(path);
      files = result;
    } catch (e) {
      if (e instanceof Error) {
        error = e.message;
      }
    }
  }
</script>

<h1>Files in Directory {path}</h1>

{#if error}
  <p class="text-danger">{error}</p>
{:else if files.length === 0}
  <p>No files found.</p>
  <button type="button" class="btn btn-primary mt-2" on:click={selectFolder}>Browse <i class="bi bi-folder"></i></button>
{:else}
  <form on:submit|preventDefault={submitForm}>
  <ul class="list-group">
    {#each files as file, i}
      <li class="list-group-item">
        <input class="form-check-input me-1" type="checkbox" on:change={() => toggleCheckbox(file)} bind:group={selected} value="{file}" id="{String(i)}">
        <label class="form-check-label" for="{String(i)}">{file}</label>
      </li>
    {/each}
  </ul>
  <button type="button" class="btn btn-primary mt-2" on:click={selectFolder}>Browse <i class="bi bi-folder"></i></button>
  <button type="submit" class="btn btn-danger mt-2">Delete <i class="bi bi-trash"></i></button>
  </form>
{/if}
