import { defineConfig } from "cypress";

export default defineConfig({
  video: true,
  viewportWidth: 1280,
  viewportHeight: 900,
  e2e: {
    baseUrl: 'http://localhost:3000',
    setupNodeEvents(on, config) {
    },
  },
});
