<script lang="ts">
  import Layout from "../components/Layout.svelte";
  import {
    DownloadImages,
    GetConfig,
    SetConfig,
    OpenDirDialogWindow,
    MessageDialog,
  } from "../../wailsjs/go/main/App.js";

  import { onMount } from "svelte";

  import { WindowReload } from "../../wailsjs/runtime";

  let imageCategory = "";
  let totalImageCount = 0;
  let imageInterval = "";
  let message = "";
  let defaultPath = "";
  let apikey = "xksdflksjdf";

  async function handleSaveSetting() {
    if (imageCategory === "" || totalImageCount === 0 || imageInterval === "") {
      return;
    }

    const conf = {
      ImageCategory: imageCategory,
      TotalImage: Number(totalImageCount),
      DefaultPath: defaultPath,
    };

    /*DownloadImages(conf).then((res) => {
    console.log(res)
  })*/

    SetConfig(conf);
    try {
      const msg = await MessageDialog("Config updated successfully");
    } catch (e) {
      const error = e instanceof Error ? e.message : "Unknown error";
      message = error;
    }
  }

  onMount(async () => {
    const conf = (await GetConfig()) as any;
    imageCategory = conf.ImageCategory;
    totalImageCount = conf.TotalImage;
    imageInterval = conf.Interval;
    defaultPath = conf.DefaultPath;
  });

  const handleSelectFolder = async () => {
    const path = await OpenDirDialogWindow();
    if (path === "") {
      defaultPath = defaultPath;
    } else {
      defaultPath = path;
    }
    //isFolderSelected = true;
  };

  const handleRestoreSetting = async () => {
    const conf = {
      ImageCategory: "nature",
      TotalImage: 10,
      DefaultPath: ".picasa/images",
    };
    imageCategory = conf.ImageCategory;
    totalImageCount = conf.TotalImage;
    defaultPath = conf.DefaultPath;
    SetConfig(conf);
    try {
      await MessageDialog("Config restored successfully");
    } catch (e) {
      const error = e instanceof Error ? e.message : "Unknown error";
      message = error;
    }
  };
</script>

<template>
  <Layout>
    <h1 class="font-bold text-gray-600 dark:text-white">Configuration</h1>
    <div class="layout">
      <div
        class="mt-4 text-[#999] border border-gray-200 dark:border-gray-500 rounded-md
"
      >
        <div class="selection">
          <div>
            <label for="imagepath"
              >Image path <p class="italic">{defaultPath}</p></label
            >
            <button
              class="text-gray-100 py-2 px-5 mb-3 rounded-md bg-gray-500"
              on:click={handleSelectFolder}
            >
              Change folder</button
            >
          </div>
          <div>
            <label for="imagepath"> Image query</label>
            <input
              type="text"
              class="border border-gray-400 w-6/12 outline-none p-2 rounded-md"
              spellcheck="false"
              autocomplete="off"
              id="image-query"
              bind:value={imageCategory}
            />
          </div>
          <div class="my-2">
            <label for="imagepath"> Total image </label>
            <input
              type="number"
              class="border border-gray-400 w-6/12 outline-none p-2 rounded-md"
              id="image-query"
              bind:value={totalImageCount}
            />
          </div>
          <div>
            <label for="imagepath">
              Image change interval <span class="text-sm"
                >(e.g: 5m, 10m, 1h, 5h)
              </span></label
            >
            <input
              type="text"
              class="border border-gray-400 p-2 w-6/12 rounded-md outline-none"
              spellcheck="false"
              autocomplete="off"
              id="image-query"
              bind:value={imageInterval}
            />
          </div>
          <div class="mt-2">
            <label for="imagepath"> Unsplash API key</label>
            <input
              type="text"
              class="border border-gray-400 p-2 w-6/12 rounded-md outline-none"
              spellcheck="false"
              autocomplete="off"
              id="image-query"
              bind:value={apikey}
            />
          </div>

          <div class="mt-5">
            <button
              class="text-gray-100 py-2 px-10 mb-3 rounded-md bg-gray-500"
              on:click={handleSaveSetting}
            >
              Save
            </button>
            <button
              class="text-gray-100 py-2 px-10 mb-3 rounded-md bg-gray-500"
              on:click={handleRestoreSetting}
            >
              Restore config
            </button>
          </div>
        </div>
      </div>
    </div>
  </Layout>
</template>

<style>
  .layout {
    width: 40rem;
    margin: auto;
  }

  .selection {
    padding: 30px 20px 20px 20px;
  }

  .selection label {
    display: block;
  }
</style>
