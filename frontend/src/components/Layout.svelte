<script>
  import { link } from "svelte-spa-router";

  import IconifyIcon from "@iconify/svelte";

  let defaultLink = "/";
</script>

<template>
  <div>
    <div class="mt-0 p-10 text-left">
      <a
        class="no-underline pl-3 text-gray-600 dark:text-white flex"
        href={""}
        use:link={defaultLink}
      >
        <IconifyIcon
          icon="bi:arrow-left"
          style="position: relative; top: 3px;padding-right: 5px;"
        />Back
      </a>
    </div>
    <slot />
  </div>
</template>

<style>
  .nav-link {
    margin-top: 30px;
    padding: 20px;
    text-align: left;
    font-size: 18px;
  }
</style>
