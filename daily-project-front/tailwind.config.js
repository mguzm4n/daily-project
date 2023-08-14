/** @type {import('tailwindcss').Config} */
export default {
  darkMode: 'class',
  content: [
    "./index.html",
    "./src/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      keyframes: {
        wiggle: {
          "0%, %100": { transform: 'rotate(-0.5deg)'},
          "50%": { transform: 'rotate(0.5deg)'}
        }
      }
    },
  },
  plugins: [],
}

