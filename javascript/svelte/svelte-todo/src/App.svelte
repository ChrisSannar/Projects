<script>
  import List from "./components/List.svelte";
  import axios from "axios";

  axios
    .post("http://localhost:4000/graphql", { query: "{ hello }" })
    .then((resp) => {
      console.log("RESP", resp);
    });
</script>

<style>
  main {
    text-align: center;
    padding: 1em;
    max-width: 240px;
    margin: 0 auto;
  }

  h1 {
    color: #ff3e00;
    text-transform: uppercase;
    font-size: 4em;
    font-weight: 100;
  }

  @media (min-width: 640px) {
    main {
      max-width: none;
    }
  }
</style>

<main>
  <List />
</main>
