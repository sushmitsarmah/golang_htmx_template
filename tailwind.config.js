/** @type {import('tailwindcss').Config} */

const daisyui = require("daisyui");

module.exports = {
  content: ["./internal/templates/**/*.templ"],
  theme: {
    extend: {},
  },
  plugins: [daisyui],
}

