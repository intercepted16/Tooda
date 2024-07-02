<script lang="ts">
  import "./app.css";
  import { AddTodo, GetTodos, RemoveTodo, CheckTodo } from "../wailsjs/go/main/App.js";
  import { onMount } from "svelte";
  let todos: { [key: string]: boolean } = {};
  onMount(async () => {
    todos = await GetTodos();
  });
  let input: HTMLInputElement;
  let errors = false;
  $: {
    if (input) {
      if (errors) {
        input.classList.add("input-error");
      } else {
        input.classList.remove("input-error");
      }
    }
  }
  const onCheckBoxClick = async (e: EventTarget) => {
    if (!(e instanceof HTMLInputElement)) {
      console.error("Event target is not an HTMLInputElement");
      return 1;
    }
    await CheckTodo(e.value);
    todos = await GetTodos();
    console.log(todos);
  };
</script>

<main class="mx-auto mt-2">
  <div class="flex flex-col justify-center items-center mx-auto">
    <h2 class="text-2xl font-bold">Todos</h2>
    <ul class="flex flex-col items-center">
      {#each Object.entries(todos) as [todo, checked]}
        <li
          class="sm:space-x-2 items-center justify-center flex flex-col p-2 space-y-2 sm:space-y-2">
          <label class="flex items-center space-x-2 mx-auto justify-items-center">
            <!-- Really annoying Sv -->
            <input
              class="checkbox"
              type="checkbox"
              value={todo}
              {checked}
              on:click={(e) => onCheckBoxClick(e.target)} />
            <p>{todo}</p>
          </label>
          <button
            class="btn variant-filled-error"
            on:click={async () => {
              await RemoveTodo(todo);
              todos = await GetTodos();
            }}>Remove todo</button>
        </li>
      {/each}
    </ul>
  </div>

  <form
    class="space-y-2 flex flex-col mt-4"
    on:submit={async (event) => {
      event.preventDefault();
      if (input.value.length === 0) {
        errors = true;
        return;
      }
      await AddTodo(input.value);
      todos = await GetTodos();
      input.value = ""; // Clear the input after adding a todo
    }}>
    <label for="todo" class="label text-lg font-bold">Add a new todo</label>
    <input
      class="input"
      bind:this={input}
      on:input={() => {
        console.log("hi");
        console.log(input.value.length);
        if (input.value.length === 0) {
          errors = true;
        } else {
          errors = false;
        }
      }}
      title="Input"
      type="text"
      placeholder="Enter your todo here..." />
    <button class="btn variant-filled-primary disabled:scale-[100%]" disabled={errors} type="submit"
      >Add todo</button>
  </form>
</main>

<style>
  main {
    max-width: 600px;
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  }

  ul {
    list-style: none;
    padding: 0;
  }
</style>
