import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
    server: {
        port: 4200, //default is 5173
        allowedHosts: ["nanacatlan.com"],
    }
});
