/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: ["class"],
  content: [
    './internal/templates/**/*.{js,jsx,templ}',
  ],
  prefix: "",
  theme: {
    extend: {
    },
  },
  plugins: [
    require("tailwindcss-animate"),
    require("daisyui")
  ],
}