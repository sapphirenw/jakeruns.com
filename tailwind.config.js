/** @type {import('tailwindcss').Config} */
module.exports = {
  mode: 'jit',
  content: ['./templates/**/*.html'],
  theme: {
    extend: {
      colors: {
        txt: {
          DEFAULT: '#1A1E23',
          '50': '#5B697B',
          '100': '#576475',
          '200': '#4E5A69',
          '300': '#45505E',
          '400': '#3D4652',
          '500': '#343C46',
          '600': '#2B323A',
          '700': '#23282F',
          '800': '#1A1E23',
          '900': '#020203'
        },
      }
    }
  },
  daisyui: {
    themes: [
      {
        base_theme: {
          "primary": "#896bff",
          "secondary": "#896bff",
          "accent": "#AD99FF",
          "neutral": "#E2E4ED",
          "base-100": "#F2F3F7",
          "base-200": "#EFF0F5",
          "base-300": "#E9EAF1",
          "info": "#FAFAFC",
          "success": "#418a2f",
          "warning": "#fbbd23",
          "error": "#D1495B",
        },
      },
    ],
  },
  plugins: [
    require('@tailwindcss/typography'),
    require("daisyui"),
  ],
}

