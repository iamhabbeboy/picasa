/** @type {import('tailwindcss').Config} */
export default {
  darkMode: "media",
  content: ["./index.html", "./src/**/*.{svelte,js,ts,jsx,tsx}"],
  theme: {
    extend: {
      filter: {
        customFilter: "brightness(0) invert(1)",
      },
    },
  },
  plugins: [],
};
