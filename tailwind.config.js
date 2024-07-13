/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: "selector",
  content: ["./**/*.templ"],
  ignoreFiles: ["node_modules/**", "bin/**"],
  theme: {
    extend: {},
  },
  plugins: [],
};
