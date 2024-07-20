/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: "selector",
  content: ["./**/*.templ", "./**/*.html"],
  ignoreFiles: ["node_modules/**", "bin/**"],
  theme: {
    extend: {
      backgroundImage: {
        "coding-pattern": "url('/public/bg-img-svg.svg')",
        "fade-pattern": "linear-gradient(0deg, rgba(241,245,249,1) 0%, rgba(241,245,249,0) 40%)",
        "fade-pattern-dark": "linear-gradient(0deg, rgba(23,23,23,1) 0%, rgba(23,23,23,0) 40%)",
      },
    },
  },
  plugins: [],
};
