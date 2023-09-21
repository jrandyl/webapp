<script>
  import { onMount } from "svelte";
  import Navbar from "./lib/Navbar.svelte";

  let data = "";

  onMount(async () => {
    try {
      const response = await fetch("http://localhost:4000/api");
      if (response.ok) {
        data = await response.text();
      } else {
        console.error("Failed to fetch data");
      }
    } catch (error) {
      console.error("Error fetching data:", error);
    }
  });
</script>

<main>
  <Navbar />
</main>

<style>
  main {
    font-size: 10px;
    background: var(--bodycolor1);
    height: 100vh;
  }
</style>
