require("esbuild").buildSync({
    entryPoints: [
        "./src/online-stats-app.ts",
        "./src/waitlist-stats-app.ts",
        "./src/organizers-companies-app.ts",
    ],
    bundle: true,
    minify: process.env.MINIFY === "true",
    outdir: "../public/assets/js",
});
