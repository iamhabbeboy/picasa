<script type="ts">
  import Image from "../assets/images/no-image.svg";

  async function handleDownloadImages() {
    // --
    console.log("Hello, world");
  }
</script>

<template>
  <div class="no-image-loaded">
    <div>
      <img src={Image} alt="" width="300" />
      <p>&nbsp;</p>
      <button on:click> Select a folder </button>
      <p>- OR -</p>
      <div>
        <label for="category">Enter Image Category</label>
        <p></p>
        <input
          type="text"
          placeholder="E.g: cars, technology, nature"
          id="category"
        />

        <p></p>

        <button on:click={handleDownloadImages}> Download Image</button>
      </div>
    </div>
  </div>
</template>

<style>
  .no-image-loaded {
    color: #999;
    display: flex;
    justify-content: center;
    align-items: center;
    height: 100vh;
  }
</style>
