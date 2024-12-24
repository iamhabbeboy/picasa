<script lang="ts">
  import {
    GetDownloadedImages,
    SelectImageDir,
    DownloadImages,
  } from '../../wailsjs/go/main/App.js';
  import { BrowserOpenURL } from '../../wailsjs/runtime';
  import rpc from '../rpc';
  import { replace, link } from 'svelte-spa-router';
  import Modal from '../components/Modal.svelte';
  import { imagePathStore } from '../store/app';
  import { onMount } from 'svelte';
  import ImageConfig from '../components/ImageConfig.svelte';
  import DownloadImage from '../../src/assets/images/download.svg';
  import ConfigImage from '../../src/assets/images/config.svg';

  import LoaderImage from '../../src/assets/images/loader.svg';

  let images: string[] = [];
  let path: string;
  let isFolderSelected: boolean = false;
  let isLoading = true;

  imagePathStore.subscribe((value) => {
    path = value;
  });

  function dispatcher(image: string) {
    imagePathStore.set(image);
  }

  function openCreditUrl(name: string) {
    BrowserOpenURL('https://unsplash.com/' + name);
  }

  async function handleOpenSelectFolder() {
    const path = await SelectImageDir();
    images = path;
    isFolderSelected = true;
  }

  async function downloadImages() {
    try {
      isLoading = true;
      const res = await DownloadImages();
    } catch (e) {
    } finally {
      const result = await GetDownloadedImages();
      images = result ?? [];
      isLoading = false;
    }
  }

  onMount(async () => {
    const result = await GetDownloadedImages();
    images = result ?? [];
    isLoading = false;

    rpc.on('shortcut.page.setting', () => {
      replace('/setting');
    });
  });
</script>

<template>
  <div class="image-config border-b dark:border-gray-600 flex justify-between">
    <a
      href="#"
      on:click={downloadImages}
      class="dark:text-gray-50 text-gray-500 text-xs flex"
    >
      <img
        src={DownloadImage}
        width="15"
        alt=""
        class="mr-1 dark:brightness-0 dark:invert-[1]"
      /> Download images</a
    >
    <a
      href="/setting"
      class=" text-xs flex text-gray-500 dark:text-gray-50"
      use:link
    >
      <img
        src={ConfigImage}
        width="15"
        alt=""
        class="mr-1 dark:brightness-0 dark:invert-[1]"
      /> Config
    </a>
  </div>
  <section class=" w-[95%] mx-auto mt-3">
    {#if isLoading}
      <div class="mx-auto w-48 h-screen flex justify-center items-center">
        <img src={LoaderImage} alt="" />
        <h4>Loading...</h4>
      </div>
    {:else}
      <div class="flex flex-wrap justify-between">
        {#each images as image}
          <div
            class="bg-gray-800 w-[300px] h-[300px] mb-3 cursor-pointer"
            on:click={() => dispatcher(image)}
            on:keydown={() => dispatcher(image)}
          >
            <img
              src={image.toString()}
              alt=""
              class="object-cover h-[100%] w-[100%]"
            />
            <!--<div class="caption">
            credit: <button on:click={() => openCreditUrl("name")}
              >Abbey photo</button
            >
          </div>-->
          </div>
        {/each}
      </div>
    {/if}
    <Modal />
  </section>

  {#if images.length === 0}
    <ImageConfig on:click={handleOpenSelectFolder} />
  {/if}
</template>

<style>
  .image-config {
    padding: 10px;
    text-align: left;
  }

  .image {
    background: transparent;
    margin: 5px 2px;
    position: relative;
  }

  .image-placeholder {
    max-width: 100%;
    max-height: 100%;
    object-fit: cover;
    cursor: pointer;
    transition: 0.5s all;
  }

  .image-placeholder:hover {
    opacity: 0.4;
  }

  .caption {
    background: rgba(0, 0, 0, 0.5);
    position: absolute;
    bottom: 5px;
    padding: 10px;
    margin: 0px;
    font-size: 12px;
    color: #999;
  }

  .caption button {
    text-decoration: underline;
    color: #ddd;
    background: transparent;
    border: 0px;
    cursor: pointer;
  }

  .caption button:hover {
    text-decoration: none;
  }
</style>
