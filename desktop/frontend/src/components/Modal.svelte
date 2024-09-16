<script type="ts">
  import { fade, slide } from "svelte/transition";
  import { imagePathStore } from "../store/app";
  let path: string = "";
  // export let isShowingModal: boolean;

  imagePathStore.subscribe((value) => {
    path = value;
  });

  const closeModal = () => {
    imagePathStore.update((value) => (value = ""));
  };

  const handleKeydown = (event: KeyboardEvent) => {
    if (event.key === "Escape") {
      closeModal();
    }
  };
</script>

<template>
  <div>
    {#if path !== ""}
      <div
        class="modal-background"
        on:click={closeModal}
        on:keydown={handleKeydown}
        transition:fade={{ duration: 300 }}
      >
        <div
          transition:fade={{ duration: 300 }}
          class="modal"
          on:click|stopPropagation
          on:keydown|stopPropagation
        >
          <button class="close-btn" on:click={closeModal}>Close</button>
          <img src={path} alt="" width="100%" height="100%" />
          <p></p>
        </div>
      </div>
    {/if}
  </div>
</template>

<style>
  .modal-background {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
    z-index: 1000;
    color: #999;
  }

  .modal {
    background-color: rgba(255, 255, 255, 0.7);
    padding: 0.9rem;
    border-radius: 8px;
    width: 700px;
    max-width: 100%;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.3);
  }

  .close-btn {
    background-color: #fff;
    border: none;
    color: white;
    padding: 0.5rem 1rem;
    cursor: pointer;
    border: 1px solid #999;
    color: #999;
  }
</style>
