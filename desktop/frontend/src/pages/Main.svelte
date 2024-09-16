<script lang="ts">
  import { GetDownloadedImages } from "../../wailsjs/go/main/App.js";
  import { BrowserOpenURL } from "../../wailsjs/runtime";
  import rpc from "../rpc";
  import { replace } from "svelte-spa-router";
  import Modal from "../components/Modal.svelte";
  import { imagePathStore } from "../store/app";
  import { onMount } from "svelte";

  let images: string[] = [];
  let path: string;

  imagePathStore.subscribe((value) => {
    path = value;
  });

  function getImages(): void {
    GetDownloadedImages().then((result: any) => {
      images = result;
    });
  }

  function dispatcher(image: string) {
    imagePathStore.set(image);
  }

  function openCreditUrl(name: string) {
    BrowserOpenURL("https://unsplash.com/" + name);
  }

  onMount(() => {
    getImages();

    rpc.on("shortcut.page.setting", () => {
      replace("/setting");
    });
  });
</script>

<template>
  <div class="image-config">
    <button>Config </button>
  </div>

  <section class="container">
    {#each images as image}
      <div
        class="image"
        on:click={() => dispatcher(image)}
        on:keydown={() => dispatcher(image)}
      >
        <img
          src={image.toString()}
          width="300"
          alt=""
          height="250"
          class="image-placeholder"
        />
        <div class="caption">
          credit: <button on:click={() => openCreditUrl("name")}
            >Abbey photo</button
          >
        </div>
      </div>
    {/each}

    <Modal />
  </section>
</template>

<style>
  .image-config {
    background: #ddd;
    padding: 10px;
    text-align: left;
  }
  .container {
    width: 95%;
    margin: 20px auto;
    overflow: hidden;
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
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
