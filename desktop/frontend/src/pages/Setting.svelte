<script lang="ts">
  import Layout from "../components/Layout.svelte";
  import {
    DownloadImages,
    GetConfig,
    SetConfig
  } from "../../wailsjs/go/main/App.js";

 import { onMount } from "svelte";

import { WindowReload } from "../../wailsjs/runtime";

let imageCategory = ""
let totalImageCount = ""
let imageInterval = ""
let message = ""


function handleSaveSetting() {
  if(imageCategory === "" || totalImageCount === "" || imageInterval === "") {
    return;
  }

  const conf = {
    ImageCategory: imageCategory,
    TotalImage: Number(totalImageCount)
  }

  /*DownloadImages(conf).then((res) => {
    console.log(res)
  })*/

  SetConfig(conf)
  message = "Config updated successfully"
}

onMount(async () => {
 const conf = await GetConfig()
  imageCategory = conf.ImageCategory
  totalImageCount = conf.TotalImage
  imageInterval = conf.Interval
})

</script>

<template>
  <Layout>
    <h1 class="font-bold text-gray-600 dark:text-white">Configuration</h1>
    <div class="layout">
      <div class="form">
        <div class="selection">
          <h3 class="my-5"> Select Folder </h3>

        <button class="text-gray-100 py-2 px-5 mb-3 rounded-md bg-gray-500" on:click={handleSaveSetting}> Click to open</button>
        </div>
      </div>

      <div class="margin-top form">
        <div class="selection">
          <div> 
            <label for="imagepath"> Image query</label>
            <input type="text" class="border border-gray-400 outline-none p-2 rounded-md" id="image-query" bind:value={imageCategory}/>
          </div>
          <div class="my-2">
            <label for="imagepath"> Total image </label>
            <input type="number" class="border border-gray-400 outline-none p-2 rounded-md" id="image-query"  bind:value={totalImageCount}/>
          </div>

          <div>
            <label for="imagepath"> Image change interval </label>
            <input type="text" class="border border-gray-400 p-2 rounded-md outline-none" id="image-query" bind:value={imageInterval} />
          </div>
        </div>

        <button class="text-gray-100 py-2 px-5 mb-3 rounded-md bg-gray-500" on:click={handleSaveSetting}> Save Setting </button>
            {#if message !== ""}
              <p class=""> Configuration saved </p>
            {/if}
      </div>
    </div>
  </Layout>
</template>

<style>

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
    padding: 30px 20px 20px 20px;
  }

  .selection label {
    display: block;
  }

</style>
