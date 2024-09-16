<script>
  import Layout from "../components/Layout.svelte";
</script>

<template>
  <Layout>
    <h1 class="config-header">Configuration</h1>
    <div class="layout">
      <div class="form">
        <div class="selection">
          <button class="btn"> Open Folder </button>
        </div>
      </div>

      <div class="margin-top form">
        <div class="selection">
          <div>
            <label for="imagepath"> Image Query</label>
            <input type="text" id="image-query" />
          </div>
          <div>
            <label for="imagepath"> Total Image </label>
            <input type="text" id="image-query" />
          </div>

          <div>
            <label for="imagepath"> Duration</label>
            <input type="text" id="image-query" />
          </div>
        </div>
      </div>

      <div class="margin-sm">
        <button class="btn"> Save Setting </button>
      </div>
    </div>
  </Layout>
</template>

<style>
  .config-header {
    color: #333;
  }

  .layout {
    width: 40rem;
    margin: auto;
  }

  .margin-top {
    margin-top: 3.8rem;
  }

  .margin-sm {
    margin-top: 2rem;
  }

  .form {
    color: #999;
    border: 1px solid #ddd;
    border-radius: 9px;
  }

  .selection {
    padding: 40px 20px;
  }

  .btn {
    padding: 10px;
    background: #eee;
    border: 1px solid #999;
    border-radius: 5px;
    font-weight: bold;
  }

  .selection label {
    display: block;
  }

  .selection input[type="text"] {
    padding: 5px;
    border: 1px solid #999;
  }
</style>
